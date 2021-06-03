package loginhub

import (
	"embed"
	"io/fs"
)

func subfs(fsys fs.FS, dir string) fs.FS {
	subFsys, err := fs.Sub(fsys, dir)
	if err != nil {
		panic(err)
	}
	return subFsys
}

//go:embed build/styles
var styleFS embed.FS

func StyleDir() fs.FS {
	return subfs(styleFS, "build/styles")
}

//go:embed assets/scripts
var scriptFS embed.FS

func ScriptDir() fs.FS {
	return subfs(scriptFS, "assets/scripts")
}

//go:embed assets/images
var imageFS embed.FS

func ImageDir() fs.FS {
	return subfs(imageFS, "assets/images")
}
