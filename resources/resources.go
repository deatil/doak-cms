package resources

import (
    "embed"
    "io/fs"
    "net/http"
)

//go:embed static/*
var Static embed.FS

//go:embed views
var Views embed.FS

// 静态文件
func StaticAssets() http.FileSystem {
    return http.FS(Static)
}

// 模板文件
func ViewFileSystem() http.FileSystem {
    fsys, err := fs.Sub(Views, "views")
    if err != nil {
        panic(err)
    }

    return http.FS(fsys)
}
