package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityInformationProtectionAction interface {
	SecurityInformationProtectionAction() BaseSecurityInformationProtectionActionImpl
}

var _ SecurityInformationProtectionAction = BaseSecurityInformationProtectionActionImpl{}

type BaseSecurityInformationProtectionActionImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseSecurityInformationProtectionActionImpl) SecurityInformationProtectionAction() BaseSecurityInformationProtectionActionImpl {
	return s
}

var _ SecurityInformationProtectionAction = RawSecurityInformationProtectionActionImpl{}

// RawSecurityInformationProtectionActionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSecurityInformationProtectionActionImpl struct {
	securityInformationProtectionAction BaseSecurityInformationProtectionActionImpl
	Type                                string
	Values                              map[string]interface{}
}

func (s RawSecurityInformationProtectionActionImpl) SecurityInformationProtectionAction() BaseSecurityInformationProtectionActionImpl {
	return s.securityInformationProtectionAction
}

func UnmarshalSecurityInformationProtectionActionImplementation(input []byte) (SecurityInformationProtectionAction, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SecurityInformationProtectionAction into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.security.addContentFooterAction") {
		var out SecurityAddContentFooterAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAddContentFooterAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.addContentHeaderAction") {
		var out SecurityAddContentHeaderAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAddContentHeaderAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.addWatermarkAction") {
		var out SecurityAddWatermarkAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityAddWatermarkAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.applyLabelAction") {
		var out SecurityApplyLabelAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityApplyLabelAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.customAction") {
		var out SecurityCustomAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityCustomAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.justifyAction") {
		var out SecurityJustifyAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityJustifyAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.metadataAction") {
		var out SecurityMetadataAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityMetadataAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.protectAdhocAction") {
		var out SecurityProtectAdhocAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityProtectAdhocAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.protectByTemplateAction") {
		var out SecurityProtectByTemplateAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityProtectByTemplateAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.protectDoNotForwardAction") {
		var out SecurityProtectDoNotForwardAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityProtectDoNotForwardAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.recommendLabelAction") {
		var out SecurityRecommendLabelAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRecommendLabelAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.removeContentFooterAction") {
		var out SecurityRemoveContentFooterAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRemoveContentFooterAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.removeContentHeaderAction") {
		var out SecurityRemoveContentHeaderAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRemoveContentHeaderAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.removeProtectionAction") {
		var out SecurityRemoveProtectionAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRemoveProtectionAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.security.removeWatermarkAction") {
		var out SecurityRemoveWatermarkAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SecurityRemoveWatermarkAction: %+v", err)
		}
		return out, nil
	}

	var parent BaseSecurityInformationProtectionActionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSecurityInformationProtectionActionImpl: %+v", err)
	}

	return RawSecurityInformationProtectionActionImpl{
		securityInformationProtectionAction: parent,
		Type:                                value,
		Values:                              temp,
	}, nil

}
