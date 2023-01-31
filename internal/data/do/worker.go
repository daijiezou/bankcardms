package do

type Worker struct {
	WorkerId   string `json:"worker_id" xorm:"not null pk default '' comment('身份证号') VARCHAR(64)"`
	Name       string `json:"name" xorm:"not null default '' comment('姓名') VARCHAR(128)"`
	Address    string `json:"address" xorm:"not null default '' comment('住址') VARCHAR(256)"`
	Sex        int    `json:"sex" xorm:"default 1 comment('1:男性；2：女性') INT"`
	CreateTime int64  `json:"create_time" xorm:"not null default 0 BIGINT"`
	UpdateTime int64  `json:"update_time" xorm:"not null default 0 BIGINT"`
	DeleteTime int    `json:"delete_time" xorm:"not null default 0 INT"`
}

func (m *Worker) TableName() string {
	return "worker"
}
