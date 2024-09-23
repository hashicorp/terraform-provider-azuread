package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId{}

// DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId is a struct representing the Resource ID for a Directory Certificate Authority Certificate Based Application Configuration
type DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId struct {
	CertificateBasedApplicationConfigurationId string
}

// NewDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationID returns a new DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId struct
func NewDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationID(certificateBasedApplicationConfigurationId string) DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId {
	return DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId{
		CertificateBasedApplicationConfigurationId: certificateBasedApplicationConfigurationId,
	}
}

// ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationID parses 'input' into a DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId
func ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationID(input string) (*DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIDInsensitively parses 'input' case-insensitively into a DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId
// note: this method should only be used for API response data and not user input
func ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIDInsensitively(input string) (*DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CertificateBasedApplicationConfigurationId, ok = input.Parsed["certificateBasedApplicationConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "certificateBasedApplicationConfigurationId", input)
	}

	return nil
}

// ValidateDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationID checks that 'input' can be parsed as a Directory Certificate Authority Certificate Based Application Configuration ID
func ValidateDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Certificate Authority Certificate Based Application Configuration ID
func (id DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId) ID() string {
	fmtString := "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/%s"
	return fmt.Sprintf(fmtString, id.CertificateBasedApplicationConfigurationId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Certificate Authority Certificate Based Application Configuration ID
func (id DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("certificateAuthorities", "certificateAuthorities", "certificateAuthorities"),
		resourceids.StaticSegment("certificateBasedApplicationConfigurations", "certificateBasedApplicationConfigurations", "certificateBasedApplicationConfigurations"),
		resourceids.UserSpecifiedSegment("certificateBasedApplicationConfigurationId", "certificateBasedApplicationConfigurationId"),
	}
}

// String returns a human-readable description of this Directory Certificate Authority Certificate Based Application Configuration ID
func (id DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationId) String() string {
	components := []string{
		fmt.Sprintf("Certificate Based Application Configuration: %q", id.CertificateBasedApplicationConfigurationId),
	}
	return fmt.Sprintf("Directory Certificate Authority Certificate Based Application Configuration (%s)", strings.Join(components, "\n"))
}
