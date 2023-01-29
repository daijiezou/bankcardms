package do

type User struct {
	UserId      int64  `json:"user_id" xorm:"not null pk autoincr BIGINT"`
	UserName    string `json:"user_name" xorm:"unique VARCHAR(256)"`
	Password    string `json:"password" xorm:"VARCHAR(256)"`
	Salt        string `json:"salt" xorm:"VARCHAR(256)"`
	DisplayName string `json:"display_name" xorm:"VARCHAR(256)"`
	Phone       string `json:"phone" xorm:"VARCHAR(256)"`
	Email       string `json:"email" xorm:"not null VARCHAR(256)"`
	UpdateTime  int64  `json:"update_time" xorm:"default 0 BIGINT"`
	CreateTime  int64  `json:"create_time" xorm:"default 0 BIGINT"`
	DeleteTime  int    `json:"delete_time" xorm:"deleted not null default 0 INT"`
}

func (m *User) TableName() string {
	return "user"
}
