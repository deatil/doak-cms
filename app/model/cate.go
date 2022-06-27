package model

// 分类
type Cate struct {
    Id      int64  `xorm:"int(10)" json:"id"`
    Pid     int64  `xorm:"int(10)" json:"pid"`
    Name    string `xorm:"varchar(50)" json:"name"`
    Slug    string `xorm:"varchar(50)" json:"slug"`
    Desc    string `xorm:"varchar(200)" json:"desc"`
    Sort    int    `xorm:"int(5)" json:"sort"`
    Tpl     string `xorm:"varchar(200)" json:"tpl"`
    ViewTpl string `xorm:"varchar(200)" json:"view_tpl"`
    Status  int    `xorm:"tinyint(1)" json:"status"`
    AddTime int64  `xorm:"created" json:"add_time"`
    AddIp   string `xorm:"varchar(50)" json:"add_ip"`
}

// 分类树
type CateTree struct {
    Cate
    Children []CateTree
}

// 转换为树形结构
func ToCateTree(data []Cate, pid int64) []CateTree {
    var list []CateTree

    for _, item := range data {
        if item.Pid == pid {
            newItem := CateTree{
                Cate: item,
                Children: ToCateTree(data, item.Id),
            }

            list = append(list, newItem)
        }
    }

    return list
}
