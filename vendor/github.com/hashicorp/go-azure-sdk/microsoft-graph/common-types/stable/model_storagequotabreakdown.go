package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageQuotaBreakdown interface {
	Entity
	StorageQuotaBreakdown() BaseStorageQuotaBreakdownImpl
}

var _ StorageQuotaBreakdown = BaseStorageQuotaBreakdownImpl{}

type BaseStorageQuotaBreakdownImpl struct {
	DisplayName  nullable.Type[string] `json:"displayName,omitempty"`
	ManageWebUrl nullable.Type[string] `json:"manageWebUrl,omitempty"`
	Used         nullable.Type[int64]  `json:"used,omitempty"`

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

func (s BaseStorageQuotaBreakdownImpl) StorageQuotaBreakdown() BaseStorageQuotaBreakdownImpl {
	return s
}

func (s BaseStorageQuotaBreakdownImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ StorageQuotaBreakdown = RawStorageQuotaBreakdownImpl{}

// RawStorageQuotaBreakdownImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawStorageQuotaBreakdownImpl struct {
	storageQuotaBreakdown BaseStorageQuotaBreakdownImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawStorageQuotaBreakdownImpl) StorageQuotaBreakdown() BaseStorageQuotaBreakdownImpl {
	return s.storageQuotaBreakdown
}

func (s RawStorageQuotaBreakdownImpl) Entity() BaseEntityImpl {
	return s.storageQuotaBreakdown.Entity()
}

var _ json.Marshaler = BaseStorageQuotaBreakdownImpl{}

func (s BaseStorageQuotaBreakdownImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseStorageQuotaBreakdownImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseStorageQuotaBreakdownImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseStorageQuotaBreakdownImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.storageQuotaBreakdown"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseStorageQuotaBreakdownImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalStorageQuotaBreakdownImplementation(input []byte) (StorageQuotaBreakdown, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling StorageQuotaBreakdown into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.serviceStorageQuotaBreakdown") {
		var out ServiceStorageQuotaBreakdown
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ServiceStorageQuotaBreakdown: %+v", err)
		}
		return out, nil
	}

	var parent BaseStorageQuotaBreakdownImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseStorageQuotaBreakdownImpl: %+v", err)
	}

	return RawStorageQuotaBreakdownImpl{
		storageQuotaBreakdown: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
