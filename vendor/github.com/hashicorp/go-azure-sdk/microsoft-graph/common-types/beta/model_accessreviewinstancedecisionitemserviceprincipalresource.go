package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessReviewInstanceDecisionItemResource = AccessReviewInstanceDecisionItemServicePrincipalResource{}

type AccessReviewInstanceDecisionItemServicePrincipalResource struct {
	// The globally unique identifier of the application to which access is granted.
	AppId nullable.Type[string] `json:"appId,omitempty"`

	// Fields inherited from AccessReviewInstanceDecisionItemResource

	// Display name of the resource
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Resource ID
	Id nullable.Type[string] `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Type of resource. Types include: Group, ServicePrincipal, DirectoryRole, AzureRole, AccessPackageAssignmentPolicy.
	Type nullable.Type[string] `json:"type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s AccessReviewInstanceDecisionItemServicePrincipalResource) AccessReviewInstanceDecisionItemResource() BaseAccessReviewInstanceDecisionItemResourceImpl {
	return BaseAccessReviewInstanceDecisionItemResourceImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		Type:        s.Type,
	}
}

var _ json.Marshaler = AccessReviewInstanceDecisionItemServicePrincipalResource{}

func (s AccessReviewInstanceDecisionItemServicePrincipalResource) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewInstanceDecisionItemServicePrincipalResource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewInstanceDecisionItemServicePrincipalResource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewInstanceDecisionItemServicePrincipalResource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewInstanceDecisionItemServicePrincipalResource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewInstanceDecisionItemServicePrincipalResource: %+v", err)
	}

	return encoded, nil
}
