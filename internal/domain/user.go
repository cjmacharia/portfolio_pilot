package domain

type User struct {
	UserID        int     `json:"_"`
	Email         string  `json:"email"`
	Name          string  `json:"name"`
	Password      string  `json:"_"`
	WalletBalance float64 `json:"wallet_balance"`
}

type UserRepository interface {
	SignUp(u *User) (int, error)
	GetUserByEmail(email string) (User, error)
	GetUserWalletBalance(userID int) (float64, error)
	UpdateUserWalletBalance(userID int, newBalance float64) error
}
