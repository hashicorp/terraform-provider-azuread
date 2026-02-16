package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDecisionIdInsightId{}

// IdentityGovernanceAccessReviewDecisionIdInsightId is a struct representing the Resource ID for a Identity Governance Access Review Decision Id Insight
type IdentityGovernanceAccessReviewDecisionIdInsightId struct {
	AccessReviewInstanceDecisionItemId string
	GovernanceInsightId                string
}

// NewIdentityGovernanceAccessReviewDecisionIdInsightID returns a new IdentityGovernanceAccessReviewDecisionIdInsightId struct
func NewIdentityGovernanceAccessReviewDecisionIdInsightID(accessReviewInstanceDecisionItemId string, governanceInsightId string) IdentityGovernanceAccessReviewDecisionIdInsightId {
	return IdentityGovernanceAccessReviewDecisionIdInsightId{
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		GovernanceInsightId:                governanceInsightId,
	}
}

// ParseIdentityGovernanceAccessReviewDecisionIdInsightID parses 'input' into a IdentityGovernanceAccessReviewDecisionIdInsightId
func ParseIdentityGovernanceAccessReviewDecisionIdInsightID(input string) (*IdentityGovernanceAccessReviewDecisionIdInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDecisionIdInsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDecisionIdInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewDecisionIdInsightIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewDecisionIdInsightId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewDecisionIdInsightIDInsensitively(input string) (*IdentityGovernanceAccessReviewDecisionIdInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDecisionIdInsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDecisionIdInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewDecisionIdInsightId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewInstanceDecisionItemId, ok = input.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", input)
	}

	if id.GovernanceInsightId, ok = input.Parsed["governanceInsightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAccessReviewDecisionIdInsightID checks that 'input' can be parsed as a Identity Governance Access Review Decision Id Insight ID
func ValidateIdentityGovernanceAccessReviewDecisionIdInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewDecisionIdInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review Decision Id Insight ID
func (id IdentityGovernanceAccessReviewDecisionIdInsightId) ID() string {
	fmtString := "/identityGovernance/accessReviews/decisions/%s/insights/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceDecisionItemId, id.GovernanceInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review Decision Id Insight ID
func (id IdentityGovernanceAccessReviewDecisionIdInsightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("accessReviews", "accessReviews", "accessReviews"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
		resourceids.StaticSegment("insights", "insights", "insights"),
		resourceids.UserSpecifiedSegment("governanceInsightId", "governanceInsightId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review Decision Id Insight ID
func (id IdentityGovernanceAccessReviewDecisionIdInsightId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Governance Insight: %q", id.GovernanceInsightId),
	}
	return fmt.Sprintf("Identity Governance Access Review Decision Id Insight (%s)", strings.Join(components, "\n"))
}
