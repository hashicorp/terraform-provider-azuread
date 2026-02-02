package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UnifiedStorageQuota{}

type UnifiedStorageQuota struct {
	Deleted      nullable.Type[int64]            `json:"deleted,omitempty"`
	ManageWebUrl nullable.Type[string]           `json:"manageWebUrl,omitempty"`
	Remaining    nullable.Type[int64]            `json:"remaining,omitempty"`
	Services     *[]ServiceStorageQuotaBreakdown `json:"services,omitempty"`
	State        nullable.Type[string]           `json:"state,omitempty"`
	Total        nullable.Type[int64]            `json:"total,omitempty"`
	Used         nullable.Type[int64]            `json:"used,omitempty"`

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

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.unifiedStorageQuota"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UnifiedStorageQuota: %+v", err)
	}

	return encoded, nil
}
