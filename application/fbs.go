package application

import (
	"ozon-seller-api/domain/repository"
)

type FBSApplication struct {
	FBSRepository repository.FBSRepository
}

func NewFBSApplication(r repository.FBSRepository) *FBSApplication {
	if r == nil {
		return nil
	}

	return &FBSApplication{r}
}

// GetUnfulfilledList is
func (a *FBSApplication) GetUnfulfilledList() {
	a.FBSRepository.GetUnfulfilledList()
}

