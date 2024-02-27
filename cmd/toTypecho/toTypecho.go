/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2024
 */

package toTypecho

import (
	"blog/ay"
	"blog/dal"
	"blog/dal/model"
	"blog/dal/query"
	"context"
	"github.com/spf13/cobra"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

var (
	Db       *gorm.DB
	Host     = "127.0.0.1"
	User     = "root"
	Password = "root"
	Database = "typecho"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "toTypecho",
		Short: "typecho数据导入",

		Run: func(cmd *cobra.Command, args []string) {
			var err error
			ay.Init()

			var option gorm.Dialector

			dsn := User + ":" + Password + "@tcp(" + Host + ":3306)/" + Database + "?charset=utf8mb4&parseTime=true&loc=Local"
			option = mysql.New(mysql.Config{DSN: dsn})

			if Db, err = gorm.Open(option, &gorm.Config{
				NamingStrategy: schema.NamingStrategy{SingularTable: true},
			}); err != nil {
				return
			}

			//toMetas()
			toContent()
		},
	}
	return cmd
}

func toContent() {
	type contentStruct struct {
		Cid          int64
		Title        string
		Slug         string
		Text         string
		AuthorId     int64 `gorm:"column:authorId;default:0"`
		Status       string
		Type         string
		Password     string
		CommentsNum  int64  `gorm:"column:commentsNum"`
		ViewsNum     int64  `gorm:"column:viewsNum"`
		AllowComment int    `gorm:"column:allowComment"`
		AllowPing    string `gorm:"column:allowPing"`
		AllowFeed    string `gorm:"column:allowFeed"`
		Created      int64
		Modified     int64
	}

	var contents []contentStruct

	Db.Table("typecho_contents").Find(&contents)

	qe := query.Use(ay.Db).Contents
	for _, v := range contents {
		types := model.ContentsTypePost
		if v.Type == "page" {
			types = model.ContentsTypePage
		}
		qe.WithContext(context.Background()).Create(&model.Contents{
			Cid:          v.Cid,
			Type:         types,
			Title:        v.Title,
			Slug:         v.Slug,
			Text:         v.Text,
			AuthorId:     v.AuthorId,
			Status:       dal.StatusSuccess,
			Password:     v.Password,
			CommentsNum:  v.CommentsNum,
			ViewsNum:     v.ViewsNum,
			AllowComment: dal.SureYes,
			AllowPing:    dal.SureYes,
			AllowFeed:    dal.SureYes,
			CreatedAt:    v.Created,
			UpdatedAt:    v.Modified,
		})
	}

}

func toMetas() {
	type metasStruct struct {
		Mid         int64  `gorm:"column:mid"`
		Name        string `gorm:"column:name"`
		Slug        string `gorm:"column:slug"`
		Type        string `gorm:"column:type"`
		Description string `gorm:"column:description"`
		Count       int    `gorm:"column:count"`
		Order       int    `gorm:"column:order"`
		Parent      int    `gorm:"column:parent"`
	}

	var metas []metasStruct

	Db.Table("typecho_metas").Find(&metas)

	qe := query.Use(ay.Db).Metas
	for _, v := range metas {
		types := model.MetasTypeCategory
		if v.Type == "tag" {
			types = model.MetasTypeTag
		}
		qe.WithContext(context.Background()).Create(&model.Metas{
			Mid:         v.Mid,
			Name:        v.Name,
			Icon:        "",
			Slug:        v.Slug,
			Type:        types,
			Description: v.Description,
			Count:       v.Count,
			Sort:        v.Order,
		})
	}
}
