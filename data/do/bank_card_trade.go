package do

type BankCardTrade struct {
	TradeId     string `json:"trade_id" xorm:"not null pk default '' unique VARCHAR(64)"`
	CardId      string `json:"card_id" xorm:"not null default '' comment('银行卡号') VARCHAR(64)"`
	TradeTime   int64  `json:"trade_time" xorm:"not null default 0 comment('交易时间') BIGINT"`
	TradeAmount int    `json:"trade_amount" xorm:"not null default 0 comment('交易金额') INT"`
	Remarks     string `json:"remarks" xorm:"not null comment('交易备注') TEXT"`
	CreateTime  int64  `json:"create_time" xorm:"default 0 BIGINT"`
	UpdateTime  int64  `json:"update_time" xorm:"default 0 BIGINT"`
	DeleteTime  int64  `json:"delete_time" xorm:"default 0 not null BIGINT"`
}

func (m *BankCardTrade) TableName() string {
	return "bank_card_trade"
}

type BankCardTradeList struct {
	ListCount      int             `json:"list_count"`
	BankCardTrades []BankCardTrade `json:"bank_card_trades"`
}
