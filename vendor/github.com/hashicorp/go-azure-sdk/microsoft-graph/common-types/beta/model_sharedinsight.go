package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SharedInsight{}

type SharedInsight struct {
	// Details about the shared item. Read only.
	LastShared *SharingDetail `json:"lastShared,omitempty"`

	LastSharedMethod Entity `json:"lastSharedMethod"`

	// Used for navigating to the item that was shared. For file attachments, the type is fileAttachment. For linked
	// attachments, the type is driveItem.
	Resource Entity `json:"resource"`

	// Reference properties of the shared document, such as the url and type of the document. Read-only
	ResourceReference *ResourceReference `json:"resourceReference,omitempty"`

	// Properties that you can use to visualize the document in your experience. Read-only
	ResourceVisualization *ResourceVisualization `json:"resourceVisualization,omitempty"`

	SharingHistory *[]SharingDetail `json:"sharingHistory,omitempty"`

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

func (s SharedInsight) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SharedInsight{}

func (s SharedInsight) MarshalJSON() ([]byte, error) {
	type wrapper SharedInsight
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SharedInsight: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SharedInsight: %+v", err)
	}

	delete(decoded, "resourceReference")
	delete(decoded, "resourceVisualization")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.sharedInsight"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SharedInsight: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SharedInsight{}

func (s *SharedInsight) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		LastShared            *SharingDetail         `json:"lastShared,omitempty"`
		ResourceReference     *ResourceReference     `json:"resourceReference,omitempty"`
		ResourceVisualization *ResourceVisualization `json:"resourceVisualization,omitempty"`
		SharingHistory        *[]SharingDetail       `json:"sharingHistory,omitempty"`
		Id                    *string                `json:"id,omitempty"`
		ODataId               *string                `json:"@odata.id,omitempty"`
		ODataType             *string                `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.LastShared = decoded.LastShared
	s.ResourceReference = decoded.ResourceReference
	s.ResourceVisualization = decoded.ResourceVisualization
	s.SharingHistory = decoded.SharingHistory
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SharedInsight into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["lastSharedMethod"]; ok {
		impl, err := UnmarshalEntityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastSharedMethod' for 'SharedInsight': %+v", err)
		}
		s.LastSharedMethod = impl
	}

	if v, ok := temp["resource"]; ok {
		impl, err := UnmarshalEntityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Resource' for 'SharedInsight': %+v", err)
		}
		s.Resource = impl
	}

	return nil
}
