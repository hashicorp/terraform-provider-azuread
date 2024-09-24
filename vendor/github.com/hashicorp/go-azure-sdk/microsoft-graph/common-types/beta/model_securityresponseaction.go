package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityResponseAction interface {
	SecurityResponseAction() BaseSecurityResponseActionImpl
}

var _ SecurityResponseAction = BaseSecurityResponseActionImpl{}

type BaseSecurityResponseActionImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseSecurityResponseActionImpl) SecurityResponseAction() BaseSecurityResponseActionImpl {
	return s
}

var _ SecurityResponseAction = RawSecurityResponseActionImpl{}

// RawSecurityResponseActionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityResponseActionImpl struct {
	securityResponseAction BaseSecurityResponseActionImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawSecurityResponseActionImpl) SecurityResponseAction() BaseSecurityResponseActionImpl {
	return s.securityResponseAction
}

func UnmarshalSecurityResponseActionImplementation(input []byte) (SecurityResponseAction, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityResponseAction into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.allowFileResponseAction") {
		var out SecurityAllowFileResponseAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAllowFileResponseAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.blockFileResponseAction") {
		var out SecurityBlockFileResponseAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityBlockFileResponseAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.collectInvestigationPackageResponseAction") {
		var out SecurityCollectInvestigationPackageResponseAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCollectInvestigationPackageResponseAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.disableUserResponseAction") {
		var out SecurityDisableUserResponseAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityDisableUserResponseAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.forceUserPasswordResetResponseAction") {
		var out SecurityForceUserPasswordResetResponseAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityForceUserPasswordResetResponseAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.hardDeleteResponseAction") {
		var out SecurityHardDeleteResponseAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityHardDeleteResponseAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.initiateInvestigationResponseAction") {
		var out SecurityInitiateInvestigationResponseAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityInitiateInvestigationResponseAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.isolateDeviceResponseAction") {
		var out SecurityIsolateDeviceResponseAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityIsolateDeviceResponseAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.markUserAsCompromisedResponseAction") {
		var out SecurityMarkUserAsCompromisedResponseAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMarkUserAsCompromisedResponseAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.moveToDeletedItemsResponseAction") {
		var out SecurityMoveToDeletedItemsResponseAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMoveToDeletedItemsResponseAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.moveToInboxResponseAction") {
		var out SecurityMoveToInboxResponseAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMoveToInboxResponseAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.moveToJunkResponseAction") {
		var out SecurityMoveToJunkResponseAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMoveToJunkResponseAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.restrictAppExecutionResponseAction") {
		var out SecurityRestrictAppExecutionResponseAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRestrictAppExecutionResponseAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.runAntivirusScanResponseAction") {
		var out SecurityRunAntivirusScanResponseAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRunAntivirusScanResponseAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.softDeleteResponseAction") {
		var out SecuritySoftDeleteResponseAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecuritySoftDeleteResponseAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.stopAndQuarantineFileResponseAction") {
		var out SecurityStopAndQuarantineFileResponseAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityStopAndQuarantineFileResponseAction: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityResponseActionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityResponseActionImpl: %+v", err)
	}

	return RawSecurityResponseActionImpl{
		securityResponseAction: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
