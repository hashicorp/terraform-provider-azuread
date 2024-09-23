package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = MicrosoftApplicationDataAccessSettings{}

type MicrosoftApplicationDataAccessSettings struct {
	// The ID of a Microsoft Entra security group for which the members are allowed to access Microsoft 365 data using only
	// Microsoft 365 apps, but not other Microsoft apps such as Edge. This is only applicable if
	// isEnabledForAllMicrosoftApplications is set to true.
	DisabledForGroup nullable.Type[string] `json:"disabledForGroup,omitempty"`

	// When set to true, all users in the organization can access in a Microsoft app any Microsoft 365 data that the user
	// has been authorized to access. The Microsoft app can be a Microsoft 365 app (for example, Excel, Outlook) or
	// non-Microsoft 365 app (for example, Edge). The default is true. It is possible to disable this access for a subset of
	// users in a Microsoft Entra security group, by specifying the group in the disabledForGroup property. When set to
	// false, all users can access authorized Microsoft 365 data only in a Microsoft 365 app.
	IsEnabledForAllMicrosoftApplications nullable.Type[bool] `json:"isEnabledForAllMicrosoftApplications,omitempty"`

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

func (s MicrosoftApplicationDataAccessSettings) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = MicrosoftApplicationDataAccessSettings{}

func (s MicrosoftApplicationDataAccessSettings) MarshalJSON() ([]byte, error) {
	type wrapper MicrosoftApplicationDataAccessSettings
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MicrosoftApplicationDataAccessSettings: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MicrosoftApplicationDataAccessSettings: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.microsoftApplicationDataAccessSettings"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MicrosoftApplicationDataAccessSettings: %+v", err)
	}

	return encoded, nil
}
