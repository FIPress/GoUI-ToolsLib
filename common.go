package goui_tools

import (
	"github.com/fipress/go-rj"
	"os"
	"path/filepath"
)

const packageConfigFile = "package.rj"

type packageConfig struct {
	Name        string
	VersionCode int
	VersionName string
}

func createPackageConfig(path string) (cfg *packageConfig) {
	cfg = &packageConfig{
		Name:        "GoUIApp Name",
		VersionCode: 1,
		VersionName: "0.0.1",
	}

	fullName := filepath.Join(path, packageConfigFile)

	rj.MarshalToFile(cfg, fullName)
	return
}

func getBinDir() (binDir string, ok bool) {
	dir, err := os.Executable()
	if err != nil {
		getLogger().Error("Get executable directory of GoUI-CLI failed: %s", err.Error())
		return
	}

	dir, err = filepath.EvalSymlinks(dir)

	if err != nil {
		getLogger().Error("Get executable directory of GoUI-CLI failed: %s", err.Error())
		return
	}

	binDir = filepath.Dir(dir)
	ok = true

	return
}
