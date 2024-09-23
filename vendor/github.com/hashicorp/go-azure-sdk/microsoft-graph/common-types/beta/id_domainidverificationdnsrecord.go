package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DomainIdVerificationDnsRecordId{}

// DomainIdVerificationDnsRecordId is a struct representing the Resource ID for a Domain Id Verification Dns Record
type DomainIdVerificationDnsRecordId struct {
	DomainId          string
	DomainDnsRecordId string
}

// NewDomainIdVerificationDnsRecordID returns a new DomainIdVerificationDnsRecordId struct
func NewDomainIdVerificationDnsRecordID(domainId string, domainDnsRecordId string) DomainIdVerificationDnsRecordId {
	return DomainIdVerificationDnsRecordId{
		DomainId:          domainId,
		DomainDnsRecordId: domainDnsRecordId,
	}
}

// ParseDomainIdVerificationDnsRecordID parses 'input' into a DomainIdVerificationDnsRecordId
func ParseDomainIdVerificationDnsRecordID(input string) (*DomainIdVerificationDnsRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DomainIdVerificationDnsRecordId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DomainIdVerificationDnsRecordId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDomainIdVerificationDnsRecordIDInsensitively parses 'input' case-insensitively into a DomainIdVerificationDnsRecordId
// note: this method should only be used for API response data and not user input
func ParseDomainIdVerificationDnsRecordIDInsensitively(input string) (*DomainIdVerificationDnsRecordId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DomainIdVerificationDnsRecordId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DomainIdVerificationDnsRecordId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DomainIdVerificationDnsRecordId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.DomainId, ok = input.Parsed["domainId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "domainId", input)
	}

	if id.DomainDnsRecordId, ok = input.Parsed["domainDnsRecordId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "domainDnsRecordId", input)
	}

	return nil
}

// ValidateDomainIdVerificationDnsRecordID checks that 'input' can be parsed as a Domain Id Verification Dns Record ID
func ValidateDomainIdVerificationDnsRecordID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDomainIdVerificationDnsRecordID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Domain Id Verification Dns Record ID
func (id DomainIdVerificationDnsRecordId) ID() string {
	fmtString := "/domains/%s/verificationDnsRecords/%s"
	return fmt.Sprintf(fmtString, id.DomainId, id.DomainDnsRecordId)
}

// Segments returns a slice of Resource ID Segments which comprise this Domain Id Verification Dns Record ID
func (id DomainIdVerificationDnsRecordId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("domains", "domains", "domains"),
		resourceids.UserSpecifiedSegment("domainId", "domainId"),
		resourceids.StaticSegment("verificationDnsRecords", "verificationDnsRecords", "verificationDnsRecords"),
		resourceids.UserSpecifiedSegment("domainDnsRecordId", "domainDnsRecordId"),
	}
}

// String returns a human-readable description of this Domain Id Verification Dns Record ID
func (id DomainIdVerificationDnsRecordId) String() string {
	components := []string{
		fmt.Sprintf("Domain: %q", id.DomainId),
		fmt.Sprintf("Domain Dns Record: %q", id.DomainDnsRecordId),
	}
	return fmt.Sprintf("Domain Id Verification Dns Record (%s)", strings.Join(components, "\n"))
}
