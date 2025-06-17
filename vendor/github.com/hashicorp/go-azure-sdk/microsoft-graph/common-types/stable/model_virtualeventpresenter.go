package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = VirtualEventPresenter{}

type VirtualEventPresenter struct {
	// Email address of the presenter.
	Email nullable.Type[string] `json:"email,omitempty"`

	// Identity information of the presenter. The supported identities are: communicationsGuestIdentity and
	// communicationsUserIdentity.
	Identity Identity `json:"identity"`

	// Other details about the presenter. This property returns null when the virtual event type is virtualEventTownhall.
	PresenterDetails *VirtualEventPresenterDetails `json:"presenterDetails,omitempty"`

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

func (s VirtualEventPresenter) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = VirtualEventPresenter{}

func (s VirtualEventPresenter) MarshalJSON() ([]byte, error) {
	type wrapper VirtualEventPresenter
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling VirtualEventPresenter: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling VirtualEventPresenter: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.virtualEventPresenter"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling VirtualEventPresenter: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &VirtualEventPresenter{}

func (s *VirtualEventPresenter) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Email            nullable.Type[string]         `json:"email,omitempty"`
		PresenterDetails *VirtualEventPresenterDetails `json:"presenterDetails,omitempty"`
		Id               *string                       `json:"id,omitempty"`
		ODataId          *string                       `json:"@odata.id,omitempty"`
		ODataType        *string                       `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Email = decoded.Email
	s.PresenterDetails = decoded.PresenterDetails
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling VirtualEventPresenter into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["identity"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Identity' for 'VirtualEventPresenter': %+v", err)
		}
		s.Identity = impl
	}

	return nil
}
