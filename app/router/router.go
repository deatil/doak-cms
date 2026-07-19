package router

import (
    "github.com/gofiber/fiber/v3"
    "github.com/gofiber/fiber/v3/middleware/static"

    "github.com/deatil/doak-cms/pkg/config"

    "github.com/deatil/doak-cms/resources"
    "github.com/deatil/doak-cms/app/data"
    "github.com/deatil/doak-cms/app/middleware"
    "github.com/deatil/doak-cms/app/controller/cms"
    "github.com/deatil/doak-cms/app/controller/admin"
)

// 路由
func Http(app *fiber.App) {
    // CMS
    HttpCms(app)

    // 管理员
    HttpAdmin(app)

    // 静态文件
    HttpStatic(app)
}

// 静态文件
func HttpStatic(app *fiber.App) {
    // 上传文件
    app.Use("/upload/", static.New("./public/upload"))

    // 静态文件
    if config.IsUseEmbed {
        app.Use("/static", static.New("", static.Config{
            FS:       resources.StaticFileSystem(),
            Browse:   true,
            Compress: true,
            Download: false,
        }))
    } else {
        app.Use("/static/", static.New("./resources/static"))
    }
}

// CMS
func HttpCms(app *fiber.App) {
    // 网站是否开启检测
    siteopenCheckMiddleware := middleware.NewSiteopenCheck()

    indexController := new(cms.Index)
    app.Get("/", siteopenCheckMiddleware, indexController.Index).Name("cms.index")

    settings := data.GetSettings()
    data.SetCmsRouter(map[string]string{
        "cate_url": settings["website_cate_url"],
        "view_url": settings["website_view_url"],
        "tag_url": settings["website_tag_url"],
        "page_url": settings["website_page_url"],
    }, func(app *fiber.App, routes map[string]string) {
        cateController := new(cms.Cate)
        app.Get(routes["cate_url"], siteopenCheckMiddleware, cateController.Index).Name("cms.cate")
    
        viewController := new(cms.View)
        app.Get(routes["view_url"], siteopenCheckMiddleware, viewController.Index).Name("cms.view")
    
        tagController := new(cms.Tag)
        app.Get(routes["tag_url"], siteopenCheckMiddleware, tagController.Index).Name("cms.tag")
    
        pageController := new(cms.Page)
        app.Get(routes["page_url"], siteopenCheckMiddleware, pageController.Index).Name("cms.page")    
    })

}

// 管理员
func HttpAdmin(app *fiber.App) {
    cfg := config.Section("app")
    prefix := cfg.Key("router-prefix").String()

    // 后台前缀
    sys := app.Group("/" + prefix)

    // 不需要授权路由
    authController := new(admin.Auth)
    sys.Get("/captcha", authController.Captcha).Name("admin.captcha")
    sys.Get("/login", authController.Login).Name("admin.login")
    sys.Post("/login", authController.LoginCheck).Name("admin.login-check")

    // 权限验证
    sysAuth := sys.Use(middleware.NewAuth())

    indexController := new(admin.Index)
    sysAuth.Get("/", indexController.Index).Name("admin.index")
    sysAuth.Get("/logout", authController.Logout).Name("admin.logout")

    // 信息设置
    profileController := new(admin.Profile)
    sysAuth.Get("/profile", profileController.Index).Name("admin.profile.index")
    sysAuth.Post("/profile", profileController.Save).Name("admin.profile.save")
    sysAuth.Get("/profile/password", profileController.Password).Name("admin.profile.password")
    sysAuth.Post("/profile/password", profileController.PasswordSave).Name("admin.profile.password-save")

    // 上传
    uploadController := new(admin.Upload)
    sysAuth.Post("/upload/file", uploadController.File).Name("admin.upload.file")
    sysAuth.Post("/upload/image", uploadController.Image).Name("admin.upload.image")

    // 文章
    artController := new(admin.Art)
    sysAuth.Get("/art", artController.Index).Name("admin.art.index")
    sysAuth.Get("/art/add", artController.Add).Name("admin.art.add")
    sysAuth.Post("/art/add", artController.AddSave).Name("admin.art.add-save")
    sysAuth.Get("/art/:id/edit", artController.Edit).Name("admin.art.edit")
    sysAuth.Post("/art/:id/edit", artController.EditSave).Name("admin.art.edit-save")
    sysAuth.Post("/art/:id/delete", artController.Delete).Name("admin.art.delete")

    // admin 检测
    sysAdmin := sysAuth.Use(middleware.NewAdminCheck())

    // 设置
    settingController := new(admin.Setting)
    sysAdmin.Get("/setting", settingController.Index).Name("admin.setting")
    sysAdmin.Post("/setting", settingController.Save).Name("admin.setting-save")

    // 分类
    cateController := new(admin.Cate)
    sysAdmin.Get("/cate", cateController.Index).Name("admin.cate.index")
    sysAdmin.Get("/cate/add", cateController.Add).Name("admin.cate.add")
    sysAdmin.Post("/cate/add", cateController.AddSave).Name("admin.cate.add-save")
    sysAdmin.Get("/cate/:id/edit", cateController.Edit).Name("admin.cate.edit")
    sysAdmin.Post("/cate/:id/edit", cateController.EditSave).Name("admin.cate.edit-save")
    sysAdmin.Post("/cate/:id/delete", cateController.Delete).Name("admin.cate.delete")

    // 单页
    pageController := new(admin.Page)
    sysAdmin.Get("/page", pageController.Index).Name("admin.page.index")
    sysAdmin.Get("/page/add", pageController.Add).Name("admin.page.add")
    sysAdmin.Post("/page/add", pageController.AddSave).Name("admin.page.add-save")
    sysAdmin.Get("/page/:id/edit", pageController.Edit).Name("admin.page.edit")
    sysAdmin.Post("/page/:id/edit", pageController.EditSave).Name("admin.page.edit-save")
    sysAdmin.Post("/page/:id/delete", pageController.Delete).Name("admin.page.delete")

    // 标签
    tagController := new(admin.Tag)
    sysAdmin.Get("/tag", tagController.Index).Name("admin.tag.index")
    sysAdmin.Get("/tag/add", tagController.Add).Name("admin.tag.add")
    sysAdmin.Post("/tag/add", tagController.AddSave).Name("admin.tag.add-save")
    sysAdmin.Get("/tag/:id/edit", tagController.Edit).Name("admin.tag.edit")
    sysAdmin.Post("/tag/:id/edit", tagController.EditSave).Name("admin.tag.edit-save")
    sysAdmin.Post("/tag/:id/delete", tagController.Delete).Name("admin.tag.delete")

    // 账号
    userController := new(admin.User)
    sysAdmin.Get("/user", userController.Index).Name("admin.user.index")
    sysAdmin.Get("/user/add", userController.Add).Name("admin.user.add")
    sysAdmin.Post("/user/add", userController.AddSave).Name("admin.user.add-save")
    sysAdmin.Get("/user/:id/edit", userController.Edit).Name("admin.user.edit")
    sysAdmin.Post("/user/:id/edit", userController.EditSave).Name("admin.user.edit-save")
    sysAdmin.Post("/user/:id/delete", userController.Delete).Name("admin.user.delete")

    // 附件
    attachController := new(admin.Attach)
    sysAdmin.Get("/attach", attachController.Index).Name("admin.attach.index")
    sysAdmin.Post("/attach/delete/:id", attachController.Delete).Name("admin.attach.delete")
    sysAdmin.Get("/attach/download/:id", attachController.Download).Name("admin.attach.download")

}


