package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = AccessReviewInstanceDecisionItem{}

type AccessReviewInstanceDecisionItem struct {
	// The identifier of the accessReviewInstance parent. Supports $select. Read-only.
	AccessReviewId *string `json:"accessReviewId,omitempty"`

	// The identifier of the user who applied the decision. 00000000-0000-0000-0000-000000000000 if the assigned reviewer
	// hasn't applied the decision or it was automatically applied. Read-only.
	AppliedBy *UserIdentity `json:"appliedBy,omitempty"`

	// The timestamp when the approval decision was applied. The DatetimeOffset type represents date and time information
	// using ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z.
	// Supports $select. Read-only.
	AppliedDateTime nullable.Type[string] `json:"appliedDateTime,omitempty"`

	// The result of applying the decision. Possible values: New, AppliedSuccessfully, AppliedWithUnknownFailure,
	// AppliedSuccessfullyButObjectNotFound, and ApplyNotSupported. Supports $select, $orderby, and $filter (eq only).
	// Read-only.
	ApplyResult nullable.Type[string] `json:"applyResult,omitempty"`

	// Result of the review. Possible values: Approve, Deny, NotReviewed, or DontKnow. Supports $select, $orderby, and
	// $filter (eq only).
	Decision nullable.Type[string] `json:"decision,omitempty"`

	// Insights are recommendations to reviewers on whether to approve or deny a decision. There can be multiple insights
	// associated with an accessReviewInstanceDecisionItem.
	Insights *[]GovernanceInsight `json:"insights,omitempty"`

	// There's exactly one accessReviewInstance associated with each decision. The instance is the parent of the decision
	// item, representing the recurrence of the access review the decision is made on.
	Instance *AccessReviewInstance `json:"instance,omitempty"`

	// Justification left by the reviewer when they made the decision.
	Justification nullable.Type[string] `json:"justification,omitempty"`

	// Every decision item in an access review represents a principal's access to a resource. This property represents
	// details of the principal. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The
	// principal is 'Bob' and the resource is 'Sales'. Principals can be of two types - userIdentity and
	// servicePrincipalIdentity. Supports $select. Read-only.
	Principal *Identity `json:"principal,omitempty"`

	// Link to the principal object. For example:
	// https://graph.microsoft.com/v1.0/users/a6c7aecb-cbfd-4763-87ef-e91b4bd509d9. Read-only.
	PrincipalLink nullable.Type[string] `json:"principalLink,omitempty"`

	// Every decision item in an access review represents a principal's membership to a resource. This property provides the
	// details of the membership. For example, whether the principal has direct access or indirect access to the resource.
	// Supports $select. Read-only.
	PrincipalResourceMembership *DecisionItemPrincipalResourceMembership `json:"principalResourceMembership,omitempty"`

	// A system-generated recommendation for the approval decision based off last interactive sign-in to tenant. Recommend
	// approve if sign-in is within 30 days of start of review. Recommend deny if sign-in is greater than 30 days of start
	// of review. Recommendation not available otherwise. Possible values: Approve, Deny, or NoInfoAvailable. Supports
	// $select, $orderby, and $filter (eq only). Read-only.
	Recommendation nullable.Type[string] `json:"recommendation,omitempty"`

	// Every decision item in an access review represents a principal's access to a resource. This property represents
	// details of the resource. For example, if a decision item represents access of User 'Bob' to Group 'Sales' - The
	// principal is Bob and the resource is 'Sales'. Resources can be of multiple types. See
	// accessReviewInstanceDecisionItemResource. Read-only.
	Resource *AccessReviewInstanceDecisionItemResource `json:"resource,omitempty"`

	// A link to the resource. For example,
	// https://graph.microsoft.com/v1.0/servicePrincipals/c86300f3-8695-4320-9f6e-32a2555f5ff8. Supports $select. Read-only.
	ResourceLink nullable.Type[string] `json:"resourceLink,omitempty"`

	// The identifier of the reviewer. 00000000-0000-0000-0000-000000000000 if the assigned reviewer hasn't reviewed.
	// Supports $select. Read-only.
	ReviewedBy *UserIdentity `json:"reviewedBy,omitempty"`

	// The timestamp when the review decision occurred. Supports $select. Read-only.
	ReviewedDateTime nullable.Type[string] `json:"reviewedDateTime,omitempty"`

	// The target of this specific decision. Decision targets can be of different types – each one with its own specific
	// properties. See accessReviewInstanceDecisionItemTarget. Read-only. This property is replaced by the principal and
	// resource properties in v1.0.
	Target *AccessReviewInstanceDecisionItemTarget `json:"target,omitempty"`

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

func (s AccessReviewInstanceDecisionItem) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = AccessReviewInstanceDecisionItem{}

func (s AccessReviewInstanceDecisionItem) MarshalJSON() ([]byte, error) {
	type wrapper AccessReviewInstanceDecisionItem
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AccessReviewInstanceDecisionItem: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AccessReviewInstanceDecisionItem: %+v", err)
	}

	delete(decoded, "accessReviewId")
	delete(decoded, "appliedBy")
	delete(decoded, "appliedDateTime")
	delete(decoded, "applyResult")
	delete(decoded, "principal")
	delete(decoded, "principalLink")
	delete(decoded, "principalResourceMembership")
	delete(decoded, "recommendation")
	delete(decoded, "resource")
	delete(decoded, "resourceLink")
	delete(decoded, "reviewedBy")
	delete(decoded, "reviewedDateTime")
	delete(decoded, "target")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.accessReviewInstanceDecisionItem"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AccessReviewInstanceDecisionItem: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &AccessReviewInstanceDecisionItem{}

func (s *AccessReviewInstanceDecisionItem) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AccessReviewId              *string                                  `json:"accessReviewId,omitempty"`
		AppliedDateTime             nullable.Type[string]                    `json:"appliedDateTime,omitempty"`
		ApplyResult                 nullable.Type[string]                    `json:"applyResult,omitempty"`
		Decision                    nullable.Type[string]                    `json:"decision,omitempty"`
		Instance                    *AccessReviewInstance                    `json:"instance,omitempty"`
		Justification               nullable.Type[string]                    `json:"justification,omitempty"`
		PrincipalLink               nullable.Type[string]                    `json:"principalLink,omitempty"`
		PrincipalResourceMembership *DecisionItemPrincipalResourceMembership `json:"principalResourceMembership,omitempty"`
		Recommendation              nullable.Type[string]                    `json:"recommendation,omitempty"`
		ResourceLink                nullable.Type[string]                    `json:"resourceLink,omitempty"`
		ReviewedDateTime            nullable.Type[string]                    `json:"reviewedDateTime,omitempty"`
		Id                          *string                                  `json:"id,omitempty"`
		ODataId                     *string                                  `json:"@odata.id,omitempty"`
		ODataType                   *string                                  `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AccessReviewId = decoded.AccessReviewId
	s.AppliedDateTime = decoded.AppliedDateTime
	s.ApplyResult = decoded.ApplyResult
	s.Decision = decoded.Decision
	s.Instance = decoded.Instance
	s.Justification = decoded.Justification
	s.PrincipalLink = decoded.PrincipalLink
	s.PrincipalResourceMembership = decoded.PrincipalResourceMembership
	s.Recommendation = decoded.Recommendation
	s.ResourceLink = decoded.ResourceLink
	s.ReviewedDateTime = decoded.ReviewedDateTime
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling AccessReviewInstanceDecisionItem into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["appliedBy"]; ok {
		impl, err := UnmarshalUserIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AppliedBy' for 'AccessReviewInstanceDecisionItem': %+v", err)
		}
		s.AppliedBy = &impl
	}

	if v, ok := temp["insights"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling Insights into list []json.RawMessage: %+v", err)
		}

		output := make([]GovernanceInsight, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalGovernanceInsightImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'Insights' for 'AccessReviewInstanceDecisionItem': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.Insights = &output
	}

	if v, ok := temp["principal"]; ok {
		impl, err := UnmarshalIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Principal' for 'AccessReviewInstanceDecisionItem': %+v", err)
		}
		s.Principal = &impl
	}

	if v, ok := temp["resource"]; ok {
		impl, err := UnmarshalAccessReviewInstanceDecisionItemResourceImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Resource' for 'AccessReviewInstanceDecisionItem': %+v", err)
		}
		s.Resource = &impl
	}

	if v, ok := temp["reviewedBy"]; ok {
		impl, err := UnmarshalUserIdentityImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ReviewedBy' for 'AccessReviewInstanceDecisionItem': %+v", err)
		}
		s.ReviewedBy = &impl
	}

	if v, ok := temp["target"]; ok {
		impl, err := UnmarshalAccessReviewInstanceDecisionItemTargetImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Target' for 'AccessReviewInstanceDecisionItem': %+v", err)
		}
		s.Target = &impl
	}

	return nil
}
