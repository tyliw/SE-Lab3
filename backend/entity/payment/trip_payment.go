package payment

import (
	"gorm.io/gorm"
	"time"
)

type TripPayment struct {
	gorm.Model
	PaymentDate 	time.Time 		`gorm:"not null"`
	TotalPrice 		float32			`gorm:"not null"`
	PaymentStatus	string			`gorm:"not null"`
	PaymentMethod	string			`gorm:"not null"`
}