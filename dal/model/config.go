/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package model

const (
	ConfigTitleKey       = "title"
	ConfigKeywordsKey    = "keywords"
	ConfigDescriptionKey = "description"
	ConfigWebUrlKey      = "web_url"
	ConfigTemplateKey    = "template"
)

type Config struct {
	K string `gorm:"not null;index" json:"k"`
	V string `gorm:"not null" json:"v"`
}
