package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessReviewNotificationRecipientScope = AccessReviewNotificationRecipientQueryScope{}

type AccessReviewNotificationRecipientQueryScope struct {
	// Represents the query for who the recipients are. For example, /groups/{group id}/members for group members and
	// /users/{user id} for a specific user.
	Query nullable.Type[string] `json:"query,omitempty"`

	// In the scenario where reviewers need to be specified dynamically, indicates the relative source of the query. This
	// property is only required if a relative query (that is, ./manager) is specified.
	QueryRoot nullable.Type[string] `json:"queryRoot,omitempty"`

	// Indicates the type of query. Allowed value is MicrosoftGraph.
	QueryType nullable.Type[string] `json:"queryType,omitempty"`

	// Fields inherited from AccessReviewNotificationRecipientScope

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AccessReviewNotificationRecipientQueryScope) AccessReviewNotificationRecipientScope() BaseAccessReviewNotificationRecipientScopeImpl {
	return BaseAccessReviewNotificationRecipientScopeImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessReviewNotificationRecipientQueryScope{}

func (s AccessReviewNotificationRecipientQueryScope) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewNotificationRecipientQueryScope
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewNotificationRecipientQueryScope: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewNotificationRecipientQueryScope: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewNotificationRecipientQueryScope"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewNotificationRecipientQueryScope: %+v", err)
	}

	return encoded, nil
}
