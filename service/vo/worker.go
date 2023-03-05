package vo

import "BankCardMS/data/do"

type WorkerDetail struct {
	WorkerInfo             *do.Worker       `json:"worker_info"`
	BankCardList           *do.BankCardList `json:"bank_card_list"`
	CurrentYearTotalIncome int64            `json:"current_year_total_income"`
}
