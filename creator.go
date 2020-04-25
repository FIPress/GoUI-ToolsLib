package goui_tools

import (
	"fmt"
	"github.com/fipress/fiputil"
	"os"
	"path/filepath"
)

const (
	fileExists = "File exists and is not a directory:"
	dirExists  = "Dir exists and not empty:"
	sampleDir  = "sample"
)

func CreateProject(name string, settings *Settings) string {
	fullPath := filepath.Join(settings.WorkingDir, name)
	fi, err := os.Stat(fullPath)
	if err == nil {
		if !fi.IsDir() {
			return fileExists + fullPath
		}

		if fi.Size() > 100 {
			return dirExists + fullPath
		}
	} else {
		err = os.MkdirAll(fullPath, 0755)
		if err != nil {
			return fmt.Sprintln("Create directory", fullPath, " error:", err)
		}
	}

	ok := false
	settings.BinDir, ok = getBinDir()
	if !ok {
		return ""
	}
	src := filepath.Join(settings.BinDir, sampleDir)
	err = fiputil.CopyDir(src, fullPath, nil)

	if err != nil {
		fmt.Println()
		return "Copy directory failed, error:" + err.Error()
	}

	createPackageConfig(fullPath)

	return ""
}
