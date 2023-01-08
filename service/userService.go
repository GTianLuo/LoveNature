package service

import (
	"context"
	"github.com/gin-gonic/gin"
	"lovenature/conf"
	"lovenature/dao"
	"lovenature/dto"
	"lovenature/model"
	"lovenature/pkg/e"
	"lovenature/pkg/util"
)

type UserService struct {
	Email    string `form:"email"`
	Password string `form:"password"`
	NickName string `form:"nickName"`
	Icon     string `form:"icon"`
	Code     string `form:"code"`
	Token    string `form:"token"`
}

func NewUserService() *UserService {
	return &UserService{}
}

func (s *UserService) SendCode() *dto.Result {
	redisClient := conf.NewRedisClient()
	code := e.Success

	//校验email的正确性
	if isTrue := util.VerifyEmailFormat(s.Email); !isTrue {
		code = e.InvalidEmail
		return dto.Fail(code, nil)
	}

	//检查验证码是否重复发送
	if cnt := redisClient.Exists(e.VerificationCodeKey + s.Email).Val(); cnt == 1 {
		//60s内已发送过验证码
		code := e.RepeatSending
		return dto.Fail(code, nil)
	}

	//获取随机验证码并发送
	vCode := util.RandomCode(6)
	redisClient.Set(e.VerificationCodeKey+s.Email, vCode, e.VerificationCodeKeyTTL)
	util.SendCode(vCode, s.Email)
	return dto.Success(code, "发送成功")
}

func (s *UserService) Register(ctx context.Context) *dto.Result {

	redisClient := conf.NewRedisClient()
	userDao := dao.NewUserDao(ctx)

	//检验密码格式
	if isTrue := util.VerifyPasswordFormat(s.Password); !isTrue {
		return dto.Fail(e.WrongPasswordFormat, nil)
	}

	//验证码校验
	vCode := redisClient.Get(e.VerificationCodeKey + s.Email).Val()
	if vCode != s.Code {
		return dto.Fail(e.WrongCode, nil)
	}

	//判断用户是否已经注册
	if isExist := userDao.IsExistByEmail(s.Email); isExist {
		return dto.Fail(e.RepeatRegister, nil)
	}

	//创建用户并持久化
	user := &model.User{
		Email:    s.Email,
		Password: util.Encryption(s.Password),
		NickName: util.GetNickName(),
	}
	if err := userDao.CreateUser(user); err != nil {
		return dto.Fail(e.Error, err)
	}
	return dto.Success(e.Success, "注册成功")
}

func (s *UserService) LoginByPassword(ctx *gin.Context) *dto.Result {
	userDao := dao.NewUserDao(ctx)
	redisClient := conf.NewRedisClient()

	//检验密码账号
	if isExist := userDao.IsExistByEmail(s.Email); !isExist {
		//邮箱未注册
		return dto.Fail(e.WrongAccountOrPassword, nil)
	}
	user := userDao.GetUser(s.Email)
	if user.Password != util.Encryption(s.Password) {
		//密码不正确
		return dto.Fail(e.WrongAccountOrPassword, nil)
	}
	//生成token
	token := util.NextToken()
	//利用token将用户的基本信息保存到redis中
	userDto := dto.BuildUser(user, token)
	redisClient.Del(e.UserLoginInfo + user.NickName) //防止用户重复登录导致生成多个token
	redisClient.HMSet(e.UserLoginInfo+user.NickName, util.StructToMap(userDto))
	redisClient.Expire(e.UserLoginInfo+user.NickName, e.UserLoginInfoTTL)
	//返回用户信息
	return dto.Success(e.Success, userDto)
}

func (s *UserService) LoginByCode(ctx *gin.Context) *dto.Result {

	userDao := dao.NewUserDao(ctx)
	redisClient := conf.NewRedisClient()

	//检验邮箱
	if isExist := userDao.IsExistByEmail(s.Email); !isExist {
		//邮箱未注册
		return dto.Fail(e.WrongAccountOrPassword, nil)
	}
	if vCode := redisClient.Get(e.VerificationCodeKey + s.Email).Val(); vCode != s.Code {
		//验证码错误
		return dto.Fail(e.WrongCode, nil)
	}
	user := userDao.GetUser(s.Email)
	//生成token
	token := util.NextToken()
	//将token保存到redis
	userDto := dto.BuildUser(user, token)
	redisClient.Del(e.UserLoginInfo + user.NickName) //防止用户重复登录导致生成多个token
	redisClient.HMSet(e.UserLoginInfo+user.NickName, util.StructToMap(userDto))
	redisClient.Expire(e.UserLoginInfo+user.NickName, e.UserLoginInfoTTL)
	//返回用户信息
	return dto.Success(e.Success, userDto)

}

func (s *UserService) Logout(ctx *gin.Context) *dto.Result {
	redisClient := conf.NewRedisClient()
	//删除redis中的登录状态
	redisClient.Del(e.UserLoginInfo + s.NickName)
	return dto.Success(e.Success, "退出成功")
}

func (s *UserService) UpdatePassword(ctx *gin.Context) *dto.Result {
	userDao := dao.NewUserDao(ctx)
	redisClient := conf.NewRedisClient()

	//校验验证码
	redisClient.Get(e.VerificationCodeKey + s.Email)
	//检查密码格式
	if isTrue := util.VerifyPasswordFormat(s.Password); !isTrue {
		return dto.Fail(e.WrongPasswordFormat, nil)
	}
	//修改密码
	if err := userDao.UpdatePassword(s.NickName, util.Encryption(s.Password)); err != nil {
		return dto.Fail(e.Error, err)
	}
	//删除当前用户的登录状态
	if err := redisClient.Del(e.UserLoginInfo + s.NickName).Err(); err != nil {
		return dto.Fail(e.Error, err)
	}
	return dto.Success(e.Success, "修改成功,请重新登录")
}
