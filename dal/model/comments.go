/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package model

import "blog/dal"

type CommentType int

const (
	CommentTypeComment CommentType = iota + 1 // 评论
)

type Comments struct {
	Coid         int64       `gorm:"column:coid;primary_key;AUTO_INCREMENT;NOT NULL" json:"coid,omitempty"` // 评论主键id
	Type         CommentType `gorm:"column:type;default:1" json:"type,omitempty"`                           // 评论类型
	ParentId     int64       `gorm:"column:parent_id;default:0" json:"parent_id,omitempty"`                 // 上级评论id
	Cid          int64       `gorm:"column:cid;default:0" json:"cid,omitempty"`                             // 文章id
	Author       string      `gorm:"column:author;default:''" json:"author,omitempty"`                      // 评论者名称
	AuthorId     int64       `gorm:"column:author_id;default:0" json:"author_id,omitempty"`                 // 评论者id
	OwnerId      int64       `gorm:"column:owner_id;default:0" json:"owner_id,omitempty"`                   // 作者id
	Mail         string      `gorm:"column:mail;default:''" json:"mail,omitempty"`                          // 邮箱
	Url          string      `gorm:"column:url;default:''" json:"url,omitempty"`                            // 网址
	Ip           string      `gorm:"column:ip;default:''" json:"ip,omitempty"`                              // ip
	Agent        string      `gorm:"column:agent;default:''" json:"agent,omitempty"`                        // ua信息
	Text         string      `gorm:"column:text" json:"text,omitempty"`                                     // 内容
	Status       dal.Status  `gorm:"column:status;default:1" json:"status,omitempty"`                       // 评论状态
	CreatedAt    int64       `gorm:"column:created_at;default:0" json:"created_at,omitempty"`               // 创建时间
	ContentTitle string      `gorm:"-"`                                                                     // 文章标题
	Slug         string      `gorm:"-"`
}
