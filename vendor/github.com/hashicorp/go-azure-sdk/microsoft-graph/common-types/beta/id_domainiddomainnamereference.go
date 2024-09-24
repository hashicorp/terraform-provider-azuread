package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DomainIdDomainNameReferenceId{}

// DomainIdDomainNameReferenceId is a struct representing the Resource ID for a Domain Id Domain Name Reference
type DomainIdDomainNameReferenceId struct {
	DomainId          string
	DirectoryObjectId string
}

// NewDomainIdDomainNameReferenceID returns a new DomainIdDomainNameReferenceId struct
func NewDomainIdDomainNameReferenceID(domainId string, directoryObjectId string) DomainIdDomainNameReferenceId {
	return DomainIdDomainNameReferenceId{
		DomainId:          domainId,
		DirectoryObjectId: directoryObjectId,
	}
}

// ParseDomainIdDomainNameReferenceID parses 'input' into a DomainIdDomainNameReferenceId
func ParseDomainIdDomainNameReferenceID(input string) (*DomainIdDomainNameReferenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DomainIdDomainNameReferenceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DomainIdDomainNameReferenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDomainIdDomainNameReferenceIDInsensitively parses 'input' case-insensitively into a DomainIdDomainNameReferenceId
// note: this method should only be used for API response data and not user input
func ParseDomainIdDomainNameReferenceIDInsensitively(input string) (*DomainIdDomainNameReferenceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DomainIdDomainNameReferenceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DomainIdDomainNameReferenceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DomainIdDomainNameReferenceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DomainId, ok = input.Parsed["domainId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "domainId", input)
	}

	if id.DirectoryObjectId, ok = input.Parsed["directoryObjectId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "directoryObjectId", input)
	}

	return nil
}

// ValidateDomainIdDomainNameReferenceID checks that 'input' can be parsed as a Domain Id Domain Name Reference ID
func ValidateDomainIdDomainNameReferenceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDomainIdDomainNameReferenceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Domain Id Domain Name Reference ID
func (id DomainIdDomainNameReferenceId) ID() string {
	fmtString := "/domains/%s/domainNameReferences/%s"
	return fmt.Sprintf(fmtString, id.DomainId, id.DirectoryObjectId)
}

// Segments returns a slice of Resource ID Segments which comprise this Domain Id Domain Name Reference ID
func (id DomainIdDomainNameReferenceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("domains", "domains", "domains"),
		resourceids.UserSpecifiedSegment("domainId", "domainId"),
		resourceids.StaticSegment("domainNameReferences", "domainNameReferences", "domainNameReferences"),
		resourceids.UserSpecifiedSegment("directoryObjectId", "directoryObjectId"),
	}
}

// String returns a human-readable description of this Domain Id Domain Name Reference ID
func (id DomainIdDomainNameReferenceId) String() string {
	components := []string{
		fmt.Sprintf("Domain: %q", id.DomainId),
		fmt.Sprintf("Directory Object: %q", id.DirectoryObjectId),
	}
	return fmt.Sprintf("Domain Id Domain Name Reference (%s)", strings.Join(components, "\n"))
}
