package main

import (
	"fmt"

	"github.com/rewgs/reaconfig/internal/reaper"
)

func main() {
	reaper := reaper.New()
	if reaper.IsInstalled() {
		fmt.Println("Reaper is installed")
	}
}
