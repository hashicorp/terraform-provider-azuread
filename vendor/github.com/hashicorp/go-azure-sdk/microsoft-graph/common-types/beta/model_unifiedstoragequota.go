package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UnifiedStorageQuota{}

type UnifiedStorageQuota struct {
	Deleted nullable.Type[int64] `json:"deleted,omitempty"`

	// A URL that can be used in a browser to manage the breakdown. Read-only.
	ManageWebUrl nullable.Type[string] `json:"manageWebUrl,omitempty"`

	// Total space remaining before reaching the quota limit in bytes.
	Remaining nullable.Type[int64] `json:"remaining,omitempty"`

	// The breakdown of services contributing to the user's quota usage.
	Services *[]ServiceStorageQuotaBreakdown `json:"services,omitempty"`

	// Indicates the state of the storage space. The possible values are: normal, nearing, critical, full, and overLimit.
	State nullable.Type[string] `json:"state,omitempty"`

	// Total allowed storage space in bytes.
	Total nullable.Type[int64] `json:"total,omitempty"`

	// Total space used in bytes.
	Used nullable.Type[int64] `json:"used,omitempty"`

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

func (s UnifiedStorageQuota) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UnifiedStorageQuota{}

func (s UnifiedStorageQuota) MarshalJSON() ([]byte, error) {
	type wrapper UnifiedStorageQuota
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UnifiedStorageQuota: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UnifiedStorageQuota: %+v", err)
	}

	delete(decoded, "manageWebUrl")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedStorageQuota"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedStorageQuota: %+v", err)
	}

	return encoded, nil
}
