package view

import (
    "github.com/gofiber/template/jet"

    "github.com/deatil/doak-cms/pkg/time"
    "github.com/deatil/doak-cms/pkg/utils"

    "github.com/deatil/doak-cms/app/url"
    "github.com/deatil/doak-cms/app/data"
)

// 设置模板方法
func SetViewFuncs(engine *jet.Engine) {
    // 常用方法
    engine.AddFunc("createTime", time.CreateFromTimestamp)
    engine.AddFunc("formatTime", time.CreateFromFormat)
    engine.AddFunc("nowTime", time.Now)
    engine.AddFunc("formatSize", utils.FormatSize)

    engine.AddFunc("adminUrl", url.AdminUrl)
    engine.AddFunc("avatarUrl", url.AvatarUrl)

    // 静态文件及附件
    engine.AddFunc("assets", url.Assets)
    engine.AddFunc("attachUrl", url.AttachUrl)
    engine.AddFunc("attachPath", url.AttachPath)
    engine.AddFunc("attachUrlWithId", url.AttachUrlWithId)
    engine.AddFunc("attachPathWithId", url.AttachPathWithId)

    // cms 链接
    engine.AddFunc("cateUrl", url.CateUrl)
    engine.AddFunc("artUrl", url.ArtUrl)
    engine.AddFunc("tagUrl", url.TagUrl)
    engine.AddFunc("pageUrl", url.PageUrl)

    // 查询数据
    engine.AddFunc("getSettings", data.GetSettings)
    engine.AddFunc("getSetting", data.GetSetting)
    engine.AddFunc("getCateList", data.GetCateList)
    engine.AddFunc("getCateInfo", data.GetCateInfo)
    engine.AddFunc("getCateInfoWithSlug", data.GetCateInfoWithSlug)
    engine.AddFunc("getArtList", data.GetArtList)
    engine.AddFunc("getArtInfo", data.GetArtInfo)
    engine.AddFunc("getPageInfo", data.GetPageInfo)
    engine.AddFunc("getTagList", data.GetTagList)
}
