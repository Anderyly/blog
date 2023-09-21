/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package model

type Links struct {
	Id          int64  `gorm:"column:mid;primary_key;AUTO_INCREMENT;NOT NULL" json:"mid,omitempty"` // 主键id
	Title       string `gorm:"column:title;default:''" json:"title,omitempty"`                      // 标题
	Image       string `gorm:"column:image;default:''" json:"image,omitempty"`                      // 图片
	Description string `gorm:"column:description;default:''" json:"description,omitempty"`          // 描述
	Link        string `gorm:"column:link;default:''" json:"link,omitempty"`                        // 地址
	Sort        int    `gorm:"column:sort" json:"sort"`                                             // 排序
}
