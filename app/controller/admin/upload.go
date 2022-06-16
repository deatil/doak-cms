package admin

import (
    "fmt"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/http"
    "github.com/deatil/doak-cms/pkg/utils"
    "github.com/deatil/doak-cms/pkg/config"
)

// 上传
type Upload struct{}

// 文件
func (this *Upload) File(ctx *fiber.Ctx) error {
    cfg := config.Section("upload")

    // 解析多部分表单。
    if form, err := ctx.MultipartForm(); err == nil {
        // 从 "文档 "键中获取所有文件。
        files := form.File["documents"]

        var url []string

        // 循环浏览文件。
        for _, file := range files {
            // fmt.Println(file.Filename, file.Size, file.Header["Content-Type"][0])

            newFilename := utils.MD5(file.Filename)
            filename := fmt.Sprintf(cfg.Key("dir").String() + "/%s", newFilename)

            url = append(url, fmt.Sprintf(cfg.Key("url").String() + "/%s", newFilename))

            // 保存文件到磁盘。
            if err := ctx.SaveFile(file, filename); err != nil {
                return http.Error(ctx, 1, "上传失败")
            }
        }

        return http.Success(ctx, "上传成功", url)
    }

    return http.Error(ctx, 1, "上传失败")
}
