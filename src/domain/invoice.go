package domain

type Buyer struct {
	id   int
	name string
}

type Seller struct {
	id   int
	name string
}

type Data struct {
	model  int
	buyer  Buyer
	seller Seller
	total  float64
}

type Invoice struct {
	ID   int64
	Data Data
}
