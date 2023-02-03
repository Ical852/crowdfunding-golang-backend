package campaign

import (
	"bwastartup/user"
	"github.com/leekchan/accounting"
	"time"
)

type Campaign struct {
	ID 					int
	UserID 				int
	Name 				string
	ShortDescription 	string
	Description 		string
	Perks 				string
	BackerCount 		int
	GoalAmount 			int
	CurrentAmount 		int
	Slug 				string
	CreatedAt 			time.Time
	UpdatedAt 			time.Time
	CampaignImages		[]CampaignImage
	User 				user.User
}

func (c Campaign) GoalAmountFormatIDR() string {
	ac := accounting.Accounting{Symbol: "Rp", Precision: 2, Thousand: ".", Decimal: ","}
	return ac.FormatMoney(c.GoalAmount)
}

type CampaignImage struct {
	ID 			int
	CampaignID 	int
	FileName 	string
	IsPrimary 	int
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
}