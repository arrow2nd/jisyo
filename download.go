package main

import (
	"compress/gzip"
	"errors"
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/urfave/cli/v2"
)

func downloadCommand() *cli.Command {
	return &cli.Command{
		Name:        "downlaod",
		Aliases:     []string{"d"},
		Usage:       "辞書をダウンロード",
		Description: "登録済みの辞書を指定のディレクトリに一括でダウンロードします",
		Action: func(ctx *cli.Context) error {
			fmt.Println("🚚 ダウンロードを開始します")

			maxDownloadJisyos := len(sharedConfig.Jisyos)
			for i, jisyo := range sharedConfig.Jisyos {
				fmt.Printf("[%d/%d] 📦 %s (%s)\n", i+1, maxDownloadJisyos, jisyo.Name, jisyo.URL)

				if _, exit := downloadJisyo(jisyo.URL); exit != nil {
					return exit
				}
			}

			showSuccess()

			return nil
		},
	}
}

func downloadJisyo(downloadUrl string) (*jisyo, cli.ExitCoder) {
	res, err := http.Get(downloadUrl)
	if err != nil {
		msg := fmt.Errorf("ダウンロードに失敗しました: %w", err)
		return nil, cli.Exit(msg, exitCodeErrNetwork.ToInt())
	}
	defer res.Body.Close()

	jisyoDir, exit := sharedConfig.GetDirPath()
	if exit != nil {
		return nil, exit
	}

	// 展開
	rawFilename := filepath.Base(res.Request.URL.Path)
	filename, err := extract(res.Body, rawFilename, jisyoDir)
	if err != nil {
		msg := fmt.Errorf("展開に失敗しました: %w", err)
		return nil, cli.Exit(msg, exitCodeErrExtract.ToInt())
	}

	return &jisyo{
		Name: filename,
		URL:  downloadUrl,
	}, nil
}

func extract(src io.Reader, filename, distDir string) (string, error) {
	var reader io.Reader = src

	// gzip
	if strings.HasSuffix(filename, ".gz") {
		g, err := gzip.NewReader(src)
		if err != nil {
			return "", err
		}
		defer g.Close()

		filename = strings.TrimSuffix(filename, ".gz")
		reader = g
	}

	// tar, zip
	// TODO: 必要になったらやるかも
	if strings.HasSuffix(filename, ".tar") || strings.HasSuffix(filename, ".zip") {
		return "", errors.New("アーカイブ形式の展開には未対応です")
	}

	// 書き込み
	distPath := filepath.Join(distDir, filename)
	dist, err := os.Create(distPath)
	if err != nil {
		return "", err
	}
	defer dist.Close()

	if _, err := io.Copy(dist, reader); err != nil {
		return "", err
	}

	return filename, nil
}
