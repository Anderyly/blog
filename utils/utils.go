/*
 * @Author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc/
 * @copyright Copyright (c) 2023
 */

package utils

import (
	"fmt"
	"golang.org/x/net/html"
	"html/template"
	"strings"
)

func SliceUnique[T any](list []T) []T {
	tpl := make(map[interface{}]int)
	var result []T
	for _, v := range list {
		if _, ok := tpl[v]; !ok {
			result = append(result, v)
			tpl[v] = 1
		}
	}
	return result
}

func GetAbstract(content string, maxChars int) string {
	reader := strings.NewReader(content)
	tokenizer := html.NewTokenizer(reader)

	var summary string
	var currentChars int

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			goto Done
		case html.TextToken:
			token := tokenizer.Token()
			text := strings.TrimSpace(token.Data)
			currentChars += len(text)
			if currentChars <= maxChars {
				summary += text
			} else {
				break
			}
		}
	}

Done:
	return summary
}

func HtmlEscape(content string) string {
	return template.HTMLEscapeString(content)

}

func FindFirstImage(htmlContent string) (string, error) {
	reader := strings.NewReader(htmlContent)
	tokenizer := html.NewTokenizer(reader)

	for {
		tokenType := tokenizer.Next()
		switch tokenType {
		case html.ErrorToken:
			return "", fmt.Errorf("未找到图片")
		case html.StartTagToken, html.SelfClosingTagToken:
			token := tokenizer.Token()
			if token.Data == "img" {
				for _, attr := range token.Attr {
					if attr.Key == "src" {
						return attr.Val, nil
					}
				}
			}
		}
	}
}

func Page(page, total, limit int) (pages []int) {
	allPage := total / limit

	if page > 3 {
		if allPage-3 < page {
			for i := 6 - (allPage - page); i >= 1; i-- {
				pages = append(pages, page-i)
			}
		} else {
			for i := 3; i >= 1; i-- {
				pages = append(pages, page-i)
			}
		}

	} else {
		for i := 1; i < page; i++ {
			pages = append(pages, i)
		}
	}
	pages = append(pages, page)

	length := len(pages)

	j := 1
	for i := length; i < 7; i++ {
		if page+j > allPage {
			break
		}
		pages = append(pages, page+j)
		j++
	}

	return

}
