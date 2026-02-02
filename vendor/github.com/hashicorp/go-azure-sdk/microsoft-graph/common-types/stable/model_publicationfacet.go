package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PublicationFacet struct {
	// The user who checked out the file.
	CheckedOutBy IdentitySet `json:"checkedOutBy"`

	// The state of publication for this document. Either published or checkout. Read-only.
	Level nullable.Type[string] `json:"level,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The unique identifier for the version that is visible to the current caller. Read-only.
	VersionId nullable.Type[string] `json:"versionId,omitempty"`
}

var _ json.Marshaler = PublicationFacet{}

func (s PublicationFacet) MarshalJSON() ([]byte, error) {
	type wrapper PublicationFacet
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PublicationFacet: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PublicationFacet: %+v", err)
	}

	delete(decoded, "level")
	delete(decoded, "versionId")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PublicationFacet: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PublicationFacet{}

func (s *PublicationFacet) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Level     nullable.Type[string] `json:"level,omitempty"`
		ODataId   *string               `json:"@odata.id,omitempty"`
		ODataType *string               `json:"@odata.type,omitempty"`
		VersionId nullable.Type[string] `json:"versionId,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Level = decoded.Level
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.VersionId = decoded.VersionId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PublicationFacet into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["checkedOutBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CheckedOutBy' for 'PublicationFacet': %+v", err)
		}
		s.CheckedOutBy = impl
	}

	return nil
}
