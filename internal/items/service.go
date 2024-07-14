package items

import "github.com/SolBaa/receipt-processor/internal/domain"

type ItemService interface {
	GetItems() ([]domain.Items, error)
}

type items struct {
	items []domain.Items
}

func NewItemService() ItemService {
	return &items{}
}

func (i *items) GetItems() ([]domain.Items, error) {
	// Implement the logic to get the items

	item := domain.Items{
		ShortDescription: "Item 1",
		Price:            "10.00",
	}
	i.items = append(i.items, item)

	return i.items, nil
}
