/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package index

import (
	"blog/app/controllers"
	"blog/ay"
	"blog/dal"
	"blog/dal/model"
	"blog/dal/query"
	"encoding/xml"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/url"
	"os"
	"time"
)

type SitemapController struct {
}

type SitemapUrl struct {
	Loc     string `xml:"loc"`
	Lastmod string `xml:"lastmod"`
}

type SitemapUrlSet struct {
	XMLName xml.Name     `xml:"urlset"`
	URLs    []SitemapUrl `xml:"url"`
}

func (con SitemapController) Index(c *gin.Context) {

	var urls []SitemapUrl

	urls = append(urls, SitemapUrl{
		Loc:     controllers.SiteConf[model.ConfigWebUrlKey] + url.QueryEscape("post-sitemap.xml"),
		Lastmod: time.Now().Format("2006-01-02T15:04:05+08:00"),
	})
	urls = append(urls, SitemapUrl{
		Loc:     controllers.SiteConf[model.ConfigWebUrlKey] + url.QueryEscape("page-sitemap.xml"),
		Lastmod: time.Now().Format("2006-01-02T15:04:05+08:00"),
	})
	urls = append(urls, SitemapUrl{
		Loc:     controllers.SiteConf[model.ConfigWebUrlKey] + url.QueryEscape("category-sitemap.xml"),
		Lastmod: time.Now().Format("2006-01-02T15:04:05+08:00"),
	})
	urls = append(urls, SitemapUrl{
		Loc:     controllers.SiteConf[model.ConfigWebUrlKey] + url.QueryEscape("post_tag-sitemap.xml"),
		Lastmod: time.Now().Format("2006-01-02T15:04:05+08:00"),
	})
	urls = append(urls, SitemapUrl{
		Loc:     controllers.SiteConf[model.ConfigWebUrlKey] + url.QueryEscape("author-sitemap.xml"),
		Lastmod: time.Now().Format("2006-01-02T15:04:05+08:00"),
	})

	sitemap := SitemapUrlSet{
		URLs: urls,
	}

	encoder := xml.NewEncoder(os.Stdout)
	encoder.Indent("", "  ")

	if err := encoder.Encode(sitemap); err != nil {
		ay.Logger.Error(err.Error())
		return
	}
	c.XML(http.StatusOK, sitemap)
}
func (con SitemapController) Post(c *gin.Context) {

	var contents []*model.Contents
	var err error
	qe := query.Use(ay.Db).Contents
	if contents, err = qe.WithContext(c).Select(qe.Slug, qe.CreatedAt, qe.UpdatedAt).Where(qe.Type.Eq(int(model.ContentsTypePost)), qe.Status.Eq(int(dal.StatusSuccess))).Find(); err != nil {
		ay.Logger.Error(err.Error())
		return
	}

	var urls []SitemapUrl
	for _, v := range contents {
		loc := controllers.SiteConf[model.ConfigWebUrlKey] + "archives/" + url.QueryEscape(v.Slug) + ".html"
		urls = append(urls, SitemapUrl{
			Loc:     loc,
			Lastmod: time.Unix(v.UpdatedAt, 0).Format("2006-01-02T15:04:05+08:00"),
		})
	}

	sitemap := SitemapUrlSet{
		URLs: urls,
	}

	encoder := xml.NewEncoder(os.Stdout)
	encoder.Indent("", "  ")

	if err := encoder.Encode(sitemap); err != nil {
		ay.Logger.Error(err.Error())
		return
	}
	c.XML(http.StatusOK, sitemap)
}

func (con SitemapController) Page(c *gin.Context) {

	var contents []*model.Contents
	var err error
	qe := query.Use(ay.Db).Contents
	if contents, err = qe.WithContext(c).Select(qe.Slug, qe.CreatedAt, qe.UpdatedAt).Where(qe.Type.Eq(int(model.ContentsTypePage)), qe.Status.Eq(int(dal.StatusSuccess))).Find(); err != nil {
		ay.Logger.Error(err.Error())
		return
	}

	var urls []SitemapUrl
	urls = append(urls, SitemapUrl{
		Loc: controllers.SiteConf[model.ConfigWebUrlKey],
	})
	for _, v := range contents {
		loc := controllers.SiteConf[model.ConfigWebUrlKey] + url.QueryEscape(v.Slug)
		urls = append(urls, SitemapUrl{
			Loc:     loc,
			Lastmod: time.Unix(v.UpdatedAt, 0).Format("2006-01-02T15:04:05+08:00"),
		})
	}

	sitemap := SitemapUrlSet{
		URLs: urls,
	}

	encoder := xml.NewEncoder(os.Stdout)
	encoder.Indent("", "  ")

	if err := encoder.Encode(sitemap); err != nil {
		ay.Logger.Error(err.Error())
		return
	}
	c.XML(http.StatusOK, sitemap)
}

