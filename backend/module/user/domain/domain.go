package domain

type Role string

const (
	AdminRole Role = "admin"
	UserRole  Role = "user"
)

type User struct {
	ID          string
	Email       string
	Password    string
	FIO         string
	PhoneNumber string
	Role        Role
}
