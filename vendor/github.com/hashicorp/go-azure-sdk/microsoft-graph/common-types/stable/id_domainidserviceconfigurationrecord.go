package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DomainIdServiceConfigurationRecordId{}

// DomainIdServiceConfigurationRecordId is a struct representing the Resource ID for a Domain Id Service Configuration Record
type DomainIdServiceConfigurationRecordId struct {
	DomainId          string
	DomainDnsRecordId string
}

// NewDomainIdServiceConfigurationRecordID returns a new DomainIdServiceConfigurationRecordId struct
func NewDomainIdServiceConfigurationRecordID(domainId string, domainDnsRecordId string) DomainIdServiceConfigurationRecordId {
	return DomainIdServiceConfigurationRecordId{
		DomainId:          domainId,
		DomainDnsRecordId: domainDnsRecordId,
	}
}

// ParseDomainIdServiceConfigurationRecordID parses 'input' into a DomainIdServiceConfigurationRecordId
func ParseDomainIdServiceConfigurationRecordID(input string) (*DomainIdServiceConfigurationRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DomainIdServiceConfigurationRecordId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DomainIdServiceConfigurationRecordId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDomainIdServiceConfigurationRecordIDInsensitively parses 'input' case-insensitively into a DomainIdServiceConfigurationRecordId
// note: this method should only be used for API response data and not user input
func ParseDomainIdServiceConfigurationRecordIDInsensitively(input string) (*DomainIdServiceConfigurationRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DomainIdServiceConfigurationRecordId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DomainIdServiceConfigurationRecordId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DomainIdServiceConfigurationRecordId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DomainId, ok = input.Parsed["domainId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "domainId", input)
	}

	if id.DomainDnsRecordId, ok = input.Parsed["domainDnsRecordId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "domainDnsRecordId", input)
	}

	return nil
}

// ValidateDomainIdServiceConfigurationRecordID checks that 'input' can be parsed as a Domain Id Service Configuration Record ID
func ValidateDomainIdServiceConfigurationRecordID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDomainIdServiceConfigurationRecordID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Domain Id Service Configuration Record ID
func (id DomainIdServiceConfigurationRecordId) ID() string {
	fmtString := "/domains/%s/serviceConfigurationRecords/%s"
	return fmt.Sprintf(fmtString, id.DomainId, id.DomainDnsRecordId)
}

// Segments returns a slice of Resource ID Segments which comprise this Domain Id Service Configuration Record ID
func (id DomainIdServiceConfigurationRecordId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("domains", "domains", "domains"),
		resourceids.UserSpecifiedSegment("domainId", "domainId"),
		resourceids.StaticSegment("serviceConfigurationRecords", "serviceConfigurationRecords", "serviceConfigurationRecords"),
		resourceids.UserSpecifiedSegment("domainDnsRecordId", "domainDnsRecordId"),
	}
}

// String returns a human-readable description of this Domain Id Service Configuration Record ID
func (id DomainIdServiceConfigurationRecordId) String() string {
	components := []string{
		fmt.Sprintf("Domain: %q", id.DomainId),
		fmt.Sprintf("Domain Dns Record: %q", id.DomainDnsRecordId),
	}
	return fmt.Sprintf("Domain Id Service Configuration Record (%s)", strings.Join(components, "\n"))
}
