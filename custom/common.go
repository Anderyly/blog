/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package custom

import (
	"blog/ay"
	"blog/utils"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"strconv"
)

func Common(r *gin.Engine) {
	r.SetFuncMap(template.FuncMap{
		"UNHtml": func(str string) template.HTML {
			return template.HTML(str)
		},
		"Add": func(a, b int) int {
			return a + b
		},
		"AddString": func(a string, b int) int {
			c, err := strconv.Atoi(a)
			if err != nil {
				log.Println(err)
			}
			return c + b
		},
		"MathDel": func(a, b float64) float64 {
			return a - b
		},
		"GetAbstract": func(content string, maxChars int) string {
			return utils.GetAbstract(content, maxChars)
		},
		"Escape": func(content string) string {
			return utils.HtmlEscape(content)
		},
		"Cover": func(content string) string {
			image, err := utils.FindFirstImage(content)
			if err != nil {
				ay.Logger.Error(err.Error())
				return ""
			}
			return image
		},
	})
}
