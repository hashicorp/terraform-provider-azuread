package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId{}

// IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId is a struct representing the Resource ID for a Identity Governance Access Review Definition Id Instance Id Decision
type IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId struct {
	AccessReviewScheduleDefinitionId   string
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
}

// NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID returns a new IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId struct
func NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID(accessReviewScheduleDefinitionId string, accessReviewInstanceId string, accessReviewInstanceDecisionItemId string) IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId {
	return IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId{
		AccessReviewScheduleDefinitionId:   accessReviewScheduleDefinitionId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
	}
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID parses 'input' into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIDInsensitively(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID checks that 'input' can be parsed as a Identity Governance Access Review Definition Id Instance Id Decision ID
func ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review Definition Id Instance Id Decision ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId) ID() string {
	fmtString := "/identityGovernance/accessReviews/definitions/%s/instances/%s/decisions/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewScheduleDefinitionId, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review Definition Id Instance Id Decision ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("accessReviews", "accessReviews", "accessReviews"),
		resourceids.StaticSegment("definitions", "definitions", "definitions"),
		resourceids.UserSpecifiedSegment("accessReviewScheduleDefinitionId", "accessReviewScheduleDefinitionId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review Definition Id Instance Id Decision ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Schedule Definition: %q", id.AccessReviewScheduleDefinitionId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
	}
	return fmt.Sprintf("Identity Governance Access Review Definition Id Instance Id Decision (%s)", strings.Join(components, "\n"))
}
