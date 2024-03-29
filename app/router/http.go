package router

import (
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/fiber/v2/middleware/filesystem"

    "github.com/deatil/doak-cms/pkg/config"

    "github.com/deatil/doak-cms/resources"
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
    app.Static("/upload/", "./public/upload")

    // 静态文件
    if config.IsUseEmbed {
        app.Use("/static", filesystem.New(filesystem.Config{
            Root: resources.StaticAssets(),
            PathPrefix: "static",
            Browse: true,
        }))
    } else {
        app.Static("/static/", "./resources/static")
    }
}

// CMS
func HttpCms(app *fiber.App) {
    // 网站是否开启检测
    siteopenCheckMiddleware := middleware.NewSiteopenCheck()

    indexController := new(cms.Index)
    app.Get("/", siteopenCheckMiddleware, indexController.Index)

    cateController := new(cms.Cate)
    app.Get("/c/:slug", siteopenCheckMiddleware, cateController.Index)

    viewController := new(cms.View)
    app.Get("/a/:id", siteopenCheckMiddleware, viewController.Index)

    tagController := new(cms.Tag)
    app.Get("/tag/:tag", siteopenCheckMiddleware, tagController.Index)

    pageController := new(cms.Page)
    app.Get("/p/:name", siteopenCheckMiddleware, pageController.Index)
}

// 管理员
func HttpAdmin(app *fiber.App) {
    cfg := config.Section("app")
    prefix := cfg.Key("router-prefix").String()

    // 后台前缀
    sys := app.Group("/" + prefix)

    // 不需要授权路由
    authController := new(admin.Auth)
    sys.Get("/captcha", authController.Captcha)
    sys.Get("/login", authController.Login)
    sys.Post("/login", authController.LoginCheck)

    // 权限验证
    sysAuth := sys.Use(middleware.NewAuth())

    indexController := new(admin.Index)
    sysAuth.Get("/", indexController.Index)
    sysAuth.Get("/logout", authController.Logout)

    // 信息设置
    profileController := new(admin.Profile)
    sysAuth.Get("/profile", profileController.Index)
    sysAuth.Post("/profile", profileController.Save)
    sysAuth.Get("/profile/password", profileController.Password)
    sysAuth.Post("/profile/password", profileController.PasswordSave)

    // 上传
    uploadController := new(admin.Upload)
    sysAuth.Post("/upload/file", uploadController.File)
    sysAuth.Post("/upload/image", uploadController.Image)

    // 文章
    artController := new(admin.Art)
    sysAuth.Get("/art", artController.Index)
    sysAuth.Get("/art/add", artController.Add)
    sysAuth.Post("/art/add", artController.AddSave)
    sysAuth.Get("/art/:id/edit", artController.Edit)
    sysAuth.Post("/art/:id/edit", artController.EditSave)
    sysAuth.Post("/art/:id/delete", artController.Delete)

    // admin 检测
    sysAdmin := sysAuth.Use(middleware.NewAdminCheck())

    // 设置
    settingController := new(admin.Setting)
    sysAdmin.Get("/setting", settingController.Index)
    sysAdmin.Post("/setting", settingController.Save)

    // 分类
    cateController := new(admin.Cate)
    sysAdmin.Get("/cate", cateController.Index)
    sysAdmin.Get("/cate/add", cateController.Add)
    sysAdmin.Post("/cate/add", cateController.AddSave)
    sysAdmin.Get("/cate/:id/edit", cateController.Edit)
    sysAdmin.Post("/cate/:id/edit", cateController.EditSave)
    sysAdmin.Post("/cate/:id/delete", cateController.Delete)

    // 单页
    pageController := new(admin.Page)
    sysAdmin.Get("/page", pageController.Index)
    sysAdmin.Get("/page/add", pageController.Add)
    sysAdmin.Post("/page/add", pageController.AddSave)
    sysAdmin.Get("/page/:id/edit", pageController.Edit)
    sysAdmin.Post("/page/:id/edit", pageController.EditSave)
    sysAdmin.Post("/page/:id/delete", pageController.Delete)

    // 标签
    tagController := new(admin.Tag)
    sysAdmin.Get("/tag", tagController.Index)
    sysAdmin.Get("/tag/add", tagController.Add)
    sysAdmin.Post("/tag/add", tagController.AddSave)
    sysAdmin.Get("/tag/:id/edit", tagController.Edit)
    sysAdmin.Post("/tag/:id/edit", tagController.EditSave)
    sysAdmin.Post("/tag/:id/delete", tagController.Delete)

    // 账号
    userController := new(admin.User)
    sysAdmin.Get("/user", userController.Index)
    sysAdmin.Get("/user/add", userController.Add)
    sysAdmin.Post("/user/add", userController.AddSave)
    sysAdmin.Get("/user/:id/edit", userController.Edit)
    sysAdmin.Post("/user/:id/edit", userController.EditSave)
    sysAdmin.Post("/user/:id/delete", userController.Delete)

    // 附件
    attachController := new(admin.Attach)
    sysAdmin.Get("/attach", attachController.Index)
    sysAdmin.Post("/attach/delete/:id", attachController.Delete)
    sysAdmin.Get("/attach/download/:id", attachController.Download)

}


