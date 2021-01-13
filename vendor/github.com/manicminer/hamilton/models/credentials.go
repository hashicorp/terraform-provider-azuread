package models

import "time"

// KeyCredential describes a key (certificate) credential for an object.
type KeyCredential struct {
	CustomKeyIdentifier *string    `json:"customKeyIdentifier,omitempty"`
	DisplayName         *string    `json:"displayName,omitempty"`
	EndDateTime         *time.Time `json:"endDateTime,omitempty"`
	KeyId               *string    `json:"keyId,omitempty"`
	StartDateTime       *time.Time `json:"startDateTime,omitempty"`
	Type                *string    `json:"type,omitempty"`
	Usage               *string    `json:"usage,omitempty"`
	Key                 *string    `json:"key,omitempty"`
}

// PasswordCredential describes a password credential for an object.
type PasswordCredential struct {
	CustomKeyIdentifier *string    `json:"customKeyIdentifier,omitempty"`
	DisplayName         *string    `json:"displayName,omitempty"`
	EndDateTime         *time.Time `json:"endDateTime,omitempty"`
	Hint                *string    `json:"hint,omitempty"`
	KeyId               *string    `json:"keyId,omitempty"`
	SecretText          *string    `json:"secretText,omitempty"`
	StartDateTime       *time.Time `json:"startDateTime,omitempty"`
}
