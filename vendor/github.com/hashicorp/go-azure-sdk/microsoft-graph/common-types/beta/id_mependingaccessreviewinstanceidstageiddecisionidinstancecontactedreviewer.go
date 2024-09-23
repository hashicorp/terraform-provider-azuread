package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId{}

// MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId is a struct representing the Resource ID for a Me Pending Access Review Instance Id Stage Id Decision Id Instance Contacted Reviewer
type MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId struct {
	AccessReviewInstanceId             string
	AccessReviewStageId                string
	AccessReviewInstanceDecisionItemId string
	AccessReviewReviewerId             string
}

// NewMePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerID returns a new MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId struct
func NewMePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerID(accessReviewInstanceId string, accessReviewStageId string, accessReviewInstanceDecisionItemId string, accessReviewReviewerId string) MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId {
	return MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId{
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewStageId:                accessReviewStageId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		AccessReviewReviewerId:             accessReviewReviewerId,
	}
}

// ParseMePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerID parses 'input' into a MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId
func ParseMePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerID(input string) (*MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerIDInsensitively(input string) (*MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.AccessReviewReviewerId, ok = input.Parsed["accessReviewReviewerId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "accessReviewReviewerId", input)
	}

	return nil
}

// ValidateMePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerID checks that 'input' can be parsed as a Me Pending Access Review Instance Id Stage Id Decision Id Instance Contacted Reviewer ID
func ValidateMePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Id Stage Id Decision Id Instance Contacted Reviewer ID
func (id MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/stages/%s/decisions/%s/instance/contactedReviewers/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewStageId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewReviewerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Id Stage Id Decision Id Instance Contacted Reviewer ID
func (id MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("stages", "stages", "stages"),
		resourceids.UserSpecifiedSegment("accessReviewStageId", "accessReviewStageId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
		resourceids.StaticSegment("instance", "instance", "instance"),
		resourceids.StaticSegment("contactedReviewers", "contactedReviewers", "contactedReviewers"),
		resourceids.UserSpecifiedSegment("accessReviewReviewerId", "accessReviewReviewerId"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Id Stage Id Decision Id Instance Contacted Reviewer ID
func (id MePendingAccessReviewInstanceIdStageIdDecisionIdInstanceContactedReviewerId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Stage: %q", id.AccessReviewStageId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Reviewer: %q", id.AccessReviewReviewerId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Id Stage Id Decision Id Instance Contacted Reviewer (%s)", strings.Join(components, "\n"))
}
