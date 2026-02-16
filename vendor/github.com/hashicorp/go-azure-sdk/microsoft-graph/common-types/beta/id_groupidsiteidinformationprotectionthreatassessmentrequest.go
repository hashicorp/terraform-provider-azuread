package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &GroupIdSiteIdInformationProtectionThreatAssessmentRequestId{}

// GroupIdSiteIdInformationProtectionThreatAssessmentRequestId is a struct representing the Resource ID for a Group Id Site Id Information Protection Threat Assessment Request
type GroupIdSiteIdInformationProtectionThreatAssessmentRequestId struct {
	GroupId                   string
	SiteId                    string
	ThreatAssessmentRequestId string
}

// NewGroupIdSiteIdInformationProtectionThreatAssessmentRequestID returns a new GroupIdSiteIdInformationProtectionThreatAssessmentRequestId struct
func NewGroupIdSiteIdInformationProtectionThreatAssessmentRequestID(groupId string, siteId string, threatAssessmentRequestId string) GroupIdSiteIdInformationProtectionThreatAssessmentRequestId {
	return GroupIdSiteIdInformationProtectionThreatAssessmentRequestId{
		GroupId:                   groupId,
		SiteId:                    siteId,
		ThreatAssessmentRequestId: threatAssessmentRequestId,
	}
}

// ParseGroupIdSiteIdInformationProtectionThreatAssessmentRequestID parses 'input' into a GroupIdSiteIdInformationProtectionThreatAssessmentRequestId
func ParseGroupIdSiteIdInformationProtectionThreatAssessmentRequestID(input string) (*GroupIdSiteIdInformationProtectionThreatAssessmentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdInformationProtectionThreatAssessmentRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdInformationProtectionThreatAssessmentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseGroupIdSiteIdInformationProtectionThreatAssessmentRequestIDInsensitively parses 'input' case-insensitively into a GroupIdSiteIdInformationProtectionThreatAssessmentRequestId
// note: this method should only be used for API response data and not user input
func ParseGroupIdSiteIdInformationProtectionThreatAssessmentRequestIDInsensitively(input string) (*GroupIdSiteIdInformationProtectionThreatAssessmentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&GroupIdSiteIdInformationProtectionThreatAssessmentRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := GroupIdSiteIdInformationProtectionThreatAssessmentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *GroupIdSiteIdInformationProtectionThreatAssessmentRequestId) FromParseResult(input resourceids.ParseResult) error {
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

	return nil
}

// ValidateGroupIdSiteIdInformationProtectionThreatAssessmentRequestID checks that 'input' can be parsed as a Group Id Site Id Information Protection Threat Assessment Request ID
func ValidateGroupIdSiteIdInformationProtectionThreatAssessmentRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseGroupIdSiteIdInformationProtectionThreatAssessmentRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Group Id Site Id Information Protection Threat Assessment Request ID
func (id GroupIdSiteIdInformationProtectionThreatAssessmentRequestId) ID() string {
	fmtString := "/groups/%s/sites/%s/informationProtection/threatAssessmentRequests/%s"
	return fmt.Sprintf(fmtString, id.GroupId, id.SiteId, id.ThreatAssessmentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this Group Id Site Id Information Protection Threat Assessment Request ID
func (id GroupIdSiteIdInformationProtectionThreatAssessmentRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("groups", "groups", "groups"),
		resourceids.UserSpecifiedSegment("groupId", "groupId"),
		resourceids.StaticSegment("sites", "sites", "sites"),
		resourceids.UserSpecifiedSegment("siteId", "siteId"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("threatAssessmentRequests", "threatAssessmentRequests", "threatAssessmentRequests"),
		resourceids.UserSpecifiedSegment("threatAssessmentRequestId", "threatAssessmentRequestId"),
	}
}

// String returns a human-readable description of this Group Id Site Id Information Protection Threat Assessment Request ID
func (id GroupIdSiteIdInformationProtectionThreatAssessmentRequestId) String() string {
	components := []string{
		fmt.Sprintf("Group: %q", id.GroupId),
		fmt.Sprintf("Site: %q", id.SiteId),
		fmt.Sprintf("Threat Assessment Request: %q", id.ThreatAssessmentRequestId),
	}
	return fmt.Sprintf("Group Id Site Id Information Protection Threat Assessment Request (%s)", strings.Join(components, "\n"))
}
