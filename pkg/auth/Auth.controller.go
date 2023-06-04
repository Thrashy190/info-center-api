package auth

import (
	"github.com/Thrashy190/info-center-api/pkg/models"
	"github.com/Thrashy190/info-center-api/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"os"
	"time"
)

type AuthenticationController struct {
	AuthService AuthenticationServiceImpl
}

func (c *AuthenticationController) Login(ctx *gin.Context) {
	var auth models.Auth

	if ctx.Bind(&auth) != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	user := c.AuthService.Login(&auth)

	if user.ID == 0 {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email or password"})
		return
	}

	if err := utils.CheckPasswordHash(auth.Password, user.Password); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid Email or password"})
		return
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("ACCESS_TOKEN_SECRET")))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to create token", "message": err.Error()})
		return
	}

	ctx.SetSameSite(http.SameSiteLaxMode)
	ctx.SetCookie("Authorization", tokenString, 3600*24*30, "", "", false, true)

	ctx.JSON(http.StatusOK, gin.H{"data": user})
}

func (c *AuthenticationController) SignUp(ctx *gin.Context) {
	var user models.Users

	if ctx.Bind(&user) != nil {
		ctx.JSON(400, gin.H{"error": "Invalid request payload"})
		return
	}

	hash, err := utils.HashPassword(user.Password)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}

	user = models.Users{
		Email:          user.Email,
		Password:       string(hash),
		Name:           user.Name,
		FatherLastName: user.FatherLastName,
		MotherLastName: user.MotherLastName,
		EmployeeNumber: user.EmployeeNumber,
	}

	if err := c.AuthService.SignUp(&user); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, gin.H{"data": user})
}

func (c *AuthenticationController) SignOut(ctx *gin.Context) {
	_, err := ctx.Cookie("Authorization")

	if err != nil {
		ctx.JSON(http.StatusUnauthorized, gin.H{"error": "Authorization Token is required"})
		ctx.Abort()
		return
	}
	ctx.SetCookie("Authorization", "", -1, "", "", false, true)
	ctx.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
