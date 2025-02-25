package routes

import (
	"github.com/SolBaa/receipt-processor/cmd/server/handler"
	"github.com/SolBaa/receipt-processor/internal/items"
	"github.com/SolBaa/receipt-processor/internal/receipt"
	"github.com/go-chi/chi/v5"
)

func Routes(r chi.Router) {
	// Create a new receipt service
	receiptService := receipt.NewReceiptService()
	itemsService := items.NewItemService()
	// Create a new handler
	h := handler.NewHandler(receiptService)
	hItems := handler.NewItemsHandler(itemsService)

	r.Route("/receipts", func(r chi.Router) {
		r.Get("/", h.GetReceipt)
		r.Get("/{id}/points", h.GetReceiptPoints)
		r.Post("/process", h.ProcessReceipt)
	})

	r.Route("/items", func(r chi.Router) {
		r.Get("/", hItems.GetItems)
	})
}
