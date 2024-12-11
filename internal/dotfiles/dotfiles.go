package dotfiles

import (
	"errors"
	"net/http"

	// "fmt"
	"io/fs"
	"log"
	"os"
	"os/exec"
	"path/filepath"

	"github.com/rewgs/reaconfig/internal/dot"
)

type Dotfiles struct {
	Src string
	Dst string
}

func (d *Dotfiles) valid() bool {
	// src
	// TODO: Verify that `src` is a valid URL.

	// dst
	_, err := os.Open(d.Dst)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			// TODO: file doesn't exist; deal with it.
		} else {
			// TODO: Some other error was thrown, so deal with it. Maybe return it?
		}
	}
	return true
}

func New(src string, dst string) *Dotfiles {
	dotfiles := Dotfiles{
		Src: src,
		Dst: dst,
	}

	if !dotfiles.valid() {
		// TODO: handle
	}

	return &dotfiles
}

// Clones dotfiles repo
func (d *Dotfiles) Clone() {
	checkIfExists := func(dir string) {
		_, err := os.Open(dir)
		if err != nil {
			// TODO: Alert user d.Dst exists, give option to abort/not clone.
		}
	}

	gitClone := func(src string, dst string) {
		cmd := exec.Command("git", "clone", src, dst)
		err := cmd.Run()
		if err != nil {
			// TODO: Something went wrong
			//
			// If `git` is not available, offer to install, or:
			http.Get(src) // TODO: continue from here
		}
		checkIfExists(dst)
	}

	checkIfExists(d.Dst)
	gitClone(d.Src, d.Dst)
}

// Scans the Dotfiles directory for a `conman` directory.
// If exists: for each config in `conman` directory, looks for corresponding
// directory in Dotfiles and creates dot.New() if present.
// If not: maybe gives a warning? Not sure.
func (d *Dotfiles) GetDots() []*dot.Dot {
	// Scans the Dotfiles directory for a `conman` directory.
	getConmanDot := func(entries []fs.DirEntry) (*dot.Dot, bool) {
		var conman *dot.Dot
		for _, entry := range entries {
			if entry.Name() == "conman" {
				conman = dot.New(entry, nil)
				return conman, true
			}
		}
		return conman, false
	}

	// Returns all configs (lua files) in Dotfiles/conman directory
	getConfigFiles := func(conman *dot.Dot) []fs.DirEntry {
		var configs []fs.DirEntry

		path, err := filepath.Abs(conman.Path.Name())
		if err != nil {
			log.Fatal(err)
		}
		// TODO: Can I just get the entries of a fs.DirEntry directly without
		// converting to a string first like I did above? ^
		entries, err := os.ReadDir(path)
		if err != nil {
			log.Fatal(err)
		}

		for _, entry := range entries {
			path, err := filepath.Abs(entry.Name())
			if err != nil {
				// TODO:
			}
			if filepath.Ext(path) == ".lua" {
				configs = append(configs, entry)
			}
		}
		return configs
	}

	var dots []*dot.Dot

	entries, err := os.ReadDir(d.Dst)
	if err != nil {
		// Dotfiles is empty or otherwise can't be read.
		log.Fatal(err)
	}

	conman, found := getConmanDot(entries)
	if !found {
		// TODO:
	} else {
		dots = append(dots, conman)
	}

	configFiles := getConfigFiles(conman)

	// For each config in `conman` directory, looks for corresponding
	// directory in Dotfiles and creates dot.New() if present.
	for _, entry := range entries {
		if entry.IsDir() {
			for _, config := range configFiles {
				// TODO: Check for alt names as well
				if entry.Name() == config.Name() {
					dot := dot.New(entry, config)
					dots = append(dots, dot)
				}
			}
		}
	}

	return dots
}
