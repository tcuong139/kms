package services

import (
	"errors"
	"kms_golang/config"
	"kms_golang/database"
	"kms_golang/models"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Claims is reused from middleware package — defined here for service use
type AuthClaims struct {
	UserID    string `json:"user_id"`
	UserName  string `json:"user_name"`
	Auth      int16  `json:"auth"`
	GuardType string `json:"guard_type"`
	jwt.RegisteredClaims
}

const TokenExpiry = 24 * time.Hour

// safeStr dereferences a *string safely
func safeStr(s *string) string {
	if s == nil {
		return ""
	}
	return *s
}

// safeInt16 dereferences a *int16 safely
func safeInt16(n *int16) int16 {
	if n == nil {
		return 0
	}
	return *n
}

// LoginWeb authenticates a web (staff/admin) user by login_id
func LoginWeb(loginID, password string) (string, *models.User, error) {
	var user models.User
	if err := database.DB.Where("login_id = ? AND delete_flg = 0", loginID).First(&user).Error; err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	if user.Password == nil {
		return "", nil, errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*user.Password), []byte(password)); err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := generateToken(user.UserID, safeStr(user.UserName), safeInt16(user.Auth), "web")
	if err != nil {
		return "", nil, err
	}
	return token, &user, nil
}

// LoginCustomer authenticates a customer
func LoginCustomer(loginID, password string) (string, *models.Customer, error) {
	var customer models.Customer
	if err := database.DB.Where("customer_loginid = ? AND delete_flag = 0", loginID).First(&customer).Error; err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	if customer.CustomerPassword == nil {
		return "", nil, errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*customer.CustomerPassword), []byte(password)); err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := generateToken(customer.CustomerCd, safeStr(customer.CustomerName), 0, "customer")
	if err != nil {
		return "", nil, err
	}
	return token, &customer, nil
}

// LoginSekosaki authenticates a sekosaki (construction company) user
func LoginSekosaki(loginID, password string) (string, *models.Sekosaki, error) {
	var sekosaki models.Sekosaki
	if err := database.DB.Where("sekosaki_login_id = ? AND delete_flag = 0", loginID).First(&sekosaki).Error; err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	if sekosaki.SekosakiPassword == nil {
		return "", nil, errors.New("invalid credentials")
	}
	if err := bcrypt.CompareHashAndPassword([]byte(*sekosaki.SekosakiPassword), []byte(password)); err != nil {
		return "", nil, errors.New("invalid credentials")
	}

	token, err := generateToken(sekosaki.SekosakiCd, safeStr(sekosaki.SekosakiName), 0, "sekosaki")
	if err != nil {
		return "", nil, err
	}
	return token, &sekosaki, nil
}

// generateToken creates a signed JWT token
func generateToken(userID, userName string, auth int16, guardType string) (string, error) {
	claims := AuthClaims{
		UserID:    userID,
		UserName:  userName,
		Auth:      auth,
		GuardType: guardType,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(TokenExpiry)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWTSecret))
}

// HashPassword hashes a plain-text password
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
