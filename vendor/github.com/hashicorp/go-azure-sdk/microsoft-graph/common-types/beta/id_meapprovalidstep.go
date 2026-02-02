package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeApprovalIdStepId{}

// MeApprovalIdStepId is a struct representing the Resource ID for a Me Approval Id Step
type MeApprovalIdStepId struct {
	ApprovalId     string
	ApprovalStepId string
}

// NewMeApprovalIdStepID returns a new MeApprovalIdStepId struct
func NewMeApprovalIdStepID(approvalId string, approvalStepId string) MeApprovalIdStepId {
	return MeApprovalIdStepId{
		ApprovalId:     approvalId,
		ApprovalStepId: approvalStepId,
	}
}

// ParseMeApprovalIdStepID parses 'input' into a MeApprovalIdStepId
func ParseMeApprovalIdStepID(input string) (*MeApprovalIdStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeApprovalIdStepId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeApprovalIdStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeApprovalIdStepIDInsensitively parses 'input' case-insensitively into a MeApprovalIdStepId
// note: this method should only be used for API response data and not user input
func ParseMeApprovalIdStepIDInsensitively(input string) (*MeApprovalIdStepId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeApprovalIdStepId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeApprovalIdStepId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeApprovalIdStepId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ApprovalId, ok = input.Parsed["approvalId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalId", input)
	}

	if id.ApprovalStepId, ok = input.Parsed["approvalStepId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "approvalStepId", input)
	}

	return nil
}

// ValidateMeApprovalIdStepID checks that 'input' can be parsed as a Me Approval Id Step ID
func ValidateMeApprovalIdStepID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeApprovalIdStepID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Approval Id Step ID
func (id MeApprovalIdStepId) ID() string {
	fmtString := "/me/approvals/%s/steps/%s"
	return fmt.Sprintf(fmtString, id.ApprovalId, id.ApprovalStepId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Approval Id Step ID
func (id MeApprovalIdStepId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("approvals", "approvals", "approvals"),
		resourceids.UserSpecifiedSegment("approvalId", "approvalId"),
		resourceids.StaticSegment("steps", "steps", "steps"),
		resourceids.UserSpecifiedSegment("approvalStepId", "approvalStepId"),
	}
}

// String returns a human-readable description of this Me Approval Id Step ID
func (id MeApprovalIdStepId) String() string {
	components := []string{
		fmt.Sprintf("Approval: %q", id.ApprovalId),
		fmt.Sprintf("Approval Step: %q", id.ApprovalStepId),
	}
	return fmt.Sprintf("Me Approval Id Step (%s)", strings.Join(components, "\n"))
}
