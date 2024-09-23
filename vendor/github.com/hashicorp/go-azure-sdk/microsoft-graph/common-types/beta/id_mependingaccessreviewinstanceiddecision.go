package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePendingAccessReviewInstanceIdDecisionId{}

// MePendingAccessReviewInstanceIdDecisionId is a struct representing the Resource ID for a Me Pending Access Review Instance Id Decision
type MePendingAccessReviewInstanceIdDecisionId struct {
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
}

// NewMePendingAccessReviewInstanceIdDecisionID returns a new MePendingAccessReviewInstanceIdDecisionId struct
func NewMePendingAccessReviewInstanceIdDecisionID(accessReviewInstanceId string, accessReviewInstanceDecisionItemId string) MePendingAccessReviewInstanceIdDecisionId {
	return MePendingAccessReviewInstanceIdDecisionId{
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
	}
}

// ParseMePendingAccessReviewInstanceIdDecisionID parses 'input' into a MePendingAccessReviewInstanceIdDecisionId
func ParseMePendingAccessReviewInstanceIdDecisionID(input string) (*MePendingAccessReviewInstanceIdDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdDecisionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdDecisionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceIdDecisionIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceIdDecisionId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceIdDecisionIDInsensitively(input string) (*MePendingAccessReviewInstanceIdDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdDecisionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdDecisionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePendingAccessReviewInstanceIdDecisionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewInstanceId, ok = input.Parsed["accessReviewInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", input)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = input.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", input)
	}

	return nil
}

// ValidateMePendingAccessReviewInstanceIdDecisionID checks that 'input' can be parsed as a Me Pending Access Review Instance Id Decision ID
func ValidateMePendingAccessReviewInstanceIdDecisionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceIdDecisionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Id Decision ID
func (id MePendingAccessReviewInstanceIdDecisionId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/decisions/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Id Decision ID
func (id MePendingAccessReviewInstanceIdDecisionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Id Decision ID
func (id MePendingAccessReviewInstanceIdDecisionId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Id Decision (%s)", strings.Join(components, "\n"))
}
