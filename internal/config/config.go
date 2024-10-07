package config

type Config struct {
    Path    string
    Active  bool
}

type config interface {
    GetPath()
    MakeActive()
    MakeInactive()
    Backup()
    checkForLargeFiles()
    checkForGitSubmodules()
    addGitSubmodule()
}

// Checks for files that are too big for a typical git repo.
// Recommended: 50 MB
// Absolute: 100 MB
func (c *Config) checkForLargeFiles() {
}
