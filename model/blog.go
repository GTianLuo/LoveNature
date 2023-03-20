package model

import (
	"time"
)

type Blog struct {
	CreatedAt      time.Time `json:"createdAt"` //创建时间
	Author         string    `json:"author"`    //作者
	Content        string    `json:"content"`
	BlogTitle      string    `json:"blogTitle"`
	Pictures       []string  `json:"pictures"`       // 图片
	GetLikesNumber int       `json:"getLikesNumber"` // 获赞数
	Location       string    `json:"location"`       //位置
}
