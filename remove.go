package main

import (
	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

func removeCommand() *cli.Command {
	return &cli.Command{
		Name:        "remove",
		Aliases:     []string{"r"},
		Usage:       "辞書を削除",
		Description: "辞書ファイルも削除されます",
		Action: func(ctx *cli.Context) error {
			urls := []string{}

			for _, jisyo := range sharedConfig.Jisyos {
				urls = append(urls, jisyo.URL)
			}

			prompt := &survey.MultiSelect{
				Message: "削除する辞書を選択",
				Options: urls,
			}

			selectedURLs := []string{}
			if err := survey.AskOne(prompt, &selectedURLs); err != nil {
				return nil // キャンセル
			}

			// 設定から削除

			// ファイルを削除

			return nil
		},
	}
}
