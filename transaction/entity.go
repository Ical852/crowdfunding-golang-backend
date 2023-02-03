package transaction

import (
	"bwastartup/campaign"
	"bwastartup/user"
	"github.com/leekchan/accounting"
	"time"
)

type Transaction struct {
	ID 			int
	CampaignID 	int
	UserID 		int
	Amount 		int
	Status 		string
	Code 		string
	PaymentUrl	string
	User 		user.User
	Campaign 	campaign.Campaign
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}

func (t Transaction) AmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(t.Amount)
}