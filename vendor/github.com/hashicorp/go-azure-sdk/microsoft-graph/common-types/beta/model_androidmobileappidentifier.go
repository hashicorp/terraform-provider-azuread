package beta

import (
	"encoding/json"
	"fmt"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MobileAppIdentifier = AndroidMobileAppIdentifier{}

type AndroidMobileAppIdentifier struct {
	// The identifier for an app, as specified in the play store.
	PackageId *string `json:"packageId,omitempty"`

	// Fields inherited from MobileAppIdentifier

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AndroidMobileAppIdentifier) MobileAppIdentifier() BaseMobileAppIdentifierImpl {
	return BaseMobileAppIdentifierImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidMobileAppIdentifier{}

func (s AndroidMobileAppIdentifier) MarshalJSON() ([]byte, error) {
	type wrapper AndroidMobileAppIdentifier
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidMobileAppIdentifier: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidMobileAppIdentifier: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidMobileAppIdentifier"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidMobileAppIdentifier: %+v", err)
	}

	return encoded, nil
}
