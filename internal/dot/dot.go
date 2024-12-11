package dot

import (
	"io/fs"

	lua "github.com/yuin/gopher-lua"
	// "github.com/rewgs/conman/internal/api"
)

// A `Dot` is an application or program configured with dotfiles.
type Dot struct {
	Path   fs.DirEntry // e.g. ~/dotfiles/nvim
	Config fs.DirEntry // e.g. ~/conman/neovim.lua
	Name   string
}

func New(path fs.DirEntry, config fs.DirEntry) *Dot {
	dot := Dot{
		Path:   path,
		Config: config,
		Name:   path.Name(),
	}

	return &dot
}

// TODO:
func (d *Dot) parseConfig() {
}

// Checks if the given Dot has an `Install()` function.
func (d *Dot) hasInstall() bool {
}

// Checks if the given Dot has an `Install()` function.
func (d *Dot) hasSetup() bool {
}

// Runs the config file's Install function.
func (d *Dot) Install() {
	Lua := lua.NewState()
	defer Lua.Close()
	// ...
}

// Runs the config file's Setup function.
func (d *Dot) Setup() {
	Lua := lua.NewState()
	defer Lua.Close()
	// ...
}
