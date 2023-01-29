package do

type BankCardStatement struct {
	CardNumber   string `json:"card_number" xorm:"not null default '' comment('银行卡号') VARCHAR(64)"`
	TradeTime    int64  `json:"trade_time" xorm:"not null default 0 comment('交易时间') BIGINT"`
	TradeAmount  int    `json:"trade_amount" xorm:"not null default 0 comment('交易金额') INT"`
	TradeRemarks string `json:"trade_remarks" xorm:"not null default '' comment('交易备注') VARCHAR(1024)"`
	CreateTime   int64  `json:"create_time" xorm:"default 0 BIGINT"`
	UpdateTime   int64  `json:"update_time" xorm:"default 0 BIGINT"`
	DeleteTime   int64  `json:"delete_time" xorm:"default 0 BIGINT"`
}

func (m *BankCardStatement) TableName() string {
	return "bank_card_statement"
}
