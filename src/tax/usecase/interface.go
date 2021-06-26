package usecase

type UseCase interface {
	// FetchTax()
}

type Reader interface {
	// Get()
	// Search()
	// List()
}

type Writer interface {
	Create() string
	// Update() error
	// Delete() error
}

type Repository interface {
	// Reader
	// Writer
}
