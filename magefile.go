//go:build mage

package main

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"strings"

	"github.com/magefile/mage/sh"
)

// Proto recompiles the proto files to go files
func Proto() error {
	wd, err := os.Getwd()
	if err != nil {
		return err
	}
	protoDir := filepath.Join(wd, "internal/proto")

	protocParameters := []string{
		fmt.Sprintf("-I=%s", protoDir),
		"--go_opt=module=github.com/timvosch/gospotlib",
		fmt.Sprintf("--go_out=%s", wd),
	}
	err = filepath.Walk(protoDir, func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || !strings.HasSuffix(path, ".proto") {
			return nil
		}
		protocParameters = append(protocParameters, path)
		return nil
	})
	if err != nil {
		return err
	}

	return sh.RunV("protoc", protocParameters...)
}

func Clean() error {
	exts := []string{".pb.go"}

	wd, err := os.Getwd()
	if err != nil {
		return err
	}

	return filepath.Walk(filepath.Join(wd, "pkg/pb"), func(path string, info fs.FileInfo, err error) error {
		if err != nil {
			return err
		}

		for _, ext := range exts {
			if !strings.HasSuffix(path, ext) {
				continue
			}
			return sh.Rm(path)
		}

		return nil
	})
}

func Docs() error {
	return sh.RunV(
		"docker", "run",
		"--rm",
		"-p", "8000:8000",
		"-v", "$PWD:/docs",
		"squidfunk/mkdocs-material",
	)
}
