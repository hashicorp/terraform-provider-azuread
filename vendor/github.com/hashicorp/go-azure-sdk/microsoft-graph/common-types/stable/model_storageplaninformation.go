package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StoragePlanInformation struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates whether there are higher storage quota plans available. Read-only.
	UpgradeAvailable nullable.Type[bool] `json:"upgradeAvailable,omitempty"`
}

var _ json.Marshaler = StoragePlanInformation{}

func (s StoragePlanInformation) MarshalJSON() ([]byte, error) {
	type wrapper StoragePlanInformation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling StoragePlanInformation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling StoragePlanInformation: %+v", err)
	}

	delete(decoded, "upgradeAvailable")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling StoragePlanInformation: %+v", err)
	}

	return encoded, nil
}
