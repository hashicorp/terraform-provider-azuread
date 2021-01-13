package models

import (
	"time"
)

// Domain describes a Domain object.
type Domain struct {
	ID                               *string   `json:"id,omitempty"`
	AuthenticationType               *string   `json:"authenticationType,omitempty"`
	IsAdminManaged                   *bool     `json:"isAdminManaged,omitempty"`
	IsDefault                        *bool     `json:"isDefault,omitempty"`
	IsInitial                        *bool     `json:"isInitial,omitempty"`
	IsRoot                           *bool     `json:"isRoot,omitempty"`
	IsVerified                       *bool     `json:"isVerified,omitempty"`
	PasswordNotificationWindowInDays *int      `json:"passwordNotificationWindowInDays,omitempty"`
	PasswordValidityPeriodInDays     *int      `json:"passwordValidityPeriodInDays,omitempty"`
	SupportedServices                *[]string `json:"supportedServices,omitempty"`

	State *DomainState `json:"state,omitempty"`
}

type DomainState struct {
	LastActionDateTime *time.Time `json:"lastActionDateTime,omitempty"`
	Operation          *string    `json:"operation,omitempty"`
	Status             *string    `json:"status,omitempty"`
}
