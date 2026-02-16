package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ IosVppAppAssignedLicense = IosVppAppAssignedUserLicense{}

type IosVppAppAssignedUserLicense struct {

	// Fields inherited from IosVppAppAssignedLicense

	// The user email address.
	UserEmailAddress nullable.Type[string] `json:"userEmailAddress,omitempty"`

	// The user ID.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// The user name.
	UserName nullable.Type[string] `json:"userName,omitempty"`

	// The user principal name.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

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

func (s IosVppAppAssignedUserLicense) IosVppAppAssignedLicense() BaseIosVppAppAssignedLicenseImpl {
	return BaseIosVppAppAssignedLicenseImpl{
		UserEmailAddress:  s.UserEmailAddress,
		UserId:            s.UserId,
		UserName:          s.UserName,
		UserPrincipalName: s.UserPrincipalName,
		Id:                s.Id,
		ODataId:           s.ODataId,
		ODataType:         s.ODataType,
	}
}

func (s IosVppAppAssignedUserLicense) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = IosVppAppAssignedUserLicense{}

func (s IosVppAppAssignedUserLicense) MarshalJSON() ([]byte, error) {
	type wrapper IosVppAppAssignedUserLicense
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling IosVppAppAssignedUserLicense: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling IosVppAppAssignedUserLicense: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.iosVppAppAssignedUserLicense"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling IosVppAppAssignedUserLicense: %+v", err)
	}

	return encoded, nil
}
