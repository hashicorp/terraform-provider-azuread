package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = NetworkaccessAlert{}

type NetworkaccessAlert struct {
	Actions             *[]NetworkaccessAlertAction     `json:"actions,omitempty"`
	AlertType           *NetworkaccessAlertType         `json:"alertType,omitempty"`
	CreationDateTime    *string                         `json:"creationDateTime,omitempty"`
	Description         nullable.Type[string]           `json:"description,omitempty"`
	DetectionTechnology nullable.Type[string]           `json:"detectionTechnology,omitempty"`
	DisplayName         *string                         `json:"displayName,omitempty"`
	Policy              *NetworkaccessFilteringPolicy   `json:"policy,omitempty"`
	RelatedResources    *[]NetworkaccessRelatedResource `json:"relatedResources,omitempty"`
	Severity            *NetworkaccessAlertSeverity     `json:"severity,omitempty"`
	VendorName          *string                         `json:"vendorName,omitempty"`

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

func (s NetworkaccessAlert) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = NetworkaccessAlert{}

func (s NetworkaccessAlert) MarshalJSON() ([]byte, error) {
	type wrapper NetworkaccessAlert
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NetworkaccessAlert: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NetworkaccessAlert: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.networkaccess.alert"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NetworkaccessAlert: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &NetworkaccessAlert{}

func (s *NetworkaccessAlert) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Actions             *[]NetworkaccessAlertAction   `json:"actions,omitempty"`
		AlertType           *NetworkaccessAlertType       `json:"alertType,omitempty"`
		CreationDateTime    *string                       `json:"creationDateTime,omitempty"`
		Description         nullable.Type[string]         `json:"description,omitempty"`
		DetectionTechnology nullable.Type[string]         `json:"detectionTechnology,omitempty"`
		DisplayName         *string                       `json:"displayName,omitempty"`
		Policy              *NetworkaccessFilteringPolicy `json:"policy,omitempty"`
		Severity            *NetworkaccessAlertSeverity   `json:"severity,omitempty"`
		VendorName          *string                       `json:"vendorName,omitempty"`
		Id                  *string                       `json:"id,omitempty"`
		ODataId             *string                       `json:"@odata.id,omitempty"`
		ODataType           *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Actions = decoded.Actions
	s.AlertType = decoded.AlertType
	s.CreationDateTime = decoded.CreationDateTime
	s.Description = decoded.Description
	s.DetectionTechnology = decoded.DetectionTechnology
	s.DisplayName = decoded.DisplayName
	s.Policy = decoded.Policy
	s.Severity = decoded.Severity
	s.VendorName = decoded.VendorName
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling NetworkaccessAlert into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["relatedResources"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RelatedResources into list []json.RawMessage: %+v", err)
		}

		output := make([]NetworkaccessRelatedResource, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalNetworkaccessRelatedResourceImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RelatedResources' for 'NetworkaccessAlert': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RelatedResources = &output
	}

	return nil
}
