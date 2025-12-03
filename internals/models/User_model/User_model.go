package user_model

type User struct {
	ID       string `json:"_id" bsom:"_id"`
	Name     string `json:"name" bson:"name" validate:"required"`
	UserID   string `json:"userid" bson:"userid"`
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=8"`
	Token    string `json:"token" bson:"token"`
}
type UserLogin struct {
	Email    string `json:"email" bson:"email" validate:"required,email"`
	Password string `json:"password" bson:"password" validate:"required,min=8"`
}

type UserRepo interface {
	RegisterUser(User User) (string, error)
	LoginUser(Email, password string) (string, error)
	LogoutUser(userID, token string) error
}
