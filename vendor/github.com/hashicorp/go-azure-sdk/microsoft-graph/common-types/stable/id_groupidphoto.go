package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdPhotoId{}

// GroupIdPhotoId is a struct representing the Resource ID for a Group Id Photo
type GroupIdPhotoId struct {
	GroupId        string
	ProfilePhotoId string
}

// NewGroupIdPhotoID returns a new GroupIdPhotoId struct
func NewGroupIdPhotoID(groupId string, profilePhotoId string) GroupIdPhotoId {
	return GroupIdPhotoId{
		GroupId:        groupId,
		ProfilePhotoId: profilePhotoId,
	}
}

// ParseGroupIdPhotoID parses 'input' into a GroupIdPhotoId
func ParseGroupIdPhotoID(input string) (*GroupIdPhotoId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdPhotoId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdPhotoId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdPhotoIDInsensitively parses 'input' case-insensitively into a GroupIdPhotoId
// note: this method should only be used for API response data and not user input
func ParseGroupIdPhotoIDInsensitively(input string) (*GroupIdPhotoId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdPhotoId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdPhotoId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdPhotoId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.ProfilePhotoId, ok = input.Parsed["profilePhotoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "profilePhotoId", input)
	}

	return nil
}

// ValidateGroupIdPhotoID checks that 'input' can be parsed as a Group Id Photo ID
func ValidateGroupIdPhotoID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdPhotoID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Photo ID
func (id GroupIdPhotoId) ID() string {
	fmtString := "/groups/%s/photos/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.ProfilePhotoId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Photo ID
func (id GroupIdPhotoId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("photos", "photos", "photos"),
		resourceids.UserSpecifiedSegment("profilePhotoId", "profilePhotoId"),
	}
}

// String returns a human-readable description of this Group Id Photo ID
func (id GroupIdPhotoId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Profile Photo: %q", id.ProfilePhotoId),
	}
	return fmt.Sprintf("Group Id Photo (%s)", strings.Join(components, "\n"))
}
