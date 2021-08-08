package transaction

import "bwa_startup/user"

type GetCampaignTransactionsiInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}

type CreateTransactionInput struct {
	Amount 		int `json:"amount" binding:"required"`
	CampaignID 	int `json:"campaign_id" binding:"required"`
	User 		user.User
}