package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DlpActionInfo interface {
	DlpActionInfo() BaseDlpActionInfoImpl
}

var _ DlpActionInfo = BaseDlpActionInfoImpl{}

type BaseDlpActionInfoImpl struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDlpActionInfoImpl) DlpActionInfo() BaseDlpActionInfoImpl {
	return s
}

var _ DlpActionInfo = RawDlpActionInfoImpl{}

// RawDlpActionInfoImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDlpActionInfoImpl struct {
	dlpActionInfo BaseDlpActionInfoImpl
	Type          string
	Values        map[string]interface{}
}

func (s RawDlpActionInfoImpl) DlpActionInfo() BaseDlpActionInfoImpl {
	return s.dlpActionInfo
}

func UnmarshalDlpActionInfoImplementation(input []byte) (DlpActionInfo, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DlpActionInfo into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.blockAccessAction") {
		var out BlockAccessAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BlockAccessAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "#microsoft.graph.notifyUserAction") {
		var out NotifyUserAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NotifyUserAction: %+v", err)
		}
		return out, nil
	}

	var parent BaseDlpActionInfoImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDlpActionInfoImpl: %+v", err)
	}

	return RawDlpActionInfoImpl{
		dlpActionInfo: parent,
		Type:          value,
		Values:        temp,
	}, nil

}
