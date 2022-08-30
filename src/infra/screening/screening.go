package screening

import "ddd-sample/src/domain/screening"

type screeningRepository struct {
}

func NewScreeningRepository() screening.ScreeningRepository {
	return &screeningRepository{}
}

func (r screeningRepository) FindByID(screeningId screening.ScreeningID) (*screening.Screening, error) {
	return nil, nil
}

func (r screeningRepository) Insert(screening *screening.Screening) error {
	return nil
}

func (r screeningRepository) Update(screening *screening.Screening) error {
	return nil
}
