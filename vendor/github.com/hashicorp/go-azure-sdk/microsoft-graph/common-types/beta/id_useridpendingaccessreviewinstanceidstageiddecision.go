package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPendingAccessReviewInstanceIdStageIdDecisionId{}

// UserIdPendingAccessReviewInstanceIdStageIdDecisionId is a struct representing the Resource ID for a User Id Pending Access Review Instance Id Stage Id Decision
type UserIdPendingAccessReviewInstanceIdStageIdDecisionId struct {
	UserId                             string
	AccessReviewInstanceId             string
	AccessReviewStageId                string
	AccessReviewInstanceDecisionItemId string
}

// NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID returns a new UserIdPendingAccessReviewInstanceIdStageIdDecisionId struct
func NewUserIdPendingAccessReviewInstanceIdStageIdDecisionID(userId string, accessReviewInstanceId string, accessReviewStageId string, accessReviewInstanceDecisionItemId string) UserIdPendingAccessReviewInstanceIdStageIdDecisionId {
	return UserIdPendingAccessReviewInstanceIdStageIdDecisionId{
		UserId:                             userId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewStageId:                accessReviewStageId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
	}
}

// ParseUserIdPendingAccessReviewInstanceIdStageIdDecisionID parses 'input' into a UserIdPendingAccessReviewInstanceIdStageIdDecisionId
func ParseUserIdPendingAccessReviewInstanceIdStageIdDecisionID(input string) (*UserIdPendingAccessReviewInstanceIdStageIdDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceIdStageIdDecisionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceIdStageIdDecisionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPendingAccessReviewInstanceIdStageIdDecisionIDInsensitively parses 'input' case-insensitively into a UserIdPendingAccessReviewInstanceIdStageIdDecisionId
// note: this method should only be used for API response data and not user input
func ParseUserIdPendingAccessReviewInstanceIdStageIdDecisionIDInsensitively(input string) (*UserIdPendingAccessReviewInstanceIdStageIdDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceIdStageIdDecisionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceIdStageIdDecisionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPendingAccessReviewInstanceIdStageIdDecisionId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AccessReviewInstanceDecisionItemId, ok = input.Parsed["accessReviewInstanceDecisionItemId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewInstanceDecisionItemId", input)
	}

	return nil
}

// ValidateUserIdPendingAccessReviewInstanceIdStageIdDecisionID checks that 'input' can be parsed as a User Id Pending Access Review Instance Id Stage Id Decision ID
func ValidateUserIdPendingAccessReviewInstanceIdStageIdDecisionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPendingAccessReviewInstanceIdStageIdDecisionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Pending Access Review Instance Id Stage Id Decision ID
func (id UserIdPendingAccessReviewInstanceIdStageIdDecisionId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/stages/%s/decisions/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Pending Access Review Instance Id Stage Id Decision ID
func (id UserIdPendingAccessReviewInstanceIdStageIdDecisionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
	}
}

// String returns a human-readable description of this User Id Pending Access Review Instance Id Stage Id Decision ID
func (id UserIdPendingAccessReviewInstanceIdStageIdDecisionId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
	}
	return fmt.Sprintf("User Id Pending Access Review Instance Id Stage Id Decision (%s)", strings.Join(components, "\n"))
}
