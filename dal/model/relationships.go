/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package model

type Relationships struct {
	Cid int64 `gorm:"column:cid;NOT NULL" json:"cid,omitempty"` // 文章id
	Mid int64 `gorm:"column:mid;NOT NULL" json:"mid,omitempty"` // 分类id
}
