package vo

import "BankCardMS/internal/data/do"

type BankCardList struct {
	ListCount int           `json:"list_count"`
	BankCards []do.BankCard `json:"workers"`
}
