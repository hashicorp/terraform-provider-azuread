package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeDataSecurityAndGovernanceActivityContentActivityId{}

// MeDataSecurityAndGovernanceActivityContentActivityId is a struct representing the Resource ID for a Me Data Security And Governance Activity Content Activity
type MeDataSecurityAndGovernanceActivityContentActivityId struct {
	ContentActivityId string
}

// NewMeDataSecurityAndGovernanceActivityContentActivityID returns a new MeDataSecurityAndGovernanceActivityContentActivityId struct
func NewMeDataSecurityAndGovernanceActivityContentActivityID(contentActivityId string) MeDataSecurityAndGovernanceActivityContentActivityId {
	return MeDataSecurityAndGovernanceActivityContentActivityId{
		ContentActivityId: contentActivityId,
	}
}

// ParseMeDataSecurityAndGovernanceActivityContentActivityID parses 'input' into a MeDataSecurityAndGovernanceActivityContentActivityId
func ParseMeDataSecurityAndGovernanceActivityContentActivityID(input string) (*MeDataSecurityAndGovernanceActivityContentActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDataSecurityAndGovernanceActivityContentActivityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDataSecurityAndGovernanceActivityContentActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeDataSecurityAndGovernanceActivityContentActivityIDInsensitively parses 'input' case-insensitively into a MeDataSecurityAndGovernanceActivityContentActivityId
// note: this method should only be used for API response data and not user input
func ParseMeDataSecurityAndGovernanceActivityContentActivityIDInsensitively(input string) (*MeDataSecurityAndGovernanceActivityContentActivityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeDataSecurityAndGovernanceActivityContentActivityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeDataSecurityAndGovernanceActivityContentActivityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeDataSecurityAndGovernanceActivityContentActivityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ContentActivityId, ok = input.Parsed["contentActivityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "contentActivityId", input)
	}

	return nil
}

// ValidateMeDataSecurityAndGovernanceActivityContentActivityID checks that 'input' can be parsed as a Me Data Security And Governance Activity Content Activity ID
func ValidateMeDataSecurityAndGovernanceActivityContentActivityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeDataSecurityAndGovernanceActivityContentActivityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Data Security And Governance Activity Content Activity ID
func (id MeDataSecurityAndGovernanceActivityContentActivityId) ID() string {
	fmtString := "/me/dataSecurityAndGovernance/activities/contentActivities/%s"
	return fmt.Sprintf(fmtString, id.ContentActivityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Data Security And Governance Activity Content Activity ID
func (id MeDataSecurityAndGovernanceActivityContentActivityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("dataSecurityAndGovernance", "dataSecurityAndGovernance", "dataSecurityAndGovernance"),
		resourceids.StaticSegment("activities", "activities", "activities"),
		resourceids.StaticSegment("contentActivities", "contentActivities", "contentActivities"),
		resourceids.UserSpecifiedSegment("contentActivityId", "contentActivityId"),
	}
}

// String returns a human-readable description of this Me Data Security And Governance Activity Content Activity ID
func (id MeDataSecurityAndGovernanceActivityContentActivityId) String() string {
	components := []string{
		fmt.Sprintf("Content Activity: %q", id.ContentActivityId),
	}
	return fmt.Sprintf("Me Data Security And Governance Activity Content Activity (%s)", strings.Join(components, "\n"))
}
