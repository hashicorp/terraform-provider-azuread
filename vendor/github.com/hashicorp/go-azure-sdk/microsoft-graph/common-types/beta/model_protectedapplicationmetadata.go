package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IntegratedApplicationMetadata = ProtectedApplicationMetadata{}

type ProtectedApplicationMetadata struct {
	// The Entra client (application) ID. Required.
	ApplicationLocation PolicyLocation `json:"applicationLocation"`

	// Fields inherited from IntegratedApplicationMetadata

	// The name of the integrated application.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The version number of the integrated application.
	Version nullable.Type[string] `json:"version,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ProtectedApplicationMetadata) IntegratedApplicationMetadata() BaseIntegratedApplicationMetadataImpl {
	return BaseIntegratedApplicationMetadataImpl{
		Name:      s.Name,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		Version:   s.Version,
	}
}

var _ json.Marshaler = ProtectedApplicationMetadata{}

func (s ProtectedApplicationMetadata) MarshalJSON() ([]byte, error) {
	type wrapper ProtectedApplicationMetadata
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ProtectedApplicationMetadata: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ProtectedApplicationMetadata: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.protectedApplicationMetadata"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ProtectedApplicationMetadata: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ProtectedApplicationMetadata{}

func (s *ProtectedApplicationMetadata) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Name      nullable.Type[string] `json:"name,omitempty"`
		ODataId   *string               `json:"@odata.id,omitempty"`
		ODataType *string               `json:"@odata.type,omitempty"`
		Version   nullable.Type[string] `json:"version,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.Version = decoded.Version

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ProtectedApplicationMetadata into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["applicationLocation"]; ok {
		impl, err := UnmarshalPolicyLocationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ApplicationLocation' for 'ProtectedApplicationMetadata': %+v", err)
		}
		s.ApplicationLocation = impl
	}

	return nil
}
