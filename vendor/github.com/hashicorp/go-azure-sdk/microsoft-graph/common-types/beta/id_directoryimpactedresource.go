package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryImpactedResourceId{}

// DirectoryImpactedResourceId is a struct representing the Resource ID for a Directory Impacted Resource
type DirectoryImpactedResourceId struct {
	ImpactedResourceId string
}

// NewDirectoryImpactedResourceID returns a new DirectoryImpactedResourceId struct
func NewDirectoryImpactedResourceID(impactedResourceId string) DirectoryImpactedResourceId {
	return DirectoryImpactedResourceId{
		ImpactedResourceId: impactedResourceId,
	}
}

// ParseDirectoryImpactedResourceID parses 'input' into a DirectoryImpactedResourceId
func ParseDirectoryImpactedResourceID(input string) (*DirectoryImpactedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryImpactedResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryImpactedResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryImpactedResourceIDInsensitively parses 'input' case-insensitively into a DirectoryImpactedResourceId
// note: this method should only be used for API response data and not user input
func ParseDirectoryImpactedResourceIDInsensitively(input string) (*DirectoryImpactedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryImpactedResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryImpactedResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryImpactedResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.ImpactedResourceId, ok = input.Parsed["impactedResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "impactedResourceId", input)
	}

	return nil
}

// ValidateDirectoryImpactedResourceID checks that 'input' can be parsed as a Directory Impacted Resource ID
func ValidateDirectoryImpactedResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryImpactedResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Impacted Resource ID
func (id DirectoryImpactedResourceId) ID() string {
	fmtString := "/directory/impactedResources/%s"
	return fmt.Sprintf(fmtString, id.ImpactedResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Impacted Resource ID
func (id DirectoryImpactedResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("impactedResources", "impactedResources", "impactedResources"),
		resourceids.UserSpecifiedSegment("impactedResourceId", "impactedResourceId"),
	}
}

// String returns a human-readable description of this Directory Impacted Resource ID
func (id DirectoryImpactedResourceId) String() string {
	components := []string{
		fmt.Sprintf("Impacted Resource: %q", id.ImpactedResourceId),
	}
	return fmt.Sprintf("Directory Impacted Resource (%s)", strings.Join(components, "\n"))
}
