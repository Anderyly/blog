package sql

import (
	"blog/ay"
	"blog/dal/model"
	"github.com/spf13/cobra"
	"gorm.io/gen"
)

func Cmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "sql",
		Short: "生成sql",

		Run: func(cmd *cobra.Command, args []string) {
			runGen()
		},
	}
	return cmd
}

func runGen() {

	// 配置文件
	ay.Init()

	g := gen.NewGenerator(gen.Config{
		OutPath: "./dal/query",
	})
	list := []interface{}{
		model.Users{},
		model.Contents{},
		model.Comments{},
		model.Metas{},
		model.Relationships{},
		model.Links{},
		model.Config{},
		model.Templates{},
		model.Plugins{},
		model.Views{},
	}

	ay.Db.DisableForeignKeyConstraintWhenMigrating = true
	if err := ay.Db.AutoMigrate(list...); err != nil {
		ay.Logger.Error(err.Error())
	}
	g.ApplyBasic(list...) // 绑定表
	g.Execute()
}
