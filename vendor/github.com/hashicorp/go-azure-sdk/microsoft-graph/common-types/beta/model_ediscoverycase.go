package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = EdiscoveryCase{}

type EdiscoveryCase struct {
	// The user who closed the case.
	ClosedBy IdentitySet `json:"closedBy"`

	// The date and time when the case was closed. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	ClosedDateTime nullable.Type[string] `json:"closedDateTime,omitempty"`

	// The date and time when the entity was created. The Timestamp type represents date and time information using ISO 8601
	// format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Returns a list of case custodian objects for this case. Nullable.
	Custodians *[]EdiscoveryCustodian `json:"custodians,omitempty"`

	// The case description.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The case name.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The external case number for customer reference.
	ExternalId nullable.Type[string] `json:"externalId,omitempty"`

	// The last user who modified the entity.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The latest date and time when the case was modified. The Timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Returns a list of case legalHold objects for this case. Nullable.
	LegalHolds *[]EdiscoveryLegalHold `json:"legalHolds,omitempty"`

	// Returns a list of case noncustodialDataSource objects for this case. Nullable.
	NoncustodialDataSources *[]EdiscoveryNoncustodialDataSource `json:"noncustodialDataSources,omitempty"`

	// Returns a list of case operation objects for this case. Nullable.
	Operations *[]EdiscoveryCaseOperation `json:"operations,omitempty"`

	// Returns a list of reviewSet objects in the case. Read-only. Nullable.
	ReviewSets *[]EdiscoveryReviewSet `json:"reviewSets,omitempty"`

	Settings *EdiscoveryCaseSettings `json:"settings,omitempty"`

	// Returns a list of sourceCollection objects associated with this case.
	SourceCollections *[]EdiscoverySourceCollection `json:"sourceCollections,omitempty"`

	// The case status. Possible values are unknown, active, pendingDelete, closing, closed, and closedWithError. For
	// details, see the following table.
	Status *EdiscoveryCaseStatus `json:"status,omitempty"`

	// Returns a list of tag objects associated to this case.
	Tags *[]EdiscoveryTag `json:"tags,omitempty"`

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

func (s EdiscoveryCase) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = EdiscoveryCase{}

func (s EdiscoveryCase) MarshalJSON() ([]byte, error) {
	type wrapper EdiscoveryCase
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EdiscoveryCase: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EdiscoveryCase: %+v", err)
	}

	delete(decoded, "reviewSets")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.ediscovery.case"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EdiscoveryCase: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &EdiscoveryCase{}

func (s *EdiscoveryCase) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ClosedDateTime          nullable.Type[string]               `json:"closedDateTime,omitempty"`
		CreatedDateTime         nullable.Type[string]               `json:"createdDateTime,omitempty"`
		Custodians              *[]EdiscoveryCustodian              `json:"custodians,omitempty"`
		Description             nullable.Type[string]               `json:"description,omitempty"`
		DisplayName             nullable.Type[string]               `json:"displayName,omitempty"`
		ExternalId              nullable.Type[string]               `json:"externalId,omitempty"`
		LastModifiedDateTime    nullable.Type[string]               `json:"lastModifiedDateTime,omitempty"`
		LegalHolds              *[]EdiscoveryLegalHold              `json:"legalHolds,omitempty"`
		NoncustodialDataSources *[]EdiscoveryNoncustodialDataSource `json:"noncustodialDataSources,omitempty"`
		ReviewSets              *[]EdiscoveryReviewSet              `json:"reviewSets,omitempty"`
		Settings                *EdiscoveryCaseSettings             `json:"settings,omitempty"`
		SourceCollections       *[]EdiscoverySourceCollection       `json:"sourceCollections,omitempty"`
		Status                  *EdiscoveryCaseStatus               `json:"status,omitempty"`
		Tags                    *[]EdiscoveryTag                    `json:"tags,omitempty"`
		Id                      *string                             `json:"id,omitempty"`
		ODataId                 *string                             `json:"@odata.id,omitempty"`
		ODataType               *string                             `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ClosedDateTime = decoded.ClosedDateTime
	s.CreatedDateTime = decoded.CreatedDateTime
	s.Custodians = decoded.Custodians
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.ExternalId = decoded.ExternalId
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.LegalHolds = decoded.LegalHolds
	s.NoncustodialDataSources = decoded.NoncustodialDataSources
	s.ReviewSets = decoded.ReviewSets
	s.Settings = decoded.Settings
	s.SourceCollections = decoded.SourceCollections
	s.Status = decoded.Status
	s.Tags = decoded.Tags
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EdiscoveryCase into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["closedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ClosedBy' for 'EdiscoveryCase': %+v", err)
		}
		s.ClosedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'EdiscoveryCase': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	if v, ok := temp["operations"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Operations into list []json.RawMessage: %+v", err)
		}

		output := make([]EdiscoveryCaseOperation, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalEdiscoveryCaseOperationImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Operations' for 'EdiscoveryCase': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Operations = &output
	}

	return nil
}
