package services

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
	"server/internal/config"
	"server/internal/database"
	"server/internal/models"
	"time"
)

type JwtDbMethods interface {
	JwtTokenCheck(c *gin.Context)
}

type JwtDbConnector struct {
	UsersTable database.UsersTable
}

func extractBearerToken(header string) (string, error) {
	if header == "" {
		return "", errors.New("no token in header")
	}

	return header, nil
}

func GenerateToken(username string, password string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user":     username,
		"password": password,
		"nbf":      time.Date(2025, 01, 01, 12, 0, 0, 0, time.UTC).Unix(),
	})

	tokenStr, err := token.SignedString([]byte(config.GetEnv(config.JwtSecret)))

	return tokenStr, err
}

func (j JwtDbConnector) JwtTokenCheck(c *gin.Context) {
	jwtToken, err := extractBearerToken(c.GetHeader(config.GetEnv(config.AuthorizationHeader)))
	if err != nil {
		return
	}

	//TODO Необходимо валидировать токен перед запрсом к базе

	//Check in DB
	isUserExist := j.UsersTable.GetOneByToken(jwtToken)

	if !isUserExist {
		c.AbortWithStatusJSON(http.StatusUnauthorized, models.SuccessResponse{
			Message: "Invalid jwt token",
		})
		return
	}

	c.Next()
}
