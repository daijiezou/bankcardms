package commonreq

type CommonListReq struct {
	Filter   string `json:"filter" form:"filter"`
	PageNum  int    `json:"page_num" form:"page_num" binding:"min=1"`
	PageSize int    `json:"page_size" form:"page_size" binding:"min=1,max=100"`
}
