package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDefinitionId{}

// IdentityGovernanceAccessReviewDefinitionId is a struct representing the Resource ID for a Identity Governance Access Review Definition
type IdentityGovernanceAccessReviewDefinitionId struct {
	AccessReviewScheduleDefinitionId string
}

// NewIdentityGovernanceAccessReviewDefinitionID returns a new IdentityGovernanceAccessReviewDefinitionId struct
func NewIdentityGovernanceAccessReviewDefinitionID(accessReviewScheduleDefinitionId string) IdentityGovernanceAccessReviewDefinitionId {
	return IdentityGovernanceAccessReviewDefinitionId{
		AccessReviewScheduleDefinitionId: accessReviewScheduleDefinitionId,
	}
}

// ParseIdentityGovernanceAccessReviewDefinitionID parses 'input' into a IdentityGovernanceAccessReviewDefinitionId
func ParseIdentityGovernanceAccessReviewDefinitionID(input string) (*IdentityGovernanceAccessReviewDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewDefinitionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewDefinitionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewDefinitionIDInsensitively(input string) (*IdentityGovernanceAccessReviewDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewScheduleDefinitionId, ok = input.Parsed["accessReviewScheduleDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewScheduleDefinitionId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAccessReviewDefinitionID checks that 'input' can be parsed as a Identity Governance Access Review Definition ID
func ValidateIdentityGovernanceAccessReviewDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review Definition ID
func (id IdentityGovernanceAccessReviewDefinitionId) ID() string {
	fmtString := "/identityGovernance/accessReviews/definitions/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewScheduleDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review Definition ID
func (id IdentityGovernanceAccessReviewDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("accessReviews", "accessReviews", "accessReviews"),
		resourceids.StaticSegment("definitions", "definitions", "definitions"),
		resourceids.UserSpecifiedSegment("accessReviewScheduleDefinitionId", "accessReviewScheduleDefinitionId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review Definition ID
func (id IdentityGovernanceAccessReviewDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Schedule Definition: %q", id.AccessReviewScheduleDefinitionId),
	}
	return fmt.Sprintf("Identity Governance Access Review Definition (%s)", strings.Join(components, "\n"))
}
