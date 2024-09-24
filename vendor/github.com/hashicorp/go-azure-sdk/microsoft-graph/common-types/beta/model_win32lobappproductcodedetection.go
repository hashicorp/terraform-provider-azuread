package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Win32LobAppDetection = Win32LobAppProductCodeDetection{}

type Win32LobAppProductCodeDetection struct {
	// The product code of Win32 Line of Business (LoB) app.
	ProductCode nullable.Type[string] `json:"productCode,omitempty"`

	// The product version of Win32 Line of Business (LoB) app.
	ProductVersion nullable.Type[string] `json:"productVersion,omitempty"`

	// Contains properties for detection operator.
	ProductVersionOperator *Win32LobAppDetectionOperator `json:"productVersionOperator,omitempty"`

	// Fields inherited from Win32LobAppDetection

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s Win32LobAppProductCodeDetection) Win32LobAppDetection() BaseWin32LobAppDetectionImpl {
	return BaseWin32LobAppDetectionImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Win32LobAppProductCodeDetection{}

func (s Win32LobAppProductCodeDetection) MarshalJSON() ([]byte, error) {
	type wrapper Win32LobAppProductCodeDetection
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Win32LobAppProductCodeDetection: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Win32LobAppProductCodeDetection: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.win32LobAppProductCodeDetection"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Win32LobAppProductCodeDetection: %+v", err)
	}

	return encoded, nil
}
