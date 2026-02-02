package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdInformationProtectionSensitivityLabelId{}

// GroupIdSiteIdInformationProtectionSensitivityLabelId is a struct representing the Resource ID for a Group Id Site Id Information Protection Sensitivity Label
type GroupIdSiteIdInformationProtectionSensitivityLabelId struct {
	GroupId            string
	SiteId             string
	SensitivityLabelId string
}

// NewGroupIdSiteIdInformationProtectionSensitivityLabelID returns a new GroupIdSiteIdInformationProtectionSensitivityLabelId struct
func NewGroupIdSiteIdInformationProtectionSensitivityLabelID(groupId string, siteId string, sensitivityLabelId string) GroupIdSiteIdInformationProtectionSensitivityLabelId {
	return GroupIdSiteIdInformationProtectionSensitivityLabelId{
		GroupId:            groupId,
		SiteId:             siteId,
		SensitivityLabelId: sensitivityLabelId,
	}
}

// ParseGroupIdSiteIdInformationProtectionSensitivityLabelID parses 'input' into a GroupIdSiteIdInformationProtectionSensitivityLabelId
func ParseGroupIdSiteIdInformationProtectionSensitivityLabelID(input string) (*GroupIdSiteIdInformationProtectionSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdInformationProtectionSensitivityLabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdInformationProtectionSensitivityLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdInformationProtectionSensitivityLabelIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdInformationProtectionSensitivityLabelId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdInformationProtectionSensitivityLabelIDInsensitively(input string) (*GroupIdSiteIdInformationProtectionSensitivityLabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdInformationProtectionSensitivityLabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdInformationProtectionSensitivityLabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdInformationProtectionSensitivityLabelId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.SensitivityLabelId, ok = input.Parsed["sensitivityLabelId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdInformationProtectionSensitivityLabelID checks that 'input' can be parsed as a Group Id Site Id Information Protection Sensitivity Label ID
func ValidateGroupIdSiteIdInformationProtectionSensitivityLabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdInformationProtectionSensitivityLabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Information Protection Sensitivity Label ID
func (id GroupIdSiteIdInformationProtectionSensitivityLabelId) ID() string {
	fmtString := "/groups/%s/sites/%s/informationProtection/sensitivityLabels/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SensitivityLabelId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Information Protection Sensitivity Label ID
func (id GroupIdSiteIdInformationProtectionSensitivityLabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("sensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Information Protection Sensitivity Label ID
func (id GroupIdSiteIdInformationProtectionSensitivityLabelId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
	}
	return fmt.Sprintf("Group Id Site Id Information Protection Sensitivity Label (%s)", strings.Join(components, "\n"))
}
