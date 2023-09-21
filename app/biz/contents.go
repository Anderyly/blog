/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package biz

import (
	"blog/ay"
	"blog/dal/model"
	"blog/dal/query"
	"blog/utils"
	"fmt"
	"gorm.io/gorm"
)

type contents struct {
	*base
}

type OptionContents func(...*model.Contents) error

func (con contents) WithSetUser() OptionContents {
	return func(list ...*model.Contents) (err error) {
		var ids []int64
		for _, v := range list {
			ids = append(ids, v.AuthorId)
		}
		ids = utils.SliceUnique(ids)

		var users []*model.Users
		if users, err = con.biz.User.ListByIds(ids); err != nil {
			return
		}

		idToUser := map[int64]*model.Users{}
		for _, v := range users {
			idToUser[v.Uid] = v
		}

		for _, v := range list {
			if user, ok := idToUser[v.AuthorId]; ok {
				v.User = user
			}
		}

		return
	}
}

func (con contents) WithSetCategory() OptionContents {
	return func(list ...*model.Contents) (err error) {
		var ids []int64
		for _, v := range list {
			ids = append(ids, v.Cid)
		}
		ids = utils.SliceUnique(ids)

		var relationships []*model.Relationships
		if relationships, err = con.biz.Relationships.List(ids); err != nil {
			return
		}
		idToRelationships := map[int64][]*model.Relationships{}
		for _, v := range relationships {
			idToRelationships[v.Cid] = append(idToRelationships[v.Cid], v)
		}

		var categoryIds []int64
		for _, v := range relationships {
			categoryIds = append(categoryIds, v.Mid)
		}

		categoryIds = utils.SliceUnique(categoryIds)

		var categories []*model.Metas
		if categories, err = con.biz.Metas.ListCategoryByIds(categoryIds); err != nil {
			return
		}

		idToCategory := map[int64]*model.Metas{}

		for _, v := range categories {
			idToCategory[v.Mid] = v
		}

		for _, content := range list {
			for _, relationship := range idToRelationships[content.Cid] {
				if row, ok := idToCategory[relationship.Mid]; ok {
					if row.Type == model.MetasTypeCategory {
						content.Category = row
					} else {
						content.Tags = append(content.Tags, row)
					}
				}
			}
		}

		return
	}
}

// ListNew 获取最新文章
func (con contents) ListNew(limit int) (list []*model.Contents, err error) {
	qe := query.Use(ay.Db).Contents
	if list, err = qe.WithContext(con.ctx).Select(qe.Slug, qe.Title).Where(qe.Type.Eq(int(model.ContentsTypePost))).Limit(limit).Order(qe.CreatedAt.Desc()).Find(); err != nil {
		return
	}
	return
}

// ListRecommend 获取推荐文章
func (con contents) ListRecommend(contentsId int64, tags []int64, limit int) (list []*model.Contents, err error) {
	qe := query.Use(ay.Db).Contents
	qw := query.Use(ay.Db).Relationships
	if list, err = qe.WithContext(con.ctx).
		Join(qw, qw.Mid.In(tags...), qe.Cid.EqCol(qw.Cid)).
		Select(qe.Slug, qe.Title).
		Where(qe.Type.Eq(int(model.ContentsTypePost)), qe.Cid.Neq(contentsId)).
		Limit(limit).
		Group(qe.Cid).
		Find(); err != nil {
		return
	}
	return
}

// ListHot 获取热门文章
func (con contents) ListHot(limit int) (list []*model.Contents, err error) {
	qe := query.Use(ay.Db).Contents
	if list, err = qe.WithContext(con.ctx).Select(qe.Slug, qe.Title).Where(qe.Type.Eq(int(model.ContentsTypePost))).Limit(limit).Order(qe.CommentsNum.Desc()).Find(); err != nil {
		return
	}
	return
}

// ListByIds 通过ids获取
func (con contents) ListByIds(ids []int64) (list []*model.Contents, err error) {
	qe := query.Use(ay.Db).Contents
	if list, err = qe.WithContext(con.ctx).Where(qe.Cid.In(ids...)).Find(); err != nil {
		return
	}
	return
}

