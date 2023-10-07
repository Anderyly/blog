/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package index

import (
	"blog/app/biz"
	"blog/app/controllers"
	"blog/ay"
	"blog/ay/lib"
	"blog/dal"
	"blog/dal/model"
	"blog/utils"
	"context"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type IndexController struct {
}

func (con IndexController) Index(c *gin.Context) {
	queryPage := c.Param("page")
	queryKeywords := c.Query("s")
	queryCategory := c.Param("category")
	queryTag := c.Param("tag")

	if queryPage == "" {
		queryPage = "0"
	}
	page, err := strconv.Atoi(queryPage)
	if err != nil {
		ay.Logger.Error(err.Error())
		return
	}

	key := queryTag
	if queryKeywords != "" {
		key = queryKeywords
	}

	args := biz.ListContentsArgs{
		Page:     page,
		Keyword:  key,
		Category: queryCategory,
	}
	b := biz.NewBiz(context.Background())
	list, count, err := b.Contents.List(args, b.Contents.WithSetCategory(), b.Contents.WithSetUser())
	if err != nil {
		ay.Logger.Error(err.Error())
		return
	}

	if page == 0 {
		page = 1
	}

	c.HTML(http.StatusOK, "default/index", gin.H{
		"siteConf": controllers.SiteConf,
		"contents": list,
		"pages":    utils.Page(page, int(count), 10),
		"page":     page,
		"category": queryCategory,
		"keywords": queryKeywords,
		"maxPage":  int(count) / 10,
		"lessPage": int(count)/10 - 3,
	})
}

func (con IndexController) Archives(c *gin.Context) {
	querySlug := c.Param("slug")
	querySlug = strings.ReplaceAll(querySlug, ".html", "")

	b := biz.NewBiz(context.Background())
	res, err := ay.IgnoreNotFoundReturn(b.Contents.GetBySlug(querySlug, b.Contents.WithSetCategory(), b.Contents.WithSetUser()))
	if err != nil {
		ay.Logger.Error(err.Error())
		return
	}

	if res == nil {
		c.HTML(http.StatusOK, "default/404", gin.H{
			"siteConf": controllers.SiteConf,
		})
		return
	}

	go func() {
		if err = b.Views.Create(&model.Views{
			Cid:       res.Cid,
			Agent:     c.Request.UserAgent(),
			Ip:        c.ClientIP(),
			CreatedAt: time.Now().Unix(),
		}); err != nil {
			ay.Logger.Error(err.Error())
			return
		}
	}()

	//lib.NewJson(c).Success(res)
	//return

	c.HTML(http.StatusOK, "default/archives", gin.H{
		"siteConf": controllers.SiteConf,
		"contents": res,
		"category": res.Category.Slug,
	})
}

type AddCommentForm struct {
	Nickname  string `form:"nickname" binding:"required"`
	Email     string `form:"email"`
	Content   string `form:"content" binding:"required"`
	ContentId int64  `form:"content_id" binding:"required"`
	ParentId  int64  `form:"parent_id"`
}

// 添加评论
func (con IndexController) AddComment(c *gin.Context) {
	var err error
	r := lib.NewJson(c)
	var data AddCommentForm
	if err = c.ShouldBind(&data); err != nil {
		r.Fail(ay.Validator{}.Translate(err))
		return
	}

	// 验证违禁词
	if commentBan, ok := controllers.SiteConf[model.ConfigCommentBan]; ok {
		commentBanArr := strings.Split(commentBan, ",")
		isIncludeBan := false
		for _, v := range commentBanArr {
			if v == "" {
				continue
			}
			if strings.Contains(data.Content, v) {
				isIncludeBan = true
				break
			}
		}
		if isIncludeBan {
			r.Fail("评论内容包含违禁词")
			return
		}
	}

	status := dal.StatusSuccess

	// 验证评论是否需要审核
	if configCommentApply, ok := controllers.SiteConf[model.ConfigCommentApply]; ok {
		var commentApply int
		if commentApply, err = strconv.Atoi(configCommentApply); err != nil {
			r.Fail(err.Error())
			return
		}
		if commentApply == int(dal.SureYes) {
			status = dal.StatusApply
		}
	}

	b := biz.NewBiz(context.Background())
	if err = b.Comments.Create(&model.Comments{
		Cid:       data.ContentId,
		Author:    data.Nickname,
		Mail:      data.Email,
		Text:      data.Content,
		ParentId:  data.ParentId,
		Status:    status,
		CreatedAt: time.Now().Unix(),
	}); err != nil {
		ay.Logger.Error(err.Error())
		return
	}

	r.Success("评论成功")

}
