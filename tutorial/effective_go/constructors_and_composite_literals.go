package main

type File struct {
	fd      int
	name    string
	dirinfo string
	nepipe  int
}

func NewFile(fd int, name string) *File {
	if fd < 0 {
		return nil
	}
	f := File{fd, name, nil, 0}
	return &f

}
