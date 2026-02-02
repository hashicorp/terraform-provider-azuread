package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Finding = EncryptedAwsStorageBucketFinding{}

type EncryptedAwsStorageBucketFinding struct {
	Accessibility *AwsAccessType               `json:"accessibility,omitempty"`
	StorageBucket *AuthorizationSystemResource `json:"storageBucket,omitempty"`

	// Fields inherited from Finding

	// Defines when the finding was created.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

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

func (s EncryptedAwsStorageBucketFinding) Finding() BaseFindingImpl {
	return BaseFindingImpl{
		CreatedDateTime: s.CreatedDateTime,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s EncryptedAwsStorageBucketFinding) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EncryptedAwsStorageBucketFinding{}

func (s EncryptedAwsStorageBucketFinding) MarshalJSON() ([]byte, error) {
	type wrapper EncryptedAwsStorageBucketFinding
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EncryptedAwsStorageBucketFinding: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EncryptedAwsStorageBucketFinding: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.encryptedAwsStorageBucketFinding"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EncryptedAwsStorageBucketFinding: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EncryptedAwsStorageBucketFinding{}

func (s *EncryptedAwsStorageBucketFinding) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Accessibility   *AwsAccessType `json:"accessibility,omitempty"`
		CreatedDateTime *string        `json:"createdDateTime,omitempty"`
		Id              *string        `json:"id,omitempty"`
		ODataId         *string        `json:"@odata.id,omitempty"`
		ODataType       *string        `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Accessibility = decoded.Accessibility
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EncryptedAwsStorageBucketFinding into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["storageBucket"]; ok {
		impl, err := UnmarshalAuthorizationSystemResourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'StorageBucket' for 'EncryptedAwsStorageBucketFinding': %+v", err)
		}
		s.StorageBucket = &impl
	}

	return nil
}
