package stable

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AccessReviewQueryScope interface {
	AccessReviewScope
	AccessReviewQueryScope() BaseAccessReviewQueryScopeImpl
}

var _ AccessReviewQueryScope = BaseAccessReviewQueryScopeImpl{}

type BaseAccessReviewQueryScopeImpl struct {
	// The query representing what will be reviewed in an access review.
	Query nullable.Type[string] `json:"query,omitempty"`

	// In the scenario where reviewers need to be specified dynamically, this property is used to indicate the relative
	// source of the query. This property is only required if a relative query is specified. For example, ./manager.
	QueryRoot nullable.Type[string] `json:"queryRoot,omitempty"`

	// Indicates the type of query. Types include MicrosoftGraph and ARM.
	QueryType nullable.Type[string] `json:"queryType,omitempty"`

	// Fields inherited from AccessReviewScope

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s BaseAccessReviewQueryScopeImpl) AccessReviewQueryScope() BaseAccessReviewQueryScopeImpl {
	return s
}

func (s BaseAccessReviewQueryScopeImpl) AccessReviewScope() BaseAccessReviewScopeImpl {
	return BaseAccessReviewScopeImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ AccessReviewQueryScope = RawAccessReviewQueryScopeImpl{}

// RawAccessReviewQueryScopeImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawAccessReviewQueryScopeImpl struct {
	accessReviewQueryScope BaseAccessReviewQueryScopeImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawAccessReviewQueryScopeImpl) AccessReviewQueryScope() BaseAccessReviewQueryScopeImpl {
	return s.accessReviewQueryScope
}

func (s RawAccessReviewQueryScopeImpl) AccessReviewScope() BaseAccessReviewScopeImpl {
	return s.accessReviewQueryScope.AccessReviewScope()
}

var _ json.Marshaler = BaseAccessReviewQueryScopeImpl{}

func (s BaseAccessReviewQueryScopeImpl) MarshalJSON() ([]byte, error) {
	type wrapper BaseAccessReviewQueryScopeImpl
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BaseAccessReviewQueryScopeImpl: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseAccessReviewQueryScopeImpl: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewQueryScope"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BaseAccessReviewQueryScopeImpl: %+v", err)
	}

	return encoded, nil
}

func UnmarshalAccessReviewQueryScopeImplementation(input []byte) (AccessReviewQueryScope, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewQueryScope into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["@odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "#microsoft.graph.accessReviewInactiveUsersQueryScope") {
		var out AccessReviewInactiveUsersQueryScope
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AccessReviewInactiveUsersQueryScope: %+v", err)
		}
		return out, nil
	}

	var parent BaseAccessReviewQueryScopeImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseAccessReviewQueryScopeImpl: %+v", err)
	}

	return RawAccessReviewQueryScopeImpl{
		accessReviewQueryScope: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
