package beta

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DlpNotification interface {
	DlpNotification() BaseDlpNotificationImpl
}

var _ DlpNotification = BaseDlpNotificationImpl{}

type BaseDlpNotificationImpl struct {
	Author nullable.Type[string] `json:"author,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseDlpNotificationImpl) DlpNotification() BaseDlpNotificationImpl {
	return s
}

var _ DlpNotification = RawDlpNotificationImpl{}

// RawDlpNotificationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDlpNotificationImpl struct {
	dlpNotification BaseDlpNotificationImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawDlpNotificationImpl) DlpNotification() BaseDlpNotificationImpl {
	return s.dlpNotification
}

func UnmarshalDlpNotificationImplementation(input []byte) (DlpNotification, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DlpNotification into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.dlpWindowsDevicesNotification") {
		var out DlpWindowsDevicesNotification
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DlpWindowsDevicesNotification: %+v", err)
		}
		return out, nil
	}

	var parent BaseDlpNotificationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDlpNotificationImpl: %+v", err)
	}

	return RawDlpNotificationImpl{
		dlpNotification: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
