package main

import (
	"fmt"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

func addCommand() *cli.Command {
	return &cli.Command{
		Name:        "add",
		Aliases:     []string{"a"},
		Usage:       "辞書を追加",
		Description: "辞書のURLを追加します",
		ArgsUsage:   "<URL>",
		Before: func(ctx *cli.Context) error {
			if ctx.Args().Len() != 1 {
				return cli.Exit("引数エラー: 1つの引数が必要です", exitCodeErrArg.ToInt())
			}
			if !strings.HasPrefix(ctx.Args().First(), "http") {
				return cli.Exit("引数エラー: URLが不正です", exitCodeErrArg.ToInt())
			}
			return nil
		},
		Action: func(ctx *cli.Context) error {
			url := ctx.Args().First()

			// 重複確認
			dupIndex := -1
			for i, j := range sharedConfig.Jisyos {
				if j.URL == url {
					do := false
					prompt := &survey.Confirm{Message: "この辞書は既に追加されています。上書きしますか？", Default: false}

					if err := survey.AskOne(prompt, &do); err != nil || !do {
						return exitCancel()
					}

					dupIndex = i
					break
				}
			}

			fmt.Println("🚚 ダウンロードを開始します")

			newJisyo, exit := downloadJisyo(url)
			if exit != nil {
				return exit
			}

			// 設定に反映
			if dupIndex == -1 {
				sharedConfig.Jisyos = append(sharedConfig.Jisyos, *newJisyo)
			} else {
				sharedConfig.Jisyos[dupIndex] = *newJisyo
			}

			showSuccess()

			return saveConfig(ctx, *sharedConfig)
		},
	}
}
