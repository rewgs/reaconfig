package reaper

import (
	"os"
	"path"
	"runtime"
)

type Path struct {
	App       *string
	Bin       *string
	Resources *string
}

// Returns the directory containing the REAPER binary.
func getAppPath() string {
	var appPath string
	switch runtime.GOOS {
	case "darwin":
		appPath = path.Join("/", "Applications", "REAPER.app")
	// TODO:
	case "linux":
	// TODO:
	case "windows":
	}
	return appPath
}

// Returns the path to the REAPER binary.
func getBinPath(appPath string) string {
	var binPath string
	switch runtime.GOOS {
	case "darwin":
		binPath = path.Join(appPath, "Contents", "MacOS", "REAPER")
	}
	return binPath
}

func getResourcesPath() string {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		// TODO:
	}

	var resourcesPath string
	switch runtime.GOOS {
	case "darwin":
		resourcesPath = path.Join(homeDir, "Library", "Application Support", "REAPER")
	case "linux":
		resourcesPath = path.Join(homeDir, ".config", "REAPER")
	case "windows":
		resourcesPath = path.Join(homeDir, "AppData", "Roaming", "REAPER")
	default:
		// TODO:
	}
	return resourcesPath
}

func getPaths() *Path {
	app := getAppPath()
	bin := getBinPath(app)
	res := getResourcesPath()

	paths := Path{
		App:       &app,
		Bin:       &bin,
		Resources: &res,
	}

	return &paths
}
