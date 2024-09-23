package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPendingAccessReviewInstanceIdDecisionId{}

// UserIdPendingAccessReviewInstanceIdDecisionId is a struct representing the Resource ID for a User Id Pending Access Review Instance Id Decision
type UserIdPendingAccessReviewInstanceIdDecisionId struct {
	UserId                             string
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
}

// NewUserIdPendingAccessReviewInstanceIdDecisionID returns a new UserIdPendingAccessReviewInstanceIdDecisionId struct
func NewUserIdPendingAccessReviewInstanceIdDecisionID(userId string, accessReviewInstanceId string, accessReviewInstanceDecisionItemId string) UserIdPendingAccessReviewInstanceIdDecisionId {
	return UserIdPendingAccessReviewInstanceIdDecisionId{
		UserId:                             userId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
	}
}

// ParseUserIdPendingAccessReviewInstanceIdDecisionID parses 'input' into a UserIdPendingAccessReviewInstanceIdDecisionId
func ParseUserIdPendingAccessReviewInstanceIdDecisionID(input string) (*UserIdPendingAccessReviewInstanceIdDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceIdDecisionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceIdDecisionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPendingAccessReviewInstanceIdDecisionIDInsensitively parses 'input' case-insensitively into a UserIdPendingAccessReviewInstanceIdDecisionId
// note: this method should only be used for API response data and not user input
func ParseUserIdPendingAccessReviewInstanceIdDecisionIDInsensitively(input string) (*UserIdPendingAccessReviewInstanceIdDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceIdDecisionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceIdDecisionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPendingAccessReviewInstanceIdDecisionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.AccessReviewInstanceId, ok = input.Parsed["accessReviewInstanceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceId", input)
	}

	if id.AccessReviewInstanceDecisionItemId, ok = input.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", input)
	}

	return nil
}

// ValidateUserIdPendingAccessReviewInstanceIdDecisionID checks that 'input' can be parsed as a User Id Pending Access Review Instance Id Decision ID
func ValidateUserIdPendingAccessReviewInstanceIdDecisionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPendingAccessReviewInstanceIdDecisionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Pending Access Review Instance Id Decision ID
func (id UserIdPendingAccessReviewInstanceIdDecisionId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/decisions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Pending Access Review Instance Id Decision ID
func (id UserIdPendingAccessReviewInstanceIdDecisionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
	}
}

// String returns a human-readable description of this User Id Pending Access Review Instance Id Decision ID
func (id UserIdPendingAccessReviewInstanceIdDecisionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
	}
	return fmt.Sprintf("User Id Pending Access Review Instance Id Decision (%s)", strings.Join(components, "\n"))
}
