package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectorySubscriptionId{}

// DirectorySubscriptionId is a struct representing the Resource ID for a Directory Subscription
type DirectorySubscriptionId struct {
	CompanySubscriptionId string
}

// NewDirectorySubscriptionID returns a new DirectorySubscriptionId struct
func NewDirectorySubscriptionID(companySubscriptionId string) DirectorySubscriptionId {
	return DirectorySubscriptionId{
		CompanySubscriptionId: companySubscriptionId,
	}
}

// ParseDirectorySubscriptionID parses 'input' into a DirectorySubscriptionId
func ParseDirectorySubscriptionID(input string) (*DirectorySubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectorySubscriptionId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectorySubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectorySubscriptionIDInsensitively parses 'input' case-insensitively into a DirectorySubscriptionId
// note: this method should only be used for API response data and not user input
func ParseDirectorySubscriptionIDInsensitively(input string) (*DirectorySubscriptionId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectorySubscriptionId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectorySubscriptionId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectorySubscriptionId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CompanySubscriptionId, ok = input.Parsed["companySubscriptionId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "companySubscriptionId", input)
	}

	return nil
}

// ValidateDirectorySubscriptionID checks that 'input' can be parsed as a Directory Subscription ID
func ValidateDirectorySubscriptionID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectorySubscriptionID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Subscription ID
func (id DirectorySubscriptionId) ID() string {
	fmtString := "/directory/subscriptions/%s"
	return fmt.Sprintf(fmtString, id.CompanySubscriptionId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Subscription ID
func (id DirectorySubscriptionId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("subscriptions", "subscriptions", "subscriptions"),
		resourceids.UserSpecifiedSegment("companySubscriptionId", "companySubscriptionId"),
	}
}

// String returns a human-readable description of this Directory Subscription ID
func (id DirectorySubscriptionId) String() string {
	components := []string{
		fmt.Sprintf("Company Subscription: %q", id.CompanySubscriptionId),
	}
	return fmt.Sprintf("Directory Subscription (%s)", strings.Join(components, "\n"))
}
