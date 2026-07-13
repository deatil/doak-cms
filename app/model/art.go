package model

// 文章
type Art struct {
    Id          int64  `xorm:"int(10)" json:"id"`
    Uuid        string `xorm:"char(32)" json:"uuid"`
    CateId      int64  `xorm:"int(10)" json:"cate_id"`
    UserId      int64  `xorm:"int(10)" json:"user_id"`
    Cover       string `xorm:"char(32)" json:"cover"`
    Title       string `xorm:"varchar(150)" json:"title"`
    Keywords    string `xorm:"varchar(100)" json:"keywords"`
    Description string `xorm:"varchar(200)" json:"description"`
    Content     string `xorm:"longtext" json:"content"`
    Tags        string `xorm:"varchar(250)" json:"tags"`
    From        string `xorm:"varchar(200)" json:"from"`
    Views       int64  `xorm:"bigint(20)" json:"views"`
    IsTop       int    `xorm:"tinyint(1)" json:"is_top"`
    Status      int    `xorm:"tinyint(1)" json:"status"`
    AddTime     int64  `xorm:"created" json:"add_time"`
    AddIp       string `xorm:"varchar(50)" json:"add_ip"`
}

// 文章带分类
type ArtCate struct {
    Art  `xorm:"extends""`
    Cate `xorm:"extends""`
}

// 文章带分类
type ArtCatename struct {
    Art
    CateName string
    CateSlug string
}
