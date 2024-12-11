// TODO: How to save which packages need to be installed via ReaPack?

// `config` 

package config

// Files to be kept in a config.
var relevantFiles []string {
    "ColorThemes/",
    "FXChains/",
    "KeyMaps/",
    "MenuSets/",
    "MouseMaps/",
    "OSC/",
    "presets/",
    "PreviousConfigFiles/",
    "ProjectTemplates/",
    "Scripts",
    "TrackTemplates/"
    "UserPlugins/" // Maybe? Or re-clone like ReaPack packages?
    // ...
    "reaper.themeconfig.ini"
    // ...
    "reaper.ini"
}

type Config struct {
	Path    string
	Default bool
	Active  bool
}

type config interface {
	Path()
	MakeActive()
	MakeInactive()
	Backup()
	checkForLargeFiles()
	checkForGitSubmodules()
    isDefault()
}

// Checks for files that are too big for a typical git repo.
// Recommended: 50 MB
// Absolute: 100 MB
func (c *Config) checkForLargeFiles() {
}

func (c *Config) isDefault() bool {
}

func New() *Config {
	config := Config{}

	return &config
}
