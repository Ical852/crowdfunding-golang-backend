package transaction

import "time"

type CampaignTransactionFormatter struct {
	ID 			int 		`json:"id"`
	Name 		string 		`json:"name"`
	Amount 		int 		`json:"amount"`
	CratedAt 	time.Time 	`json:"crated_at"`
}

func FormatCampaignTransaction(transaction Transaction) CampaignTransactionFormatter {
	formatter := CampaignTransactionFormatter{
		ID:       transaction.ID,
		Name:     transaction.User.Name,
		Amount:   transaction.Amount,
		CratedAt: transaction.CreatedAt,
	}

	return formatter
}

func FormatCampaignTransactions(transactions []Transaction) []CampaignTransactionFormatter {
	transactionsFormatted := []CampaignTransactionFormatter{}

	for _, transaction := range transactions {
		formatTransaction := FormatCampaignTransaction(transaction)
		transactionsFormatted = append(transactionsFormatted, formatTransaction)
	}

	return transactionsFormatted
}

type CampaignFormatter struct {
	Name 		string `json:"name"`
	ImageUrl 	string `json:"image_url"`
}

type UserTransactionFormatter struct {
	ID 			int 				`json:"id"`
	Amount 		int 				`json:"amount"`
	Status 		string 				`json:"status"`
	CreatedAt 	time.Time 			`json:"created_at"`
	Campaign 	CampaignFormatter 	`json:"campaign"`
}

func FormatUserTransaction(transaction Transaction) UserTransactionFormatter {
	imageUrl := ""

	for _, image := range transaction.Campaign.CampaignImages {
		if image.IsPrimary == 1 {
			imageUrl = image.FileName
		}
	}

	formatter := UserTransactionFormatter{
		ID:        transaction.ID,
		Amount:    transaction.Amount,
		Status:    transaction.Status,
		CreatedAt: transaction.CreatedAt,
		Campaign:  CampaignFormatter{
			Name:     transaction.Campaign.Name,
			ImageUrl: imageUrl,
		},
	}

	return formatter
}

func FormatUserTransactions(transactions []Transaction) []UserTransactionFormatter {
	formatter := []UserTransactionFormatter{}
	for _, transaction := range transactions {
		formattedTransaction := FormatUserTransaction(transaction)
		formatter = append(formatter, formattedTransaction)
	}

	return formatter
}

type TransactionFormatter struct {
	ID 			int 		`json:"id"`
	CampaignID 	int 		`json:"campaign_id"`
	UserID 		int 		`json:"user_id"`
	Amount 		int 		`json:"amount"`
	Status 		string 		`json:"status"`
	Code 		string 		`json:"code"`
	PaymentUrl 	string 		`json:"payment_url"`
	CreatedAt 	time.Time 	`json:"created_at"`
}

func FormatTransaction(transaction Transaction) TransactionFormatter {
	formatter := TransactionFormatter{
		ID:         transaction.ID,
		CampaignID: transaction.CampaignID,
		UserID:     transaction.UserID,
		Amount:     transaction.Amount,
		Status:     transaction.Status,
		Code:       transaction.Code,
		PaymentUrl: transaction.PaymentUrl,
		CreatedAt:  transaction.CreatedAt,
	}

	return formatter
}