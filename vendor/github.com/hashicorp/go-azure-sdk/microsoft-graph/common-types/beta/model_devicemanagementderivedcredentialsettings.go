package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DeviceManagementDerivedCredentialSettings{}

type DeviceManagementDerivedCredentialSettings struct {
	// The display name for the profile.
	DisplayName *string `json:"displayName,omitempty"`

	// The URL that will be accessible to end users as they retrieve a derived credential using the Company Portal.
	HelpUrl nullable.Type[string] `json:"helpUrl,omitempty"`

	// Supported values for the derived credential issuer.
	Issuer *DeviceManagementDerivedCredentialIssuer `json:"issuer,omitempty"`

	// Supported values for the notification type to use.
	NotificationType *DeviceManagementDerivedCredentialNotificationType `json:"notificationType,omitempty"`

	// The nominal percentage of time before certificate renewal is initiated by the client.
	RenewalThresholdPercentage *int64 `json:"renewalThresholdPercentage,omitempty"`

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

func (s DeviceManagementDerivedCredentialSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DeviceManagementDerivedCredentialSettings{}

func (s DeviceManagementDerivedCredentialSettings) MarshalJSON() ([]byte, error) {
	type wrapper DeviceManagementDerivedCredentialSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceManagementDerivedCredentialSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceManagementDerivedCredentialSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.deviceManagementDerivedCredentialSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceManagementDerivedCredentialSettings: %+v", err)
	}

	return encoded, nil
}
