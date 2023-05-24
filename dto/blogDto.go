package dto

import (
	"lovenature/model"
	"time"
)

type BlogDto struct {
	CreatedAt      time.Time `json:"createdAt"` //创建时间
	BlogId         string    `json:"blogId,omitempty"`
	Email          string    `json:"email,omitempty"`
	Author         UserDto   `json:"author"` //作者
	Content        string    `json:"content"`
	BlogTitle      string    `json:"blogTitle"`
	Pictures       []string  `json:"pictures"`            // 图片
	GetLikesNumber int       `json:"getLikesNumber"`      // 获赞数
	Location       string    `json:"location"`            //位置
	Highlight      string    `json:"highlight,omitempty"` //高亮
}

func BuildBlogList(blogs []*model.Blog, highLight []string, users []*UserDto) (blogDtos []*BlogDto) {
	for i, blog := range blogs {
		blogDto := &BlogDto{
			CreatedAt:      blog.CreatedAt,
			BlogId:         blog.BlogId,
			Author:         *users[i],
			Content:        blog.Content,
			BlogTitle:      blog.BlogTitle,
			Pictures:       blog.Pictures,
			GetLikesNumber: blog.GetLikesNumber,
			Location:       blog.Location,
			Highlight:      highLight[i],
		}
		blogDtos = append(blogDtos, blogDto)
	}
	return
}
