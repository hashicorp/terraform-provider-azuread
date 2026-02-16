package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SubjectRightsRequest{}

type SubjectRightsRequest struct {
	// Collection of users who can approve the request. Currently only supported for requests of type delete.
	Approvers *[]User `json:"approvers,omitempty"`

	// Identity that the request is assigned to.
	AssignedTo Identity `json:"assignedTo"`

	// The date and time when the request was closed. The timestamp type represents date and time information using ISO 8601
	// format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	ClosedDateTime nullable.Type[string] `json:"closedDateTime,omitempty"`

	// Collection of users who can collaborate on the request.
	Collaborators *[]User `json:"collaborators,omitempty"`

	// KQL based content query that should be used for search. This property is defined only for APIs accessed using the
	// /security query path and not the /privacy query path.
	ContentQuery nullable.Type[string] `json:"contentQuery,omitempty"`

	// Identity information for the entity that created the request.
	CreatedBy IdentitySet `json:"createdBy"`

	// The date and time when the request was created. The timestamp type represents date and time information using ISO
	// 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Information about the data subject.
	DataSubject *DataSubject `json:"dataSubject,omitempty"`

	// The type of the data subject. Possible values are: customer, currentEmployee, formerEmployee, prospectiveEmployee,
	// student, teacher, faculty, other, unknownFutureValue.
	DataSubjectType *DataSubjectType `json:"dataSubjectType,omitempty"`

	// Description for the request.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The name of the request.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The external ID for the request that is immutable after creation and is used for tracking the request for the
	// external system. This property is defined only for APIs accessed using the /security query path and not the /privacy
	// query path.
	ExternalId nullable.Type[string] `json:"externalId,omitempty"`

	// Collection of history change events.
	History *[]SubjectRightsRequestHistory `json:"history,omitempty"`

	// Include all versions of the documents. By default, the current copies of the documents are returned. If SharePoint
	// sites have versioning enabled, including all versions includes the historical copies of the documents. This property
	// is defined only for APIs accessed using the /security query path and not the /privacy query path.
	IncludeAllVersions nullable.Type[bool] `json:"includeAllVersions,omitempty"`

	// Include content authored by the data subject. This property is defined only for APIs accessed using the /security
	// query path and not the /privacy query path.
	IncludeAuthoredContent nullable.Type[bool] `json:"includeAuthoredContent,omitempty"`

	// Insight about the request.
	Insight *SubjectRightsRequestDetail `json:"insight,omitempty"`

	// The date and time when the request is internally due. The timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	InternalDueDateTime nullable.Type[string] `json:"internalDueDateTime,omitempty"`

	// Identity information for the entity that last modified the request.
	LastModifiedBy IdentitySet `json:"lastModifiedBy"`

	// The date and time when the request was last modified. The timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// The mailbox locations that should be searched. This property is defined only for APIs accessed using the /security
	// query path and not the /privacy query path.
	MailboxLocations SubjectRightsRequestMailboxLocation `json:"mailboxLocations"`

	// List of notes associated with the request.
	Notes *[]AuthoredNote `json:"notes,omitempty"`

	// Pause the request after estimate has finished. By default, the data estimate runs and then pauses, allowing you to
	// preview results and then select the option to retrieve data in the UI. You can set this property to false if you want
	// it to perform the estimate and then automatically begin with the retrieval of the content. This property is defined
	// only for APIs accessed using the /security query path and not the /privacy query path.
	PauseAfterEstimate nullable.Type[bool] `json:"pauseAfterEstimate,omitempty"`

	// List of regulations that this request fulfill.
	Regulations *[]string `json:"regulations,omitempty"`

	// The SharePoint and OneDrive site locations that should be searched. This property is defined only for APIs accessed
	// using the /security query path and not the /privacy query path.
	SiteLocations SubjectRightsRequestSiteLocation `json:"siteLocations"`

	// Information about the different stages for the request.
	Stages *[]SubjectRightsRequestStageDetail `json:"stages,omitempty"`

	// The status of the request. Possible values are: active, closed, unknownFutureValue.
	Status *SubjectRightsRequestStatus `json:"status,omitempty"`

	// Information about the Microsoft Teams team that was created for the request.
	Team *Team `json:"team,omitempty"`

	// The type of the request. Possible values are: export, delete, access, tagForAction, unknownFutureValue.
	Type *SubjectRightsRequestType `json:"type,omitempty"`

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

func (s SubjectRightsRequest) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SubjectRightsRequest{}

func (s SubjectRightsRequest) MarshalJSON() ([]byte, error) {
	type wrapper SubjectRightsRequest
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SubjectRightsRequest: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SubjectRightsRequest: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.subjectRightsRequest"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SubjectRightsRequest: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &SubjectRightsRequest{}

func (s *SubjectRightsRequest) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Approvers              *[]User                            `json:"approvers,omitempty"`
		ClosedDateTime         nullable.Type[string]              `json:"closedDateTime,omitempty"`
		Collaborators          *[]User                            `json:"collaborators,omitempty"`
		ContentQuery           nullable.Type[string]              `json:"contentQuery,omitempty"`
		CreatedDateTime        nullable.Type[string]              `json:"createdDateTime,omitempty"`
		DataSubject            *DataSubject                       `json:"dataSubject,omitempty"`
		DataSubjectType        *DataSubjectType                   `json:"dataSubjectType,omitempty"`
		Description            nullable.Type[string]              `json:"description,omitempty"`
		DisplayName            nullable.Type[string]              `json:"displayName,omitempty"`
		ExternalId             nullable.Type[string]              `json:"externalId,omitempty"`
		History                *[]SubjectRightsRequestHistory     `json:"history,omitempty"`
		IncludeAllVersions     nullable.Type[bool]                `json:"includeAllVersions,omitempty"`
		IncludeAuthoredContent nullable.Type[bool]                `json:"includeAuthoredContent,omitempty"`
		Insight                *SubjectRightsRequestDetail        `json:"insight,omitempty"`
		InternalDueDateTime    nullable.Type[string]              `json:"internalDueDateTime,omitempty"`
		LastModifiedDateTime   nullable.Type[string]              `json:"lastModifiedDateTime,omitempty"`
		Notes                  *[]AuthoredNote                    `json:"notes,omitempty"`
		PauseAfterEstimate     nullable.Type[bool]                `json:"pauseAfterEstimate,omitempty"`
		Regulations            *[]string                          `json:"regulations,omitempty"`
		Stages                 *[]SubjectRightsRequestStageDetail `json:"stages,omitempty"`
		Status                 *SubjectRightsRequestStatus        `json:"status,omitempty"`
		Team                   *Team                              `json:"team,omitempty"`
		Type                   *SubjectRightsRequestType          `json:"type,omitempty"`
		Id                     *string                            `json:"id,omitempty"`
		ODataId                *string                            `json:"@odata.id,omitempty"`
		ODataType              *string                            `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Approvers = decoded.Approvers
	s.ClosedDateTime = decoded.ClosedDateTime
	s.Collaborators = decoded.Collaborators
	s.ContentQuery = decoded.ContentQuery
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DataSubject = decoded.DataSubject
	s.DataSubjectType = decoded.DataSubjectType
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.ExternalId = decoded.ExternalId
	s.History = decoded.History
	s.IncludeAllVersions = decoded.IncludeAllVersions
	s.IncludeAuthoredContent = decoded.IncludeAuthoredContent
	s.Insight = decoded.Insight
	s.InternalDueDateTime = decoded.InternalDueDateTime
	s.LastModifiedDateTime = decoded.LastModifiedDateTime
	s.Notes = decoded.Notes
	s.PauseAfterEstimate = decoded.PauseAfterEstimate
	s.Regulations = decoded.Regulations
	s.Stages = decoded.Stages
	s.Status = decoded.Status
	s.Team = decoded.Team
	s.Type = decoded.Type
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SubjectRightsRequest into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["assignedTo"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AssignedTo' for 'SubjectRightsRequest': %+v", err)
		}
		s.AssignedTo = impl
	}

	if v, ok := temp["createdBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'CreatedBy' for 'SubjectRightsRequest': %+v", err)
		}
		s.CreatedBy = impl
	}

	if v, ok := temp["lastModifiedBy"]; ok {
		impl, err := UnmarshalIdentitySetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'LastModifiedBy' for 'SubjectRightsRequest': %+v", err)
		}
		s.LastModifiedBy = impl
	}

	if v, ok := temp["mailboxLocations"]; ok {
		impl, err := UnmarshalSubjectRightsRequestMailboxLocationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'MailboxLocations' for 'SubjectRightsRequest': %+v", err)
		}
		s.MailboxLocations = impl
	}

	if v, ok := temp["siteLocations"]; ok {
		impl, err := UnmarshalSubjectRightsRequestSiteLocationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'SiteLocations' for 'SubjectRightsRequest': %+v", err)
		}
		s.SiteLocations = impl
	}

	return nil
}
