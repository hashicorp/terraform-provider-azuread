package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessReviewInstanceDecisionItemTarget = AccessReviewInstanceDecisionItemServicePrincipalTarget{}

type AccessReviewInstanceDecisionItemServicePrincipalTarget struct {
	// The appId for the service principal entity being reviewed.
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// The display name of the service principal whose access is being reviewed.
	ServicePrincipalDisplayName nullable.Type[string] `json:"servicePrincipalDisplayName,omitempty"`

	// The identifier of the service principal whose access is being reviewed.
	ServicePrincipalId nullable.Type[string] `json:"servicePrincipalId,omitempty"`

	// Fields inherited from AccessReviewInstanceDecisionItemTarget

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AccessReviewInstanceDecisionItemServicePrincipalTarget) AccessReviewInstanceDecisionItemTarget() BaseAccessReviewInstanceDecisionItemTargetImpl {
	return BaseAccessReviewInstanceDecisionItemTargetImpl{
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessReviewInstanceDecisionItemServicePrincipalTarget{}

func (s AccessReviewInstanceDecisionItemServicePrincipalTarget) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewInstanceDecisionItemServicePrincipalTarget
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewInstanceDecisionItemServicePrincipalTarget: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewInstanceDecisionItemServicePrincipalTarget: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewInstanceDecisionItemServicePrincipalTarget"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewInstanceDecisionItemServicePrincipalTarget: %+v", err)
	}

	return encoded, nil
}
