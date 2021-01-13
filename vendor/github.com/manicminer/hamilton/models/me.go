package models

// Me describes the authenticated user.
type Me struct {
	ID                *string `json:"id"`
	DisplayName       *string `json:"displayName"`
	UserPrincipalName *string `json:"userPrincipalName"`
}
