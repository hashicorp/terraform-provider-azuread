package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ExternalConnectorsExternalItem{}

type ExternalConnectorsExternalItem struct {
	// An array of access control entries. Each entry specifies the access granted to a user or group. Required.
	Acl []ExternalConnectorsAcl `json:"acl"`

	// Returns a list of activities performed on the item. Write-only.
	Activities *[]ExternalConnectorsExternalActivity `json:"activities,omitempty"`

	// A plain-text representation of the contents of the item. The text in this property is full-text indexed. Optional.
	Content *ExternalConnectorsExternalItemContent `json:"content,omitempty"`

	// A property bag with the properties of the item. The properties MUST conform to the schema defined for the
	// externalConnection. Required.
	Properties ExternalConnectorsProperties `json:"properties"`

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

func (s ExternalConnectorsExternalItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ExternalConnectorsExternalItem{}

func (s ExternalConnectorsExternalItem) MarshalJSON() ([]byte, error) {
	type wrapper ExternalConnectorsExternalItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternalConnectorsExternalItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalConnectorsExternalItem: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.externalConnectors.externalItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternalConnectorsExternalItem: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ExternalConnectorsExternalItem{}

func (s *ExternalConnectorsExternalItem) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Acl        []ExternalConnectorsAcl                `json:"acl"`
		Content    *ExternalConnectorsExternalItemContent `json:"content,omitempty"`
		Properties ExternalConnectorsProperties           `json:"properties"`
		Id         *string                                `json:"id,omitempty"`
		ODataId    *string                                `json:"@odata.id,omitempty"`
		ODataType  *string                                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Acl = decoded.Acl
	s.Content = decoded.Content
	s.Properties = decoded.Properties
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ExternalConnectorsExternalItem into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["activities"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Activities into list []json.RawMessage: %+v", err)
		}

		output := make([]ExternalConnectorsExternalActivity, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalExternalConnectorsExternalActivityImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Activities' for 'ExternalConnectorsExternalItem': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Activities = &output
	}

	return nil
}
