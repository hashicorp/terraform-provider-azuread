package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPermissionGrantId{}

// UserIdPermissionGrantId is a struct representing the Resource ID for a User Id Permission Grant
type UserIdPermissionGrantId struct {
	UserId                            string
	ResourceSpecificPermissionGrantId string
}

// NewUserIdPermissionGrantID returns a new UserIdPermissionGrantId struct
func NewUserIdPermissionGrantID(userId string, resourceSpecificPermissionGrantId string) UserIdPermissionGrantId {
	return UserIdPermissionGrantId{
		UserId:                            userId,
		ResourceSpecificPermissionGrantId: resourceSpecificPermissionGrantId,
	}
}

// ParseUserIdPermissionGrantID parses 'input' into a UserIdPermissionGrantId
func ParseUserIdPermissionGrantID(input string) (*UserIdPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPermissionGrantId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPermissionGrantIDInsensitively parses 'input' case-insensitively into a UserIdPermissionGrantId
// note: this method should only be used for API response data and not user input
func ParseUserIdPermissionGrantIDInsensitively(input string) (*UserIdPermissionGrantId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPermissionGrantId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPermissionGrantId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPermissionGrantId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ResourceSpecificPermissionGrantId, ok = input.Parsed["resourceSpecificPermissionGrantId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "resourceSpecificPermissionGrantId", input)
	}

	return nil
}

// ValidateUserIdPermissionGrantID checks that 'input' can be parsed as a User Id Permission Grant ID
func ValidateUserIdPermissionGrantID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPermissionGrantID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Permission Grant ID
func (id UserIdPermissionGrantId) ID() string {
	fmtString := "/users/%s/permissionGrants/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ResourceSpecificPermissionGrantId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Permission Grant ID
func (id UserIdPermissionGrantId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("permissionGrants", "permissionGrants", "permissionGrants"),
		resourceids.UserSpecifiedSegment("resourceSpecificPermissionGrantId", "resourceSpecificPermissionGrantId"),
	}
}

// String returns a human-readable description of this User Id Permission Grant ID
func (id UserIdPermissionGrantId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Resource Specific Permission Grant: %q", id.ResourceSpecificPermissionGrantId),
	}
	return fmt.Sprintf("User Id Permission Grant (%s)", strings.Join(components, "\n"))
}
