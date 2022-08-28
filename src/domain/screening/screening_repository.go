package screening

type ScreeningRepository interface {
	FindByID(screeningId ID) (*Screening, error)
	Insert(screening *Screening) error
	Update(screening *Screening) error
}
