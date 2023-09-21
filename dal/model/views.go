/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package model

type Views struct {
	Id        int64  `gorm:"column:coid;primary_key;AUTO_INCREMENT;NOT NULL" json:"coid,omitempty"` // 主键id
	Cid       int64  `gorm:"column:cid;default:0" json:"cid,omitempty"`                             // 文章id
	Ip        string `gorm:"column:ip;default:''" json:"ip,omitempty"`                              // ip
	Agent     string `gorm:"column:agent;default:''" json:"agent,omitempty"`                        // ua信息
	CreatedAt int64  `gorm:"column:created_at;default:0" json:"created_at,omitempty"`               // 创建时间
}
