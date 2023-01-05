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
	util.SendCode(vCode)
	return dto.Success(code, "发送成功")
}

func (s *UserService) Register(ctx context.Context) *dto.Result {

	redisClient := conf.NewRedisClient()
	userDao := dao.NewUserDao(ctx)

	//检验密码格式
	if isTrue := util.VerifyPasswordFormat(s.Password); !isTrue {
		return dto.Fail(e.WrongAccountOrPassword, nil)
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
	//将token保存到redis
	redisClient.Set(e.UserLoginToken+s.Email, token, e.USerLoginTokenTTL)
	//返回用户信息
	userDto := dto.BuildUser(user, token)
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
	redisClient.Set(e.UserLoginToken+s.Email, token, e.USerLoginTokenTTL)
	//返回用户信息
	userDto := dto.BuildUser(user, token)
	return dto.Success(e.Success, userDto)

}
