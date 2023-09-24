package main

import (
	"os"

	"github.com/urfave/cli/v2"
)

var (
	version              = "unknown"
	sharedConfig *config = nil
)

func main() {
	app := &cli.App{
		Name:  "jisyo",
		Usage: "📚 SKK辞書マネージャ",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config",
				Usage:   "設定ファイルのパス",
				Aliases: []string{"c"},
			},
		},
		Commands: []*cli.Command{
			initCommand(),
			downloadCommand(),
			addCommand(),
			removeCommand(),
			listCommand(),
			configCommand(),
		},
		Before: func(ctx *cli.Context) error {
			// 設定ファイルのパスが未指定なら、デフォルト値を入れる
			if ctx.String("config") == "" {
				path, err := defaultConfigPath()
				if err != nil {
					return err
				}

				ctx.Set("config", path)
			}

			// init なら設定ファイルを読み込まない
			if ctx.Args().First() == "init" {
				return nil
			}

			// 設定ファイルを読み込む
			if cfg, exit := loadConfig(ctx); exit != nil {
				return exit
			} else {
				dirPath, exit := expandTilde(cfg.DirPath)
				if exit != nil {
					return exit
				}

				sharedConfig = cfg
				sharedConfig.DirPath = dirPath
			}

			return nil
		},
		Version: version,
	}

	if err := app.Run(os.Args); err != nil {
		cli.Exit(err, exitCodeErr.ToInt())
	}
}
