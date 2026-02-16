package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCForensicStorageAccount{}

type CloudPCForensicStorageAccount struct {
	// Indicates the access tier of the storage account. Possible values are hot, cool, premium, cold, and
	// unknownFutureValue. Default value is hot. Read-only.
	AccessTier *CloudPCStorageAccountAccessTier `json:"accessTier,omitempty"`

	// Indicates whether immutability policies are configured for the storage account. When true, the storage account only
	// accepts hot as the snapshot access tier. When false, the storage account accepts all valid access tiers. Read-Only.
	ImmutableStorage nullable.Type[bool] `json:"immutableStorage,omitempty"`

	// Indicates the ID of the storage account. Read-only.
	StorageAccountId nullable.Type[string] `json:"storageAccountId,omitempty"`

	// Indicates the name of the storage account. Read-only.
	StorageAccountName nullable.Type[string] `json:"storageAccountName,omitempty"`

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

func (s CloudPCForensicStorageAccount) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCForensicStorageAccount{}

func (s CloudPCForensicStorageAccount) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCForensicStorageAccount
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCForensicStorageAccount: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCForensicStorageAccount: %+v", err)
	}

	delete(decoded, "accessTier")
	delete(decoded, "storageAccountId")
	delete(decoded, "storageAccountName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcForensicStorageAccount"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCForensicStorageAccount: %+v", err)
	}

	return encoded, nil
}
