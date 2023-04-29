package middleware

import (
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"gorm.io/gorm"
)

func Authentication() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientToken := ctx.Request.Header.Get("Authorization")
		if clientToken == "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "No Authorization Header Provided"})
		}

		splitToken := strings.Split(clientToken, "Bearer")
		if len(splitToken) != 2 {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": "Bearer token not in proper format"})
		}

		reqToken := strings.TrimSpace(splitToken[1])

		claims, err := ValidateToken(reqToken)
		if err != "" {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err})
			ctx.Abort()
			return
		}
		ctx.Set("user_pid", claims.Pid)
		ctx.Next()
	}
}

type SignedDetails struct {
	Pid string `json:"user_pid"`
	jwt.StandardClaims
}

func ValidateToken(SignedToken string) (claims *SignedDetails, msg string) {
	token, err := jwt.ParseWithClaims(SignedToken, &SignedDetails{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_KEY")), nil
	})
	if err != nil {
		msg = "Authenticaltion Failed!"
		return claims, msg
	}
	claims, ok := token.Claims.(*SignedDetails)
	if !ok {
		msg = "the token is invalid"
		return claims, msg
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		msg = "the token is expired"
		return claims, msg
	}
	return claims, msg
}

func DbMiddleware(db *gorm.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Set("DB", db)
		ctx.Next()
	}
}
