package handler

import (
	"encoding/json"
	"net/http"

	"github.com/SolBaa/receipt-processor/internal/items"
)

type HandlerItems struct {
	ItemService items.ItemService
}

func NewItemsHandler(is items.ItemService) *HandlerItems {
	return &HandlerItems{
		ItemService: is,
	}
}

func (h *HandlerItems) GetItems(w http.ResponseWriter, r *http.Request) {
	// Implement the logic to get the items
	items, err := h.ItemService.GetItems()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(items)
}
