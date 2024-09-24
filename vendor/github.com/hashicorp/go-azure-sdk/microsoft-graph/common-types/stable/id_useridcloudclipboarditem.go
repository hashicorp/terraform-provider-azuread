package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdCloudClipboardItemId{}

// UserIdCloudClipboardItemId is a struct representing the Resource ID for a User Id Cloud Clipboard Item
type UserIdCloudClipboardItemId struct {
	UserId               string
	CloudClipboardItemId string
}

// NewUserIdCloudClipboardItemID returns a new UserIdCloudClipboardItemId struct
func NewUserIdCloudClipboardItemID(userId string, cloudClipboardItemId string) UserIdCloudClipboardItemId {
	return UserIdCloudClipboardItemId{
		UserId:               userId,
		CloudClipboardItemId: cloudClipboardItemId,
	}
}

// ParseUserIdCloudClipboardItemID parses 'input' into a UserIdCloudClipboardItemId
func ParseUserIdCloudClipboardItemID(input string) (*UserIdCloudClipboardItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCloudClipboardItemId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCloudClipboardItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdCloudClipboardItemIDInsensitively parses 'input' case-insensitively into a UserIdCloudClipboardItemId
// note: this method should only be used for API response data and not user input
func ParseUserIdCloudClipboardItemIDInsensitively(input string) (*UserIdCloudClipboardItemId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdCloudClipboardItemId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdCloudClipboardItemId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdCloudClipboardItemId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.CloudClipboardItemId, ok = input.Parsed["cloudClipboardItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudClipboardItemId", input)
	}

	return nil
}

// ValidateUserIdCloudClipboardItemID checks that 'input' can be parsed as a User Id Cloud Clipboard Item ID
func ValidateUserIdCloudClipboardItemID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdCloudClipboardItemID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Cloud Clipboard Item ID
func (id UserIdCloudClipboardItemId) ID() string {
	fmtString := "/users/%s/cloudClipboard/items/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.CloudClipboardItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Cloud Clipboard Item ID
func (id UserIdCloudClipboardItemId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("cloudClipboard", "cloudClipboard", "cloudClipboard"),
		resourceids.StaticSegment("items", "items", "items"),
		resourceids.UserSpecifiedSegment("cloudClipboardItemId", "cloudClipboardItemId"),
	}
}

// String returns a human-readable description of this User Id Cloud Clipboard Item ID
func (id UserIdCloudClipboardItemId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Cloud Clipboard Item: %q", id.CloudClipboardItemId),
	}
	return fmt.Sprintf("User Id Cloud Clipboard Item (%s)", strings.Join(components, "\n"))
}
