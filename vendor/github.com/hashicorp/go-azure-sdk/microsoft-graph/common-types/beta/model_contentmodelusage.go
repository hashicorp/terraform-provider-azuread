package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ContentModelUsage struct {
	// Identity of the user, device, or application that first applied the contentModel to the library.
	CreatedBy IdentitySet `json:"createdBy"`

	// Date and time of the contentModel is first applied.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// The ID of the drive where the contentModel is applied.
	DriveId nullable.Type[string] `json:"driveId,omitempty"`

	// Identity of the user, device, or application that last applied the contentModel to the library.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// Date and time of the contentModel is last applied.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The ID of the contentModel.
	ModelId nullable.Type[string] `json:"modelId,omitempty"`

	// The version of the current applied contentModel.
	ModelVersion nullable.Type[string] `json:"modelVersion,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`
}

var _ json.Unmarshaler = &ContentModelUsage{}

func (s *ContentModelUsage) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		CreatedDateTime      nullable.Type[string] `json:"createdDateTime,omitempty"`
		DriveId              nullable.Type[string] `json:"driveId,omitempty"`
		LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`
		ModelId              nullable.Type[string] `json:"modelId,omitempty"`
		ModelVersion         nullable.Type[string] `json:"modelVersion,omitempty"`
		ODataId              *string               `json:"@odata.id,omitempty"`
		ODataType            *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.CreatedDateTime = decoded.CreatedDateTime
	s.DriveId = decoded.DriveId
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.ModelId = decoded.ModelId
	s.ModelVersion = decoded.ModelVersion
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ContentModelUsage into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'ContentModelUsage': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'ContentModelUsage': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	return nil
}
