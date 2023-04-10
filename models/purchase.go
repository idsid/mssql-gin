package models

type Purchase struct {
	ItemCode   string  `gorm:"column:ItemCode"`
	Dscription string  `gorm:"column:Dscription"`
	Quantity   float64 `gorm:"column:Quantity"`
}
