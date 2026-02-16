package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessPackageResourceAttribute struct {
	// Information about how to set the attribute, currently a accessPackageUserDirectoryAttributeStore type.
	Destination AccessPackageResourceAttributeDestination `json:"destination"`

	IsEditable                     nullable.Type[bool] `json:"isEditable,omitempty"`
	IsPersistedOnAssignmentRemoval nullable.Type[bool] `json:"isPersistedOnAssignmentRemoval,omitempty"`

	// The name of the attribute in the end system. If the destination is accessPackageUserDirectoryAttributeStore, then a
	// user property such as jobTitle or a directory schema extension for the user object type, such as
	// extension2b676109c7c74ae2b41549205f1947edpersonalTitle.
	Name nullable.Type[string] `json:"name,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Information about how to populate the attribute value when an accessPackageAssignmentRequest is being fulfilled,
	// currently a accessPackageResourceAttributeQuestion type.
	Source AccessPackageResourceAttributeSource `json:"source"`
}

var _ json.Unmarshaler = &AccessPackageResourceAttribute{}

func (s *AccessPackageResourceAttribute) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		IsEditable                     nullable.Type[bool]   `json:"isEditable,omitempty"`
		IsPersistedOnAssignmentRemoval nullable.Type[bool]   `json:"isPersistedOnAssignmentRemoval,omitempty"`
		Name                           nullable.Type[string] `json:"name,omitempty"`
		ODataId                        *string               `json:"@odata.id,omitempty"`
		ODataType                      *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.IsEditable = decoded.IsEditable
	s.IsPersistedOnAssignmentRemoval = decoded.IsPersistedOnAssignmentRemoval
	s.Name = decoded.Name
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessPackageResourceAttribute into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["destination"]; ok {
		impl, err := UnmarshalAccessPackageResourceAttributeDestinationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Destination' for 'AccessPackageResourceAttribute': %+v", err)
		}
		s.Destination = impl
	}

	if v, ok := temp["source"]; ok {
		impl, err := UnmarshalAccessPackageResourceAttributeSourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Source' for 'AccessPackageResourceAttribute': %+v", err)
		}
		s.Source = impl
	}

	return nil
}
