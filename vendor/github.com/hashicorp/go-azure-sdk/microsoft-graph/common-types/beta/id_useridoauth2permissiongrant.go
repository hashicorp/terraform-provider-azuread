package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOAuth2PermissionGrantId{}

// UserIdOAuth2PermissionGrantId is a struct representing the Resource ID for a User Id O Auth 2 Permission Grant
type UserIdOAuth2PermissionGrantId struct {
	UserId                  string
	OAuth2PermissionGrantId string
}

// NewUserIdOAuth2PermissionGrantID returns a new UserIdOAuth2PermissionGrantId struct
func NewUserIdOAuth2PermissionGrantID(userId string, oAuth2PermissionGrantId string) UserIdOAuth2PermissionGrantId {
	return UserIdOAuth2PermissionGrantId{
		UserId:                  userId,
		OAuth2PermissionGrantId: oAuth2PermissionGrantId,
	}
}

// ParseUserIdOAuth2PermissionGrantID parses 'input' into a UserIdOAuth2PermissionGrantId
func ParseUserIdOAuth2PermissionGrantID(input string) (*UserIdOAuth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOAuth2PermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOAuth2PermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOAuth2PermissionGrantIDInsensitively parses 'input' case-insensitively into a UserIdOAuth2PermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseUserIdOAuth2PermissionGrantIDInsensitively(input string) (*UserIdOAuth2PermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOAuth2PermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOAuth2PermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOAuth2PermissionGrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OAuth2PermissionGrantId, ok = input.Parsed["oAuth2PermissionGrantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "oAuth2PermissionGrantId", input)
	}

	return nil
}

// ValidateUserIdOAuth2PermissionGrantID checks that 'input' can be parsed as a User Id O Auth 2 Permission Grant ID
func ValidateUserIdOAuth2PermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOAuth2PermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id O Auth 2 Permission Grant ID
func (id UserIdOAuth2PermissionGrantId) ID() string {
	fmtString := "/users/%s/oauth2PermissionGrants/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OAuth2PermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id O Auth 2 Permission Grant ID
func (id UserIdOAuth2PermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("oauth2PermissionGrants", "oauth2PermissionGrants", "oauth2PermissionGrants"),
		resourceids.UserSpecifiedSegment("oAuth2PermissionGrantId", "oAuth2PermissionGrantId"),
	}
}

// String returns a human-readable description of this User Id O Auth 2 Permission Grant ID
func (id UserIdOAuth2PermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("O Auth 2 Permission Grant: %q", id.OAuth2PermissionGrantId),
	}
	return fmt.Sprintf("User Id O Auth 2 Permission Grant (%s)", strings.Join(components, "\n"))
}
