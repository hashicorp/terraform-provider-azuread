package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPendingAccessReviewInstanceIdContactedReviewerId{}

// UserIdPendingAccessReviewInstanceIdContactedReviewerId is a struct representing the Resource ID for a User Id Pending Access Review Instance Id Contacted Reviewer
type UserIdPendingAccessReviewInstanceIdContactedReviewerId struct {
	UserId                 string
	AccessReviewInstanceId string
	AccessReviewReviewerId string
}

// NewUserIdPendingAccessReviewInstanceIdContactedReviewerID returns a new UserIdPendingAccessReviewInstanceIdContactedReviewerId struct
func NewUserIdPendingAccessReviewInstanceIdContactedReviewerID(userId string, accessReviewInstanceId string, accessReviewReviewerId string) UserIdPendingAccessReviewInstanceIdContactedReviewerId {
	return UserIdPendingAccessReviewInstanceIdContactedReviewerId{
		UserId:                 userId,
		AccessReviewInstanceId: accessReviewInstanceId,
		AccessReviewReviewerId: accessReviewReviewerId,
	}
}

// ParseUserIdPendingAccessReviewInstanceIdContactedReviewerID parses 'input' into a UserIdPendingAccessReviewInstanceIdContactedReviewerId
func ParseUserIdPendingAccessReviewInstanceIdContactedReviewerID(input string) (*UserIdPendingAccessReviewInstanceIdContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceIdContactedReviewerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceIdContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPendingAccessReviewInstanceIdContactedReviewerIDInsensitively parses 'input' case-insensitively into a UserIdPendingAccessReviewInstanceIdContactedReviewerId
// note: this method should only be used for API response data and not user input
func ParseUserIdPendingAccessReviewInstanceIdContactedReviewerIDInsensitively(input string) (*UserIdPendingAccessReviewInstanceIdContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceIdContactedReviewerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceIdContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPendingAccessReviewInstanceIdContactedReviewerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.AccessReviewInstanceId, ok = input.Parsed["accessReviewInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", input)
	}

	if id.AccessReviewReviewerId, ok = input.Parsed["accessReviewReviewerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewReviewerId", input)
	}

	return nil
}

// ValidateUserIdPendingAccessReviewInstanceIdContactedReviewerID checks that 'input' can be parsed as a User Id Pending Access Review Instance Id Contacted Reviewer ID
func ValidateUserIdPendingAccessReviewInstanceIdContactedReviewerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPendingAccessReviewInstanceIdContactedReviewerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Pending Access Review Instance Id Contacted Reviewer ID
func (id UserIdPendingAccessReviewInstanceIdContactedReviewerId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/contactedReviewers/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewReviewerId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Pending Access Review Instance Id Contacted Reviewer ID
func (id UserIdPendingAccessReviewInstanceIdContactedReviewerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("contactedReviewers", "contactedReviewers", "contactedReviewers"),
		resourceids.UserSpecifiedSegment("accessReviewReviewerId", "accessReviewReviewerId"),
	}
}

// String returns a human-readable description of this User Id Pending Access Review Instance Id Contacted Reviewer ID
func (id UserIdPendingAccessReviewInstanceIdContactedReviewerId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Reviewer: %q", id.AccessReviewReviewerId),
	}
	return fmt.Sprintf("User Id Pending Access Review Instance Id Contacted Reviewer (%s)", strings.Join(components, "\n"))
}
