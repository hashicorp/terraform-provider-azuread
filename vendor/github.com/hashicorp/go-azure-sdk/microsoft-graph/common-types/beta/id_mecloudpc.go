package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeCloudPCId{}

// MeCloudPCId is a struct representing the Resource ID for a Me Cloud PC
type MeCloudPCId struct {
	CloudPCId string
}

// NewMeCloudPCID returns a new MeCloudPCId struct
func NewMeCloudPCID(cloudPCId string) MeCloudPCId {
	return MeCloudPCId{
		CloudPCId: cloudPCId,
	}
}

// ParseMeCloudPCID parses 'input' into a MeCloudPCId
func ParseMeCloudPCID(input string) (*MeCloudPCId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCloudPCId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCloudPCId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeCloudPCIDInsensitively parses 'input' case-insensitively into a MeCloudPCId
// note: this method should only be used for API response data and not user input
func ParseMeCloudPCIDInsensitively(input string) (*MeCloudPCId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeCloudPCId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeCloudPCId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeCloudPCId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CloudPCId, ok = input.Parsed["cloudPCId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "cloudPCId", input)
	}

	return nil
}

// ValidateMeCloudPCID checks that 'input' can be parsed as a Me Cloud PC ID
func ValidateMeCloudPCID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeCloudPCID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Cloud PC ID
func (id MeCloudPCId) ID() string {
	fmtString := "/me/cloudPCs/%s"
	return fmt.Sprintf(fmtString, id.CloudPCId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Cloud PC ID
func (id MeCloudPCId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("cloudPCs", "cloudPCs", "cloudPCs"),
		resourceids.UserSpecifiedSegment("cloudPCId", "cloudPCId"),
	}
}

// String returns a human-readable description of this Me Cloud PC ID
func (id MeCloudPCId) String() string {
	components := []string{
		fmt.Sprintf("Cloud PC: %q", id.CloudPCId),
	}
	return fmt.Sprintf("Me Cloud PC (%s)", strings.Join(components, "\n"))
}
