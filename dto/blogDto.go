package dto

import (
	"lovenature/model"
	"time"
)

type BlogDto struct {
	BlogId         string    `json:"blogId,omitempty"`
	CreatedAt      time.Time `json:"createdAt,omitempty"`      //创建时间
	Author         string    `json:"author,omitempty"`         //作者
	Pictures       []string  `json:"pictures,omitempty"`       // 图片
	GetLikesNumber int       `json:"getLikesNumber,omitempty"` // 获赞数
}

func BuildBlogDto(blogId string, blog *model.Blog) *BlogDto {
	return &BlogDto{
		BlogId:         blogId,
		CreatedAt:      blog.CreatedAt,
		Author:         blog.Author,
		Pictures:       blog.Pictures,
		GetLikesNumber: blog.GetLikesNumber,
	}
}
