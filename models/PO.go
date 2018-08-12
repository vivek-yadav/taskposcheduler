package models

type PO struct {
	PoId     int `json:"po_id"`
	ItemId   int `json:"item_id"`
	Quantity int `json:"quantity"`
}

type POList []*PO
