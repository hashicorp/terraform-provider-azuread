package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdManagedAppLogCollectionRequestId{}

// UserIdManagedAppLogCollectionRequestId is a struct representing the Resource ID for a User Id Managed App Log Collection Request
type UserIdManagedAppLogCollectionRequestId struct {
	UserId                           string
	ManagedAppLogCollectionRequestId string
}

// NewUserIdManagedAppLogCollectionRequestID returns a new UserIdManagedAppLogCollectionRequestId struct
func NewUserIdManagedAppLogCollectionRequestID(userId string, managedAppLogCollectionRequestId string) UserIdManagedAppLogCollectionRequestId {
	return UserIdManagedAppLogCollectionRequestId{
		UserId:                           userId,
		ManagedAppLogCollectionRequestId: managedAppLogCollectionRequestId,
	}
}

// ParseUserIdManagedAppLogCollectionRequestID parses 'input' into a UserIdManagedAppLogCollectionRequestId
func ParseUserIdManagedAppLogCollectionRequestID(input string) (*UserIdManagedAppLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedAppLogCollectionRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedAppLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdManagedAppLogCollectionRequestIDInsensitively parses 'input' case-insensitively into a UserIdManagedAppLogCollectionRequestId
// note: this method should only be used for API response data and not user input
func ParseUserIdManagedAppLogCollectionRequestIDInsensitively(input string) (*UserIdManagedAppLogCollectionRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdManagedAppLogCollectionRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdManagedAppLogCollectionRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdManagedAppLogCollectionRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ManagedAppLogCollectionRequestId, ok = input.Parsed["managedAppLogCollectionRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "managedAppLogCollectionRequestId", input)
	}

	return nil
}

// ValidateUserIdManagedAppLogCollectionRequestID checks that 'input' can be parsed as a User Id Managed App Log Collection Request ID
func ValidateUserIdManagedAppLogCollectionRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdManagedAppLogCollectionRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Managed App Log Collection Request ID
func (id UserIdManagedAppLogCollectionRequestId) ID() string {
	fmtString := "/users/%s/managedAppLogCollectionRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ManagedAppLogCollectionRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Managed App Log Collection Request ID
func (id UserIdManagedAppLogCollectionRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("managedAppLogCollectionRequests", "managedAppLogCollectionRequests", "managedAppLogCollectionRequests"),
		resourceids.UserSpecifiedSegment("managedAppLogCollectionRequestId", "managedAppLogCollectionRequestId"),
	}
}

// String returns a human-readable description of this User Id Managed App Log Collection Request ID
func (id UserIdManagedAppLogCollectionRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Managed App Log Collection Request: %q", id.ManagedAppLogCollectionRequestId),
	}
	return fmt.Sprintf("User Id Managed App Log Collection Request (%s)", strings.Join(components, "\n"))
}
