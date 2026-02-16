package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{}

// MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId is a struct representing the Resource ID for a Me Pending Access Review Instance Id Decision Id Instance Contacted Reviewer
type MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId struct {
	AccessReviewInstanceId             string
	AccessReviewInstanceDecisionItemId string
	AccessReviewReviewerId             string
}

// NewMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID returns a new MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId struct
func NewMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID(accessReviewInstanceId string, accessReviewInstanceDecisionItemId string, accessReviewReviewerId string) MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId {
	return MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{
		AccessReviewInstanceId:             accessReviewInstanceId,
		AccessReviewInstanceDecisionItemId: accessReviewInstanceDecisionItemId,
		AccessReviewReviewerId:             accessReviewReviewerId,
	}
}

// ParseMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID parses 'input' into a MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId
func ParseMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID(input string) (*MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerIDInsensitively parses 'input' case-insensitively into a MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId
// note: this method should only be used for API response data and not user input
func ParseMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerIDInsensitively(input string) (*MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

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

// ValidateMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID checks that 'input' can be parsed as a Me Pending Access Review Instance Id Decision Id Instance Contacted Reviewer ID
func ValidateMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Pending Access Review Instance Id Decision Id Instance Contacted Reviewer ID
func (id MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId) ID() string {
	fmtString := "/me/pendingAccessReviewInstances/%s/decisions/%s/instance/contactedReviewers/%s"
	return fmt.Sprintf(fmtString, id.AccessReviewInstanceId, id.AccessReviewInstanceDecisionItemId, id.AccessReviewReviewerId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Pending Access Review Instance Id Decision Id Instance Contacted Reviewer ID
func (id MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("pendingAccessReviewInstances", "pendingAccessReviewInstances", "pendingAccessReviewInstances"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceId", "accessReviewInstanceId"),
		resourceids.StaticSegment("decisions", "decisions", "decisions"),
		resourceids.UserSpecifiedSegment("accessReviewInstanceDecisionItemId", "accessReviewInstanceDecisionItemId"),
		resourceids.StaticSegment("instance", "instance", "instance"),
		resourceids.StaticSegment("contactedReviewers", "contactedReviewers", "contactedReviewers"),
		resourceids.UserSpecifiedSegment("accessReviewReviewerId", "accessReviewReviewerId"),
	}
}

// String returns a human-readable description of this Me Pending Access Review Instance Id Decision Id Instance Contacted Reviewer ID
func (id MePendingAccessReviewInstanceIdDecisionIdInstanceContactedReviewerId) String() string {
	components := []string{
		fmt.Sprintf("Access Review Instance: %q", id.AccessReviewInstanceId),
		fmt.Sprintf("Access Review Instance Decision Item: %q", id.AccessReviewInstanceDecisionItemId),
		fmt.Sprintf("Access Review Reviewer: %q", id.AccessReviewReviewerId),
	}
	return fmt.Sprintf("Me Pending Access Review Instance Id Decision Id Instance Contacted Reviewer (%s)", strings.Join(components, "\n"))
}
