package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId{}

// GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId is a struct representing the Resource ID for a Group Id Site Id Information Protection Threat Assessment Request Id Result
type GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId struct {
	GroupId                   string
	SiteId                    string
	ThreatAssessmentRequestId string
	ThreatAssessmentResultId  string
}

// NewGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultID returns a new GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId struct
func NewGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultID(groupId string, siteId string, threatAssessmentRequestId string, threatAssessmentResultId string) GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId {
	return GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId{
		GroupId:                   groupId,
		SiteId:                    siteId,
		ThreatAssessmentRequestId: threatAssessmentRequestId,
		ThreatAssessmentResultId:  threatAssessmentResultId,
	}
}

// ParseGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultID parses 'input' into a GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId
func ParseGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultID(input string) (*GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultIDInsensitively(input string) (*GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.GroupId, ok = input.Parsed["groupId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "groupId", input)
	}

	if id.SiteId, ok = input.Parsed["siteId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "siteId", input)
	}

	if id.ThreatAssessmentRequestId, ok = input.Parsed["threatAssessmentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentRequestId", input)
	}

	if id.ThreatAssessmentResultId, ok = input.Parsed["threatAssessmentResultId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentResultId", input)
	}

	return nil
}

// ValidateGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultID checks that 'input' can be parsed as a Group Id Site Id Information Protection Threat Assessment Request Id Result ID
func ValidateGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Information Protection Threat Assessment Request Id Result ID
func (id GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId) ID() string {
	fmtString := "/groups/%s/sites/%s/informationProtection/threatAssessmentRequests/%s/results/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ThreatAssessmentRequestId, id.ThreatAssessmentResultId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Information Protection Threat Assessment Request Id Result ID
func (id GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("threatAssessmentRequests", "threatAssessmentRequests", "threatAssessmentRequests"),
		resourceids.UserSpecifiedSegment("threatAssessmentRequestId", "threatAssessmentRequestId"),
		resourceids.StaticSegment("results", "results", "results"),
		resourceids.UserSpecifiedSegment("threatAssessmentResultId", "threatAssessmentResultId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Information Protection Threat Assessment Request Id Result ID
func (id GroupIdSiteIdInformationProtectionThreatAssessmentRequestIdResultId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Threat Assessment Request: %q", id.ThreatAssessmentRequestId),
		fmt.Sprintf("Threat Assessment Result: %q", id.ThreatAssessmentResultId),
	}
	return fmt.Sprintf("Group Id Site Id Information Protection Threat Assessment Request Id Result (%s)", strings.Join(components, "\n"))
}
