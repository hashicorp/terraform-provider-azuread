package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfilePublicationId{}

// UserIdProfilePublicationId is a struct representing the Resource ID for a User Id Profile Publication
type UserIdProfilePublicationId struct {
	UserId            string
	ItemPublicationId string
}

// NewUserIdProfilePublicationID returns a new UserIdProfilePublicationId struct
func NewUserIdProfilePublicationID(userId string, itemPublicationId string) UserIdProfilePublicationId {
	return UserIdProfilePublicationId{
		UserId:            userId,
		ItemPublicationId: itemPublicationId,
	}
}

// ParseUserIdProfilePublicationID parses 'input' into a UserIdProfilePublicationId
func ParseUserIdProfilePublicationID(input string) (*UserIdProfilePublicationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfilePublicationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfilePublicationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfilePublicationIDInsensitively parses 'input' case-insensitively into a UserIdProfilePublicationId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfilePublicationIDInsensitively(input string) (*UserIdProfilePublicationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfilePublicationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfilePublicationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfilePublicationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ItemPublicationId, ok = input.Parsed["itemPublicationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemPublicationId", input)
	}

	return nil
}

// ValidateUserIdProfilePublicationID checks that 'input' can be parsed as a User Id Profile Publication ID
func ValidateUserIdProfilePublicationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfilePublicationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Publication ID
func (id UserIdProfilePublicationId) ID() string {
	fmtString := "/users/%s/profile/publications/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ItemPublicationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Publication ID
func (id UserIdProfilePublicationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("publications", "publications", "publications"),
		resourceids.UserSpecifiedSegment("itemPublicationId", "itemPublicationId"),
	}
}

// String returns a human-readable description of this User Id Profile Publication ID
func (id UserIdProfilePublicationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Item Publication: %q", id.ItemPublicationId),
	}
	return fmt.Sprintf("User Id Profile Publication (%s)", strings.Join(components, "\n"))
}
