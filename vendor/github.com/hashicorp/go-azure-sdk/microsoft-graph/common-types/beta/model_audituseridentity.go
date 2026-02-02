package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ UserIdentity = AuditUserIdentity{}

type AuditUserIdentity struct {
	// For user sign ins, the identifier of the tenant that the user is a member of.
	HomeTenantId nullable.Type[string] `json:"homeTenantId,omitempty"`

	// For user sign ins, the name of the tenant that the user is a member of. Only populated in cases where the home tenant
	// has provided affirmative consent to Microsoft Entra ID to show the tenant content.
	HomeTenantName nullable.Type[string] `json:"homeTenantName,omitempty"`

	// Fields inherited from UserIdentity

	// Indicates the client IP address associated with the user performing the activity (audit log only).
	IPAddress nullable.Type[string] `json:"ipAddress,omitempty"`

	// The userPrincipalName attribute of the user.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// Fields inherited from Identity

	// The display name of the identity. This property is read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The identifier of the identity. This property is read-only.
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AuditUserIdentity) UserIdentity() BaseUserIdentityImpl {
	return BaseUserIdentityImpl{
		IPAddress:         s.IPAddress,
		UserPrincipalName: s.UserPrincipalName,
		DisplayName:       s.DisplayName,
		Id:                s.Id,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s AuditUserIdentity) Identity() BaseIdentityImpl {
	return BaseIdentityImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

var _ json.Marshaler = AuditUserIdentity{}

func (s AuditUserIdentity) MarshalJSON() ([]byte, error) {
	type wrapper AuditUserIdentity
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AuditUserIdentity: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AuditUserIdentity: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.auditUserIdentity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AuditUserIdentity: %+v", err)
	}

	return encoded, nil
}
