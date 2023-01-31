package do

type BankCard struct {
	CardId        string `json:"card_id" xorm:"not null pk default '银行卡号' comment('银行卡号') VARCHAR(64)"`
	CardImagePath string `json:"card_image_path" xorm:"not null VARCHAR(256)"`
	CardOwner     int    `json:"card_owner" xorm:"not null default 0 comment('银行卡的所有人，关联worker_id') INT"`
	CardName      string `json:"card_name" xorm:"not null default '' VARCHAR(64)"`
}

func (m *BankCard) TableName() string {
	return "bank_card"
}
