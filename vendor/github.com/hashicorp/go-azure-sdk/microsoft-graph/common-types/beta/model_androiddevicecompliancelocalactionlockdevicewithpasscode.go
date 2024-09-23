package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AndroidDeviceComplianceLocalActionBase = AndroidDeviceComplianceLocalActionLockDeviceWithPasscode{}

type AndroidDeviceComplianceLocalActionLockDeviceWithPasscode struct {
	// Passcode to reset to Android device. This property is read-only.
	Passcode nullable.Type[string] `json:"passcode,omitempty"`

	// Number of sign in failures before wiping device, the value can be 4-11. Valid values 4 to 11
	PasscodeSignInFailureCountBeforeWipe nullable.Type[int64] `json:"passcodeSignInFailureCountBeforeWipe,omitempty"`

	// Fields inherited from AndroidDeviceComplianceLocalActionBase

	// Number of minutes to wait till a local action is enforced. Valid values 0 to 2147483647
	GracePeriodInMinutes *int64 `json:"gracePeriodInMinutes,omitempty"`

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

func (s AndroidDeviceComplianceLocalActionLockDeviceWithPasscode) AndroidDeviceComplianceLocalActionBase() BaseAndroidDeviceComplianceLocalActionBaseImpl {
	return BaseAndroidDeviceComplianceLocalActionBaseImpl{
		GracePeriodInMinutes: s.GracePeriodInMinutes,
		Id:                   s.Id,
		ODataId:              s.ODataId,
		ODataType:            s.ODataType,
	}
}

func (s AndroidDeviceComplianceLocalActionLockDeviceWithPasscode) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AndroidDeviceComplianceLocalActionLockDeviceWithPasscode{}

func (s AndroidDeviceComplianceLocalActionLockDeviceWithPasscode) MarshalJSON() ([]byte, error) {
	type wrapper AndroidDeviceComplianceLocalActionLockDeviceWithPasscode
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AndroidDeviceComplianceLocalActionLockDeviceWithPasscode: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AndroidDeviceComplianceLocalActionLockDeviceWithPasscode: %+v", err)
	}

	delete(decoded, "passcode")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.androidDeviceComplianceLocalActionLockDeviceWithPasscode"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AndroidDeviceComplianceLocalActionLockDeviceWithPasscode: %+v", err)
	}

	return encoded, nil
}
