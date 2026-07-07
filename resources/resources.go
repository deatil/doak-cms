package resources

import (
    "embed"
    "io/fs"
    "net/http"
)

//go:embed static
var Static embed.FS

//go:embed views
var Views embed.FS

func StaticFileSystem() fs.FS {
    fsys, err := fs.Sub(Static, "static")
    if err != nil {
        panic(err) 
    }

    return fsys
}

// 模板文件
func ViewFileSystem() http.FileSystem {
    fsys, err := fs.Sub(Views, "views")
    if err != nil {
        panic(err)
    }

    return http.FS(fsys)
}
