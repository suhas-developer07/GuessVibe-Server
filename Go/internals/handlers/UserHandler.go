package handlers

import (
	"net/http"

	"github.com/labstack/echo/v4"
	models "github.com/suhas-developer07/GuessVibe-Server/internals/models/User_model"
	services "github.com/suhas-developer07/GuessVibe-Server/internals/service/user"
)

type UserHandler struct {
	UserServices *services.UserService
}

func (h *UserHandler) UserRegisterHandler(c echo.Context) error {
	var req models.User
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, "Invalid request")
	}
	id, err := h.UserServices.RegisterUser(req)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status: "error",
			Error:  "Failed to register user: " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "User registered successfully",
		Data:    map[string]int64{"user_id": id},
	})

}
func (h *UserHandler) UserLoginHandler(c echo.Context) error {
	var req models.UserLogin
	if err := c.Bind(&req); err != nil {
		return c.JSON(400, "Invalid request")
	}
	token, err := h.UserServices.LoginUser(req.Email, req.Password)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status: "error",
			Error:  "Failed to login user: " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "sucess",
		Message: "User LoggedIn Sucessfully",
		Data:    map[string]string{"token": token},
	})
}
func (h *UserHandler) LogoutUserHandler(c echo.Context) error {
	var Userlogout struct {
		UserID string `jsomn:"userid"`
	}
	err := c.Bind(&Userlogout)
	if err != nil {
		return c.JSON(400, models.ErrorResponse{
			Status: "error",
			Error:  "Failed to logout user: " + err.Error(),
		})
	}
	err= h.UserServices.LogoutUser(Userlogout.UserID,"")
	if err != nil {
		return c.JSON(http.StatusInternalServerError, models.ErrorResponse{
			Status: "error",
			Error:  "Failed to logout user: " + err.Error(),
		})
	}
	return c.JSON(http.StatusOK, models.SuccessResponse{
		Status:  "success",
		Message: "User logged out successfully",
	})
}
