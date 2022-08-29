package screening

type ScreeningRepository interface {
	FindByID(screeningId ScreeningID) (*Screening, error)
	Insert(screening *Screening) error
	Update(screening *Screening) error
}
