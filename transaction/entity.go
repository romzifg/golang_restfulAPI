package transaction

import (
	"bwa_startup/campaign"
	"bwa_startup/user"
	"time"
)

type Transaction struct {
	ID         int 		 `json:"id"`
	CampaignID int 		 `json:"campaign_id"`
	UserID     int 		 `json:"user_id"`
	Amount     int 		 `json:"amount"`
	Status     string 	 `json:"status"`
	Code       string 	 `json:"code"`
	User 	   user.User
	Campaign   campaign.Campaign
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}