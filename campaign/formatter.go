package campaign

import "strings"

type CampaignFormatter struct {
	ID 					int 	`json:"id"`
	UserID 				int 	`json:"user_id"`
	Name 				string 	`json:"name"`
	ShortDescription 	string 	`json:"short_description"`
	ImageUrl 			string 	`json:"image_url"`
	GoalAmount 			int 	`json:"goal_amount"`
	CurrentAmount 		int 	`json:"current_amount"`
	Slug				string 	`json:"slug"`
}

func FormatCampaign(campaign Campaign) CampaignFormatter {
	campaignFormatter := CampaignFormatter{
		ID:               campaign.ID,
		UserID:           campaign.UserID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		Slug: 			  campaign.Slug,
		ImageUrl: 		  "",
	}

	if len(campaign.CampaignImages) > 0 {
		campaignFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	return campaignFormatter
}

func FormatCampaigns(campaigns []Campaign) []CampaignFormatter {
	var campaignsFormatter []CampaignFormatter

	for _, campaign := range campaigns {
		campaignFormatter := FormatCampaign(campaign)
		campaignsFormatter = append(campaignsFormatter, campaignFormatter)
	}

	return campaignsFormatter
}

type CampaignUserFormatter struct {
	Name string `json:"name"`
	ImageUrl string `json:"image_url"`
}

type CampaignImageFormatter struct {
	ImageUrl 	string 	`json:"image_url"`
	IsPrimary 	bool 	`json:"is_primary"`
}

type CampaignDetailFormatter struct {
	ID 					int 						`json:"id"`
	Name 				string 						`json:"name"`
	ShortDescription 	string 						`json:"short_description"`
	Description 		string 						`json:"description"`
	ImageUrl 			string 						`json:"image_url"`
	GoalAmount 			int 						`json:"goal_amount"`
	CurrentAmount 		int 						`json:"current_amount"`
	BackerCount			int							`json:"backer_count"`
	UserID 				int 						`json:"user_id"`
	Slug 				string 						`json:"slug"`
	Perks 				[]string 					`json:"perks"`
	User 				CampaignUserFormatter 		`json:"user"`
	Images				[]CampaignImageFormatter 	`json:"images"`
}

func FormatDetailCampaign(campaign Campaign) CampaignDetailFormatter {
	campaignDetailFormatter := CampaignDetailFormatter{
		ID:               campaign.ID,
		Name:             campaign.Name,
		ShortDescription: campaign.ShortDescription,
		Description:      campaign.Description,
		GoalAmount:       campaign.GoalAmount,
		CurrentAmount:    campaign.CurrentAmount,
		BackerCount: 	  campaign.BackerCount,
		UserID:           campaign.UserID,
		Slug:             campaign.Slug,
		ImageUrl: 		  "",
	}

	if len(campaign.CampaignImages) > 0 {
		campaignDetailFormatter.ImageUrl = campaign.CampaignImages[0].FileName
	}

	campaignUserFormatter := CampaignUserFormatter{
		Name:     campaign.User.Name,
		ImageUrl: campaign.User.AvatarFileName,
	}
	campaignDetailFormatter.User = campaignUserFormatter

	var perks []string
	resPerks := strings.Split(campaign.Perks, ",")
	for _, perk := range resPerks {
		perks = append(perks, strings.TrimSpace(perk))
	}
	campaignDetailFormatter.Perks = perks

	var campaignImagesFormatter []CampaignImageFormatter
	for _, image := range campaign.CampaignImages {
		isPrimary := false
		if image.IsPrimary == 1 {
			isPrimary = true
		}

		campaignImage := CampaignImageFormatter{
			ImageUrl: image.FileName,
			IsPrimary: isPrimary,
		}

		campaignImagesFormatter = append(campaignImagesFormatter, campaignImage)
	}

	campaignDetailFormatter.Images = campaignImagesFormatter

	return campaignDetailFormatter
}