package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdOnenoteOperationId{}

// UserIdOnenoteOperationId is a struct representing the Resource ID for a User Id Onenote Operation
type UserIdOnenoteOperationId struct {
	UserId             string
	OnenoteOperationId string
}

// NewUserIdOnenoteOperationID returns a new UserIdOnenoteOperationId struct
func NewUserIdOnenoteOperationID(userId string, onenoteOperationId string) UserIdOnenoteOperationId {
	return UserIdOnenoteOperationId{
		UserId:             userId,
		OnenoteOperationId: onenoteOperationId,
	}
}

// ParseUserIdOnenoteOperationID parses 'input' into a UserIdOnenoteOperationId
func ParseUserIdOnenoteOperationID(input string) (*UserIdOnenoteOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteOperationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdOnenoteOperationIDInsensitively parses 'input' case-insensitively into a UserIdOnenoteOperationId
// note: this method should only be used for API response data and not user input
func ParseUserIdOnenoteOperationIDInsensitively(input string) (*UserIdOnenoteOperationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdOnenoteOperationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdOnenoteOperationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdOnenoteOperationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.OnenoteOperationId, ok = input.Parsed["onenoteOperationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "onenoteOperationId", input)
	}

	return nil
}

// ValidateUserIdOnenoteOperationID checks that 'input' can be parsed as a User Id Onenote Operation ID
func ValidateUserIdOnenoteOperationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdOnenoteOperationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Onenote Operation ID
func (id UserIdOnenoteOperationId) ID() string {
	fmtString := "/users/%s/onenote/operations/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.OnenoteOperationId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Onenote Operation ID
func (id UserIdOnenoteOperationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("onenote", "onenote", "onenote"),
		resourceids.StaticSegment("operations", "operations", "operations"),
		resourceids.UserSpecifiedSegment("onenoteOperationId", "onenoteOperationId"),
	}
}

// String returns a human-readable description of this User Id Onenote Operation ID
func (id UserIdOnenoteOperationId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Onenote Operation: %q", id.OnenoteOperationId),
	}
	return fmt.Sprintf("User Id Onenote Operation (%s)", strings.Join(components, "\n"))
}
