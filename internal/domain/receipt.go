package domain

type Receipt struct {
	ID           string  `json:"-"`
	Retailer     string  `json:"retailer"`
	PurchaseDate string  `json:"purchaseDate"`
	PurchaseTime string  `json:"purchaseTime"`
	Items        []Items `json:"items"`
	Total        string  `json:"total"`
}

type Items struct {
	ShortDescription string `json:"shortDescription"`
	Price            string `json:"price"`
}

type ReceiptPoints struct {
}

type ReceiptPointsResponse struct {
	Points int `json:"points"`
}

type ReceiptProcessResponse struct {
	ID string `json:"id"`
}
