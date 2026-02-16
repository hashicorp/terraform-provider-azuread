package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdProfileAddressId{}

// UserIdProfileAddressId is a struct representing the Resource ID for a User Id Profile Address
type UserIdProfileAddressId struct {
	UserId        string
	ItemAddressId string
}

// NewUserIdProfileAddressID returns a new UserIdProfileAddressId struct
func NewUserIdProfileAddressID(userId string, itemAddressId string) UserIdProfileAddressId {
	return UserIdProfileAddressId{
		UserId:        userId,
		ItemAddressId: itemAddressId,
	}
}

// ParseUserIdProfileAddressID parses 'input' into a UserIdProfileAddressId
func ParseUserIdProfileAddressID(input string) (*UserIdProfileAddressId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileAddressId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileAddressId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdProfileAddressIDInsensitively parses 'input' case-insensitively into a UserIdProfileAddressId
// note: this method should only be used for API response data and not user input
func ParseUserIdProfileAddressIDInsensitively(input string) (*UserIdProfileAddressId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdProfileAddressId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdProfileAddressId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdProfileAddressId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ItemAddressId, ok = input.Parsed["itemAddressId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "itemAddressId", input)
	}

	return nil
}

// ValidateUserIdProfileAddressID checks that 'input' can be parsed as a User Id Profile Address ID
func ValidateUserIdProfileAddressID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdProfileAddressID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Profile Address ID
func (id UserIdProfileAddressId) ID() string {
	fmtString := "/users/%s/profile/addresses/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ItemAddressId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Profile Address ID
func (id UserIdProfileAddressId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("profile", "profile", "profile"),
		resourceids.StaticSegment("addresses", "addresses", "addresses"),
		resourceids.UserSpecifiedSegment("itemAddressId", "itemAddressId"),
	}
}

// String returns a human-readable description of this User Id Profile Address ID
func (id UserIdProfileAddressId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Item Address: %q", id.ItemAddressId),
	}
	return fmt.Sprintf("User Id Profile Address (%s)", strings.Join(components, "\n"))
}
