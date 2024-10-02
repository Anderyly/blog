/*
 * @author anderyly
 * @email admin@aaayun.cc
 * @link http://blog.aaayun.cc
 * @copyright Copyright (c) 2022
 *
 */

package lib

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

var _ Dir = (*dir)(nil)

type Dir interface {
	IsPath(path string) (bool, error)                              // 判断路径是否存在
	Create(filePath string) error                                  // 创建文件夹 注意 linux系统/为根目录
	Get(localDir string) (int, error)                              // 获取文件夹下文件数量
	LReadName(name string, level int) (filePath string, err error) // 从当前目录向上查找文件
}

type dir struct {
}

func NewDir() Dir {
	return &dir{}
}

func (con *dir) IsPath(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func (con *dir) Create(filePath string) error {
	if ok, _ := con.IsPath(filePath); !ok {
		err := os.MkdirAll(filePath, os.ModePerm)
		if err != nil {
			return err
		}
		return err
	}
	return nil
}

func (con *dir) Get(localDir string) (num int, err error) {
	err = filepath.Walk(localDir, func(filename string, fi os.FileInfo, err error) error {
		if fi.IsDir() {
			return nil
		}
		num++
		return nil
	})
	return
}

func (con *dir) LReadName(name string, level int) (filePath string, err error) {
	for i := 0; i <= level; i++ {
		filePath = fmt.Sprintf("%s%s", strings.Repeat("../", i), name)
		var ok bool
		ok, err = con.IsPath(filePath)
		if err != nil {
			return
		}
		if ok {
			break
		}
	}
	return
}
