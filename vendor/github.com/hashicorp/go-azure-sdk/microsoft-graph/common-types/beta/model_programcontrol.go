package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ProgramControl{}

type ProgramControl struct {
	// The controlId of the control, in particular the identifier of an access review. Required on create.
	ControlId *string `json:"controlId,omitempty"`

	// The programControlType identifies the type of program control - for example, a control linking to guest access
	// reviews. Required on create.
	ControlTypeId *string `json:"controlTypeId,omitempty"`

	// The creation date and time of the program control.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The name of the control.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// The user who created the program control.
	Owner *UserIdentity `json:"owner,omitempty"`

	// The program this control is part of.
	Program *Program `json:"program,omitempty"`

	// The programId of the program this control is a part of. Required on create.
	ProgramId *string `json:"programId,omitempty"`

	// The resource, a group or an app, targeted by this program control's access review.
	Resource *ProgramResource `json:"resource,omitempty"`

	// The life cycle status of the control.
	Status nullable.Type[string] `json:"status,omitempty"`

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

func (s ProgramControl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ProgramControl{}

func (s ProgramControl) MarshalJSON() ([]byte, error) {
	type wrapper ProgramControl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ProgramControl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ProgramControl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.programControl"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ProgramControl: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ProgramControl{}

func (s *ProgramControl) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ControlId       *string               `json:"controlId,omitempty"`
		ControlTypeId   *string               `json:"controlTypeId,omitempty"`
		CreatedDateTime *string               `json:"createdDateTime,omitempty"`
		DisplayName     nullable.Type[string] `json:"displayName,omitempty"`
		Program         *Program              `json:"program,omitempty"`
		ProgramId       *string               `json:"programId,omitempty"`
		Resource        *ProgramResource      `json:"resource,omitempty"`
		Status          nullable.Type[string] `json:"status,omitempty"`
		Id              *string               `json:"id,omitempty"`
		ODataId         *string               `json:"@odata.id,omitempty"`
		ODataType       *string               `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ControlId = decoded.ControlId
	s.ControlTypeId = decoded.ControlTypeId
	s.CreatedDateTime = decoded.CreatedDateTime
	s.DisplayName = decoded.DisplayName
	s.Program = decoded.Program
	s.ProgramId = decoded.ProgramId
	s.Resource = decoded.Resource
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ProgramControl into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["owner"]; ok {
		impl, err := UnmarshalUserIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Owner' for 'ProgramControl': %+v", err)
		}
		s.Owner = &impl
	}

	return nil
}
