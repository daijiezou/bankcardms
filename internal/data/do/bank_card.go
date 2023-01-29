package do

type BankCard struct {
	CardId        int    `json:"card_id" xorm:"not null pk autoincr INT"`
	CardNumber    string `json:"card_number" xorm:"not null default '' unique VARCHAR(64)"`
	CardName      string `json:"card_name" xorm:"not null default '' VARCHAR(64)"`
	CardOwner     int    `json:"card_owner" xorm:"not null default 0 comment('银行卡的所有人，关联worker_id') INT"`
	CardImagePath string `json:"card_image_path" xorm:"VARCHAR(256)"`
}

func (m *BankCard) TableName() string {
	return "bank_card"
}
