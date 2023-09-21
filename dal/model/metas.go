/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package model

type MetasType int

const (
	MetasTypeCategory MetasType = iota + 1 // 分类
	MetasTypeTag                           // 标签
)

type Metas struct {
	Mid         int64     `gorm:"column:mid;primary_key;AUTO_INCREMENT;NOT NULL" json:"mid,omitempty"` // 主键id
	Name        string    `gorm:"column:name;default:''" json:"name,omitempty"`                        // 名称
	Icon        string    `gorm:"column:icon;default:''" json:"icon,omitempty"`                        // icon
	Slug        string    `gorm:"column:slug;default:''" json:"slug,omitempty"`                        // 别名
	Type        MetasType `gorm:"column:type;NOT NULL;default:1" json:"type,omitempty"`                // 类型
	Description string    `gorm:"column:description;default:''" json:"description,omitempty"`          // 描述
	Count       int       `gorm:"column:count;default:0" json:"count,omitempty"`                       // 文章数量
	Sort        int       `gorm:"column:sort" json:"sort"`                                             // 排序
}
