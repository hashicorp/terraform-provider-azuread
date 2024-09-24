package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId{}

// IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId is a struct representing the Resource ID for a Identity Governance Access Review Definition Id Instance Id Decision Id Instance Contacted Reviewer
type IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId struct {
	AccessReviewScheduleDefinitionId   string
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
	AccessReviewReviewerId             string
}

// NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerID returns a new IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId struct
func NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerID(accessReviewScheduleDefinitionId string, accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, accessReviewReviewerId string) IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId {
	return IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId{
		AccessReviewScheduleDefinitionId:   accessReviewScheduleDefinitionId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		AccessReviewReviewerId:             accessReviewReviewerId,
	}
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerID parses 'input' into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerID(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerIDInsensitively(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AccessReviewReviewerId, ok = input.Parsed["accessReviewReviewerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewReviewerId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerID checks that 'input' can be parsed as a Identity Governance Access Review Definition Id Instance Id Decision Id Instance Contacted Reviewer ID
func ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review Definition Id Instance Id Decision Id Instance Contacted Reviewer ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId) ID() string {
	fmtString := "/identityGovernance/accessReviews/definitions/%s/instances/%s/decisions/%s/instance/contactedReviewers/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewScheduleDefinitionId, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewReviewerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review Definition Id Instance Id Decision Id Instance Contacted Reviewer ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("contactedReviewers", "contactedReviewers", "contactedReviewers"),
		resourceids.UserSpecifiedSegment("accessReviewReviewerId", "accessReviewReviewerId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review Definition Id Instance Id Decision Id Instance Contacted Reviewer ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdDecisionIdInstanceContactedReviewerId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Schedule Definition: %q", id.AccessReviewScheduleDefinitionId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Reviewer: %q", id.AccessReviewReviewerId),
	}
	return fmt.Sprintf("Identity Governance Access Review Definition Id Instance Id Decision Id Instance Contacted Reviewer (%s)", strings.Join(components, "\n"))
}
