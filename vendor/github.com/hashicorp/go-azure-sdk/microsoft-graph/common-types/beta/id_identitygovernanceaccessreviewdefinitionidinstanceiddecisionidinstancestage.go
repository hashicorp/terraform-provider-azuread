package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId{}

// IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId is a struct representing the Resource ID for a Identity Governance Access Review Definition Id Instance Id Decision Id Instance Stage
type IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId struct {
	AccessReviewScheduleDefinitionId   string
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
	AccessReviewStageId                string
}

// NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageID returns a new IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId struct
func NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageID(accessReviewScheduleDefinitionId string, accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, accessReviewStageId string) IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId {
	return IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId{
		AccessReviewScheduleDefinitionId:   accessReviewScheduleDefinitionId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		AccessReviewStageId:                accessReviewStageId,
	}
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageID parses 'input' into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageID(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageIDInsensitively(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewScheduleDefinitionId, ok = input.Parsed["accessReviewScheduleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewScheduleDefinitionId", input)
	}

	if id.AccessReviewInstanceId, ok = input.Parsed["accessReviewInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", input)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = input.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", input)
	}

	if id.AccessReviewStageId, ok = input.Parsed["accessReviewStageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageID checks that 'input' can be parsed as a Identity Governance Access Review Definition Id Instance Id Decision Id Instance Stage ID
func ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review Definition Id Instance Id Decision Id Instance Stage ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId) ID() string {
	fmtString := "/identityGovernance/accessReviews/definitions/%s/instances/%s/decisions/%s/instance/stages/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewScheduleDefinitionId, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewStageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review Definition Id Instance Id Decision Id Instance Stage ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("accessReviews", "accessReviews", "accessReviews"),
		resourceids.StaticSegment("definitions", "definitions", "definitions"),
		resourceids.UserSpecifiedSegment("accessReviewScheduleDefinitionId", "accessReviewScheduleDefinitionId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
		resourceids.StaticSegment("instance", "instance", "instance"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review Definition Id Instance Id Decision Id Instance Stage ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceStageId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Schedule Definition: %q", id.AccessReviewScheduleDefinitionId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
	}
	return fmt.Sprintf("Identity Governance Access Review Definition Id Instance Id Decision Id Instance Stage (%s)", strings.Join(components, "\n"))
}
