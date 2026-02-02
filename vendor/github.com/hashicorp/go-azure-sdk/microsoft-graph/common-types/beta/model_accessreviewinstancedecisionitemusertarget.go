package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessReviewInstanceDecisionItemTarget = AccessReviewInstanceDecisionItemUserTarget{}

type AccessReviewInstanceDecisionItemUserTarget struct {
	// The name of user.
	UserDisplayName nullable.Type[string] `json:"userDisplayName,omitempty"`

	// The identifier of user.
	UserId nullable.Type[string] `json:"userId,omitempty"`

	// The user principal name.
	UserPrincipalName nullable.Type[string] `json:"userPrincipalName,omitempty"`

	// Fields inherited from AccessReviewInstanceDecisionItemTarget

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AccessReviewInstanceDecisionItemUserTarget) AccessReviewInstanceDecisionItemTarget() BaseAccessReviewInstanceDecisionItemTargetImpl {
	return BaseAccessReviewInstanceDecisionItemTargetImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessReviewInstanceDecisionItemUserTarget{}

func (s AccessReviewInstanceDecisionItemUserTarget) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewInstanceDecisionItemUserTarget
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewInstanceDecisionItemUserTarget: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewInstanceDecisionItemUserTarget: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewInstanceDecisionItemUserTarget"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewInstanceDecisionItemUserTarget: %+v", err)
	}

	return encoded, nil
}
