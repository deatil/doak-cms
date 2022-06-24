package admin

import (
    "fmt"
    "regexp"
    "strings"
    "github.com/gofiber/fiber/v2"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/http"
    "github.com/deatil/doak-cms/pkg/utils"
    "github.com/deatil/doak-cms/pkg/config"

    "github.com/deatil/doak-cms/app/model"
)

/**
 * 上传
 *
 * @create 2022-6-19
 * @author deatil
 */
type Upload struct{}

// 图片
func (this *Upload) Image(ctx *fiber.Ctx) error {
    cfg := config.Section("upload")

    // 接收上传的文件
    file, err := ctx.FormFile("file")
    if err != nil {
        return http.Error(ctx, "上传失败")
    }

    filsnames := strings.Split(file.Filename, ".")
    if len(filsnames) != 2 {
        return http.Error(ctx, "上传失败")
    }

    // 验证后缀
    fileext := filsnames[1]
    matchExt := cfg.Key("image-ext").String()
    if match, _ := regexp.MatchString("(?i)^(" + matchExt + ")$", fileext); !match {
        return http.Error(ctx, "上传的图片格式错误")
    }

    // 文件 md5
    filemd5 := ""

    // 打开临时文件
    attachFileopen, err := file.Open()
    defer attachFileopen.Close()
    if err != nil {
        return http.Error(ctx, "上传失败")
    }

    filemd5, _ = utils.FileMD5WithStream(attachFileopen)

    // 已存在信息返回
    var data model.Attach
    _, err = db.Engine().
        Where("md5 = ?", filemd5).
        Get(&data)
    if data.Id != "" {
        fileurl := fmt.Sprintf(cfg.Key("url").String() + "/%s", data.Path)
        return http.Success(ctx, "上传成功", map[string]string{
            "id": data.Id,
            "url": fileurl,
        })
    }

    newFilename := utils.MD5(utils.Uniqueid())

    filepath := fmt.Sprintf("images/%s.%s", newFilename, fileext)
    filename := fmt.Sprintf(cfg.Key("dir").String() + "/%s", filepath)

    // 保存文件到磁盘
    if err = ctx.SaveFile(file, filename); err != nil {
        return http.Error(ctx, "上传失败")
    }

    insertData := &model.Attach{
        Id: utils.Uniqueid(),
        Name: file.Filename,
        Path: filepath,
        Ext: fileext,
        Size: file.Size,
        Md5: filemd5,
        Status: 1,
        AddIp: ctx.IP(),
    }
    _, err = db.Engine().Insert(insertData)
    if err != nil {
        return http.Error(ctx, "添加失败")
    }

    fileurl := fmt.Sprintf(cfg.Key("url").String() + "/%s", filepath)
    return http.Success(ctx, "上传成功", map[string]string{
        "id": insertData.Id,
        "url": fileurl,
    })
}

// 文件
func (this *Upload) File(ctx *fiber.Ctx) error {
    cfg := config.Section("upload")

    // 接收上传的文件
    file, err := ctx.FormFile("file")
    if err != nil {
        return http.Error(ctx, "上传失败")
    }

    filsnames := strings.Split(file.Filename, ".")
    if len(filsnames) != 2 {
        return http.Error(ctx, "上传失败")
    }

    // 验证后缀
    fileext := filsnames[1]
    matchExt := cfg.Key("file-ext").String()
    if match, _ := regexp.MatchString("(?i)^(" + matchExt + ")$", fileext); !match {
        return http.Error(ctx, "上传的文件格式错误")
    }

    // 文件 md5
    filemd5 := ""

    // 打开临时文件
    attachFileopen, err := file.Open()
    defer attachFileopen.Close()
    if err != nil {
        return http.Error(ctx, "上传失败")
    }

    filemd5, _ = utils.FileMD5WithStream(attachFileopen)

    // 已存在信息返回
    var data model.Attach
    _, err = db.Engine().
        Where("md5 = ?", filemd5).
        Get(&data)
    if data.Id != "" {
        fileurl := fmt.Sprintf(cfg.Key("url").String() + "/%s", data.Path)
        return http.Success(ctx, "上传成功", map[string]string{
            "id": data.Id,
            "url": fileurl,
        })
    }

    newFilename := utils.MD5(utils.Uniqueid())

    filepath := fmt.Sprintf("attach/%s.%s", newFilename, fileext)
    filename := fmt.Sprintf(cfg.Key("dir").String() + "/%s", filepath)

    // 保存文件到磁盘
    if err = ctx.SaveFile(file, filename); err != nil {
        return http.Error(ctx, "上传失败")
    }

    insertData := &model.Attach{
        Id: utils.Uniqueid(),
        Name: file.Filename,
        Path: filepath,
        Ext: fileext,
        Size: file.Size,
        Md5: filemd5,
        Status: 1,
        AddIp: ctx.IP(),
    }
    _, err = db.Engine().Insert(insertData)
    if err != nil {
        return http.Error(ctx, "添加失败")
    }

    fileurl := fmt.Sprintf(cfg.Key("url").String() + "/%s", filepath)
    return http.Success(ctx, "上传成功", map[string]string{
        "id": insertData.Id,
        "url": fileurl,
    })
}
