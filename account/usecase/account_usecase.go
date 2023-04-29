package usecase

import (
	"fmt"
	"log"
	"mini-wallet/domain"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type (
	accountUseCase struct {
		accountRepo domain.AccountRepository
	}
)

const (
	secretKey       = "very-secret-key"
	tokenExpiration = 24 * time.Hour
)

func NewAccountUseCase(accountRepo domain.AccountRepository) domain.AccountUseCase {
	return &accountUseCase{accountRepo}
}

func (c *accountUseCase) CreateAccount(customerXID string) (string, error) {

	token, err := createToken(customerXID)
	if err != nil {
		log.Printf("error when creating token: %v", err)
		return "", err
	}

	account := domain.Account{
		CustomerXID: customerXID,
		Token:       token,
	}
	err = c.accountRepo.CreateAccount(account)
	if err != nil {
		log.Printf("error when creating account: %v", err)
		return "", err
	}

	return token, nil
}

func createToken(customerXID string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"customer_xid": customerXID,
		"exp":          time.Now().Add(tokenExpiration).Unix(),
	})

	tokenString, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func (c *accountUseCase) ValidateAccountToken(tokenString string) (string, error) {
	customerXID, err := c.validateToken(tokenString)
	if err != nil {
		log.Printf("error validating token: %v", err)
		return "", err
	}
	return customerXID, nil
}

func (c *accountUseCase) validateToken(tokenString string) (string, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secretKey), nil
	})
	if err != nil {
		return "", fmt.Errorf("failed to parse token: %w", err)
	}
	if !token.Valid {
		return "", fmt.Errorf("invalid token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("invalid claims")
	}
	customerXID, ok := claims["customer_xid"].(string)
	if !ok {
		return "", fmt.Errorf("invalid customer_xid claim")
	}
	exp := claims["exp"].(float64)
	if time.Now().Unix() > int64(exp) {
		return "", fmt.Errorf("token has expired")
	}
	return customerXID, nil
}
