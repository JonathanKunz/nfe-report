package ports

type TaxRepository struct {
	Get func(int)
}
