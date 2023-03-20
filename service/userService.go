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
	if vCode != s.Code || vCode == "" {
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
	if vCode := redisClient.Get(e.VerificationCodeKey + s.Email).Val(); vCode != s.Code || vCode == "" {
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
	//检查密码格式
	if isTrue := util.VerifyPasswordFormat(s.Password); !isTrue {
		return dto.Fail(e.WrongPasswordFormat, nil)
	}
	//校验验证码
	if code := redisClient.Get(e.VerificationCodeKey + s.Email).Val(); code != s.Code || code == "" {
		return dto.Fail(e.WrongCode, nil)
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

func (s *UserService) UploadIcon(ctx *gin.Context) *dto.Result {
	userDao := dao.NewUserDao(ctx)
	redisClient := conf.NewRedisClient()
	iconFile, header, err := ctx.Request.FormFile("iconFile")
	defer iconFile.Close()
	if err != nil {
		return dto.Fail(e.InvalidParam, err)
	}
	//校验文件
	if header.Size > (8 << 18) {
		return dto.Fail(e.IconTooBig, nil)
	}
	if typ := header.Header.Get("Content-Type"); typ != "image/png" &&
		typ != "image/gif" &&
		typ != "image/jpeg" &&
		typ != "image/jpg" &&
		typ != "image/bmp" {
		return dto.Fail(e.WrongPictureFormat, nil)
	}
	//若原先头像不是默认头像的话删除头像
	userI, _ := ctx.Get("user")
	if icon := userI.(*dto.UserDto).Icon; icon != "default.jpg" {
		if err := util.DelImg(icon); err != nil {
			return dto.Fail(e.Error, err)
		}
	}
	//上传图片
	url, err := util.UploadImg(iconFile, header.Size)
	if err != nil {
		return dto.Fail(e.Error, err)
	}
	//修改数据库
	if err = userDao.UploadIcon(s.NickName, url); err != nil {
		dto.Fail(e.Error, nil)
	}
	//修改缓存
	if err = redisClient.HSet(e.UserLoginInfo+s.NickName, "Icon", url).Err(); err != nil {
		dto.Fail(e.Error, nil)
	}
	//ctx
	return dto.Success(e.Success, url)
}

func (s *UserService) UpdateNickName(ctx *gin.Context, newNickName string) *dto.Result {
	redisClient := conf.NewRedisClient()
	userDao := dao.NewUserDao(ctx)
	//判断昵称是否合法
	if newNickName == "" {
		return dto.Fail(e.NilNickName, nil)
	}
	if newNickName == s.NickName {
		return dto.Success(e.Success, "修改成功")
	}
	//检查昵称是否已存在
	if userDao.IsExistByNickName(newNickName) {
		return dto.Fail(e.NickNameAlreadyExist, nil)
	}
	//修改昵称
	//修改redis中的昵称
	redisClient.HSet(e.UserLoginInfo+s.NickName, "NickName", newNickName)
	//重命名key
	redisClient.Rename(e.UserLoginInfo+s.NickName, e.UserLoginInfo+newNickName)
	//修改mysql
	if err := userDao.UpdateNickName(s.NickName, newNickName); err != nil {
		return dto.Fail(e.Error, err)
	}
	return dto.Success(e.Success, "修改成功")
}
