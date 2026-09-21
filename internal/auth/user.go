package auth

import "fmt"

type Role string

const (
	RoleUser  Role = "user"
	RoleAdmin Role = "admin"
)

type User struct {
	ID    string
	Roles []Role
}

func (u *User) HasRole(role Role) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}

func UserFromClaims(claims *KeycloakClaims, clientID string) (User, error) {
	user := User{
		ID: claims.Subject,
	}

	client, ok := claims.ResourceAccess[clientID]
	if !ok {
		return User{}, fmt.Errorf("roles for client %q not found", clientID)
	}

	for _, role := range client.Roles {
		switch role {
		case string(RoleUser):
			user.Roles = append(user.Roles, RoleUser)

		case string(RoleAdmin):
			user.Roles = append(user.Roles, RoleAdmin)

		default:
		}
	}

	if len(user.Roles) == 0 {
		return User{}, fmt.Errorf("user has no recognized roles for client %q", clientID)
	}

	return user, nil
}
