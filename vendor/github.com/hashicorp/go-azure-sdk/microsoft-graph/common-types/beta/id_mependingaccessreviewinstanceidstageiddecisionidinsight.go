package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId{}

// MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId is a struct representing the Resource ID for a Me Pending Access Review Instance Id Stage Id Decision Id Insight
type MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId struct {
	AccessReviewInstanceId             string
	AccessReviewStageId                string
	AccessReviewInstanceDecisionItemId string
	GovernanceInsightId                string
}

// NewMePendingAccessReviewInstanceIdStageIdDecisionIdInsightID returns a new MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId struct
func NewMePendingAccessReviewInstanceIdStageIdDecisionIdInsightID(accessReviewInstanceId string, accessReviewStageId string, accessReviewInstanceDecisionItemId string, governanceInsightId string) MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId {
	return MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId{
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewStageId:                accessReviewStageId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		GovernanceInsightId:                governanceInsightId,
	}
}

// ParseMePendingAccessReviewInstanceIdStageIdDecisionIdInsightID parses 'input' into a MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId
func ParseMePendingAccessReviewInstanceIdStageIdDecisionIdInsightID(input string) (*MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceIdStageIdDecisionIdInsightIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceIdStageIdDecisionIdInsightIDInsensitively(input string) (*MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewInstanceId, ok = input.Parsed["accessReviewInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", input)
	}

	if id.AccessReviewStageId, ok = input.Parsed["accessReviewStageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", input)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = input.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", input)
	}

	if id.GovernanceInsightId, ok = input.Parsed["governanceInsightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", input)
	}

	return nil
}

// ValidateMePendingAccessReviewInstanceIdStageIdDecisionIdInsightID checks that 'input' can be parsed as a Me Pending Access Review Instance Id Stage Id Decision Id Insight ID
func ValidateMePendingAccessReviewInstanceIdStageIdDecisionIdInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceIdStageIdDecisionIdInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Id Stage Id Decision Id Insight ID
func (id MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/stages/%s/decisions/%s/insights/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId, id.GovernanceInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Id Stage Id Decision Id Insight ID
func (id MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
		resourceids.StaticSegment("insights", "insights", "insights"),
		resourceids.UserSpecifiedSegment("governanceInsightId", "governanceInsightId"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Id Stage Id Decision Id Insight ID
func (id MePendingAccessReviewInstanceIdStageIdDecisionIdInsightId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Governance Insight: %q", id.GovernanceInsightId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Id Stage Id Decision Id Insight (%s)", strings.Join(components, "\n"))
}
