package handler

import (
	"net/http"

	"github.com/SolBaa/receipt-processor/internal/receipt"
)

func GetReceipt(w http.ResponseWriter, r *http.Request) {
	// Get receipt
	msj := receipt.GetReceipt()
	w.Write([]byte(msj))
}

func ProcessReceipt(w http.ResponseWriter, r *http.Request) {
	// Process receipt
	receipt.ProcessReceipt()
}

func GetReceiptPoints(w http.ResponseWriter, r *http.Request) {
	// Get receipt points
	msj := receipt.GetReceiptPoints()
	w.Write([]byte(msj))
}