type ListContentsArgs struct {
	Page     int
	Keyword  string
	Category string
}

// List 文章列表
func (con contents) List(args ListContentsArgs, opts ...OptionContents) (list []*model.Contents, count int64, err error) {
	qe := query.Use(ay.Db).Contents
	do := qe.WithContext(con.ctx)
	fmt.Println(args)
	if args.Keyword != "" {
		do = do.Where(qe.Title.Like("%" + args.Keyword + "%"))
	}
	if args.Category != "" {
		qw := query.Use(ay.Db).Metas
		qr := query.Use(ay.Db).Relationships
		do = do.Join(qr, qe.Cid.EqCol(qr.Cid)).
			Join(qw, qw.Mid.EqCol(qr.Mid)).
			Where(qw.Slug.Eq(args.Category))
	}
	if list, count, err = do.Order(qe.CreatedAt.Desc()).FindByPage(args.Page, 10); err != nil {
		return
	}

	for _, opt := range opts {
		if err = opt(list...); err != nil {
			return
		}
	}
	return
}

// GetBySlug 通过slug获取文章
func (con contents) GetBySlug(slug string, opts ...OptionContents) (res *model.Contents, err error) {
	qe := query.Use(ay.Db).Contents
	do := qe.WithContext(con.ctx)
	if res, err = do.Where(qe.Slug.Eq(slug)).Take(); err != nil {
		return
	}

	for _, opt := range opts {
		if err = opt(res); err != nil {
			return
		}
	}
	return
}

// Count 获取文章总数
func (con contents) Count(contentsType model.ContentsType) (count int64, err error) {
	qe := query.Use(ay.Db).Contents
	if count, err = qe.WithContext(con.ctx).Where(qe.Type.Eq(int(contentsType))).Count(); err != nil {
		return
	}
	return
}

// GetLast 获取上一篇文章
func (con contents) GetLast(contentsId int64) (row *model.Contents, err error) {
	qe := query.Use(ay.Db).Contents
	if row, err = qe.WithContext(con.ctx).Where(qe.Cid.Lt(contentsId)).Order(qe.Cid.Desc()).Take(); err != nil {
		return
	}
	return
}

// GetNext 获取下一篇文章
func (con contents) GetNext(contentsId int64) (row *model.Contents, err error) {
	qe := query.Use(ay.Db).Contents
	if row, err = qe.WithContext(con.ctx).Where(qe.Cid.Gt(contentsId)).Order(qe.Cid).Take(); err != nil {
		return
	}
	return
}

func (con contents) Create(data *model.Contents) (err error) {
	qe := query.Use(ay.Db).Contents
	return qe.WithContext(con.ctx).Create(data)
}

// Update 更新文章
func (con contents) Update(data *model.Contents) (err error) {
	qe := query.Use(ay.Db).Contents
	_, err = qe.WithContext(con.ctx).Where(qe.Cid.Eq(data.Cid)).Updates(data)
	return
}

// Delete 删除文章
func (con contents) Delete(contentsId int64) (err error) {
	qe := query.Use(ay.Db).Contents
	_, err = qe.WithContext(con.ctx).Where(qe.Cid.Eq(contentsId)).Delete()
	return
}

// AddViews 增加浏览量
func (con contents) AddViews(db *gorm.DB, contentsId int64) (err error) {
	qe := query.Use(db).Contents
	field := qe.ViewsNum.ColumnName().String()
	_, err = qe.WithContext(con.ctx).Where(qe.Cid.Eq(contentsId)).Updates(
		map[string]interface{}{
			field: gorm.Expr(fmt.Sprintf("%s + ?", field), 1),
		})
	return
}

// AddComments 增加评论数
func (con contents) AddComments(contentsId int64) (err error) {
	qe := query.Use(ay.Db).Contents
	field := qe.CommentsNum.ColumnName().String()
	_, err = qe.WithContext(con.ctx).Where(qe.Cid.Eq(contentsId)).Updates(
		map[string]interface{}{
			field: gorm.Expr(fmt.Sprintf("%s + ?", field), 1),
		})
	return
}
