package goui_tools

import (
	"github.com/fipress/go-rj"
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
