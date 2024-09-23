package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId{}

// IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId is a struct representing the Resource ID for a Identity Governance Access Review Definition Id Instance Id Decision Id Insight
type IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId struct {
	AccessReviewScheduleDefinitionId   string
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
	GovernanceInsightId                string
}

// NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightID returns a new IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId struct
func NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightID(accessReviewScheduleDefinitionId string, accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, governanceInsightId string) IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId {
	return IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId{
		AccessReviewScheduleDefinitionId:   accessReviewScheduleDefinitionId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		GovernanceInsightId:                governanceInsightId,
	}
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightID parses 'input' into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightID(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightIDInsensitively(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.GovernanceInsightId, ok = input.Parsed["governanceInsightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightID checks that 'input' can be parsed as a Identity Governance Access Review Definition Id Instance Id Decision Id Insight ID
func ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review Definition Id Instance Id Decision Id Insight ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId) ID() string {
	fmtString := "/identityGovernance/accessReviews/definitions/%s/instances/%s/decisions/%s/insights/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewScheduleDefinitionId, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.GovernanceInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review Definition Id Instance Id Decision Id Insight ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("accessReviews", "accessReviews", "accessReviews"),
		resourceids.StaticSegment("definitions", "definitions", "definitions"),
		resourceids.UserSpecifiedSegment("accessReviewScheduleDefinitionId", "accessReviewScheduleDefinitionId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
		resourceids.StaticSegment("insights", "insights", "insights"),
		resourceids.UserSpecifiedSegment("governanceInsightId", "governanceInsightId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review Definition Id Instance Id Decision Id Insight ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInsightId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Schedule Definition: %q", id.AccessReviewScheduleDefinitionId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Governance Insight: %q", id.GovernanceInsightId),
	}
	return fmt.Sprintf("Identity Governance Access Review Definition Id Instance Id Decision Id Insight (%s)", strings.Join(components, "\n"))
}
