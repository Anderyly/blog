package biz

import (
	"context"
)

type Biz struct {
	Metas         *metas
	Contents      *contents
	Comments      *comments
	Relationships *relationships
	User          *user
	Links         *links
	Config        *config
	Views         *views
	Templates     *templates
}

type base struct {
	ctx context.Context
	biz *Biz
}

func NewBiz(ctx context.Context) (b *Biz) {
	b = &Biz{}
	base := &base{ctx: ctx, biz: b}
	b.Metas = &metas{base}
	b.Contents = &contents{base}
	b.Comments = &comments{base}
	b.Relationships = &relationships{base}
	b.User = &user{base}
	b.Links = &links{base}
	b.Config = &config{base}
	b.Views = &views{base}
	b.Templates = &templates{base}

	return
}
