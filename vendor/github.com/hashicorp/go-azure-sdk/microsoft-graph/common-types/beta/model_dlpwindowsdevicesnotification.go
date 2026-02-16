package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DlpNotification = DlpWindowsDevicesNotification{}

type DlpWindowsDevicesNotification struct {
	ContentName   nullable.Type[string] `json:"contentName,omitempty"`
	LastModfiedBy nullable.Type[string] `json:"lastModfiedBy,omitempty"`

	// Fields inherited from DlpNotification

	Author nullable.Type[string] `json:"author,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s DlpWindowsDevicesNotification) DlpNotification() BaseDlpNotificationImpl {
	return BaseDlpNotificationImpl{
		Author:    s.Author,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = DlpWindowsDevicesNotification{}

func (s DlpWindowsDevicesNotification) MarshalJSON() ([]byte, error) {
	type wrapper DlpWindowsDevicesNotification
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DlpWindowsDevicesNotification: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DlpWindowsDevicesNotification: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.dlpWindowsDevicesNotification"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DlpWindowsDevicesNotification: %+v", err)
	}

	return encoded, nil
}
