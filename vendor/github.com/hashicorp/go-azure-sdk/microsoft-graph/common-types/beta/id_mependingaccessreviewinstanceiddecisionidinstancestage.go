package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePendingAccessReviewInstanceIdDecisionIdInstanceStageId{}

// MePendingAccessReviewInstanceIdDecisionIdInstanceStageId is a struct representing the Resource ID for a Me Pending Access Review Instance Id Decision Id Instance Stage
type MePendingAccessReviewInstanceIdDecisionIdInstanceStageId struct {
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
	AccessReviewStageId                string
}

// NewMePendingAccessReviewInstanceIdDecisionIdInstanceStageID returns a new MePendingAccessReviewInstanceIdDecisionIdInstanceStageId struct
func NewMePendingAccessReviewInstanceIdDecisionIdInstanceStageID(accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, accessReviewStageId string) MePendingAccessReviewInstanceIdDecisionIdInstanceStageId {
	return MePendingAccessReviewInstanceIdDecisionIdInstanceStageId{
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		AccessReviewStageId:                accessReviewStageId,
	}
}

// ParseMePendingAccessReviewInstanceIdDecisionIdInstanceStageID parses 'input' into a MePendingAccessReviewInstanceIdDecisionIdInstanceStageId
func ParseMePendingAccessReviewInstanceIdDecisionIdInstanceStageID(input string) (*MePendingAccessReviewInstanceIdDecisionIdInstanceStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdDecisionIdInstanceStageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdDecisionIdInstanceStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceIdDecisionIdInstanceStageIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceIdDecisionIdInstanceStageId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceIdDecisionIdInstanceStageIDInsensitively(input string) (*MePendingAccessReviewInstanceIdDecisionIdInstanceStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdDecisionIdInstanceStageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdDecisionIdInstanceStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePendingAccessReviewInstanceIdDecisionIdInstanceStageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewInstanceId, ok = input.Parsed["accessReviewInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", input)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = input.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", input)
	}

	if id.AccessReviewStageId, ok = input.Parsed["accessReviewStageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", input)
	}

	return nil
}

// ValidateMePendingAccessReviewInstanceIdDecisionIdInstanceStageID checks that 'input' can be parsed as a Me Pending Access Review Instance Id Decision Id Instance Stage ID
func ValidateMePendingAccessReviewInstanceIdDecisionIdInstanceStageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceIdDecisionIdInstanceStageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Id Decision Id Instance Stage ID
func (id MePendingAccessReviewInstanceIdDecisionIdInstanceStageId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/decisions/%s/instance/stages/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewStageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Id Decision Id Instance Stage ID
func (id MePendingAccessReviewInstanceIdDecisionIdInstanceStageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
		resourceids.StaticSegment("instance", "instance", "instance"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageId"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Id Decision Id Instance Stage ID
func (id MePendingAccessReviewInstanceIdDecisionIdInstanceStageId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Id Decision Id Instance Stage (%s)", strings.Join(components, "\n"))
}
