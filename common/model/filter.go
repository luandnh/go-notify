package model

type GeneralFilter struct {
	Page     []byte `json:"page"`
	PageSize int    `json:"page_size"`
}
