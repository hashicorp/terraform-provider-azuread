package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = UsedInsight{}

type UsedInsight struct {
	// Information about when the item was last viewed or modified by the user. Read only.
	LastUsed *UsageDetails `json:"lastUsed,omitempty"`

	// Used for navigating to the item that was used. For file attachments, the type is fileAttachment. For linked
	// attachments, the type is driveItem.
	Resource Entity `json:"resource"`

	// Reference properties of the used document, such as the url and type of the document. Read-only
	ResourceReference *ResourceReference `json:"resourceReference,omitempty"`

	// Properties that you can use to visualize the document in your experience. Read-only
	ResourceVisualization *ResourceVisualization `json:"resourceVisualization,omitempty"`

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

func (s UsedInsight) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UsedInsight{}

func (s UsedInsight) MarshalJSON() ([]byte, error) {
	type wrapper UsedInsight
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UsedInsight: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UsedInsight: %+v", err)
	}

	delete(decoded, "resourceReference")
	delete(decoded, "resourceVisualization")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.usedInsight"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UsedInsight: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &UsedInsight{}

func (s *UsedInsight) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		LastUsed              *UsageDetails          `json:"lastUsed,omitempty"`
		ResourceReference     *ResourceReference     `json:"resourceReference,omitempty"`
		ResourceVisualization *ResourceVisualization `json:"resourceVisualization,omitempty"`
		Id                    *string                `json:"id,omitempty"`
		ODataId               *string                `json:"@odata.id,omitempty"`
		ODataType             *string                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.LastUsed = decoded.LastUsed
	s.ResourceReference = decoded.ResourceReference
	s.ResourceVisualization = decoded.ResourceVisualization
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling UsedInsight into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["resource"]; ok {
		impl, err := UnmarshalEntityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Resource' for 'UsedInsight': %+v", err)
		}
		s.Resource = impl
	}

	return nil
}
