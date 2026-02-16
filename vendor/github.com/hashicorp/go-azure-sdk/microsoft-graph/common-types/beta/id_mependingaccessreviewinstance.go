package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePendingAccessReviewInstanceId{}

// MePendingAccessReviewInstanceId is a struct representing the Resource ID for a Me Pending Access Review Instance
type MePendingAccessReviewInstanceId struct {
	AccessReviewInstanceId string
}

// NewMePendingAccessReviewInstanceID returns a new MePendingAccessReviewInstanceId struct
func NewMePendingAccessReviewInstanceID(accessReviewInstanceId string) MePendingAccessReviewInstanceId {
	return MePendingAccessReviewInstanceId{
		AccessReviewInstanceId: accessReviewInstanceId,
	}
}

// ParseMePendingAccessReviewInstanceID parses 'input' into a MePendingAccessReviewInstanceId
func ParseMePendingAccessReviewInstanceID(input string) (*MePendingAccessReviewInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceIDInsensitively(input string) (*MePendingAccessReviewInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePendingAccessReviewInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewInstanceId, ok = input.Parsed["accessReviewInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", input)
	}

	return nil
}

// ValidateMePendingAccessReviewInstanceID checks that 'input' can be parsed as a Me Pending Access Review Instance ID
func ValidateMePendingAccessReviewInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance ID
func (id MePendingAccessReviewInstanceId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance ID
func (id MePendingAccessReviewInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance ID
func (id MePendingAccessReviewInstanceId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance (%s)", strings.Join(components, "\n"))
}
