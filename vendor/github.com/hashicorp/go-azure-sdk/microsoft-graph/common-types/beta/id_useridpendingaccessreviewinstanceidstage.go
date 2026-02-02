package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPendingAccessReviewInstanceIdStageId{}

// UserIdPendingAccessReviewInstanceIdStageId is a struct representing the Resource ID for a User Id Pending Access Review Instance Id Stage
type UserIdPendingAccessReviewInstanceIdStageId struct {
	UserId                 string
	AccessReviewInstanceId string
	AccessReviewStageId    string
}

// NewUserIdPendingAccessReviewInstanceIdStageID returns a new UserIdPendingAccessReviewInstanceIdStageId struct
func NewUserIdPendingAccessReviewInstanceIdStageID(userId string, accessReviewInstanceId string, accessReviewStageId string) UserIdPendingAccessReviewInstanceIdStageId {
	return UserIdPendingAccessReviewInstanceIdStageId{
		UserId:                 userId,
		AccessReviewInstanceId: accessReviewInstanceId,
		AccessReviewStageId:    accessReviewStageId,
	}
}

// ParseUserIdPendingAccessReviewInstanceIdStageID parses 'input' into a UserIdPendingAccessReviewInstanceIdStageId
func ParseUserIdPendingAccessReviewInstanceIdStageID(input string) (*UserIdPendingAccessReviewInstanceIdStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceIdStageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceIdStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPendingAccessReviewInstanceIdStageIDInsensitively parses 'input' case-insensitively into a UserIdPendingAccessReviewInstanceIdStageId
// note: this method should only be used for API response data and not user input
func ParseUserIdPendingAccessReviewInstanceIdStageIDInsensitively(input string) (*UserIdPendingAccessReviewInstanceIdStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceIdStageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceIdStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPendingAccessReviewInstanceIdStageId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.AccessReviewInstanceId, ok = input.Parsed["accessReviewInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", input)
	}

	if id.AccessReviewStageId, ok = input.Parsed["accessReviewStageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", input)
	}

	return nil
}

// ValidateUserIdPendingAccessReviewInstanceIdStageID checks that 'input' can be parsed as a User Id Pending Access Review Instance Id Stage ID
func ValidateUserIdPendingAccessReviewInstanceIdStageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPendingAccessReviewInstanceIdStageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Pending Access Review Instance Id Stage ID
func (id UserIdPendingAccessReviewInstanceIdStageId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/stages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewStageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Pending Access Review Instance Id Stage ID
func (id UserIdPendingAccessReviewInstanceIdStageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageId"),
	}
}

// String returns a human-readable description of this User Id Pending Access Review Instance Id Stage ID
func (id UserIdPendingAccessReviewInstanceIdStageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
	}
	return fmt.Sprintf("User Id Pending Access Review Instance Id Stage (%s)", strings.Join(components, "\n"))
}
