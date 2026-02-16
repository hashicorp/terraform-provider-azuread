package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePhotoId{}

// MePhotoId is a struct representing the Resource ID for a Me Photo
type MePhotoId struct {
	ProfilePhotoId string
}

// NewMePhotoID returns a new MePhotoId struct
func NewMePhotoID(profilePhotoId string) MePhotoId {
	return MePhotoId{
		ProfilePhotoId: profilePhotoId,
	}
}

// ParseMePhotoID parses 'input' into a MePhotoId
func ParseMePhotoID(input string) (*MePhotoId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePhotoId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePhotoId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePhotoIDInsensitively parses 'input' case-insensitively into a MePhotoId
// note: this method should only be used for API response data and not user input
func ParseMePhotoIDInsensitively(input string) (*MePhotoId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePhotoId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePhotoId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePhotoId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ProfilePhotoId, ok = input.Parsed["profilePhotoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "profilePhotoId", input)
	}

	return nil
}

// ValidateMePhotoID checks that 'input' can be parsed as a Me Photo ID
func ValidateMePhotoID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePhotoID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Photo ID
func (id MePhotoId) ID() string {
	fmtString := "/me/photos/%s"
	return fmt.Sprintf(fmtString, id.ProfilePhotoId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Photo ID
func (id MePhotoId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("photos", "photos", "photos"),
		resourceids.UserSpecifiedSegment("profilePhotoId", "profilePhotoId"),
	}
}

// String returns a human-readable description of this Me Photo ID
func (id MePhotoId) String() string {
	components := []string{
		fmt.Sprintf("Profile Photo: %q", id.ProfilePhotoId),
	}
	return fmt.Sprintf("Me Photo (%s)", strings.Join(components, "\n"))
}
