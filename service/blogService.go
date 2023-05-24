package service

import (
	"github.com/gin-gonic/gin"
	"lovenature/dao"
	"lovenature/dto"
	"lovenature/model"
	"lovenature/pkg/e"
	"lovenature/pkg/util"
	"time"
)

func NewBlogService() *BlogService {
	return &BlogService{}
}

type BlogService struct {
	BlogId string `form:"blogId"`
	//Pictures string `form:"pictures"` // 图片
	BlogTitle string `form:"blogTitle"`
	Email     string `form:"email"`
	Content   string `form:"content"`
	Location  string `form:"location"`
}

func (service *BlogService) PostBlog(ctx *gin.Context) *dto.Result {

	blogDao := dao.NewBlogDao(ctx)
	//获取图片
	form, _ := ctx.MultipartForm()
	fileHeaders := form.File["pictures"]
	//校验图片类型和大小
	for _, header := range fileHeaders {
		//校验文件
		if header.Size > (8 << 18) {
			return dto.Fail(e.IconTooBig, nil)
		}
		if typ := header.Header.Get("Content-Type"); typ != "image/png" &&
			typ != "image/gif" &&
			typ != "image/jpeg" &&
			typ != "image/jpg" &&
			typ != "image/jfif" &&
			typ != "image/bmp" {
			return dto.Fail(e.WrongPictureFormat, nil)
		}
	}
	var pictureUrls []string
	for _, header := range fileHeaders {
		file, err := header.Open()
		if err != nil {
			return dto.Fail(e.Error, err)
		}
		if url, err := util.UploadImg(file, header.Size); err != nil {
			return dto.Fail(e.Error, err)
		} else {
			pictureUrls = append(pictureUrls, url)
		}
		_ = file.Close()
	}

	//封装model
	blog := &model.Blog{
		CreatedAt:      time.Now(),
		Email:          service.Email,
		Location:       service.Location,
		BlogTitle:      service.BlogTitle,
		Content:        service.Content,
		Pictures:       pictureUrls,
		GetLikesNumber: 0,
	}
	//存储在es中
	_, err := blogDao.IndexBlog(blog)
	if err != nil {
		return dto.Fail(e.Error, err)
	}
	return dto.Success(e.Success, "发布成功")
}

func (service *BlogService) SearchByKeyWord(ctx *gin.Context, keyword string, page int) *dto.Result {

	blogDao := dao.NewBlogDao(ctx)
	userDao := dao.NewUserDao(ctx)
	//es中搜索
	err, blogs, highlights := blogDao.SearchByKeyWord(keyword, page)

	if err != nil {
		return dto.Fail(e.Error, err)
	}
	if len(blogs) == 0 {
		return dto.Fail(e.NoMoreBlogs, nil)
	}
	var emails []string
	var users []model.User
	for _, blog := range blogs {
		emails = append(emails, blog.Email)
	}
	for _, email := range emails {
		users = append(users, *userDao.GetUser(email)) //userDao.GetUser(email)
	}
	//mysql查找作者信息
	userDtos := dto.BuildUserList(users)
	return dto.Success(e.Success, dto.BuildBlogList(blogs, highlights, userDtos))
}

func (service *BlogService) GetBlogList(ctx *gin.Context, way string, page int) *dto.Result {
	blogDao := dao.NewBlogDao(ctx)
	userDao := dao.NewUserDao(ctx)
	blogs, err := blogDao.GetBlogList(way, page)
	if err != nil {
		return dto.Fail(e.Error, err)
	}
	if len(blogs) == 0 {
		return dto.Fail(e.NoMoreBlogs, nil)
	}
	var emails []string
	for _, blog := range blogs {
		emails = append(emails, blog.Email)
	}
	var users []model.User
	for _, email := range emails {
		users = append(users, *userDao.GetUser(email)) //userDao.GetUser(email)
	}
	userDtos := dto.BuildUserList(users)
	return dto.Success(e.Success, dto.BuildBlogList(blogs, make([]string, len(blogs)+1), userDtos))
}
