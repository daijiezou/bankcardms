package do

type BankCard struct {
	CardId        string `json:"card_id" xorm:"not null pk default '银行卡号' comment('银行卡号') VARCHAR(64)"`
	CardImagePath string `json:"card_image_path" xorm:"not null VARCHAR(256)"`
	CardOwner     string `json:"card_owner" xorm:"not null default '' comment('银行卡的所有人，关联worker_id') VARCHAR(64)"`
	BankName      string `json:"bank_name" xorm:"not null default '' VARCHAR(64)"`
	Remarks       string `json:"remarks" xorm:"not null TEXT"`
	CreateTime    int64  `json:"create_time" xorm:"not null default 0 BIGINT"`
	UpdateTime    int64  `json:"update_time" xorm:"not null default 0 BIGINT"`
	DeleteTime    int64  `json:"delete_time" xorm:"deleted not null default 0 BIGINT"`
}

func (m *BankCard) TableName() string {
	return "bank_card"
}

type BankCardList struct {
	ListCount int              `json:"list_count"`
	BankCards []BankCardDetail `json:"bank_cards"`
}

type BankCardDetail struct {
	CardId        string `json:"card_id" xorm:"not null pk default '银行卡号' comment('银行卡号') VARCHAR(64)"`
	CardImagePath string `json:"card_image_path" xorm:"not null VARCHAR(256)"`
	CardOwner     string `json:"card_owner" xorm:"not null default '' comment('银行卡的所有人，关联worker_id') VARCHAR(64)"`
	Name          string `json:"name" xorm:"not null default"`
	BankName      string `json:"bank_name" xorm:"not null default '' VARCHAR(64)"`
	Remarks       string `json:"remarks" xorm:"not null TEXT"`
	CreateTime    int64  `json:"create_time" xorm:"not null default 0 BIGINT"`
	UpdateTime    int64  `json:"update_time" xorm:"not null default 0 BIGINT"`
	DeleteTime    int64  `json:"delete_time" xorm:"not null default 0 BIGINT"`
}
