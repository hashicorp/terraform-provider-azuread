package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdSponsorId{}

// UserIdSponsorId is a struct representing the Resource ID for a User Id Sponsor
type UserIdSponsorId struct {
	UserId            string
	DirectoryObjectId string
}

// NewUserIdSponsorID returns a new UserIdSponsorId struct
func NewUserIdSponsorID(userId string, directoryObjectId string) UserIdSponsorId {
	return UserIdSponsorId{
		UserId:            userId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseUserIdSponsorID parses 'input' into a UserIdSponsorId
func ParseUserIdSponsorID(input string) (*UserIdSponsorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdSponsorId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdSponsorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdSponsorIDInsensitively parses 'input' case-insensitively into a UserIdSponsorId
// note: this method should only be used for API response data and not user input
func ParseUserIdSponsorIDInsensitively(input string) (*UserIdSponsorId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdSponsorId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdSponsorId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdSponsorId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateUserIdSponsorID checks that 'input' can be parsed as a User Id Sponsor ID
func ValidateUserIdSponsorID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdSponsorID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Sponsor ID
func (id UserIdSponsorId) ID() string {
	fmtString := "/users/%s/sponsors/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Sponsor ID
func (id UserIdSponsorId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("sponsors", "sponsors", "sponsors"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this User Id Sponsor ID
func (id UserIdSponsorId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("User Id Sponsor (%s)", strings.Join(components, "\n"))
}
