package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId{}

// UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId is a struct representing the Resource ID for a User Id Pending Access Review Instance Id Decision Id Instance Stage
type UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId struct {
	UserId                             string
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
	AccessReviewStageId                string
}

// NewUserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageID returns a new UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId struct
func NewUserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageID(userId string, accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, accessReviewStageId string) UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId {
	return UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId{
		UserId:                             userId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		AccessReviewStageId:                accessReviewStageId,
	}
}

// ParseUserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageID parses 'input' into a UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId
func ParseUserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageID(input string) (*UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageIDInsensitively parses 'input' case-insensitively into a UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId
// note: this method should only be used for API response data and not user input
func ParseUserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageIDInsensitively(input string) (*UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AccessReviewStageId, ok = input.Parsed["accessReviewStageId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewStageId", input)
	}

	return nil
}

// ValidateUserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageID checks that 'input' can be parsed as a User Id Pending Access Review Instance Id Decision Id Instance Stage ID
func ValidateUserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Pending Access Review Instance Id Decision Id Instance Stage ID
func (id UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/decisions/%s/instance/stages/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewStageId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Pending Access Review Instance Id Decision Id Instance Stage ID
func (id UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
		resourceids.StaticSegment("instance", "instance", "instance"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageId"),
	}
}

// String returns a human-readable description of this User Id Pending Access Review Instance Id Decision Id Instance Stage ID
func (id UserIdPendingAccessReviewInstanceIdDecisionIdInstanceStageId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
	}
	return fmt.Sprintf("User Id Pending Access Review Instance Id Decision Id Instance Stage (%s)", strings.Join(components, "\n"))
}
