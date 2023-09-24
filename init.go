package main

import (
	"os"

	"github.com/AlecAivazis/survey/v2"
	"github.com/urfave/cli/v2"
)

func initCommand() *cli.Command {
	return &cli.Command{
		Name:        "init",
		Usage:       "設定ファイルを初期化",
		Description: "既に存在している場合、上書きされます",
		Aliases:     []string{"i"},
		Action: func(ctx *cli.Context) error {
			path := ctx.String("config")

			// 既にある
			if _, err := os.Stat(path); err == nil {
				do := false
				prompt := &survey.Confirm{
					Message: "既に設定ファイルが存在します。初期化しますか？",
					Default: false,
				}

				if err := survey.AskOne(prompt, &do); err != nil || !do {
					return exitCancel()
				}
			}

			// 辞書ディレクトリ
			prompt := &survey.Input{Message: "辞書の保存先"}
			dirPath := ""
			if err := survey.AskOne(prompt, &dirPath, survey.WithValidator(survey.Required)); err != nil {
				return exitCancel()
			}

			newConfig := config{
				DirPath: dirPath,
			}

			showSuccess()

			return saveConfig(ctx, newConfig)
		},
	}

}
