/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2022
 */

package custom

import (
	"blog/app/biz"
	"blog/app/controllers"
	"blog/ay"
	"blog/ay/lib"
	"blog/dal/model"
	"blog/utils"
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"strconv"
	"time"
)

func Set(r *gin.Engine) {
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
		"DateFormat": func(t int64, template string) string {
			return time.Unix(t, 0).Format(template)
		},
		// 获取分类
		"ListCategory": func() (list []*model.Metas) {
			b := biz.NewBiz(context.Background())
			var err error
			list, err = b.Metas.ListCategory()
			if err != nil {
				ay.Logger.Error(err.Error())
			}
			return
		},
		// 获取最新文章
		"ListContentsNew": func(limit int) (list []*model.Contents) {
			b := biz.NewBiz(context.Background())
			var err error
			list, err = b.Contents.ListNew(limit)
			if err != nil {
				ay.Logger.Error(err.Error())
			}
			return
		},
		// 获取热门文章
		"ListContentsHot": func(limit int) (list []*model.Contents) {
			b := biz.NewBiz(context.Background())
			var err error
			list, err = b.Contents.ListHot(limit)
			if err != nil {
				ay.Logger.Error(err.Error())
			}
			return
		},
		// 获取推荐文章
		"ListContentsRecommend": func(contentsId int64, tags []*model.Metas, limit int) (list []*model.Contents) {
			b := biz.NewBiz(context.Background())
			var err error
			var ids []int64
			for _, v := range tags {
				ids = append(ids, v.Mid)
			}
			list, err = b.Contents.ListRecommend(contentsId, ids, limit)
			if err != nil {
				ay.Logger.Error(err.Error())
			}
			return
		},
		// 获取上一篇文章
		"GetLastContents": func(contentId int64) (res *model.Contents) {
			b := biz.NewBiz(context.Background())
			var err error
			res, err = ay.IgnoreNotFoundReturn(b.Contents.GetLast(contentId))
			if err != nil {
				ay.Logger.Error(err.Error())
			}
			return
		},
		// 获取下一篇文章
		"GetNextContents": func(contentId int64) (res *model.Contents) {
			b := biz.NewBiz(context.Background())
			var err error
			res, err = ay.IgnoreNotFoundReturn(b.Contents.GetNext(contentId))
			if err != nil {
				ay.Logger.Error(err.Error())
			}
			return
		},
		// 获取评论总数
		"GetContentsCount": func(contentsType int) (count int64) {
			b := biz.NewBiz(context.Background())
			var err error
			count, err = b.Contents.Count(model.ContentsType(contentsType))
			if err != nil {
				ay.Logger.Error(err.Error())
			}
			return
		},
		// 获取评论
		"ListContentComments": func(contentsId int64) []*model.Comments {
			b := biz.NewBiz(context.Background())
			list, err := b.Comments.ListByContentId(contentsId)
			if err != nil {
				ay.Logger.Error(err.Error())
				return nil
			}
			return list
		},
		// 获取最新评论
		"ListCommentsNew": func(limit int) (list []*model.Comments) {
			b := biz.NewBiz(context.Background())
			var err error
			list, err = b.Comments.ListNew(limit, b.Comments.WithSetContents())
			if err != nil {
				ay.Logger.Error(err.Error())
			}
			return
		},
		// 获取评论总数
		"GetCommentCount": func(contentId int64) (count int64) {
			b := biz.NewBiz(context.Background())
			var err error
			count, err = b.Comments.Count(biz.CommentsCountCondition{ContentId: contentId})
			if err != nil {
				ay.Logger.Error(err.Error())
			}
			return
		},
		"ListLinks": func() []*model.Links {
			b := biz.NewBiz(context.Background())
			list, err := b.Links.List()
			if err != nil {
				ay.Logger.Error(err.Error())
			}
			return list
		},
		"GetAbstract": func(content string, maxChars int) string {
			return utils.GetAbstract(content, maxChars)
		},
		"Escape": func(content string) string {
			return utils.HtmlEscape(content)
		},
		"GetPageLink": func(page int) string {
			return controllers.SiteConf[model.ConfigWebUrlKey] + "page/" + strconv.Itoa(page)
		},
		"Cover": func(content string) string {
			image, err := utils.FindFirstImage(content)
			if err != nil {
				i := 1
				return controllers.SiteConf[model.ConfigWebUrlKey] + "static/templates/default/random/" + strconv.Itoa(i) + ".jpg"
			}
			return image
		},
		"GetArchivesLink": func(slug string) string {

			return controllers.SiteConf[model.ConfigWebUrlKey] + "archives/" + slug + ".html"
		},
		"GetAuthorLink": func(account string) string {
			return controllers.SiteConf[model.ConfigWebUrlKey] + "author/" + account
		},
		"GetTagLink": func(slug string) string {
			return controllers.SiteConf[model.ConfigWebUrlKey] + "archives/tag/" + slug
		},
		"GetCategoryLink": func(slug string) string {
			return controllers.SiteConf[model.ConfigWebUrlKey] + "category/" + slug
		},
		"GetUser": func() *model.Users {
			b := biz.NewBiz(context.Background())
			user, err := ay.IgnoreNotFoundReturn(b.User.GetById(1))
			if err != nil {
				ay.Logger.Error(err.Error())
			}
			return user
		},
		"GetAvatar": func(mail string) string {
			h := lib.NewStr().Md5(mail)
			return controllers.SiteConf[model.ConfigAvatarUrlKey] + "/avatar/" + h + "?s=200&d=mm&r=g"
		},
		"GetTemplateConfig": func() (data map[string]interface{}) {
			b := biz.NewBiz(context.Background())
			templates, err := ay.IgnoreNotFoundReturn(b.Templates.Get(controllers.SiteConf[model.ConfigTemplateKey]))
			if err != nil {
				ay.Logger.Error(err.Error())
				return
			}
			data = map[string]interface{}{}
			err = json.Unmarshal([]byte(templates.V), &data)
			if err != nil {
				ay.Logger.Error(err.Error())
				return
			}
			return
		},
	})
}
