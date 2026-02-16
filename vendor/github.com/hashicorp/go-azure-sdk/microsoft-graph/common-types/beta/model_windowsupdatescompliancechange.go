package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdatesComplianceChange interface {
	Entity
	WindowsUpdatesComplianceChange() BaseWindowsUpdatesComplianceChangeImpl
}

var _ WindowsUpdatesComplianceChange = BaseWindowsUpdatesComplianceChangeImpl{}

type BaseWindowsUpdatesComplianceChangeImpl struct {
	// The date and time when a compliance change was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// True indicates that a compliance change is revoked, preventing further application. Revoking a compliance change is a
	// final action.
	IsRevoked nullable.Type[bool] `json:"isRevoked,omitempty"`

	// The date and time when the compliance change was revoked.
	RevokedDateTime nullable.Type[string] `json:"revokedDateTime,omitempty"`

	// The policy this compliance change is a member of.
	UpdatePolicy *WindowsUpdatesUpdatePolicy `json:"updatePolicy,omitempty"`

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

func (s BaseWindowsUpdatesComplianceChangeImpl) WindowsUpdatesComplianceChange() BaseWindowsUpdatesComplianceChangeImpl {
	return s
}

func (s BaseWindowsUpdatesComplianceChangeImpl) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ WindowsUpdatesComplianceChange = RawWindowsUpdatesComplianceChangeImpl{}

// RawWindowsUpdatesComplianceChangeImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawWindowsUpdatesComplianceChangeImpl struct {
	windowsUpdatesComplianceChange BaseWindowsUpdatesComplianceChangeImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawWindowsUpdatesComplianceChangeImpl) WindowsUpdatesComplianceChange() BaseWindowsUpdatesComplianceChangeImpl {
	return s.windowsUpdatesComplianceChange
}

func (s RawWindowsUpdatesComplianceChangeImpl) Entity() BaseEntityImpl {
	return s.windowsUpdatesComplianceChange.Entity()
}

var _ json.Marshaler = BaseWindowsUpdatesComplianceChangeImpl{}

func (s BaseWindowsUpdatesComplianceChangeImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseWindowsUpdatesComplianceChangeImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseWindowsUpdatesComplianceChangeImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseWindowsUpdatesComplianceChangeImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.windowsUpdates.complianceChange"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseWindowsUpdatesComplianceChangeImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalWindowsUpdatesComplianceChangeImplementation(input []byte) (WindowsUpdatesComplianceChange, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling WindowsUpdatesComplianceChange into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.windowsUpdates.contentApproval") {
		var out WindowsUpdatesContentApproval
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WindowsUpdatesContentApproval: %+v", err)
		}
		return out, nil
	}

	var parent BaseWindowsUpdatesComplianceChangeImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseWindowsUpdatesComplianceChangeImpl: %+v", err)
	}

	return RawWindowsUpdatesComplianceChangeImpl{
		windowsUpdatesComplianceChange: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
