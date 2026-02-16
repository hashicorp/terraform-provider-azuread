package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = DelegatedAdminServiceManagementDetail{}

type DelegatedAdminServiceManagementDetail struct {
	// The URL of the management portal for the managed service. Read-only.
	ServiceManagementUrl *string `json:"serviceManagementUrl,omitempty"`

	// The name of a managed service. Read-only.
	ServiceName *string `json:"serviceName,omitempty"`

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

func (s DelegatedAdminServiceManagementDetail) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DelegatedAdminServiceManagementDetail{}

func (s DelegatedAdminServiceManagementDetail) MarshalJSON() ([]byte, error) {
	type wrapper DelegatedAdminServiceManagementDetail
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DelegatedAdminServiceManagementDetail: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DelegatedAdminServiceManagementDetail: %+v", err)
	}

	delete(decoded, "serviceManagementUrl")
	delete(decoded, "serviceName")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.delegatedAdminServiceManagementDetail"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DelegatedAdminServiceManagementDetail: %+v", err)
	}

	return encoded, nil
}
