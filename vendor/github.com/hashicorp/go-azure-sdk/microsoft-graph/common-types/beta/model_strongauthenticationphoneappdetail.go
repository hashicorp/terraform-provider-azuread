package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = StrongAuthenticationPhoneAppDetail{}

type StrongAuthenticationPhoneAppDetail struct {
	AuthenticationType               *string               `json:"authenticationType,omitempty"`
	AuthenticatorFlavor              nullable.Type[string] `json:"authenticatorFlavor,omitempty"`
	DeviceId                         nullable.Type[string] `json:"deviceId,omitempty"`
	DeviceName                       nullable.Type[string] `json:"deviceName,omitempty"`
	DeviceTag                        nullable.Type[string] `json:"deviceTag,omitempty"`
	DeviceToken                      nullable.Type[string] `json:"deviceToken,omitempty"`
	HashFunction                     nullable.Type[string] `json:"hashFunction,omitempty"`
	LastAuthenticatedDateTime        nullable.Type[string] `json:"lastAuthenticatedDateTime,omitempty"`
	NotificationType                 nullable.Type[string] `json:"notificationType,omitempty"`
	OathSecretKey                    nullable.Type[string] `json:"oathSecretKey,omitempty"`
	OathTokenMetadata                *OathTokenMetadata    `json:"oathTokenMetadata,omitempty"`
	OathTokenTimeDriftInSeconds      *int64                `json:"oathTokenTimeDriftInSeconds,omitempty"`
	PhoneAppVersion                  nullable.Type[string] `json:"phoneAppVersion,omitempty"`
	TenantDeviceId                   nullable.Type[string] `json:"tenantDeviceId,omitempty"`
	TokenGenerationIntervalInSeconds nullable.Type[int64]  `json:"tokenGenerationIntervalInSeconds,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s StrongAuthenticationPhoneAppDetail) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = StrongAuthenticationPhoneAppDetail{}

func (s StrongAuthenticationPhoneAppDetail) MarshalJSON() ([]byte, error) {
	type wrapper StrongAuthenticationPhoneAppDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling StrongAuthenticationPhoneAppDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling StrongAuthenticationPhoneAppDetail: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.strongAuthenticationPhoneAppDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling StrongAuthenticationPhoneAppDetail: %+v", err)
	}

	return encoded, nil
}
