package nospaceleftondevice

func newDir(parent *Dir) *Dir {
	return &Dir{
		parent:  parent,
		files:   []*File{},
		subDirs: map[string]*Dir{},
	}
}

type HasSize interface {
	getSize() int
}

type File struct {
	name string
	size int
}

func (f *File) getSize() int {
	return f.size
}

type Dir struct {
	parent  *Dir
	files   []*File
	subDirs map[string]*Dir
}

func (d *Dir) getChildren() []HasSize {
	children := []HasSize{}
	for _, subDir := range d.subDirs {
		children = append(children, subDir)
	}
	for _, file := range d.files {
		children = append(children, file)
	}
	return children
}

func (d *Dir) getSize() int {
	size := 0
	for _, child := range d.getChildren() {
		size += child.getSize()
	}
	return size
}
