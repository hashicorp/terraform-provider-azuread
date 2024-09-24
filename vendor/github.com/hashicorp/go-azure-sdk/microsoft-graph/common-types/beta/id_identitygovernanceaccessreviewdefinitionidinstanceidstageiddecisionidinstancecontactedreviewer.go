package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId{}

// IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId is a struct representing the Resource ID for a Identity Governance Access Review Definition Id Instance Id Stage Id Decision Id Instance Contacted Reviewer
type IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId struct {
	AccessReviewScheduleDefinitionId   string
	AccessReviewInstanceId             string
	AccessReviewStageId                string
	AccessReviewInstanceDecisionItemId string
	AccessReviewReviewerId             string
}

// NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerID returns a new IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId struct
func NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerID(accessReviewScheduleDefinitionId string, accessReviewInstanceId string, accessReviewStageId string, accessReviewInstanceDecisionItemId string, accessReviewReviewerId string) IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId {
	return IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId{
		AccessReviewScheduleDefinitionId:   accessReviewScheduleDefinitionId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewStageId:                accessReviewStageId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		AccessReviewReviewerId:             accessReviewReviewerId,
	}
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerID parses 'input' into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerID(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerIDInsensitively(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AccessReviewReviewerId, ok = input.Parsed["accessReviewReviewerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewReviewerId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerID checks that 'input' can be parsed as a Identity Governance Access Review Definition Id Instance Id Stage Id Decision Id Instance Contacted Reviewer ID
func ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review Definition Id Instance Id Stage Id Decision Id Instance Contacted Reviewer ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId) ID() string {
	fmtString := "/identityGovernance/accessReviews/definitions/%s/instances/%s/stages/%s/decisions/%s/instance/contactedReviewers/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewScheduleDefinitionId, id.AccessReviewInstanceId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewReviewerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review Definition Id Instance Id Stage Id Decision Id Instance Contacted Reviewer ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId) Segments() []resourceids.Segment {
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
		resourceids.StaticSegment("instance", "instance", "instance"),
		resourceids.StaticSegment("contactedReviewers", "contactedReviewers", "contactedReviewers"),
		resourceids.UserSpecifiedSegment("accessReviewReviewerId", "accessReviewReviewerId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review Definition Id Instance Id Stage Id Decision Id Instance Contacted Reviewer ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdStageIdDecisionIdInstanceContactedReviewerId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Schedule Definition: %q", id.AccessReviewScheduleDefinitionId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Reviewer: %q", id.AccessReviewReviewerId),
	}
	return fmt.Sprintf("Identity Governance Access Review Definition Id Instance Id Stage Id Decision Id Instance Contacted Reviewer (%s)", strings.Join(components, "\n"))
}
