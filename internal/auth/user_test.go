package auth_test

import (
	"testing"

	"webshop/internal/auth"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestAuth(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Auth Suite")
}

var _ = Describe("UserFromClaims", func() {
	DescribeTable("allows users without recognized client roles",
		func(access map[string]auth.ClientRoles) {
			user, err := auth.UserFromClaims(&auth.KeycloakClaims{
				Subject:        "user-123",
				ResourceAccess: access,
			}, "webshop")

			Expect(err).NotTo(HaveOccurred())
			Expect(user.ID).To(Equal("user-123"))
			Expect(user.Roles).To(BeEmpty())
		},
		Entry("when the client has no roles", map[string]auth.ClientRoles{"webshop": {}}),
		Entry("when the client is absent", map[string]auth.ClientRoles{}),
		Entry("when the client has only unrecognized roles", map[string]auth.ClientRoles{"webshop": {Roles: []string{"support"}}}),
		Entry("when resource access is absent", nil),
	)
})
