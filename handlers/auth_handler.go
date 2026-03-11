package handlers

import (
	"kms_golang/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

// PostLoginWeb handles POST /login (staff/admin login)
func PostLoginWeb(c *gin.Context) {
	var req struct {
		LoginID  string `json:"login_id" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	token, user, err := services.LoginWeb(req.LoginID, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "ログインIDまたはパスワードが正しくありません"})
		return
	}

	userName := ""
	if user.UserName != nil {
		userName = *user.UserName
	}
	var auth int16
	if user.Auth != nil {
		auth = *user.Auth
	}

	c.JSON(http.StatusOK, gin.H{
		"token":     token,
		"user_id":   user.UserID,
		"user_name": userName,
		"auth":      auth,
		"guard":     "web",
	})
}

// PostLoginCustomer handles POST /login-customer
func PostLoginCustomer(c *gin.Context) {
	var req struct {
		LoginID  string `json:"login_id" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	token, customer, err := services.LoginCustomer(req.LoginID, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "ログインIDまたはパスワードが正しくありません"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":       token,
		"customer_cd": customer.CustomerCd,
		"customer_name": func() string {
			if customer.CustomerName == nil {
				return ""
			}
			return *customer.CustomerName
		}(),
		"guard": "customer",
	})
}

// PostLoginSekosaki handles POST /login-sekosaki
func PostLoginSekosaki(c *gin.Context) {
	var req struct {
		LoginID  string `json:"login_id" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{"errors": err.Error()})
		return
	}

	token, sekosaki, err := services.LoginSekosaki(req.LoginID, req.Password)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "ログインIDまたはパスワードが正しくありません"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"token":       token,
		"sekosaki_cd": sekosaki.SekosakiCd,
		"sekosaki_name": func() string {
			if sekosaki.SekosakiName == nil {
				return ""
			}
			return *sekosaki.SekosakiName
		}(),
		"guard": "sekosaki",
	})
}

// Logout handles GET /logout (all guards)
func Logout(c *gin.Context) {
	// JWT is stateless; client discards the token
	c.JSON(http.StatusOK, gin.H{"message": "ログアウトしました"})
}
