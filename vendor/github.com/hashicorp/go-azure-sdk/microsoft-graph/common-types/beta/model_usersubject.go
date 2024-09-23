package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ConditionalAccessWhatIfSubject = UserSubject{}

type UserSubject struct {
	ExternalTenantId nullable.Type[string]                      `json:"externalTenantId,omitempty"`
	ExternalUserType *ConditionalAccessGuestOrExternalUserTypes `json:"externalUserType,omitempty"`
	UserId           nullable.Type[string]                      `json:"userId,omitempty"`

	// Fields inherited from ConditionalAccessWhatIfSubject

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s UserSubject) ConditionalAccessWhatIfSubject() BaseConditionalAccessWhatIfSubjectImpl {
	return BaseConditionalAccessWhatIfSubjectImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = UserSubject{}

func (s UserSubject) MarshalJSON() ([]byte, error) {
	type wrapper UserSubject
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling UserSubject: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling UserSubject: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.userSubject"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling UserSubject: %+v", err)
	}

	return encoded, nil
}
