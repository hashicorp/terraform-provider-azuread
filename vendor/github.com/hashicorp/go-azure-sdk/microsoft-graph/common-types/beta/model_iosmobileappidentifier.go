package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppIdentifier = IosMobileAppIdentifier{}

type IosMobileAppIdentifier struct {
	// The identifier for an app, as specified in the app store.
	BundleId *string `json:"bundleId,omitempty"`

	// Fields inherited from MobileAppIdentifier

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s IosMobileAppIdentifier) MobileAppIdentifier() BaseMobileAppIdentifierImpl {
	return BaseMobileAppIdentifierImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosMobileAppIdentifier{}

func (s IosMobileAppIdentifier) MarshalJSON() ([]byte, error) {
	type wrapper IosMobileAppIdentifier
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosMobileAppIdentifier: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosMobileAppIdentifier: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosMobileAppIdentifier"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosMobileAppIdentifier: %+v", err)
	}

	return encoded, nil
}
