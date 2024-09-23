package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePendingAccessReviewInstanceIdStageIdDecisionId{}

// MePendingAccessReviewInstanceIdStageIdDecisionId is a struct representing the Resource ID for a Me Pending Access Review Instance Id Stage Id Decision
type MePendingAccessReviewInstanceIdStageIdDecisionId struct {
	AccessReviewInstanceId             string
	AccessReviewStageId                string
	AccessReviewInstanceDecisionItemId string
}

// NewMePendingAccessReviewInstanceIdStageIdDecisionID returns a new MePendingAccessReviewInstanceIdStageIdDecisionId struct
func NewMePendingAccessReviewInstanceIdStageIdDecisionID(accessReviewInstanceId string, accessReviewStageId string, accessReviewInstanceDecisionItemId string) MePendingAccessReviewInstanceIdStageIdDecisionId {
	return MePendingAccessReviewInstanceIdStageIdDecisionId{
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewStageId:                accessReviewStageId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
	}
}

// ParseMePendingAccessReviewInstanceIdStageIdDecisionID parses 'input' into a MePendingAccessReviewInstanceIdStageIdDecisionId
func ParseMePendingAccessReviewInstanceIdStageIdDecisionID(input string) (*MePendingAccessReviewInstanceIdStageIdDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdStageIdDecisionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdStageIdDecisionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceIdStageIdDecisionIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceIdStageIdDecisionId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceIdStageIdDecisionIDInsensitively(input string) (*MePendingAccessReviewInstanceIdStageIdDecisionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdStageIdDecisionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdStageIdDecisionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePendingAccessReviewInstanceIdStageIdDecisionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMePendingAccessReviewInstanceIdStageIdDecisionID checks that 'input' can be parsed as a Me Pending Access Review Instance Id Stage Id Decision ID
func ValidateMePendingAccessReviewInstanceIdStageIdDecisionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceIdStageIdDecisionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Id Stage Id Decision ID
func (id MePendingAccessReviewInstanceIdStageIdDecisionId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/stages/%s/decisions/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Id Stage Id Decision ID
func (id MePendingAccessReviewInstanceIdStageIdDecisionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Id Stage Id Decision ID
func (id MePendingAccessReviewInstanceIdStageIdDecisionId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Id Stage Id Decision (%s)", strings.Join(components, "\n"))
}
