package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId{}

// IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId is a struct representing the Resource ID for a Identity Governance Access Review Decision Id Instance Contacted Reviewer
type IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId struct {
	AccessReviewInstanceDecisionItemId string
	AccessReviewReviewerId             string
}

// NewIdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerID returns a new IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId struct
func NewIdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerID(accessReviewInstanceDecisionItemId string, accessReviewReviewerId string) IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId {
	return IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId{
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		AccessReviewReviewerId:             accessReviewReviewerId,
	}
}

// ParseIdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerID parses 'input' into a IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId
func ParseIdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerID(input string) (*IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerIDInsensitively(input string) (*IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewInstanceDecisionItemId, ok = input.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", input)
	}

	if id.AccessReviewReviewerId, ok = input.Parsed["accessReviewReviewerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewReviewerId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerID checks that 'input' can be parsed as a Identity Governance Access Review Decision Id Instance Contacted Reviewer ID
func ValidateIdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review Decision Id Instance Contacted Reviewer ID
func (id IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId) ID() string {
	fmtString := "/identityGovernance/accessReviews/decisions/%s/instance/contactedReviewers/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceDecisionItemId, id.AccessReviewReviewerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review Decision Id Instance Contacted Reviewer ID
func (id IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("accessReviews", "accessReviews", "accessReviews"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
		resourceids.StaticSegment("instance", "instance", "instance"),
		resourceids.StaticSegment("contactedReviewers", "contactedReviewers", "contactedReviewers"),
		resourceids.UserSpecifiedSegment("accessReviewReviewerId", "accessReviewReviewerId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review Decision Id Instance Contacted Reviewer ID
func (id IdentityGovernanceAccessReviewDecisionIdInstanceContactedReviewerId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Reviewer: %q", id.AccessReviewReviewerId),
	}
	return fmt.Sprintf("Identity Governance Access Review Decision Id Instance Contacted Reviewer (%s)", strings.Join(components, "\n"))
}
