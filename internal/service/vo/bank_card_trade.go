package vo

import "BankCardMS/internal/data/do"

type BankCardTradeList struct {
	ListCount      int                `json:"list_count"`
	BankCardTrades []do.BankCardTrade `json:"workers"`
}
