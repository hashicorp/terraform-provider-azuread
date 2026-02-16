package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPendingAccessReviewInstanceIdDecisionIdInsightId{}

// UserIdPendingAccessReviewInstanceIdDecisionIdInsightId is a struct representing the Resource ID for a User Id Pending Access Review Instance Id Decision Id Insight
type UserIdPendingAccessReviewInstanceIdDecisionIdInsightId struct {
	UserId                             string
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
	GovernanceInsightId                string
}

// NewUserIdPendingAccessReviewInstanceIdDecisionIdInsightID returns a new UserIdPendingAccessReviewInstanceIdDecisionIdInsightId struct
func NewUserIdPendingAccessReviewInstanceIdDecisionIdInsightID(userId string, accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, governanceInsightId string) UserIdPendingAccessReviewInstanceIdDecisionIdInsightId {
	return UserIdPendingAccessReviewInstanceIdDecisionIdInsightId{
		UserId:                             userId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		GovernanceInsightId:                governanceInsightId,
	}
}

// ParseUserIdPendingAccessReviewInstanceIdDecisionIdInsightID parses 'input' into a UserIdPendingAccessReviewInstanceIdDecisionIdInsightId
func ParseUserIdPendingAccessReviewInstanceIdDecisionIdInsightID(input string) (*UserIdPendingAccessReviewInstanceIdDecisionIdInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceIdDecisionIdInsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceIdDecisionIdInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPendingAccessReviewInstanceIdDecisionIdInsightIDInsensitively parses 'input' case-insensitively into a UserIdPendingAccessReviewInstanceIdDecisionIdInsightId
// note: this method should only be used for API response data and not user input
func ParseUserIdPendingAccessReviewInstanceIdDecisionIdInsightIDInsensitively(input string) (*UserIdPendingAccessReviewInstanceIdDecisionIdInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceIdDecisionIdInsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceIdDecisionIdInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPendingAccessReviewInstanceIdDecisionIdInsightId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.GovernanceInsightId, ok = input.Parsed["governanceInsightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", input)
	}

	return nil
}

// ValidateUserIdPendingAccessReviewInstanceIdDecisionIdInsightID checks that 'input' can be parsed as a User Id Pending Access Review Instance Id Decision Id Insight ID
func ValidateUserIdPendingAccessReviewInstanceIdDecisionIdInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPendingAccessReviewInstanceIdDecisionIdInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Pending Access Review Instance Id Decision Id Insight ID
func (id UserIdPendingAccessReviewInstanceIdDecisionIdInsightId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/decisions/%s/insights/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.GovernanceInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Pending Access Review Instance Id Decision Id Insight ID
func (id UserIdPendingAccessReviewInstanceIdDecisionIdInsightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
		resourceids.StaticSegment("insights", "insights", "insights"),
		resourceids.UserSpecifiedSegment("governanceInsightId", "governanceInsightId"),
	}
}

// String returns a human-readable description of this User Id Pending Access Review Instance Id Decision Id Insight ID
func (id UserIdPendingAccessReviewInstanceIdDecisionIdInsightId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Governance Insight: %q", id.GovernanceInsightId),
	}
	return fmt.Sprintf("User Id Pending Access Review Instance Id Decision Id Insight (%s)", strings.Join(components, "\n"))
}
