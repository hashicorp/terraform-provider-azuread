package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TelecomExpenseManagementPartner{}

type TelecomExpenseManagementPartner struct {
	// Whether the partner's AAD app has been authorized to access Intune.
	AppAuthorized *bool `json:"appAuthorized,omitempty"`

	// Display name of the TEM partner.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Whether Intune's connection to the TEM service is currently enabled or disabled.
	Enabled *bool `json:"enabled,omitempty"`

	// Timestamp of the last request sent to Intune by the TEM partner.
	LastConnectionDateTime *string `json:"lastConnectionDateTime,omitempty"`

	// URL of the TEM partner's administrative control panel, where an administrator can configure their TEM service.
	Url nullable.Type[string] `json:"url,omitempty"`

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

func (s TelecomExpenseManagementPartner) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TelecomExpenseManagementPartner{}

func (s TelecomExpenseManagementPartner) MarshalJSON() ([]byte, error) {
	type wrapper TelecomExpenseManagementPartner
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TelecomExpenseManagementPartner: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TelecomExpenseManagementPartner: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.telecomExpenseManagementPartner"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TelecomExpenseManagementPartner: %+v", err)
	}

	return encoded, nil
}
