package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID       uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username     string    `gorm:"type:varchar(100);unique;not null"`
	PasswordHash string    `gorm:"type:varchar(255);not null"`
	Email        string    `gorm:"type:varchar(100);unique;not null"`
	CreatedAt    time.Time
	Role         string `gorm:"type:varchar(50);not null"`
}

type Product struct {
	ProductID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	Price       float64   `gorm:"type:numeric;not null"`
	Stock       int       `gorm:"not null"`
	CategoryID  uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt   time.Time
}

type Category struct {
	CategoryID  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
}

type Order struct {
	OrderID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	OrderDate   time.Time
	Status      string  `gorm:"type:varchar(50);not null"`
	TotalAmount float64 `gorm:"type:numeric;not null"`
}

type OrderItem struct {
	OrderItemID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrderID     uuid.UUID `gorm:"type:uuid;not null"`
	ProductID   uuid.UUID `gorm:"type:uuid;not null"`
	Quantity    int       `gorm:"not null"`
	Price       float64   `gorm:"type:numeric;not null"`
}

type ShoppingCart struct {
	CartID    uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt time.Time
}

type CartItem struct {
	CartItemID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CartID     uuid.UUID `gorm:"type:uuid;not null"`
	ProductID  uuid.UUID `gorm:"type:uuid;not null"`
	Quantity   int       `gorm:"not null"`
}

type Payment struct {
	PaymentID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrderID       uuid.UUID `gorm:"type:uuid;not null"`
	Amount        float64   `gorm:"type:numeric;not null"`
	PaymentDate   time.Time
	PaymentMethod string `gorm:"type:varchar(50);not null"`
}

type Review struct {
	ReviewID  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProductID uuid.UUID `gorm:"type:uuid;not null"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Rating    int       `gorm:"not null"`
	Comment   string    `gorm:"type:text"`
	CreatedAt time.Time
}

type Session struct {
	SessionID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	CreatedAt time.Time
	ExpiresAt time.Time
}

type Role struct {
	RoleID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	RoleName string    `gorm:"type:varchar(50);not null"`
}

type UserAddress struct {
	AddressID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	Street    string    `gorm:"type:varchar(255)"`
	City      string    `gorm:"type:varchar(100)"`
	State     string    `gorm:"type:varchar(100)"`
	ZipCode   string    `gorm:"type:varchar(20)"`
}

type ProductImage struct {
	ImageID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProductID uuid.UUID `gorm:"type:uuid;not null"`
	ImageURL  string    `gorm:"type:text;not null"`
	CreatedAt time.Time
}

type AuditLog struct {
	LogID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Action    string    `gorm:"type:text"`
	UserID    uuid.UUID `gorm:"type:uuid"`
	Timestamp time.Time
}

type Cache struct {
	CacheKey       string `gorm:"type:varchar(255);primaryKey"`
	CacheValue     string `gorm:"type:text"`
	ExpirationTime time.Time
}

func Migrate(db *gorm.DB) {
	db.AutoMigrate(
		&User{},
		&Product{},
		&Category{},
		&Order{},
		&OrderItem{},
		&ShoppingCart{},
		&CartItem{},
		&Payment{},
		&Review{},
		&Session{},
		&Role{},
		&UserAddress{},
		&ProductImage{},
		&AuditLog{},
		&Cache{},
	)
}
