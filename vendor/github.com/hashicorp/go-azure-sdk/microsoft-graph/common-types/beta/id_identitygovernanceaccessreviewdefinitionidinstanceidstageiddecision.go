package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId{}

// IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId is a struct representing the Resource ID for a Identity Governance Access Review Definition Id Instance Id Stage Id Decision
type IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId struct {
	AccessReviewScheduleDefinitionId   string
	AccessReviewInstanceId             string
	AccessReviewStageId                string
	AccessReviewInstanceDecisionItemId string
}

// NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID returns a new IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId struct
func NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID(accessReviewScheduleDefinitionId string, accessReviewInstanceId string, accessReviewStageId string, accessReviewInstanceDecisionItemId string) IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId {
	return IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId{
		AccessReviewScheduleDefinitionId:   accessReviewScheduleDefinitionId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewStageId:                accessReviewStageId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
	}
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID parses 'input' into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIDInsensitively(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewScheduleDefinitionId, ok = input.Parsed["accessReviewScheduleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewScheduleDefinitionId", input)
	}

	if id.AccessReviewInstanceId, ok = input.Parsed["accessReviewInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", input)
	}

	if id.AccessReviewStageId, ok = input.Parsed["accessReviewStageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", input)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = input.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID checks that 'input' can be parsed as a Identity Governance Access Review Definition Id Instance Id Stage Id Decision ID
func ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review Definition Id Instance Id Stage Id Decision ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId) ID() string {
	fmtString := "/identityGovernance/accessReviews/definitions/%s/instances/%s/stages/%s/decisions/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewScheduleDefinitionId, id.AccessReviewInstanceId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review Definition Id Instance Id Stage Id Decision ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("accessReviews", "accessReviews", "accessReviews"),
		resourceids.StaticSegment("definitions", "definitions", "definitions"),
		resourceids.UserSpecifiedSegment("accessReviewScheduleDefinitionId", "accessReviewScheduleDefinitionId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review Definition Id Instance Id Stage Id Decision ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Schedule Definition: %q", id.AccessReviewScheduleDefinitionId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
	}
	return fmt.Sprintf("Identity Governance Access Review Definition Id Instance Id Stage Id Decision (%s)", strings.Join(components, "\n"))
}
