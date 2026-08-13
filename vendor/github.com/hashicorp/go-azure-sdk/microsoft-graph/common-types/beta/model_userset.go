package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2023, 2026 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserSet interface {
	UserSet() BaseUserSetImpl
}

var _ UserSet = BaseUserSetImpl{}

type BaseUserSetImpl struct {
	// For a user in an approval stage, this property indicates whether the user is a backup fallback approver.
	IsBackup nullable.Type[bool] `json:"isBackup,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseUserSetImpl) UserSet() BaseUserSetImpl {
	return s
}

var _ UserSet = RawUserSetImpl{}

// RawUserSetImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawUserSetImpl struct {
	userSet BaseUserSetImpl
	Type    string
	Values  map[string]interface{}
}

func (s RawUserSetImpl) UserSet() BaseUserSetImpl {
	return s.userSet
}

func (s RawUserSetImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalUserSetImplementation(input []byte) (UserSet, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling UserSet into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.connectedOrganizationMembers") {
		var out ConnectedOrganizationMembers
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectedOrganizationMembers: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.externalSponsors") {
		var out ExternalSponsors
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExternalSponsors: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.groupMembers") {
		var out GroupMembers
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GroupMembers: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.internalSponsors") {
		var out InternalSponsors
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InternalSponsors: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.requestorManager") {
		var out RequestorManager
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RequestorManager: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.singleUser") {
		var out SingleUser
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SingleUser: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.targetUserSponsors") {
		var out TargetUserSponsors
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TargetUserSponsors: %+v", err)
		}
		return out, nil
	}

	var parent BaseUserSetImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseUserSetImpl: %+v", err)
	}

	return RawUserSetImpl{
		userSet: parent,
		Type:    value,
		Values:  temp,
	}, nil

}
