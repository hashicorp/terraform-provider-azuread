package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserIdentity interface {
	Identity
	UserIdentity() BaseUserIdentityImpl
}

var _ UserIdentity = BaseUserIdentityImpl{}

type BaseUserIdentityImpl struct {
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

func (s BaseUserIdentityImpl) UserIdentity() BaseUserIdentityImpl {
	return s
}

func (s BaseUserIdentityImpl) Identity() BaseIdentityImpl {
	return BaseIdentityImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
	}
}

var _ UserIdentity = RawUserIdentityImpl{}

// RawUserIdentityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawUserIdentityImpl struct {
	userIdentity BaseUserIdentityImpl
	Type         string
	Values       map[string]interface{}
}

func (s RawUserIdentityImpl) UserIdentity() BaseUserIdentityImpl {
	return s.userIdentity
}

func (s RawUserIdentityImpl) Identity() BaseIdentityImpl {
	return s.userIdentity.Identity()
}

var _ json.Marshaler = BaseUserIdentityImpl{}

func (s BaseUserIdentityImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseUserIdentityImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseUserIdentityImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseUserIdentityImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userIdentity"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseUserIdentityImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalUserIdentityImplementation(input []byte) (UserIdentity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling UserIdentity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.auditUserIdentity") {
		var out AuditUserIdentity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AuditUserIdentity: %+v", err)
		}
		return out, nil
	}

	var parent BaseUserIdentityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseUserIdentityImpl: %+v", err)
	}

	return RawUserIdentityImpl{
		userIdentity: parent,
		Type:         value,
		Values:       temp,
	}, nil

}
