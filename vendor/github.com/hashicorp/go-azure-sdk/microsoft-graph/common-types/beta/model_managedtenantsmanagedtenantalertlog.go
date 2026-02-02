package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ManagedTenantsManagedTenantAlertLog{}

type ManagedTenantsManagedTenantAlertLog struct {
	Alert              *ManagedTenantsManagedTenantAlert `json:"alert,omitempty"`
	Content            *ManagedTenantsAlertLogContent    `json:"content,omitempty"`
	CreatedByUserId    nullable.Type[string]             `json:"createdByUserId,omitempty"`
	CreatedDateTime    nullable.Type[string]             `json:"createdDateTime,omitempty"`
	LastActionByUserId nullable.Type[string]             `json:"lastActionByUserId,omitempty"`
	LastActionDateTime nullable.Type[string]             `json:"lastActionDateTime,omitempty"`

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

func (s ManagedTenantsManagedTenantAlertLog) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ManagedTenantsManagedTenantAlertLog{}

func (s ManagedTenantsManagedTenantAlertLog) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsManagedTenantAlertLog
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsManagedTenantAlertLog: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsManagedTenantAlertLog: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.managedTenants.managedTenantAlertLog"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsManagedTenantAlertLog: %+v", err)
	}

	return encoded, nil
}
