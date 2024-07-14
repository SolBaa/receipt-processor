package receipt

import (
	"fmt"
	"math"
	"strconv"
	"sync"
	"time"
	"unicode"

	"github.com/SolBaa/receipt-processor/internal/domain"
	"github.com/google/uuid"
)

// var (
// 	receipts []domain.Receipt
// 	mu       sync.Mutex // Mutex para asegurar la seguridad de concurrencia
// )

type ReceiptService interface {
	GetReceipts() ([]domain.Receipt, error)
	ProcessReceipt(domain.Receipt) (domain.ReceiptProcessResponse, error)
	GetReceiptPoints(string) (domain.ReceiptPointsResponse, error)
}

type receiptService struct {
	receipts []domain.Receipt
	mu       sync.Mutex // Mutex para asegurar la seguridad de concurrencia
}

func NewReceiptService() ReceiptService {
	return &receiptService{}
}

func (s *receiptService) GetReceipts() ([]domain.Receipt, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if len(s.receipts) == 0 {
		return nil, nil
	}
	return s.receipts, nil
}

func (s *receiptService) ProcessReceipt(data domain.Receipt) (domain.ReceiptProcessResponse, error) {
	// Generar un UUID aleatorio
	id := uuid.New().String()
	res := domain.ReceiptProcessResponse{
		ID: id,
	}

	// Crear y guardar el recibo en memoria
	receipt := domain.Receipt{
		ID:           id,
		Retailer:     data.Retailer,
		PurchaseDate: data.PurchaseDate,
		PurchaseTime: data.PurchaseTime,
		Items:        data.Items,
		Total:        data.Total,
	}

	s.mu.Lock()
	s.receipts = append(s.receipts, receipt)
	s.mu.Unlock()

	return res, nil
}

// GetReceiptPoints calcula los puntos de un recibo dado su ID
func (s *receiptService) GetReceiptPoints(id string) (domain.ReceiptPointsResponse, error) {
	fmt.Println(id)
	var res domain.ReceiptPointsResponse
	var points int

	for _, receipt := range s.receipts {
		if receipt.ID == id {
			// Calcular puntos
			fmt.Println(receipt)

			// One point for every alphanumeric character in the retailer name
			for _, char := range receipt.Retailer {
				if unicode.IsLetter(char) || unicode.IsDigit(char) {
					points++
				}
			}

			// 50 points if the total is a round dollar amount with no cents.
			total, err := strconv.ParseFloat(receipt.Total, 64)
			if err != nil {
				return res, err
			}
			if math.Floor(total) == total {
				points += 50
			}

			// 25 points if the total is a multiple of 0.25
			if total/0.25 == math.Floor(total/0.25) {
				points += 25
			}

			// 5 points for every two items on the receipt
			points += len(receipt.Items) / 2 * 5

			// If the trimmed length of the item description is a multiple of 3, multiply the price by 0.2 and round up to the nearest integer.
			for _, item := range receipt.Items {
				trimmed := len(item.ShortDescription)
				if trimmed%3 == 0 {
					price, err := strconv.ParseFloat(item.Price, 64)
					if err != nil {
						return res, err
					}
					points += int(math.Ceil(price * 0.2))
				}
			}

			// 6 points if the day in the purchase date is odd.
			purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
			if err != nil {
				return res, err
			}
			if purchaseDate.Day()%2 != 0 {
				points += 6
			}

			// 10 points if the time of purchase is after 2:00pm and before 4:00pm.
			purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
			if err != nil {
				return res, err
			}
			if (purchaseTime.Hour() > 14 || (purchaseTime.Hour() == 14 && purchaseTime.Minute() >= 0)) && (purchaseTime.Hour() < 16 || (purchaseTime.Hour() == 16 && purchaseTime.Minute() == 0)) {
				points += 10
			}
		}
	}
	fmt.Println(points)
	res.Points = points
	return res, nil
}
