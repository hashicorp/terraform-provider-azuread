package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsWorkloadAction struct {
	// The unique identifier for the workload action. Required. Read-only.
	ActionId nullable.Type[string] `json:"actionId,omitempty"`

	// The category for the workload action. Possible values are: automated, manual, unknownFutureValue. Optional.
	// Read-only.
	Category *ManagedTenantsWorkloadActionCategory `json:"category,omitempty"`

	// The description for the workload action. Optional. Read-only.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The display name for the workload action. Optional. Read-only.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	Licenses *[]string `json:"licenses,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The service associated with workload action. Optional. Read-only.
	Service nullable.Type[string] `json:"service,omitempty"`

	// The collection of settings associated with the workload action. Optional. Read-only.
	Settings *[]ManagedTenantsSetting `json:"settings,omitempty"`
}

var _ json.Marshaler = ManagedTenantsWorkloadAction{}

func (s ManagedTenantsWorkloadAction) MarshalJSON() ([]byte, error) {
	type wrapper ManagedTenantsWorkloadAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManagedTenantsWorkloadAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManagedTenantsWorkloadAction: %+v", err)
	}

	delete(decoded, "actionId")
	delete(decoded, "category")
	delete(decoded, "description")
	delete(decoded, "displayName")
	delete(decoded, "service")
	delete(decoded, "settings")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManagedTenantsWorkloadAction: %+v", err)
	}

	return encoded, nil
}
