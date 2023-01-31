package vo

type BankCardList struct {
	ListCount int        `json:"list_count"`
	BankCards []BankCard `json:"workers"`
}

type BankCard struct {
	CardId                 string `json:"card_id" xorm:"not null pk default '银行卡号' comment('银行卡号') VARCHAR(64)"`
	CardImagePath          string `json:"card_image_path" xorm:"not null VARCHAR(256)"`
	CardOwner              string `json:"card_owner" xorm:"not null default '' comment('银行卡的所有人，关联worker_id') VARCHAR(64)"`
	BankName               string `json:"bank_name" xorm:"not null default '' VARCHAR(64)"`
	Remarks                string `json:"remarks" xorm:"not null TEXT"`
	CurrentYearTotalIncome int    `json:"current_year_total_income"`
	CreateTime             int64  `json:"create_time" xorm:"not null default 0 BIGINT"`
	UpdateTime             int64  `json:"update_time" xorm:"not null default 0 BIGINT"`
}
