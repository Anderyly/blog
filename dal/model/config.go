/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package model

const (
	ConfigTitleKey       = "title"         // 网站标题
	ConfigKeywordsKey    = "keywords"      // 网站关键词
	ConfigDescriptionKey = "description"   // 网站描述
	ConfigWebUrlKey      = "web_url"       // 网站地址
	ConfigTemplateKey    = "template"      // 当前模板
	ConfigCommentBan     = "comment_ban"   // 评论违禁词
	ConfigCommentApply   = "comment_apply" // 评论审核
	ConfigAvatarUrlKey   = "avatar_url"    // 头像源
)

type Config struct {
	K string `gorm:"not null;index" json:"k"`
	V string `gorm:"not null" json:"v"`
}
