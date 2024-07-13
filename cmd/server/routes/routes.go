package routes

import (
	"github.com/SolBaa/receipt-processor/cmd/server/handler"
	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	r.Route("/receipts", func(r chi.Router) {
		r.Get("/", handler.GetReceipt)
		r.Get("/{id}/points", handler.GetReceiptPoints)
		r.Post("/process", handler.ProcessReceipt)
	})
}
