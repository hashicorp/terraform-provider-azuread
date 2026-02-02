package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{}

// UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId is a struct representing the Resource ID for a User Id Pending Access Review Instance Id Decision Id Instance Contacted Reviewer
type UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId struct {
	UserId                             string
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
	AccessReviewReviewerId             string
}

// NewUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID returns a new UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId struct
func NewUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID(userId string, accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, accessReviewReviewerId string) UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId {
	return UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{
		UserId:                             userId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		AccessReviewReviewerId:             accessReviewReviewerId,
	}
}

// ParseUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID parses 'input' into a UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId
func ParseUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID(input string) (*UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerIDInsensitively parses 'input' case-insensitively into a UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId
// note: this method should only be used for API response data and not user input
func ParseUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerIDInsensitively(input string) (*UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AccessReviewReviewerId, ok = input.Parsed["accessReviewReviewerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewReviewerId", input)
	}

	return nil
}

// ValidateUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID checks that 'input' can be parsed as a User Id Pending Access Review Instance Id Decision Id Instance Contacted Reviewer ID
func ValidateUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Pending Access Review Instance Id Decision Id Instance Contacted Reviewer ID
func (id UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/decisions/%s/instance/contactedReviewers/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewReviewerId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Pending Access Review Instance Id Decision Id Instance Contacted Reviewer ID
func (id UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
		resourceids.StaticSegment("instance", "instance", "instance"),
		resourceids.StaticSegment("contactedReviewers", "contactedReviewers", "contactedReviewers"),
		resourceids.UserSpecifiedSegment("accessReviewReviewerId", "accessReviewReviewerId"),
	}
}

// String returns a human-readable description of this User Id Pending Access Review Instance Id Decision Id Instance Contacted Reviewer ID
func (id UserIdPendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Reviewer: %q", id.AccessReviewReviewerId),
	}
	return fmt.Sprintf("User Id Pending Access Review Instance Id Decision Id Instance Contacted Reviewer (%s)", strings.Join(components, "\n"))
}
