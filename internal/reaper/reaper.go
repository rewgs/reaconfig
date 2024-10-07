package reaper

import (
    "path"
    "os"
    "runtime"
)

type Reaper struct {
    ResourcesPath   string
}

func getResourcesPath() string {
    var resourcesPath string

    switch runtime.GOOS {
    case "darwin":
        homeDir, err := os.UserHomeDir()
        if err != nil {
            // TODO:
        }
        resourcesPath = path.Join(homeDir, "Library", "Application Support", "REAPER")
    // TODO:
    // case "linux":
    // TODO:
    // case "windows":
    default:
        // TODO:
    }
    return resourcesPath
}
