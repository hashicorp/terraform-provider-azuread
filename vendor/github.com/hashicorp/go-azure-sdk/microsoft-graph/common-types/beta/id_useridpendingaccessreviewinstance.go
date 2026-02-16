package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPendingAccessReviewInstanceId{}

// UserIdPendingAccessReviewInstanceId is a struct representing the Resource ID for a User Id Pending Access Review Instance
type UserIdPendingAccessReviewInstanceId struct {
	UserId                 string
	AccessReviewInstanceId string
}

// NewUserIdPendingAccessReviewInstanceID returns a new UserIdPendingAccessReviewInstanceId struct
func NewUserIdPendingAccessReviewInstanceID(userId string, accessReviewInstanceId string) UserIdPendingAccessReviewInstanceId {
	return UserIdPendingAccessReviewInstanceId{
		UserId:                 userId,
		AccessReviewInstanceId: accessReviewInstanceId,
	}
}

// ParseUserIdPendingAccessReviewInstanceID parses 'input' into a UserIdPendingAccessReviewInstanceId
func ParseUserIdPendingAccessReviewInstanceID(input string) (*UserIdPendingAccessReviewInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPendingAccessReviewInstanceIDInsensitively parses 'input' case-insensitively into a UserIdPendingAccessReviewInstanceId
// note: this method should only be used for API response data and not user input
func ParseUserIdPendingAccessReviewInstanceIDInsensitively(input string) (*UserIdPendingAccessReviewInstanceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPendingAccessReviewInstanceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.AccessReviewInstanceId, ok = input.Parsed["accessReviewInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", input)
	}

	return nil
}

// ValidateUserIdPendingAccessReviewInstanceID checks that 'input' can be parsed as a User Id Pending Access Review Instance ID
func ValidateUserIdPendingAccessReviewInstanceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPendingAccessReviewInstanceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Pending Access Review Instance ID
func (id UserIdPendingAccessReviewInstanceId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Pending Access Review Instance ID
func (id UserIdPendingAccessReviewInstanceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
	}
}

// String returns a human-readable description of this User Id Pending Access Review Instance ID
func (id UserIdPendingAccessReviewInstanceId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
	}
	return fmt.Sprintf("User Id Pending Access Review Instance (%s)", strings.Join(components, "\n"))
}
