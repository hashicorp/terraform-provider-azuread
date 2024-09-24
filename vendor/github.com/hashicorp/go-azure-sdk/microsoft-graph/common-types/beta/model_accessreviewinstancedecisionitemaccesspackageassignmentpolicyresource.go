package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ AccessReviewInstanceDecisionItemResource = AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource{}

type AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource struct {
	// Display name of the access package to which access is granted.
	AccessPackageDisplayName nullable.Type[string] `json:"accessPackageDisplayName,omitempty"`

	// Identifier of the access package to which access is granted.
	AccessPackageId nullable.Type[string] `json:"accessPackageId,omitempty"`

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

func (s AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource) AccessReviewInstanceDecisionItemResource() BaseAccessReviewInstanceDecisionItemResourceImpl {
	return BaseAccessReviewInstanceDecisionItemResourceImpl{
		DisplayName: s.DisplayName,
		Id:          s.Id,
		ODataId:     s.ODataId,
		ODataType:   s.ODataType,
		Type:        s.Type,
	}
}

var _ json.Marshaler = AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource{}

func (s AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewInstanceDecisionItemAccessPackageAssignmentPolicyResource: %+v", err)
	}

	return encoded, nil
}
