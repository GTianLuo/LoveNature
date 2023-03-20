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
	Author    string `form:"author"`
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
		Author:         service.Author,
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

func (service *BlogService) SearchByKeyWord(ctx *gin.Context, keyword string) *dto.Result {
	//es中搜索
	return nil
}

func (service *BlogService) GetBlogList(ctx *gin.Context, way string, page int) *dto.Result {
	//es中搜索
	blogDao := dao.NewBlogDao(ctx)
	blogs, err := blogDao.GetBlogList(way, page)
	if err != nil {
		return dto.Fail(e.Error, err)
	}
	if len(blogs) == 0 {
		return dto.Fail(e.NoMoreBlogs, nil)
	}
	return dto.Success(e.Success, blogs)
}
