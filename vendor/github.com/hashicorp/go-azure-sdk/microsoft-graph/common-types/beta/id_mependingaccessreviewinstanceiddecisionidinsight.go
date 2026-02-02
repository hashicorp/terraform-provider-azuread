package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePendingAccessReviewInstanceIdDecisionIdInsightId{}

// MePendingAccessReviewInstanceIdDecisionIdInsightId is a struct representing the Resource ID for a Me Pending Access Review Instance Id Decision Id Insight
type MePendingAccessReviewInstanceIdDecisionIdInsightId struct {
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
	GovernanceInsightId                string
}

// NewMePendingAccessReviewInstanceIdDecisionIdInsightID returns a new MePendingAccessReviewInstanceIdDecisionIdInsightId struct
func NewMePendingAccessReviewInstanceIdDecisionIdInsightID(accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, governanceInsightId string) MePendingAccessReviewInstanceIdDecisionIdInsightId {
	return MePendingAccessReviewInstanceIdDecisionIdInsightId{
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		GovernanceInsightId:                governanceInsightId,
	}
}

// ParseMePendingAccessReviewInstanceIdDecisionIdInsightID parses 'input' into a MePendingAccessReviewInstanceIdDecisionIdInsightId
func ParseMePendingAccessReviewInstanceIdDecisionIdInsightID(input string) (*MePendingAccessReviewInstanceIdDecisionIdInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdDecisionIdInsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdDecisionIdInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceIdDecisionIdInsightIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceIdDecisionIdInsightId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceIdDecisionIdInsightIDInsensitively(input string) (*MePendingAccessReviewInstanceIdDecisionIdInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdDecisionIdInsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdDecisionIdInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePendingAccessReviewInstanceIdDecisionIdInsightId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewInstanceId, ok = input.Parsed["accessReviewInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", input)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = input.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", input)
	}

	if id.GovernanceInsightId, ok = input.Parsed["governanceInsightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", input)
	}

	return nil
}

// ValidateMePendingAccessReviewInstanceIdDecisionIdInsightID checks that 'input' can be parsed as a Me Pending Access Review Instance Id Decision Id Insight ID
func ValidateMePendingAccessReviewInstanceIdDecisionIdInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceIdDecisionIdInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Id Decision Id Insight ID
func (id MePendingAccessReviewInstanceIdDecisionIdInsightId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/decisions/%s/insights/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.GovernanceInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Id Decision Id Insight ID
func (id MePendingAccessReviewInstanceIdDecisionIdInsightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
		resourceids.StaticSegment("insights", "insights", "insights"),
		resourceids.UserSpecifiedSegment("governanceInsightId", "governanceInsightId"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Id Decision Id Insight ID
func (id MePendingAccessReviewInstanceIdDecisionIdInsightId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Governance Insight: %q", id.GovernanceInsightId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Id Decision Id Insight (%s)", strings.Join(components, "\n"))
}
