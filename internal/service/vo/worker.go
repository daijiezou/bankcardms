package vo

import "BankCardMS/internal/data/do"

type WorkerList struct {
	ListCount int         `json:"list_count"`
	Workers   []do.Worker `json:"workers"`
}
