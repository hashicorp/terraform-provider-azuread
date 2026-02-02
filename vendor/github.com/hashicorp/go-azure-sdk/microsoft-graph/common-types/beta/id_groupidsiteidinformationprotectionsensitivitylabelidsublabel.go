package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId{}

// GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId is a struct representing the Resource ID for a Group Id Site Id Information Protection Sensitivity Label Id Sublabel
type GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId struct {
	GroupId             string
	SiteId              string
	SensitivityLabelId  string
	SensitivityLabelId1 string
}

// NewGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID returns a new GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId struct
func NewGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID(groupId string, siteId string, sensitivityLabelId string, sensitivityLabelId1 string) GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId {
	return GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId{
		GroupId:             groupId,
		SiteId:              siteId,
		SensitivityLabelId:  sensitivityLabelId,
		SensitivityLabelId1: sensitivityLabelId1,
	}
}

// ParseGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID parses 'input' into a GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId
func ParseGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID(input string) (*GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelIDInsensitively(input string) (*GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId) FromParseResult(input resourceids.ParseResult) error {
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

	if id.SensitivityLabelId1, ok = input.Parsed["sensitivityLabelId1"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "sensitivityLabelId1", input)
	}

	return nil
}

// ValidateGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID checks that 'input' can be parsed as a Group Id Site Id Information Protection Sensitivity Label Id Sublabel ID
func ValidateGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Information Protection Sensitivity Label Id Sublabel ID
func (id GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId) ID() string {
	fmtString := "/groups/%s/sites/%s/informationProtection/sensitivityLabels/%s/sublabels/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.SensitivityLabelId, id.SensitivityLabelId1)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Information Protection Sensitivity Label Id Sublabel ID
func (id GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("sensitivityLabels", "sensitivityLabels", "sensitivityLabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId", "sensitivityLabelId"),
		resourceids.StaticSegment("sublabels", "sublabels", "sublabels"),
		resourceids.UserSpecifiedSegment("sensitivityLabelId1", "sensitivityLabelId1"),
	}
}

// String returns a human-readable description of this Group Id Site Id Information Protection Sensitivity Label Id Sublabel ID
func (id GroupIdSiteIdInformationProtectionSensitivityLabelIdSublabelId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Sensitivity Label: %q", id.SensitivityLabelId),
		fmt.Sprintf("Sensitivity Label Id 1: %q", id.SensitivityLabelId1),
	}
	return fmt.Sprintf("Group Id Site Id Information Protection Sensitivity Label Id Sublabel (%s)", strings.Join(components, "\n"))
}
