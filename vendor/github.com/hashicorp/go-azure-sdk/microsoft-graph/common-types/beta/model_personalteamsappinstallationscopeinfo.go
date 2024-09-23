package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TeamsAppInstallationScopeInfo = PersonalTeamsAppInstallationScopeInfo{}

type PersonalTeamsAppInstallationScopeInfo struct {
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// Fields inherited from TeamsAppInstallationScopeInfo

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	Scope *TeamsAppInstallationScopes `json:"scope,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s PersonalTeamsAppInstallationScopeInfo) TeamsAppInstallationScopeInfo() BaseTeamsAppInstallationScopeInfoImpl {
	return BaseTeamsAppInstallationScopeInfoImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
		Scope:     s.Scope,
	}
}

var _ json.Marshaler = PersonalTeamsAppInstallationScopeInfo{}

func (s PersonalTeamsAppInstallationScopeInfo) MarshalJSON() ([]byte, error) {
	type wrapper PersonalTeamsAppInstallationScopeInfo
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PersonalTeamsAppInstallationScopeInfo: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PersonalTeamsAppInstallationScopeInfo: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.personalTeamsAppInstallationScopeInfo"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PersonalTeamsAppInstallationScopeInfo: %+v", err)
	}

	return encoded, nil
}
