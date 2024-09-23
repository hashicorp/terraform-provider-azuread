package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPhotoId{}

// UserIdPhotoId is a struct representing the Resource ID for a User Id Photo
type UserIdPhotoId struct {
	UserId         string
	ProfilePhotoId string
}

// NewUserIdPhotoID returns a new UserIdPhotoId struct
func NewUserIdPhotoID(userId string, profilePhotoId string) UserIdPhotoId {
	return UserIdPhotoId{
		UserId:         userId,
		ProfilePhotoId: profilePhotoId,
	}
}

// ParseUserIdPhotoID parses 'input' into a UserIdPhotoId
func ParseUserIdPhotoID(input string) (*UserIdPhotoId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPhotoId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPhotoId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPhotoIDInsensitively parses 'input' case-insensitively into a UserIdPhotoId
// note: this method should only be used for API response data and not user input
func ParseUserIdPhotoIDInsensitively(input string) (*UserIdPhotoId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPhotoId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPhotoId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPhotoId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ProfilePhotoId, ok = input.Parsed["profilePhotoId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "profilePhotoId", input)
	}

	return nil
}

// ValidateUserIdPhotoID checks that 'input' can be parsed as a User Id Photo ID
func ValidateUserIdPhotoID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPhotoID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Photo ID
func (id UserIdPhotoId) ID() string {
	fmtString := "/users/%s/photos/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ProfilePhotoId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Photo ID
func (id UserIdPhotoId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("photos", "photos", "photos"),
		resourceids.UserSpecifiedSegment("profilePhotoId", "profilePhotoId"),
	}
}

// String returns a human-readable description of this User Id Photo ID
func (id UserIdPhotoId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Profile Photo: %q", id.ProfilePhotoId),
	}
	return fmt.Sprintf("User Id Photo (%s)", strings.Join(components, "\n"))
}