func (con SitemapController) Category(c *gin.Context) {

	var res []*model.Metas
	var err error
	qe := query.Use(ay.Db).Metas
	if res, err = qe.WithContext(c).Select(qe.Mid, qe.Slug).Where(qe.Type.Eq(int(model.MetasTypeCategory))).Find(); err != nil {
		ay.Logger.Error(err.Error())
		return
	}

	var urls []SitemapUrl
	for _, v := range res {
		qw := query.Use(ay.Db).Relationships
		qc := query.Use(ay.Db).Contents
		var content *model.Contents
		do := qc.WithContext(c).Select(qc.Cid, qc.UpdatedAt).Join(qw, qc.Cid.EqCol(qw.Cid)).Where(qw.Mid.Eq(v.Mid)).Order(qc.UpdatedAt.Desc())
		if content, err = ay.IgnoreNotFoundReturn(do.Take()); err != nil {
			ay.Logger.Error(err.Error())
			return
		}
		if content == nil {
			content.UpdatedAt = 0
		}
		loc := controllers.SiteConf[model.ConfigWebUrlKey] + "category/" + url.QueryEscape(v.Slug)
		urls = append(urls, SitemapUrl{
			Loc:     loc,
			Lastmod: time.Unix(content.UpdatedAt, 0).Format("2006-01-02T15:04:05+08:00"),
		})
	}

	sitemap := SitemapUrlSet{
		URLs: urls,
	}

	encoder := xml.NewEncoder(os.Stdout)
	encoder.Indent("", "  ")

	if err := encoder.Encode(sitemap); err != nil {
		ay.Logger.Error(err.Error())
		return
	}
	c.XML(http.StatusOK, sitemap)
}

func (con SitemapController) Tag(c *gin.Context) {

	var res []*model.Metas
	var err error
	qe := query.Use(ay.Db).Metas
	if res, err = qe.WithContext(c).Select(qe.Mid, qe.Slug).Where(qe.Type.Eq(int(model.MetasTypeTag))).Find(); err != nil {
		ay.Logger.Error(err.Error())
		return
	}

	var urls []SitemapUrl
	for _, v := range res {

		loc := controllers.SiteConf[model.ConfigWebUrlKey] + "archives/tag/" + url.QueryEscape(v.Slug)
		urls = append(urls, SitemapUrl{
			Loc:     loc,
			Lastmod: time.Now().Format("2006-01-02T15:04:05+08:00"),
		})
	}

	sitemap := SitemapUrlSet{
		URLs: urls,
	}

	encoder := xml.NewEncoder(os.Stdout)
	encoder.Indent("", "  ")

	if err := encoder.Encode(sitemap); err != nil {
		ay.Logger.Error(err.Error())
		return
	}
	c.XML(http.StatusOK, sitemap)
}

func (con SitemapController) Author(c *gin.Context) {

	var res []*model.Users
	var err error
	qe := query.Use(ay.Db).Users
	if res, err = qe.WithContext(c).Select(qe.Account).Find(); err != nil {
		ay.Logger.Error(err.Error())
		return
	}

	var urls []SitemapUrl
	for _, v := range res {

		loc := controllers.SiteConf[model.ConfigWebUrlKey] + "author/" + url.QueryEscape(v.Account)
		urls = append(urls, SitemapUrl{
			Loc:     loc,
			Lastmod: time.Now().Format("2006-01-02T15:04:05+08:00"),
		})
	}

	sitemap := SitemapUrlSet{
		URLs: urls,
	}

	encoder := xml.NewEncoder(os.Stdout)
	encoder.Indent("", "  ")

	if err := encoder.Encode(sitemap); err != nil {
		ay.Logger.Error(err.Error())
		return
	}
	c.XML(http.StatusOK, sitemap)
}
