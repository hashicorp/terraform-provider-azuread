package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TenantSetupInfo{}

type TenantSetupInfo struct {
	DefaultRolesSettings  *PrivilegedRoleSettings `json:"defaultRolesSettings,omitempty"`
	FirstTimeSetup        nullable.Type[bool]     `json:"firstTimeSetup,omitempty"`
	RelevantRolesSettings *[]string               `json:"relevantRolesSettings,omitempty"`
	SetupStatus           *SetupStatus            `json:"setupStatus,omitempty"`
	SkipSetup             nullable.Type[bool]     `json:"skipSetup,omitempty"`
	UserRolesActions      nullable.Type[string]   `json:"userRolesActions,omitempty"`

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

func (s TenantSetupInfo) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TenantSetupInfo{}

func (s TenantSetupInfo) MarshalJSON() ([]byte, error) {
	type wrapper TenantSetupInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TenantSetupInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TenantSetupInfo: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.tenantSetupInfo"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TenantSetupInfo: %+v", err)
	}

	return encoded, nil
}
