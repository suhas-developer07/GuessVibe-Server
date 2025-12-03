package user_model

type User struct {
	ID       string `json:"_id" bsom:"_id"`
	UserID   string `json:"userid" bson:"userid" validate:"required,min=3,max=30"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=8"`
	Token    string `json:"token" bson:"token"`
}
type UserLogin struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=8"`
}

type UserRepo interface {
	RegisterUser(User User) (int64, error)
	LoginUser(Email, password string) (string, error)
	LogoutUser(userID, token string) error
}
