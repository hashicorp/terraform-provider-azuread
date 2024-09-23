package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId{}

// IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId is a struct representing the Resource ID for a Identity Governance Access Review Definition Id Instance Id Contacted Reviewer
type IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId struct {
	AccessReviewScheduleDefinitionId string
	AccessReviewInstanceId           string
	AccessReviewReviewerId           string
}

// NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerID returns a new IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId struct
func NewIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerID(accessReviewScheduleDefinitionId string, accessReviewInstanceId string, accessReviewReviewerId string) IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId {
	return IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId{
		AccessReviewScheduleDefinitionId: accessReviewScheduleDefinitionId,
		AccessReviewInstanceId:           accessReviewInstanceId,
		AccessReviewReviewerId:           accessReviewReviewerId,
	}
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerID parses 'input' into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerID(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerIDInsensitively(input string) (*IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewScheduleDefinitionId, ok = input.Parsed["accessReviewScheduleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewScheduleDefinitionId", input)
	}

	if id.AccessReviewInstanceId, ok = input.Parsed["accessReviewInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", input)
	}

	if id.AccessReviewReviewerId, ok = input.Parsed["accessReviewReviewerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewReviewerId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerID checks that 'input' can be parsed as a Identity Governance Access Review Definition Id Instance Id Contacted Reviewer ID
func ValidateIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review Definition Id Instance Id Contacted Reviewer ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId) ID() string {
	fmtString := "/identityGovernance/accessReviews/definitions/%s/instances/%s/contactedReviewers/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewScheduleDefinitionId, id.AccessReviewInstanceId, id.AccessReviewReviewerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review Definition Id Instance Id Contacted Reviewer ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("accessReviews", "accessReviews", "accessReviews"),
		resourceids.StaticSegment("definitions", "definitions", "definitions"),
		resourceids.UserSpecifiedSegment("accessReviewScheduleDefinitionId", "accessReviewScheduleDefinitionId"),
		resourceids.StaticSegment("instances", "instances", "instances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("contactedReviewers", "contactedReviewers", "contactedReviewers"),
		resourceids.UserSpecifiedSegment("accessReviewReviewerId", "accessReviewReviewerId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review Definition Id Instance Id Contacted Reviewer ID
func (id IdentityGovernanceAccessReviewDefinitionIdInstanceIdContactedReviewerId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Schedule Definition: %q", id.AccessReviewScheduleDefinitionId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Reviewer: %q", id.AccessReviewReviewerId),
	}
	return fmt.Sprintf("Identity Governance Access Review Definition Id Instance Id Contacted Reviewer (%s)", strings.Join(components, "\n"))
}
