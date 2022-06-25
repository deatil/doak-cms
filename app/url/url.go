package url

import (
    "fmt"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/config"

    "github.com/deatil/doak-cms/app/model"
)

// 资源
func Assets(path string) string {
    return fmt.Sprintf("/static/%s", path)
}

// =======================

// 管理页面链接
func AdminUrl(path ...string) string {
    cfg := config.Section("app")
    prefix := cfg.Key("router-prefix").String()

    newPath := "/" + prefix
    if len(path) > 0 {
        return newPath + "/" + path[0]
    }

    return newPath
}

// =======================

// 分类链接
func CateUrl(path string) string {
    if path == "" {
        path = "/c"
        return path
    }

    return fmt.Sprintf("/c/%s", path)
}

// 文章链接
func ArtUrl(path string) string {
    if path == "" {
        path = "/a"
        return path
    }

    return fmt.Sprintf("/a/%s", path)
}

// 标签链接
func TagUrl(path string) string {
    return fmt.Sprintf("/tag/%s", path)
}

// 单页链接
func PageUrl(path string) string {
    return fmt.Sprintf("/p/%s", path)
}

// =======================

// 附件链接
func AttachUrl(path string) string {
    cfg := config.Section("upload")

    fileurl := fmt.Sprintf(cfg.Key("url").String() + "/%s", path)

    return fileurl
}

// 附件具体位置
func AttachPath(path string) string {
    cfg := config.Section("upload")

    filename := fmt.Sprintf(cfg.Key("dir").String() + "/%s", path)

    return filename
}

// =======================

// 附件链接
func AttachUrlWithId(id string) string {
    cfg := config.Section("upload")

    path := ""
    var data model.Attach
    db.Engine().
        Where("id = ?", id).
        Get(&data)
    if data.Id != "" {
        path = data.Path
    }

    fileurl := fmt.Sprintf(cfg.Key("url").String() + "/%s", path)

    return fileurl
}

// 附件具体位置
func AttachPathWithId(id string) string {
    cfg := config.Section("upload")

    path := ""
    var data model.Attach
    db.Engine().
        Where("id = ?", id).
        Get(&data)
    if data.Id != "" {
        path = data.Path
    }

    filename := fmt.Sprintf(cfg.Key("dir").String() + "/%s", path)

    return filename
}

// 头像链接
func AvatarUrl(id string) string {
    if id == "" {
        return Assets("admin/image/avatar-default.jpg")
    }

    avatar := AttachUrlWithId(id)

    return avatar
}
