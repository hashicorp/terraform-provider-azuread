package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ StorageQuotaBreakdown = ServiceStorageQuotaBreakdown{}

type ServiceStorageQuotaBreakdown struct {

	// Fields inherited from StorageQuotaBreakdown

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

func (s ServiceStorageQuotaBreakdown) StorageQuotaBreakdown() BaseStorageQuotaBreakdownImpl {
	return BaseStorageQuotaBreakdownImpl{
		DisplayName:  s.DisplayName,
		ManageWebUrl: s.ManageWebUrl,
		Used:         s.Used,
		Id:           s.Id,
		ODataId:      s.ODataId,
		ODataType:    s.ODataType,
	}
}

func (s ServiceStorageQuotaBreakdown) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ServiceStorageQuotaBreakdown{}

func (s ServiceStorageQuotaBreakdown) MarshalJSON() ([]byte, error) {
	type wrapper ServiceStorageQuotaBreakdown
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ServiceStorageQuotaBreakdown: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceStorageQuotaBreakdown: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.serviceStorageQuotaBreakdown"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ServiceStorageQuotaBreakdown: %+v", err)
	}

	return encoded, nil
}
