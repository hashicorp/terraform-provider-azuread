package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePendingAccessReviewInstanceIdContactedReviewerId{}

// MePendingAccessReviewInstanceIdContactedReviewerId is a struct representing the Resource ID for a Me Pending Access Review Instance Id Contacted Reviewer
type MePendingAccessReviewInstanceIdContactedReviewerId struct {
	AccessReviewInstanceId string
	AccessReviewReviewerId string
}

// NewMePendingAccessReviewInstanceIdContactedReviewerID returns a new MePendingAccessReviewInstanceIdContactedReviewerId struct
func NewMePendingAccessReviewInstanceIdContactedReviewerID(accessReviewInstanceId string, accessReviewReviewerId string) MePendingAccessReviewInstanceIdContactedReviewerId {
	return MePendingAccessReviewInstanceIdContactedReviewerId{
		AccessReviewInstanceId: accessReviewInstanceId,
		AccessReviewReviewerId: accessReviewReviewerId,
	}
}

// ParseMePendingAccessReviewInstanceIdContactedReviewerID parses 'input' into a MePendingAccessReviewInstanceIdContactedReviewerId
func ParseMePendingAccessReviewInstanceIdContactedReviewerID(input string) (*MePendingAccessReviewInstanceIdContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdContactedReviewerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceIdContactedReviewerIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceIdContactedReviewerId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceIdContactedReviewerIDInsensitively(input string) (*MePendingAccessReviewInstanceIdContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdContactedReviewerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePendingAccessReviewInstanceIdContactedReviewerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.AccessReviewInstanceId, ok = input.Parsed["accessReviewInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", input)
	}

	if id.AccessReviewReviewerId, ok = input.Parsed["accessReviewReviewerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewReviewerId", input)
	}

	return nil
}

// ValidateMePendingAccessReviewInstanceIdContactedReviewerID checks that 'input' can be parsed as a Me Pending Access Review Instance Id Contacted Reviewer ID
func ValidateMePendingAccessReviewInstanceIdContactedReviewerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceIdContactedReviewerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Id Contacted Reviewer ID
func (id MePendingAccessReviewInstanceIdContactedReviewerId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/contactedReviewers/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewReviewerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Id Contacted Reviewer ID
func (id MePendingAccessReviewInstanceIdContactedReviewerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("contactedReviewers", "contactedReviewers", "contactedReviewers"),
		resourceids.UserSpecifiedSegment("accessReviewReviewerId", "accessReviewReviewerId"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Id Contacted Reviewer ID
func (id MePendingAccessReviewInstanceIdContactedReviewerId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Reviewer: %q", id.AccessReviewReviewerId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Id Contacted Reviewer (%s)", strings.Join(components, "\n"))
}
