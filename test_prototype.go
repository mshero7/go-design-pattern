package main

import (
	"fmt"
	p "go-example/design-pattern/prototype"
)

func testPrototype() {
	file1 := &p.File{Name: "File1"}
	file2 := &p.File{Name: "File2"}
	file3 := &p.File{Name: "File3"}

	folder1 := &p.Folder{
		Children: []p.Inode{file1},
		Name:     "Folder1",
	}

	folder2 := &p.Folder{
		Children: []p.Inode{folder1, file2, file3},
		Name:     "Folder2",
	}
	fmt.Println("\nPrinting hierarchy for Folder2")
	folder2.Print("  ")

	cloneFolder := folder2.Clone()
	fmt.Println("\nPrinting hierarchy for clone Folder")
	cloneFolder.Print("  ")
}
