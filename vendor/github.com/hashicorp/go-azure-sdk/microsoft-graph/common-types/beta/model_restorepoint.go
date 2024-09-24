package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = RestorePoint{}

type RestorePoint struct {
	// Expiration date time of the restore point.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// Date time when the restore point was created.
	ProtectionDateTime nullable.Type[string] `json:"protectionDateTime,omitempty"`

	// The site, drive, or mailbox units that are protected under a protection policy.
	ProtectionUnit *ProtectionUnitBase `json:"protectionUnit,omitempty"`

	// The type of the restore point. The possible values are: none, fastRestore, unknownFutureValue.
	Tags *RestorePointTags `json:"tags,omitempty"`

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

func (s RestorePoint) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = RestorePoint{}

func (s RestorePoint) MarshalJSON() ([]byte, error) {
	type wrapper RestorePoint
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RestorePoint: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RestorePoint: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.restorePoint"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RestorePoint: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &RestorePoint{}

func (s *RestorePoint) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`
		ProtectionDateTime nullable.Type[string] `json:"protectionDateTime,omitempty"`
		Tags               *RestorePointTags     `json:"tags,omitempty"`
		Id                 *string               `json:"id,omitempty"`
		ODataId            *string               `json:"@odata.id,omitempty"`
		ODataType          *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ExpirationDateTime = decoded.ExpirationDateTime
	s.ProtectionDateTime = decoded.ProtectionDateTime
	s.Tags = decoded.Tags
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling RestorePoint into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["protectionUnit"]; ok {
		impl, err := UnmarshalProtectionUnitBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProtectionUnit' for 'RestorePoint': %+v", err)
		}
		s.ProtectionUnit = &impl
	}

	return nil
}
