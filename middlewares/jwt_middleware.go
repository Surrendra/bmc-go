package middlewares

import (
	"BaliMediaCenter/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"net/http"
	"os"
	"strings"
	"time"
)

// Secret key for signing JWT tokens (should be securely loaded from environment in production)
var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// Claims defines the structure for JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

// GenerateJWT generates a new JWT token for a username
func GenerateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(60 * 24 * time.Minute)

	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}

	// Create the token and sign it using the secret key
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// AuthMiddleware is the middleware to authenticate JWT tokens
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.GetHeader("Authorization")
		// call responseHelper

		fmt.Println(tokenString)
		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Authorization token required"})
			c.Abort()
			return
		}
		tokenString = strings.TrimPrefix(tokenString, "Bearer ")
		claims := &Claims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return jwtSecret, nil
		})
		fmt.Println(token, err)
		if err != nil || !token.Valid {
			c.JSON(http.StatusUnauthorized, gin.H{"msg": "Invalid token"})
			c.Abort()
			return
		}

		// Store the username in the request context
		fmt.Println(claims)
		c.Set("username", claims.Username)
		c.Next()
	}
}

func GetSessionUser(c *gin.Context) (models.User, string) {
	code := c.MustGet("username")
	user := models.User{}
	models.DB.Where("code = ?", code).First(&user)
	if user.Code != code {
		return user, "User Not Found"
	}
	return user, ""
}
