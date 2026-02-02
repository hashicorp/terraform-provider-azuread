package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SubjectSet = SingleServicePrincipal{}

type SingleServicePrincipal struct {
	// Description of this service principal.
	Description nullable.Type[string] `json:"description,omitempty"`

	// ID of the servicePrincipal.
	ServicePrincipalId nullable.Type[string] `json:"servicePrincipalId,omitempty"`

	// Fields inherited from SubjectSet

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s SingleServicePrincipal) SubjectSet() BaseSubjectSetImpl {
	return BaseSubjectSetImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SingleServicePrincipal{}

func (s SingleServicePrincipal) MarshalJSON() ([]byte, error) {
	type wrapper SingleServicePrincipal
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SingleServicePrincipal: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SingleServicePrincipal: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.singleServicePrincipal"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SingleServicePrincipal: %+v", err)
	}

	return encoded, nil
}
