package main

import (
	"archive/zip"
	"io"
	"os"
	"path/filepath"
	"strings"
)

func zipFolder(folderPath, outputFilepath string) error {
	zipfile, err := os.Create(outputFilepath)
	if err != nil {
		return err
	}
	defer zipfile.Close()

	archive := zip.NewWriter(zipfile)
	defer archive.Close()

	baseFolder := filepath.Clean(folderPath)

	err = filepath.Walk(baseFolder, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		relPath := strings.TrimPrefix(path, baseFolder)
		relPath = strings.TrimPrefix(relPath, string(filepath.Separator))

		if info.IsDir() {
			if relPath == "" {
				return nil
			}
			_, err := archive.Create(relPath + "/")
			return err
		}

		file, err := os.Open(path)
		if err != nil {
			return err
		}
		defer file.Close()

		writer, err := archive.Create(relPath)
		if err != nil {
			return err
		}

		_, err = io.Copy(writer, file)
		return err
	})

	return err
}
