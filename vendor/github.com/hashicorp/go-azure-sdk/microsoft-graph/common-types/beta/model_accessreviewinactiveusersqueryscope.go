package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessReviewQueryScope = AccessReviewInactiveUsersQueryScope{}

type AccessReviewInactiveUsersQueryScope struct {
	// Defines the duration of inactivity. Inactivity is based on the last sign in date of the user compared to the access
	// review instance's start date. If this property is not specified, it's assigned the default value PT0S.
	InactiveDuration nullable.Type[string] `json:"inactiveDuration,omitempty"`

	// Fields inherited from AccessReviewQueryScope

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

func (s AccessReviewInactiveUsersQueryScope) AccessReviewQueryScope() BaseAccessReviewQueryScopeImpl {
	return BaseAccessReviewQueryScopeImpl{
		Query:     s.Query,
		QueryRoot: s.QueryRoot,
		QueryType: s.QueryType,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

func (s AccessReviewInactiveUsersQueryScope) AccessReviewScope() BaseAccessReviewScopeImpl {
	return BaseAccessReviewScopeImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessReviewInactiveUsersQueryScope{}

func (s AccessReviewInactiveUsersQueryScope) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewInactiveUsersQueryScope
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewInactiveUsersQueryScope: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewInactiveUsersQueryScope: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewInactiveUsersQueryScope"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewInactiveUsersQueryScope: %+v", err)
	}

	return encoded, nil
}
