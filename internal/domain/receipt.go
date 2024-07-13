package domain

type Receipt struct {
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
