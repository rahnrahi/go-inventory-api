package entities

type Product struct {
	UUIDBaseModel
	Name        string  `json:"name"`
	Price       float64 `json:"price"`
	Description string  `json:"description"`
}
