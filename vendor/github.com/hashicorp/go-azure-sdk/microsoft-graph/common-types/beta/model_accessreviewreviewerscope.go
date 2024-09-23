package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessReviewScope = AccessReviewReviewerScope{}

type AccessReviewReviewerScope struct {
	// The query specifying who will be the reviewer.
	Query nullable.Type[string] `json:"query,omitempty"`

	// In the scenario where reviewers need to be specified dynamically, this property is used to indicate the relative
	// source of the query. This property is only required if a relative query, for example, ./manager, is specified.
	// Possible value: decisions.
	QueryRoot nullable.Type[string] `json:"queryRoot,omitempty"`

	// The type of query. Examples include MicrosoftGraph and ARM.
	QueryType nullable.Type[string] `json:"queryType,omitempty"`

	// Fields inherited from AccessReviewScope

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AccessReviewReviewerScope) AccessReviewScope() BaseAccessReviewScopeImpl {
	return BaseAccessReviewScopeImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessReviewReviewerScope{}

func (s AccessReviewReviewerScope) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewReviewerScope
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewReviewerScope: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewReviewerScope: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewReviewerScope"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewReviewerScope: %+v", err)
	}

	return encoded, nil
}
