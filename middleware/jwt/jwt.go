package jwt

import (
	"fmt"
	"net/http"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

type middleware struct {
	secret []byte
}

// Middleware provide to public
type Middleware interface {
	Middleware() gin.HandlerFunc
}

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
func NewMiddleware(secret []byte) Middleware {
	return &middleware{secret}
}

func (m *middleware) Middleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("X-TOKEN")
		if token == "" {
			respondWithError(http.StatusBadRequest, "Invalid Token", c)
			return
		}
		claims, err := m.decodeJWT(token)
		if err != nil {
			respondWithError(http.StatusBadRequest, err.Error(), c)

			return
		}
		c.Set("claims", claims)
		c.Next()
	}
}

func respondWithError(code int, message string, c *gin.Context) {
	resp := map[string]string{"error": message}
	c.JSON(code, resp)
	c.Abort()
}

func (m middleware) decodeJWT(tokenString string) (interface{}, error) {

	token, err := jwt.ParseWithClaims(tokenString, &claims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return m.secret, nil
	})

	if claims, ok := token.Claims.(*claims); ok && token.Valid {
		return claims, nil
	}
	return nil, err
}
