package transaction

import "bwa_startup/user"

type GetCampaignTransactionsiInput struct {
	ID   int `uri:"id" binding:"required"`
	User user.User
}