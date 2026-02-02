package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCTenantEncryptionSetting struct {
	// Indicates the date and time when last sync tenant encryption setting.
	LastSyncDateTime nullable.Type[string] `json:"lastSyncDateTime,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Indicates the Cloud PC disk encryption type for a tenant. It is a tenant-level setting that applies globally to all
	// Cloud PCs in the tenant. Possible values are: platformManagedKey, customerManagedKey, unknownFutureValue. Read-only.
	TenantDiskEncryptionType *CloudPCDiskEncryptionType `json:"tenantDiskEncryptionType,omitempty"`
}

var _ json.Marshaler = CloudPCTenantEncryptionSetting{}

func (s CloudPCTenantEncryptionSetting) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCTenantEncryptionSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCTenantEncryptionSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCTenantEncryptionSetting: %+v", err)
	}

	delete(decoded, "tenantDiskEncryptionType")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCTenantEncryptionSetting: %+v", err)
	}

	return encoded, nil
}
