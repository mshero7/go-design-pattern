package prototype

import "fmt"

type Folder struct {
	Children []Inode
	Name     string
}

func (f *Folder) Print(indentation string) {
	fmt.Println(indentation + f.Name)

	for _, j := range f.Children {
		j.Print(indentation + indentation)
	}
}

func (f *Folder) Clone() Inode {
	cloneFolder := &Folder{Name: f.Name + "_clone"}

	var tempChildern []Inode

	for _, j := range f.Children {
		copy := j.Clone()
		tempChildern = append(tempChildern, copy)
	}

	cloneFolder.Children = tempChildern
	return cloneFolder
}
