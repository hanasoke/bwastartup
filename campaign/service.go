package campaign

type Service interface {
	GetCampaigns(UserID int) ([]Campaign, error)
}

type service struct {
	repository Repository
}

// FindCampaigns implements Service.
func (*service) FindCampaigns(UserID int) ([]Campaign, error) {
	panic("unimplemented")
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (s *service) GetCampaigns(userID int) ([]Campaign, error) {
	if userID != 0 {
		campaigns, err := s.repository.FindByUserID(userID)
		if err != nil {
			return campaigns, err
		}

		return campaigns, nil
	}

	campaigns, err := s.repository.FindAll()
	if err != nil {
		return campaigns, err
	}

	return campaigns, nil
}
