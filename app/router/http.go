package router

import (
    "github.com/gofiber/fiber/v2"

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
    // 静态文件
    app.Static("/static/", "./public/static")
    app.Static("/favicon.ico", "./public/favicon.ico")
}

// CMS
func HttpCms(app *fiber.App) {
    indexController := new(cms.Index)
    app.Get("/", indexController.Index)

    cateController := new(cms.Cate)
    app.Get("/cate/:slug", cateController.Index)

    viewController := new(cms.View)
    app.Get("/v/:id", viewController.Index)
}

// 管理员
func HttpAdmin(app *fiber.App) {
    authController := new(admin.Auth)
    app.Get("/sys/captcha", authController.Captcha)
    app.Get("/sys/login", authController.Login)
    app.Post("/sys/login", authController.LoginCheck)

    // 权限验证
    sys := app.Group("/sys", middleware.NewAuth())

    indexController := new(admin.Index)
    sys.Get("/", indexController.Index)
    sys.Get("/logout", authController.Logout)

    // 设置
    settingController := new(admin.Setting)
    sys.Get("/setting", settingController.Index)
    sys.Post("/setting/save", settingController.Save)

    // 信息设置
    profileController := new(admin.Profile)
    sys.Get("/profile", profileController.Index)
    sys.Post("/profile/save", profileController.Save)

    // 分类
    cateController := new(admin.Cate)
    sys.Get("/cate", cateController.Index)
    sys.Get("/cate/add", cateController.Add)
    sys.Post("/cate/add", cateController.AddSave)
    sys.Get("/cate/:id/edit", cateController.Edit)
    sys.Post("/cate/:id/edit", cateController.EditSave)
    sys.Post("/cate/:id/delete", cateController.Delete)

    // 文章
    artController := new(admin.Art)
    sys.Get("/art", artController.Index)
    sys.Get("/art/add", artController.Add)
    sys.Post("/art/add", artController.AddSave)
    sys.Get("/art/:id/edit", artController.Edit)
    sys.Post("/art/:id/edit", artController.EditSave)
    sys.Post("/art/:id/delete", artController.Delete)

    // 标签
    tagController := new(admin.Tag)
    sys.Get("/tag", tagController.Index)
    sys.Get("/tag/:id/edit", tagController.Edit)
    sys.Post("/tag/:id/edit", tagController.EditSave)
    sys.Post("/tag/:id/delete", tagController.Delete)

    // 账号
    userController := new(admin.User)
    sys.Get("/user", userController.Index)
    sys.Get("/user/add", userController.Add)
    sys.Post("/user/add", userController.AddSave)
    sys.Get("/user/:id/edit", userController.Edit)
    sys.Post("/user/:id/edit", userController.EditSave)
    sys.Post("/user/:id/delete", userController.Delete)

    // 上传
    uploadController := new(admin.Upload)
    sys.Post("/upload/file", uploadController.File)

}


