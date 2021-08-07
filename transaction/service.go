package transaction

type service struct {
	repository Repository
}

type Service interface {
	GetTransactionsByCampaignID(input GetCampaignTransactionsiInput) ([]Transaction, error)
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetTransactionsByCampaignID(input GetCampaignTransactionsiInput) ([]Transaction, error) {
	trasactions, err := s.repository.GetByCampaignID(input.ID)
	if err != nil {
		return trasactions, err
	}

	return trasactions, nil
}