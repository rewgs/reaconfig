package reaper

type Reaper struct {
	Name string
	Path *Path // pointer to struct containing relevant path
}

func (r *Reaper) IsInstalled() bool {
	if r.Path.App != nil && r.Path.Bin != nil {
		return true
	}
	return false
}

// TODO:
// func (r *Reaper) isOpen() bool {
// 	proc, err := os.FindProcess(r.BinPath)
// }

func New() *Reaper {
	reaper := Reaper{
		Name: "REAPER",
		Path: getPaths(),
	}
	return &reaper
}
