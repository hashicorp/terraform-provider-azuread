package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = OperationApprovalPolicy{}

type OperationApprovalPolicy struct {
	// The Microsoft Entra ID (Azure AD) security group IDs for the approvers for the policy. This property is required when
	// the policy is created, and is defined by the IT Admins to define the possible approvers for the policy.
	ApproverGroupIds *[]string `json:"approverGroupIds,omitempty"`

	// Indicates the description of the policy. Maximum length of the description is 1024 characters. This property is not
	// required, but can be used by the IT Admin to describe the policy.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Indicates the display name of the policy. Maximum length of the display name is 128 characters. This property is
	// required when the policy is created, and is defined by the IT Admins to identify the policy.
	DisplayName *string `json:"displayName,omitempty"`

	// Indicates the last DateTime that the policy was modified. The value cannot be modified and is automatically populated
	// whenever values in the request are updated. For example, when the 'policyType' property changes from apps to scripts.
	// The Timestamp type represents date and time information using ISO 8601 format and is always in UTC time. For example,
	// midnight UTC on Jan 1, 2014 would look like this: '2014-01-01T00:00:00Z'. Returned by default. Read-only. This
	// property is read-only.
	LastModifiedDateTime *string `json:"lastModifiedDateTime,omitempty"`

	// The set of available platforms for the OperationApprovalPolicy. Allows configuration of a policy to specific
	// platform(s) for approval. If no specific platform is required or applicable, the platform is `notApplicable`.
	PolicyPlatform *OperationApprovalPolicyPlatform `json:"policyPlatform,omitempty"`

	// Indicates areas of the Intune UX that could support MAA UX for the current logged in IT Admin. This property is
	// required, and is defined by the IT Admins in order to correctly show the expected experience.
	PolicySet *OperationApprovalPolicySet `json:"policySet,omitempty"`

	// The set of available policy types that can be configured for approval. The policy type must always be defined in an
	// OperationApprovalRequest.
	PolicyType *OperationApprovalPolicyType `json:"policyType,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s OperationApprovalPolicy) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = OperationApprovalPolicy{}

func (s OperationApprovalPolicy) MarshalJSON() ([]byte, error) {
	type wrapper OperationApprovalPolicy
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OperationApprovalPolicy: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OperationApprovalPolicy: %+v", err)
	}

	delete(decoded, "lastModifiedDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.operationApprovalPolicy"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OperationApprovalPolicy: %+v", err)
	}

	return encoded, nil
}
