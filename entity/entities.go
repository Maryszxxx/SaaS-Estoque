package entity

type User struct {
	ID       int64
	Name     string
	Password string
}

type Product struct {
	ID       int64
	Name     string
	Price    float64
	Category string
}
