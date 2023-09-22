/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package biz

import (
	"blog/ay"
	"blog/dal"
	"blog/dal/model"
	"blog/dal/query"
	"blog/utils"
)

type comments struct {
	*base
}

type OptionComments func(...*model.Comments) error

func (con comments) WithSetContents() OptionComments {
	return func(list ...*model.Comments) (err error) {
		var contentIds []int64
		for _, v := range list {
			contentIds = append(contentIds, v.Cid)
		}
		contentIds = utils.SliceUnique(contentIds)

		var contents []*model.Contents
		if contents, err = con.biz.Contents.ListByIds(contentIds); err != nil {
			return
		}

		idToContent := map[int64]*model.Contents{}
		for _, v := range contents {
			idToContent[v.Cid] = v
		}

		for _, v := range list {
			if res, ok := idToContent[v.Cid]; ok {
				v.ContentTitle = res.Title
				v.Slug = res.Slug
			}
		}

		return
	}
}

func (con comments) WithSetChild() OptionComments {
	return func(list ...*model.Comments) (err error) {
		var commentIds []int64
		for _, v := range list {
			commentIds = append(commentIds, v.Coid)
		}

		commentIds = utils.SliceUnique(commentIds)

		var comments []*model.Comments
		if comments, err = con.ListByParentIds(commentIds, con.WithSetAtName()); err != nil {
			return
		}
		idToComment := map[int64]*model.Comments{}
		for _, v := range comments {
			idToComment[v.ParentId] = v
		}

		for _, v := range list {
			if res, ok := idToComment[v.Coid]; ok {
				v.ChildComments = append(v.ChildComments, res)
			}
		}

		return
	}
}

func (con comments) WithSetAtName() OptionComments {
	return func(list ...*model.Comments) (err error) {
		var commentIds []int64
		for _, v := range list {
			commentIds = append(commentIds, v.ParentId)
		}

		commentIds = utils.SliceUnique(commentIds)

		var comments []*model.Comments
		if comments, err = con.ListByIds(commentIds); err != nil {
			return
		}
		idToComment := map[int64]*model.Comments{}
		for _, v := range comments {
			idToComment[v.Coid] = v
		}

		for _, v := range list {
			if res, ok := idToComment[v.ParentId]; ok {
				v.LastAtName = res.Author
			}
		}

		return
	}
}

// 获取最新评论
func (con comments) ListNew(limit int, opts ...OptionComments) (list []*model.Comments, err error) {
	qe := query.Use(ay.Db).Comments
	if list, err = qe.WithContext(con.ctx).Where(qe.Status.Eq(int(dal.StatusSuccess))).Limit(limit).Order(qe.CreatedAt.Desc()).Find(); err != nil {
		return
	}

	for _, opt := range opts {
		if err = opt(list...); err != nil {
			return
		}
	}
	return
}

type CommentsCountCondition struct {
	ContentId int64
}

// 获取总数
func (con comments) Count(args CommentsCountCondition) (count int64, err error) {
	qe := query.Use(ay.Db).Comments
	do := qe.WithContext(con.ctx).Where(qe.Status.Eq(int(dal.StatusSuccess)))
	if args.ContentId != 0 {
		do = do.Where(qe.Cid.Eq(args.ContentId))
	}
	if count, err = do.Count(); err != nil {
		return
	}
	return
}

// 通过文章id获取评论
func (con comments) ListByContentId(contentId int64, opts ...OptionComments) (list []*model.Comments, err error) {

	qe := query.Use(ay.Db).Comments
	if list, err = qe.WithContext(con.ctx).
		Where(qe.ParentId.Eq(0), qe.Cid.Eq(contentId), qe.Status.Eq(int(dal.StatusSuccess))).
		Find(); err != nil {
		return
	}

	for _, opt := range opts {
		if err = opt(list...); err != nil {
			return
		}
	}

	return
}

// 通过ids获取所有评论
func (con comments) ListByParentIds(ids []int64, opts ...OptionComments) (list []*model.Comments, err error) {
	qe := query.Use(ay.Db).Comments
	if list, err = qe.WithContext(con.ctx).
		Where(qe.ParentId.In(ids...), qe.Status.Eq(int(dal.StatusSuccess))).
		Find(); err != nil {
		return
	}

	for _, opt := range opts {
		if err = opt(list...); err != nil {
			return
		}
	}
	return
}

// 通过ids获取所有评论
func (con comments) ListByIds(ids []int64, opts ...OptionComments) (list []*model.Comments, err error) {
	qe := query.Use(ay.Db).Comments
	if list, err = qe.WithContext(con.ctx).
		Where(qe.Coid.In(ids...), qe.Status.Eq(int(dal.StatusSuccess))).
		Find(); err != nil {
		return
	}

	for _, opt := range opts {
		if err = opt(list...); err != nil {
			return
		}
	}
	return
}
