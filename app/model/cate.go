package model

// 分类
type Cate struct {
    Id      int64  `xorm:"int(11)"`
    Pid     int64  `xorm:"int(11)"`
    Name    string `xorm:"varchar(50)"`
    Slug    string `xorm:"varchar(50)"`
    Desc    string `xorm:"varchar(200)"`
    Sort    int    `xorm:"int(5)"`
    Status  int    `xorm:"tinyint(1)"`
    AddTime int64  `xorm:"created"`
    AddIp   string `xorm:"varchar(50)"`
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
