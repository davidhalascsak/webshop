package auth

type KeycloakClaims struct {
	Subject        string                 `json:"sub"`
	Issuer         string                 `json:"iss"`
	Audience       []string               `json:"aud"`
	ExpiresAt      int64                  `json:"exp"`
	ResourceAccess map[string]ClientRoles `json:"resource_access"`
}

type ClientRoles struct {
	Roles []string `json:"roles"`
}
