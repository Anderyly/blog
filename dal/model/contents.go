/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package model

import "blog/dal"

type ContentsType int

const (
	ContentsTypePost ContentsType = iota + 1 // 文章
	ContentsTypePage                         // 页面
)

type Contents struct {
	Cid          int64        `gorm:"column:cid;primary_key;AUTO_INCREMENT;NOT NULL" json:"cid,omitempty"` // 主键id
	Type         ContentsType `gorm:"column:type;default:1" json:"type,omitempty"`                         // 类型
	Title        string       `gorm:"column:title;default:''" json:"title,omitempty"`                      // 标题
	Slug         string       `gorm:"column:slug;default:''" json:"slug,omitempty"`                        // 别名
	Text         string       `gorm:"column:text" json:"text,omitempty"`                                   // 内容
	AuthorId     int64        `gorm:"column:author_id;default:0" json:"author_id,omitempty"`               // 作者id
	Status       dal.Status   `gorm:"column:status;default:1" json:"status,omitempty"`                     // 状态
	Password     string       `gorm:"column:password;default:''" json:"password,omitempty"`                // 密码
	CommentsNum  int64        `gorm:"column:comments_num;default:0" json:"comments_num,omitempty"`         // 评论数量
	ViewsNum     int64        `gorm:"column:views_num;default:0" json:"views_num,omitempty"`               // 查看数量
	AllowComment dal.Sure     `gorm:"column:allow_comment;default:1" json:"allow_comment,omitempty"`       // 允许评论
	AllowPing    dal.Sure     `gorm:"column:allow_ping;default:0" json:"allow_ping,omitempty"`             // 允许
	AllowFeed    dal.Sure     `gorm:"column:allow_feed;default:0" json:"allow_feed,omitempty"`             // 允许
	CreatedAt    int64        `gorm:"column:created_at;default:0" json:"created_at,omitempty"`             // 创建时间
	UpdatedAt    int64        `gorm:"column:updated_at;default:0" json:"updated_at,omitempty"`             // 修改时间
	Category     *Metas       `gorm:"-"`
	Tags         []*Metas     `gorm:"-"`
	User         *Users       `gorm:"-"`
}
