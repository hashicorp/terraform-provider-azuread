package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId{}

// UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId is a struct representing the Resource ID for a User Id Pending Access Review Instance Id Stage Id Decision Id Insight
type UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId struct {
	UserId                             string
	AccessReviewInstanceId             string
	AccessReviewStageId                string
	AccessReviewInstanceDecisionItemId string
	GovernanceInsightId                string
}

// NewUserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightID returns a new UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId struct
func NewUserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightID(userId string, accessReviewInstanceId string, accessReviewStageId string, accessReviewInstanceDecisionItemId string, governanceInsightId string) UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId {
	return UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId{
		UserId:                             userId,
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewStageId:                accessReviewStageId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		GovernanceInsightId:                governanceInsightId,
	}
}

// ParseUserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightID parses 'input' into a UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId
func ParseUserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightID(input string) (*UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightIDInsensitively parses 'input' case-insensitively into a UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId
// note: this method should only be used for API response data and not user input
func ParseUserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightIDInsensitively(input string) (*UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.GovernanceInsightId, ok = input.Parsed["governanceInsightId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "governanceInsightId", input)
	}

	return nil
}

// ValidateUserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightID checks that 'input' can be parsed as a User Id Pending Access Review Instance Id Stage Id Decision Id Insight ID
func ValidateUserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Pending Access Review Instance Id Stage Id Decision Id Insight ID
func (id UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId) ID() string {
	fmtString := "/users/%s/pendingAccessReviewInstances/%s/stages/%s/decisions/%s/insights/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.AccessReviewInstanceId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId, id.GovernanceInsightId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Pending Access Review Instance Id Stage Id Decision Id Insight ID
func (id UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
		resourceids.StaticSegment("insights", "insights", "insights"),
		resourceids.UserSpecifiedSegment("governanceInsightId", "governanceInsightId"),
	}
}

// String returns a human-readable description of this User Id Pending Access Review Instance Id Stage Id Decision Id Insight ID
func (id UserIdPendingAccessReviewInstanceIdStageIdDecisionIdInsightId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Governance Insight: %q", id.GovernanceInsightId),
	}
	return fmt.Sprintf("User Id Pending Access Review Instance Id Stage Id Decision Id Insight (%s)", strings.Join(components, "\n"))
}
