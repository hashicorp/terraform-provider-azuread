package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagedTenantEmailNotification{}

type ManagedTenantsManagedTenantEmailNotification struct {
	Alert              *ManagedTenantsManagedTenantAlert `json:"alert,omitempty"`
	CreatedByUserId    nullable.Type[string]             `json:"createdByUserId,omitempty"`
	CreatedDateTime    nullable.Type[string]             `json:"createdDateTime,omitempty"`
	EmailAddresses     *[]ManagedTenantsEmail            `json:"emailAddresses,omitempty"`
	EmailBody          nullable.Type[string]             `json:"emailBody,omitempty"`
	LastActionByUserId nullable.Type[string]             `json:"lastActionByUserId,omitempty"`
	LastActionDateTime nullable.Type[string]             `json:"lastActionDateTime,omitempty"`
	Subject            nullable.Type[string]             `json:"subject,omitempty"`

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

func (s ManagedTenantsManagedTenantEmailNotification) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagedTenantEmailNotification{}

func (s ManagedTenantsManagedTenantEmailNotification) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagedTenantEmailNotification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagedTenantEmailNotification: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagedTenantEmailNotification: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managedTenantEmailNotification"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagedTenantEmailNotification: %+v", err)
	}

	return encoded, nil
}
