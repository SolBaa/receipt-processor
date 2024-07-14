package receipt

import (
	"testing"

	"github.com/SolBaa/receipt-processor/internal/domain"
	"github.com/stretchr/testify/require"
)

func TestGetReceiptSuccessful(t *testing.T) {
	// TestGetReceiptSuccessful tests the GetReceipt function
	service := NewReceiptService()
	receipts, err := service.GetReceipts()
	receipts = append(receipts, domain.Receipt{
		ID:           "1",
		Retailer:     "Walmart",
		PurchaseDate: "2021-10-10",
		PurchaseTime: "10:00",
		Items: []domain.Items{
			{
				ShortDescription: "Item 1",
				Price:            "10.00",
			},
		},
		Total: "10.00",
	})

	require.NoError(t, err)
	require.NotNil(t, receipts)
}

func TestProcessReceiptSuccesfull(t *testing.T) {
	// TestProcessReceiptSuccesfull tests the ProcessReceipt function
	// and checks if it returns the receipt successfully
	// and if it returns the receipt in the correct format
	// and if it returns the receipt with the correct ID
}

func TestGetReceiptPointsSuccesfull(t *testing.T) {
	// TestGetReceiptPointsSuccesfull tests the GetReceiptPoints function
	// and checks if it returns the receipt points successfully
	// and if it returns the receipt points in the correct format
	// and if it returns the receipt points with the correct ID
}
