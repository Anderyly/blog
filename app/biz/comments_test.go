/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package biz

import (
	"blog/ay"
	"context"
	"fmt"
	"testing"
)

func Test_comments_ListByContentId(t *testing.T) {
	ay.Init()
	b := NewBiz(context.Background())
	list, err := b.Comments.ListByContentId(175, b.Comments.WithSetChild())
	fmt.Println(err)
	for _, v := range list {
		t.Log(v.Coid, v.ChildComments)
	}

}
