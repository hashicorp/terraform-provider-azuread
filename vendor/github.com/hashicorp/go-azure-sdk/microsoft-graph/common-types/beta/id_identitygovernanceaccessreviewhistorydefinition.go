package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewHistoryDefinitionId{}

// IdentityGovernanceAccessReviewHistoryDefinitionId is a struct representing the Resource ID for a Identity Governance Access Review History Definition
type IdentityGovernanceAccessReviewHistoryDefinitionId struct {
	AccessReviewHistoryDefinitionId string
}

// NewIdentityGovernanceAccessReviewHistoryDefinitionID returns a new IdentityGovernanceAccessReviewHistoryDefinitionId struct
func NewIdentityGovernanceAccessReviewHistoryDefinitionID(accessReviewHistoryDefinitionId string) IdentityGovernanceAccessReviewHistoryDefinitionId {
	return IdentityGovernanceAccessReviewHistoryDefinitionId{
		AccessReviewHistoryDefinitionId: accessReviewHistoryDefinitionId,
	}
}

// ParseIdentityGovernanceAccessReviewHistoryDefinitionID parses 'input' into a IdentityGovernanceAccessReviewHistoryDefinitionId
func ParseIdentityGovernanceAccessReviewHistoryDefinitionID(input string) (*IdentityGovernanceAccessReviewHistoryDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewHistoryDefinitionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewHistoryDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewHistoryDefinitionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewHistoryDefinitionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewHistoryDefinitionIDInsensitively(input string) (*IdentityGovernanceAccessReviewHistoryDefinitionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewHistoryDefinitionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewHistoryDefinitionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewHistoryDefinitionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewHistoryDefinitionId, ok = input.Parsed["accessReviewHistoryDefinitionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewHistoryDefinitionId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAccessReviewHistoryDefinitionID checks that 'input' can be parsed as a Identity Governance Access Review History Definition ID
func ValidateIdentityGovernanceAccessReviewHistoryDefinitionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewHistoryDefinitionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review History Definition ID
func (id IdentityGovernanceAccessReviewHistoryDefinitionId) ID() string {
	fmtString := "/identityGovernance/accessReviews/historyDefinitions/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewHistoryDefinitionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review History Definition ID
func (id IdentityGovernanceAccessReviewHistoryDefinitionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("accessReviews", "accessReviews", "accessReviews"),
		resourceids.StaticSegment("historyDefinitions", "historyDefinitions", "historyDefinitions"),
		resourceids.UserSpecifiedSegment("accessReviewHistoryDefinitionId", "accessReviewHistoryDefinitionId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review History Definition ID
func (id IdentityGovernanceAccessReviewHistoryDefinitionId) String() string {
	components := []string{
		fmt.Sprintf("Access Review History Definition: %q", id.AccessReviewHistoryDefinitionId),
	}
	return fmt.Sprintf("Identity Governance Access Review History Definition (%s)", strings.Join(components, "\n"))
}
