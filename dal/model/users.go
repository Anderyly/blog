/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package model

type Role int

const (
	RoleAdmin  Role = iota + 1 // 管理员
	RoleMember                 // 普通成员
)

type Users struct {
	Uid         int64  `gorm:"column:uid;primary_key;AUTO_INCREMENT;NOT NULL" json:"uid,omitempty"` // 主键id
	Account     string `gorm:"column:account;not null" json:"account,omitempty"`                    // 账号
	Password    string `gorm:"column:password;not null" json:"-"`                                   // 密码
	Mail        string `gorm:"column:mail;default:''" json:"mail,omitempty"`                        // 邮箱
	Url         string `gorm:"column:url;default:''" json:"url,omitempty"`                          // 网址
	Nickname    string `gorm:"column:nickname;default:''" json:"nickname,omitempty"`                // 昵称
	Role        Role   `gorm:"column:group;default:2" json:"group,omitempty"`                       // 角色
	ActivatedAt uint32 `gorm:"column:activated_at;default:0" json:"activated_at,omitempty"`         // 上次活动时间
	CreatedAt   int64  `gorm:"column:created_at;default:0" json:"created_at,omitempty"`             // 创建时间
}
