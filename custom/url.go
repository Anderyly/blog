/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package custom

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"strconv"
)

func Url(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		"GetArchivesLink": func(slug string) string {
			return "/archives/" + slug + ".html"
		},
		"GetAuthorLink": func(authorId int64) string {
			return "/author/" + strconv.FormatInt(authorId, 10)
		},
	})
}
