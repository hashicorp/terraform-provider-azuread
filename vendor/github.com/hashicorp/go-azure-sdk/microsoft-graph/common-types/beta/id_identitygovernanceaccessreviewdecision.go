package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDecisionId{}

// IdentityGovernanceAccessReviewDecisionId is a struct representing the Resource ID for a Identity Governance Access Review Decision
type IdentityGovernanceAccessReviewDecisionId struct {
	AccessReviewInstanceDecisionItemId string
}

// NewIdentityGovernanceAccessReviewDecisionID returns a new IdentityGovernanceAccessReviewDecisionId struct
func NewIdentityGovernanceAccessReviewDecisionID(accessReviewInstanceDecisionItemId string) IdentityGovernanceAccessReviewDecisionId {
	return IdentityGovernanceAccessReviewDecisionId{
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
	}
}

// ParseIdentityGovernanceAccessReviewDecisionID parses 'input' into a IdentityGovernanceAccessReviewDecisionId
func ParseIdentityGovernanceAccessReviewDecisionID(input string) (*IdentityGovernanceAccessReviewDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDecisionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDecisionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewDecisionIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewDecisionId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewDecisionIDInsensitively(input string) (*IdentityGovernanceAccessReviewDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDecisionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDecisionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewDecisionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewInstanceDecisionItemId, ok = input.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAccessReviewDecisionID checks that 'input' can be parsed as a Identity Governance Access Review Decision ID
func ValidateIdentityGovernanceAccessReviewDecisionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewDecisionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review Decision ID
func (id IdentityGovernanceAccessReviewDecisionId) ID() string {
	fmtString := "/identityGovernance/accessReviews/decisions/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceDecisionItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review Decision ID
func (id IdentityGovernanceAccessReviewDecisionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("accessReviews", "accessReviews", "accessReviews"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review Decision ID
func (id IdentityGovernanceAccessReviewDecisionId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
	}
	return fmt.Sprintf("Identity Governance Access Review Decision (%s)", strings.Join(components, "\n"))
}
