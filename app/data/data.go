package data

import (
    "github.com/spf13/cast"
    jsoniter "github.com/json-iterator/go"

    "github.com/deatil/doak-cms/pkg/db"
    "github.com/deatil/doak-cms/pkg/redis"

    "github.com/deatil/doak-cms/app/model"
)

// 设置数据列表
func GetSettings() map[string]string {
    settings := make(map[string]string)

    // 缓存获取数据
    res, _ := redis.Storage().Get("settings")
    if len(res) > 0 {
        jsoniter.Unmarshal(res, &settings)
    } else {
        configs := make([]model.Config, 0)
        db.Engine().Find(&configs)

        if len(configs) > 0 {
            for _, v := range configs {
                settings[v.Key] = v.Value
            }
        }

        // 永久缓存
        data, _ := jsoniter.Marshal(settings)
        redis.Storage().Set("settings", data, 0)
    }

    return settings
}

// 设置数据单个
func GetSetting(key string) string {
    data := GetSettings()
    if info, ok := data[key]; ok {
        return info
    }

    return ""
}

// =================

// 分类列表
func GetCateList(where string, orderby string, limit int) []model.Cate {
    if limit < 0 {
        limit = 1
    }

    // 列表
    list := make([]model.Cate, 0)
    db.Engine().
        Limit(limit, 0).
        Where(where).
        OrderBy(orderby).
        Find(&list)

    return list
}

// 分类详情
func GetCateInfo(id int) model.Cate {
    cateId := cast.ToInt64(id)

    // 分类信息
    var data model.Cate
    _, err := db.Engine().
        Where("id = ?", cateId).
        Get(&data)
    if err != nil || data.Id == 0 {
        return data
    }

    return data
}

// 分类详情
func GetCateInfoWithSlug(slug string) model.Cate {
    // 分类信息
    var data model.Cate
    _, err := db.Engine().
        Where("slug = ?", slug).
        Get(&data)
    if err != nil || data.Id == 0 {
        return data
    }

    return data
}

// =================

// 文章列表
func GetArtList(where string, orderby string, limit int) []model.ArtCatename {
    if limit < 0 {
        limit = 1
    }

    // 列表
    list := make([]model.Art, 0)

    db.Engine().
        Limit(limit, 0).
        Where(where).
        OrderBy(orderby).
        Find(&list)

    // 分类列表
    cates := make([]model.Cate, 0)
    db.Engine().
        Desc("sort").
        Asc("id").
        Find(&cates)

    newCates := make(map[int64]model.Cate)
    if len(cates) > 0 {
        for _, cv := range cates {
            newCates[cv.Id] = cv
        }
    }

    newArts := make([]model.ArtCatename, 0)
    if len(list) > 0 {
        for _, v := range list {
            cateName := ""
            cateSlug := ""
            if newCate, ok := newCates[v.CateId]; ok {
                cateName = newCate.Name
                cateSlug = newCate.Slug
            }

            newArts = append(newArts, model.ArtCatename{
                Art: v,
                CateName: cateName,
                CateSlug: cateSlug,
            })
        }
    }

    return newArts
}

// 文章详情
func GetArtInfo(id int) model.Art {
    artId := cast.ToInt64(id)

    // 分类信息
    var data model.Art
    _, err := db.Engine().
        Where("id = ?", artId).
        Get(&data)
    if err != nil || data.Id == 0 {
        return data
    }

    return data
}

// =================

// 单页详情
func GetPageInfo(slug string) model.Page {
    // 分类信息
    var data model.Page
    _, err := db.Engine().
        Where("slug = ?", slug).
        Get(&data)
    if err != nil || data.Id == 0 {
        return data
    }

    return data
}

// =================

// 标签列表
func GetTagList(where string, orderby string, limit int) []model.Tag {
    if limit < 0 {
        limit = 1
    }

    // 列表
    list := make([]model.Tag, 0)
    db.Engine().
        Limit(limit, 0).
        Where(where).
        OrderBy(orderby).
        Find(&list)

    return list
}
