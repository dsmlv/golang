package models

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	UserID       uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Username     string       `gorm:"type:varchar(100);unique;not null"`
	PasswordHash string       `gorm:"type:varchar(255);not null"`
	Email        string       `gorm:"type:varchar(100);unique;not null"`
	Address      UserAddress  `gorm:"foreignKey:UserID"`
	RoleID       uuid.UUID    `gorm:"type:uuid;not null"`
	Role         Role         `gorm:"foreignKey:RoleID"`
	ShoppingCart ShoppingCart `gorm:"foreignKey:UserID"`
	Orders       []Order      `gorm:"foreignKey:UserID"`
	Reviews      []Review     `gorm:"foreignKey:UserID"`

	CreatedAt time.Time
}

type Product struct {
	ProductID   uuid.UUID      `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string         `gorm:"type:varchar(100);not null"`
	Description string         `gorm:"type:text"`
	Price       float64        `gorm:"type:numeric;not null"`
	Stock       int            `gorm:"not null"`
	CategoryID  uuid.UUID      `gorm:"type:uuid;not null"`
	Category    Category       `gorm:"foreignKey:CategoryID"`
	Images      []ProductImage `gorm:"foreignKey:ProductID"`
	Reviews     []Review       `gorm:"foreignKey:ProductID"`
	CreatedAt   time.Time
}

type Category struct {
	CategoryID  uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Name        string    `gorm:"type:varchar(100);not null"`
	Description string    `gorm:"type:text"`
	Product     []Product `gorm:"foreignKey:CategoryID"`
}

type Order struct {
	OrderID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID      uuid.UUID `gorm:"type:uuid;not null"`
	OrderDate   time.Time
	Status      string      `gorm:"type:varchar(50);not null"`
	TotalAmount float64     `gorm:"type:numeric;not null"`
	User        User        `gorm:"foreignKey:UserID"`
	Items       []OrderItem `gorm:"foreignKey:OrderID"`
}

type OrderItem struct {
	OrderItemID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	OrderID     uuid.UUID `gorm:"type:uuid;not null"`
	ProductID   uuid.UUID `gorm:"type:uuid;not null"`
	Quantity    int       `gorm:"not null"`
	Price       float64   `gorm:"type:numeric;not null"`
	Order       Order     `gorm:"foreignKey:OrderID"`
	Product     Product   `gorm:"foreignKey:ProductID"`
}

type ShoppingCart struct {
	CartID    uuid.UUID  `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID  `gorm:"type:uuid;not null;unique"`
	Items     []CartItem `gorm:"foreignKey:CartID"`
	CreatedAt time.Time
}

type CartItem struct {
	CartItemID uuid.UUID    `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	CartID     uuid.UUID    `gorm:"type:uuid;not null"`
	ProductID  uuid.UUID    `gorm:"type:uuid;not null"`
	Quantity   int          `gorm:"not null"`
	Cart       ShoppingCart `gorm:"foreignKey:CartID"`
	Product    Product      `gorm:"foreignKey:ProductID"`
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
	User      User      `gorm:"foreignKey:UserID"`
	Product   Product   `gorm:"foreignKey:ProductID"`
	CreatedAt time.Time
}

type Session struct {
	SessionID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null"`
	User      User      `gorm:"foreignKey:UserID"`
	CreatedAt time.Time
	ExpiresAt time.Time
}

type Role struct {
	RoleID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	RoleName string    `gorm:"type:varchar(50);not null"`
}

type UserAddress struct {
	AddressID uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	UserID    uuid.UUID `gorm:"type:uuid;not null,unique"`
	Street    string    `gorm:"type:varchar(255)"`
	City      string    `gorm:"type:varchar(100)"`
	State     string    `gorm:"type:varchar(100)"`
	ZipCode   string    `gorm:"type:varchar(20)"`
}

type ProductImage struct {
	ImageID   uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	ProductID uuid.UUID `gorm:"type:uuid;not null"`
	ImageURL  string    `gorm:"type:text;not null"`
	Product   Product   `gorm:"foreignKey:ProductID"`
	CreatedAt time.Time
}

type AuditLog struct {
	LogID     uuid.UUID `gorm:"type:uuid;default:uuid_generate_v4();primaryKey"`
	Action    string    `gorm:"type:text"`
	UserID    uuid.UUID `gorm:"type:uuid"`
	User      User      `gorm:"foreignKey:UserID"`
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
