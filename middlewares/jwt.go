package middlewares

import (
	"errors"
	"github.com/arifbugaresa/mnc-wallet/utils/response"
	"github.com/arifbugaresa/mnc-wallet/utils/session"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"strings"
)

type Claims struct {
	Role string `json:"role"`
	jwt.StandardClaims
}

func (c Claims) GenerateJwtToken() (token string, err error) {
	claims := &Claims{
		Role:           c.Role,
		StandardClaims: jwt.StandardClaims{},
	}

	generatedTokenJwt := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	token, err = generatedTokenJwt.SignedString([]byte(viper.GetString("jwt_secret_key")))
	if err != nil {
		return
	}

	return
}

func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := GetJwtTokenFromHeader(ctx)
		if err != nil {
			response.GenerateErrorResponse(ctx, err.Error())
			return
		}

		// check token in session
		redisSession, err := session.RedisClient.Get(ctx, token).Result()
		if err != nil {
			response.GenerateErrorResponse(ctx, "token expired")
			return
		}

		if redisSession == "" {
			response.GenerateErrorResponse(ctx, "invalid token")
			return
		}

		ctx.Next()
	}
}

func GetJwtTokenFromHeader(c *gin.Context) (tokenString string, err error) {
	authHeader := c.Request.Header.Get("Authorization")

	if authHeader == "" {
		return tokenString, errors.New("authorization header is required")
	}

	parts := strings.Split(authHeader, " ")
	if len(parts) != 2 || parts[0] != "Bearer" {
		return tokenString, errors.New("invalid Authorization header format")
	}

	return parts[1], nil
}
