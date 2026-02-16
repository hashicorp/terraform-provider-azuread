package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePendingAccessReviewInstanceIdStageId{}

// MePendingAccessReviewInstanceIdStageId is a struct representing the Resource ID for a Me Pending Access Review Instance Id Stage
type MePendingAccessReviewInstanceIdStageId struct {
	AccessReviewInstanceId string
	AccessReviewStageId    string
}

// NewMePendingAccessReviewInstanceIdStageID returns a new MePendingAccessReviewInstanceIdStageId struct
func NewMePendingAccessReviewInstanceIdStageID(accessReviewInstanceId string, accessReviewStageId string) MePendingAccessReviewInstanceIdStageId {
	return MePendingAccessReviewInstanceIdStageId{
		AccessReviewInstanceId: accessReviewInstanceId,
		AccessReviewStageId:    accessReviewStageId,
	}
}

// ParseMePendingAccessReviewInstanceIdStageID parses 'input' into a MePendingAccessReviewInstanceIdStageId
func ParseMePendingAccessReviewInstanceIdStageID(input string) (*MePendingAccessReviewInstanceIdStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdStageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceIdStageIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceIdStageId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceIdStageIDInsensitively(input string) (*MePendingAccessReviewInstanceIdStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdStageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePendingAccessReviewInstanceIdStageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewInstanceId, ok = input.Parsed["accessReviewInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", input)
	}

	if id.AccessReviewStageId, ok = input.Parsed["accessReviewStageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", input)
	}

	return nil
}

// ValidateMePendingAccessReviewInstanceIdStageID checks that 'input' can be parsed as a Me Pending Access Review Instance Id Stage ID
func ValidateMePendingAccessReviewInstanceIdStageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceIdStageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Id Stage ID
func (id MePendingAccessReviewInstanceIdStageId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/stages/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewStageId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Id Stage ID
func (id MePendingAccessReviewInstanceIdStageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageId"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Id Stage ID
func (id MePendingAccessReviewInstanceIdStageId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Id Stage (%s)", strings.Join(components, "\n"))
}
