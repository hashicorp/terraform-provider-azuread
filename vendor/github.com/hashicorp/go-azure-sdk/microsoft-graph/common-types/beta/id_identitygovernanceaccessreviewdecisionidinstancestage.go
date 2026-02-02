package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceAccessReviewDecisionIdInstanceStageId{}

// IdentityGovernanceAccessReviewDecisionIdInstanceStageId is a struct representing the Resource ID for a Identity Governance Access Review Decision Id Instance Stage
type IdentityGovernanceAccessReviewDecisionIdInstanceStageId struct {
	AccessReviewInstanceDecisionItemId string
	AccessReviewStageId                string
}

// NewIdentityGovernanceAccessReviewDecisionIdInstanceStageID returns a new IdentityGovernanceAccessReviewDecisionIdInstanceStageId struct
func NewIdentityGovernanceAccessReviewDecisionIdInstanceStageID(accessReviewInstanceDecisionItemId string, accessReviewStageId string) IdentityGovernanceAccessReviewDecisionIdInstanceStageId {
	return IdentityGovernanceAccessReviewDecisionIdInstanceStageId{
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		AccessReviewStageId:                accessReviewStageId,
	}
}

// ParseIdentityGovernanceAccessReviewDecisionIdInstanceStageID parses 'input' into a IdentityGovernanceAccessReviewDecisionIdInstanceStageId
func ParseIdentityGovernanceAccessReviewDecisionIdInstanceStageID(input string) (*IdentityGovernanceAccessReviewDecisionIdInstanceStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDecisionIdInstanceStageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDecisionIdInstanceStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceAccessReviewDecisionIdInstanceStageIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceAccessReviewDecisionIdInstanceStageId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceAccessReviewDecisionIdInstanceStageIDInsensitively(input string) (*IdentityGovernanceAccessReviewDecisionIdInstanceStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceAccessReviewDecisionIdInstanceStageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceAccessReviewDecisionIdInstanceStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceAccessReviewDecisionIdInstanceStageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewInstanceDecisionItemId, ok = input.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", input)
	}

	if id.AccessReviewStageId, ok = input.Parsed["accessReviewStageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", input)
	}

	return nil
}

// ValidateIdentityGovernanceAccessReviewDecisionIdInstanceStageID checks that 'input' can be parsed as a Identity Governance Access Review Decision Id Instance Stage ID
func ValidateIdentityGovernanceAccessReviewDecisionIdInstanceStageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceAccessReviewDecisionIdInstanceStageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Access Review Decision Id Instance Stage ID
func (id IdentityGovernanceAccessReviewDecisionIdInstanceStageId) ID() string {
	fmtString := "/identityGovernance/accessReviews/decisions/%s/instance/stages/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceDecisionItemId, id.AccessReviewStageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Access Review Decision Id Instance Stage ID
func (id IdentityGovernanceAccessReviewDecisionIdInstanceStageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("accessReviews", "accessReviews", "accessReviews"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
		resourceids.StaticSegment("instance", "instance", "instance"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageId"),
	}
}

// String returns a human-readable description of this Identity Governance Access Review Decision Id Instance Stage ID
func (id IdentityGovernanceAccessReviewDecisionIdInstanceStageId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
	}
	return fmt.Sprintf("Identity Governance Access Review Decision Id Instance Stage (%s)", strings.Join(components, "\n"))
}
