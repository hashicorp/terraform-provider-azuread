package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId{}

// IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId is a struct representing the Resource ID for a Identity Governance Access Review Definition Id Instance Id Stage Id Decision Id Insight
type IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId struct {
	AccessReviewScheduleDefinitionId   string
	AccessReviewInstanceId             string
	AccessReviewStageId                string
	AccessReviewInstanceDecisionItemId string
	GovernanceInsightId                string
}

// NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightID returns a new IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId struct
func NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightID(accessReviewScheduleDefinitionId string, accessReviewInstanceId string, accessReviewStageId string, accessReviewInstanceDecisionItemId string, governanceInsightId string) IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId {
	return IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId{
		AccessReviewScheduleDefinitionId:   accessReviewScheduleDefinitionId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewStageId:                accessReviewStageId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		GovernanceInsightId:                governanceInsightId,
	}
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightID parses 'input' into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightID(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightIDInsensitively(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.GovernanceInsightId, ok = input.Parsed["governanceInsightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightID checks that 'input' can be parsed as a Identity Governance Access Review Definition Id Instance Id Stage Id Decision Id Insight ID
func ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review Definition Id Instance Id Stage Id Decision Id Insight ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId) ID() string {
	fmtString := "/identityGovernance/accessReviews/definitions/%s/instances/%s/stages/%s/decisions/%s/insights/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewScheduleDefinitionId, id.AccessReviewInstanceId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId, id.GovernanceInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review Definition Id Instance Id Stage Id Decision Id Insight ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("insights", "insights", "insights"),
		resourceids.UserSpecifiedSegment("governanceInsightId", "governanceInsightId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review Definition Id Instance Id Stage Id Decision Id Insight ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInsightId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Schedule Definition: %q", id.AccessReviewScheduleDefinitionId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Governance Insight: %q", id.GovernanceInsightId),
	}
	return fmt.Sprintf("Identity Governance Access Review Definition Id Instance Id Stage Id Decision Id Insight (%s)", strings.Join(components, "\n"))
}
