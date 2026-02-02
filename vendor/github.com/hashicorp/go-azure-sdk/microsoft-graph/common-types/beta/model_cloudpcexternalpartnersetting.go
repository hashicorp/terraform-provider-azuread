package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = CloudPCExternalPartnerSetting{}

type CloudPCExternalPartnerSetting struct {
	// Enable or disable the connection to an external partner. If true, an external partner API will accept incoming calls
	// from external partners. Required. Supports $filter (eq).
	EnableConnection bool `json:"enableConnection"`

	// Last data sync time for this external partner. The Timestamp type represents the date and time information using ISO
	// 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 looks like this:
	// '2014-01-01T00:00:00Z'.
	LastSyncDateTime nullable.Type[string] `json:"lastSyncDateTime,omitempty"`

	// The external partner ID.
	PartnerId *string `json:"partnerId,omitempty"`

	Status *CloudPCExternalPartnerStatus `json:"status,omitempty"`

	// Status details message.
	StatusDetails nullable.Type[string] `json:"statusDetails,omitempty"`

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

func (s CloudPCExternalPartnerSetting) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = CloudPCExternalPartnerSetting{}

func (s CloudPCExternalPartnerSetting) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCExternalPartnerSetting
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCExternalPartnerSetting: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCExternalPartnerSetting: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.cloudPcExternalPartnerSetting"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCExternalPartnerSetting: %+v", err)
	}

	return encoded, nil
}
