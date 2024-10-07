package reaper

import (
    "path"
    "os"
    "runtime"
)

type Reaper struct {
    ResourcesPath   string
}

type reaper interface {
    isOpen()                bool
    isInstalled()           bool
}

func (r *Reaper) isOpen() bool {
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
