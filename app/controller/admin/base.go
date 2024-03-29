package admin

import (
    "strings"

    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/resources"
    "github.com/deatil/doak-cms/pkg/utils"
    "github.com/deatil/doak-cms/pkg/config"
    "github.com/deatil/doak-cms/app/auth"
    "github.com/deatil/doak-cms/app/response"
)

/**
 * 基类
 *
 * @create 2022-6-19
 * @author deatil
 */
type Base struct{}

// 主题
func (this *Base) Theme(tpl string) string {
    tpl = "admin/" + tpl

    return tpl
}

// 视图
func (this *Base) View(ctx *fiber.Ctx, tpl string, data fiber.Map) error {
    // 登录数据
    data["user"] = auth.GetUserInfo(ctx)

    // 管理员
    data["isAdmin"] = auth.IsAdmin(ctx)

    return ctx.Render(this.Theme(tpl), data)
}

// 模板列表
func (this *Base) ListTplFiles(prefix string) []string {
    if config.IsUseEmbed {
        return this.ListEmbedTplFiles(prefix)
    } else {
        return this.ListDirTplFiles(prefix)
    }
}

// 模板列表
func (this *Base) ListDirTplFiles(prefix string) []string {
    cfg := config.Section("view")

    views := cfg.Key("views").String()
    ext := cfg.Key("ext").String()

    dir := views + "/" + response.CmsTheme("")

    files := utils.ListFiles(dir)

    newFiles := make([]string, 0)
    if len(files) > 0 {
        for _, file := range files {
            if strings.HasPrefix(file, prefix) {
                file = strings.TrimSuffix(file, ext)
                newFiles = append(newFiles, file)
            }
        }
    }

    return newFiles
}

// embed 模板列表
func (this *Base) ListEmbedTplFiles(prefix string) []string {
    cfg := config.Section("view")

    ext := cfg.Key("ext").String()
    dir := response.CmsTheme("")

    vfs := resources.ViewFileSystem()

    f, err := vfs.Open(dir)
    if err != nil {
        return []string{}
    }

    files := utils.ListEmbedFiles(f)

    newFiles := make([]string, 0)
    if len(files) > 0 {
        for _, file := range files {
            if strings.HasPrefix(file, prefix) {
                file = strings.TrimSuffix(file, ext)
                newFiles = append(newFiles, file)
            }
        }
    }

    return newFiles
}
