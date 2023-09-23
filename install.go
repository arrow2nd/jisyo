package main

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func installCommand() *cli.Command {
	return &cli.Command{
		Name:        "install",
		Aliases:     []string{"i"},
		Usage:       "辞書をインストール",
		Description: "登録済みの辞書を指定のディレクトリに一括でダウンロードします",
		Action: func(c *cli.Context) error {
			fmt.Println(c.String("config"))
			return nil
		},
	}
}
