package api_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"webshop/internal/api"
	"webshop/internal/auth"

	"github.com/gin-gonic/gin"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAPI(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "API Suite")
}

var _ = Describe("RequireRole", func() {
	BeforeEach(func() {
		gin.SetMode(gin.TestMode)
	})

	DescribeTable("authorizes requests according to the user's roles",
		func(user *auth.User, roles []auth.Role, wantStatus int) {
			engine := gin.New()
			engine.Use(api.ErrorMiddleware())
			engine.Use(func(c *gin.Context) {
				if user != nil {
					c.Set("user", *user)
				}
				c.Next()
			})
			engine.GET("/", api.RequireRole(roles...), func(c *gin.Context) {
				c.Status(http.StatusOK)
			})

			recorder := httptest.NewRecorder()
			request := httptest.NewRequest(http.MethodGet, "/", nil)
			engine.ServeHTTP(recorder, request)

			Expect(recorder.Code).To(Equal(wantStatus))
		},
		Entry("allows a matching role",
			&auth.User{Roles: []auth.Role{auth.RoleAdmin}},
			[]auth.Role{auth.RoleAdmin},
			http.StatusOK,
		),
		Entry("allows any matching role",
			&auth.User{Roles: []auth.Role{auth.RoleUser}},
			[]auth.Role{auth.RoleUser, auth.RoleAdmin},
			http.StatusOK,
		),
		Entry("forbids a user without a matching role",
			&auth.User{Roles: []auth.Role{auth.RoleUser}},
			[]auth.Role{auth.RoleAdmin},
			http.StatusForbidden,
		),
		Entry("rejects a request without an authenticated user",
			nil,
			[]auth.Role{auth.RoleAdmin},
			http.StatusUnauthorized,
		),
	)
})
