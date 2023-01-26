package member

import (
	_ "github.com/go-sql-driver/mysql"
	"xorm.io/xorm"
)

// 会员等级结构
type Level struct {
	value int
	name  string
}

// 联系人结构
type Contact struct {
	name      string
	telephone string
	address   string
}

// 会员结构
type Member struct {
	NickName    string
	RealName    string
	Avatar      string
	Gender      int
	Telephone   string
	Level       int
	Address     string
	Remark      string
	CardId      string
	PrimaryFrom int
	Age         int
	Contacts    []*Contact
}
