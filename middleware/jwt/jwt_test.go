package jwt_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	. "github.com/salapao2136/middleware/middleware/jwt"
)

var _ = Describe("Jwt", func() {
	var secretKey []byte

	BeforeEach(func() {
		secretKey = []byte("")
	})

	Context("Middleware", func() {
		It(`should be return error message "{"error":"signature is invalid"} when not signature"`, func() {
			middle := NewMiddleware(secretKey)
			r := gin.New()
			r.Use(middle.Middleware())

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/", nil)
			req.Header.Set("X-TOKEN", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjaGFubmVsIjp7ImlkIjoiMTIiLCJzZWNyZXQiOiJzZWNyZXQiLCJ0b2tlbiI6InRva2VuIn19.62O7XhYgKWqfNsSiXtYgs9zTbS4fE4pjQTJ3UWuI-94")
			r.ServeHTTP(w, req)
			Expect(w.Code).Should(Equal(http.StatusBadRequest))
			Expect(w.Body.String()).To(Equal(`{"error":"signature is invalid"}`))
		})

		It(`should be able add value to context"`, func() {
			// secretKey = []byte("")
			middle := NewMiddleware(secretKey)
			r := gin.New()
			r.Use(middle.Middleware())

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/", nil)
			req.Header.Set("X-TOKEN", "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjaGFubmVsIjp7ImlkIjoiMTIiLCJzZWNyZXQiOiJzZWNyZXQiLCJ0b2tlbiI6InRva2VuIn19.vTAbRhoZHkrfoiYHjisA5CNNlJGqHPzbLmFV1_FdCuQ")
			r.ServeHTTP(w, req)

			Expect(w.Code).Should(Equal(http.StatusNotFound))
			Expect(w.Body.String()).To(Equal(`404 page not found`))
		})

		It(`should be return error message. when mock token purpose hack it`, func() {
			secretKey = []byte("inval")
			middle := NewMiddleware(secretKey)
			r := gin.New()
			r.Use(middle.Middleware())

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/", nil)
			req.Header.Set("X-TOKEN", "aeyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJjaGFubmVsIjp7ImlkIjoiMTIiLCJzZWNyZXQiOiJzZWNyZXQiLCJ0b2tlbiI6InRva2VuIn19.vTAbRhoZHkrfoiYHjisA5CNNlJGqHPzbLmFV1_FdCuQ")
			r.ServeHTTP(w, req)

			Expect(w.Code).Should(Equal(http.StatusBadRequest))
			Expect(w.Body.String()).To(Equal(`{"error":"illegal base64 data at input byte 37"}`))
		})

		It(`should be return error message. when token not found`, func() {
			secretKey = []byte("inval")
			middle := NewMiddleware(secretKey)
			r := gin.New()
			r.Use(middle.Middleware())

			w := httptest.NewRecorder()
			req, _ := http.NewRequest("GET", "/", nil)
			r.ServeHTTP(w, req)

			Expect(w.Code).Should(Equal(http.StatusBadRequest))
			Expect(w.Body.String()).To(Equal(`{"error":"Invalid Token"}`))
		})
	})
})
