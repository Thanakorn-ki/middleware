package middleware

import (
	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type middleware struct{}

/* Token
eyJ0eXAiOiJKV1QiLCJhbGciOiJIUzI1NiJ9.eyJpZCI6IjU1NTU1IiwidXNlcklkIjoiVUlEMjMyMTMyMTMyMzIxIiwianRpIjoiYWRkMGNmMzMtMzljMC00NzBmLWJhZjQtOTBkNWJmNzllZDRhIiwiaWF0IjoxNTQ1MDM5ODIxLCJleHAiOjE1NDUwNDM0MjF9.iMSjDk1eO0NjYS7rMckqAf_qhY9ECnFHvyijYO_sKeU

Expect decode
{ "id": "55555"  }
signature is "secret"
*/

type claims struct {
	ID string `json:"id"`
	jwt.StandardClaims
}

// NewMiddleware is init Middleware
func NewMiddleware() *middleware {
	return &middleware{}
}

func (m *middleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.DefaultQuery("token", "")
		if token == "" {
			respondWithError(401, "invalid token", c)
			return
		}
		m.decodeJWT(token, c)
	}
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.Abort()
}

func (m middleware) decodeJWT(tokenString string, c *gin.Context) {

	// secretKey := os.Getenv("SECRET_KEY")

	var secret []byte
	// if secretKey == "" {
	secret = []byte("")
	// } else {
	// 	secret = []byte(secretKey)
	// }

	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return secret, nil
	})

	if claims, ok := token.Claims.(*claims); ok && token.Valid {
		c.Set("claims", claims)
		c.Next()
	} else {
		respondWithError(401, err.Error(), c)
		return
	}
}
