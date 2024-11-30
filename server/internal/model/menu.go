package model

// Menu 菜单管理
type Menu struct {
	Model `gorm:"embedded"` // embed id and time

	ParentID   int    `gorm:"column:parent_id;type:int(11);NOT NULL" json:"parentID"`
	Name       string `gorm:"column:name;type:text;NOT NULL" json:"name"`
	Type       string `gorm:"column:type;type:text" json:"type"`
	Path       string `gorm:"column:path;type:text;NOT NULL" json:"path"`
	Component  string `gorm:"column:component;type:text;NOT NULL" json:"component"`
	Perm       string `gorm:"column:perm;type:text" json:"perm"`
	Sort       int    `gorm:"column:sort;type:int(11);NOT NULL" json:"sort"`
	Visible    int    `gorm:"column:visible;type:int(11);NOT NULL" json:"visible"`
	Icon       string `gorm:"column:icon;type:text;NOT NULL" json:"icon"`
	Redirect   string `gorm:"column:redirect;type:text;NOT NULL" json:"redirect"`
	AlwaysShow int    `gorm:"column:always_show;type:int(11);NOT NULL" json:"alwaysShow"`
	KeepAlive  int    `gorm:"column:keep_alive;type:int(11);NOT NULL" json:"keepAlive"`
	Params     string `gorm:"column:params;type:text" json:"params"`
}

// TableName table name
func (m *Menu) TableName() string {
	return "t_menu"
}

type MenuMeta struct {
	Title      string  `json:"title"`      // 标题
	Icon       string  `json:"icon"`       // 图标
	Hidden     bool    `json:"hidden"`     // 隐藏
	AlwaysShow bool    `json:"alwaysShow"` // 始终显示
	Params     *string `json:"params"`     // 路由参数
}

type MenuItem struct {
	Path      string     `json:"path"`      // 路由路径
	Name      string     `json:"name"`      // 路由名称
	Component string     `json:"component"` // 组件路径
	Redirect  string     `json:"redirect"`  // 跳转路径
	Meta      MenuMeta   `json:"meta"`      // meta
	Children  []Children `json:"children"`  // 子级
}

type ChildrenMeta struct {
	Title      string  `json:"title"`      // 标题
	Icon       string  `json:"icon"`       // 图标
	Hidden     bool    `json:"hidden"`     // 隐藏
	AlwaysShow bool    `json:"alwaysShow"` // 始终显示
	Params     *string `json:"params"`     // 路由参数
	KeepAlive  bool    `json:"keepAlive"`  // 始终显示
}

type Children struct {
	Path      string       `json:"path"`      // 路由路径
	Name      string       `json:"name"`      // 路由名称
	Component string       `json:"component"` // 组件路径
	Meta      ChildrenMeta `json:"meta"`      // meta
}
