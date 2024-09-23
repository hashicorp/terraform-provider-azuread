package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId{}

// DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId is a struct representing the Resource ID for a Directory Certificate Authority Certificate Based Application Configuration Id Trusted Certificate Authority
type DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId struct {
	CertificateBasedApplicationConfigurationId string
	CertificateAuthorityAsEntityId             string
}

// NewDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityID returns a new DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId struct
func NewDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityID(certificateBasedApplicationConfigurationId string, certificateAuthorityAsEntityId string) DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId {
	return DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId{
		CertificateBasedApplicationConfigurationId: certificateBasedApplicationConfigurationId,
		CertificateAuthorityAsEntityId:             certificateAuthorityAsEntityId,
	}
}

// ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityID parses 'input' into a DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId
func ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityID(input string) (*DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityIDInsensitively parses 'input' case-insensitively into a DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId
// note: this method should only be used for API response data and not user input
func ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityIDInsensitively(input string) (*DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CertificateBasedApplicationConfigurationId, ok = input.Parsed["certificateBasedApplicationConfigurationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "certificateBasedApplicationConfigurationId", input)
	}

	if id.CertificateAuthorityAsEntityId, ok = input.Parsed["certificateAuthorityAsEntityId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "certificateAuthorityAsEntityId", input)
	}

	return nil
}

// ValidateDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityID checks that 'input' can be parsed as a Directory Certificate Authority Certificate Based Application Configuration Id Trusted Certificate Authority ID
func ValidateDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Certificate Authority Certificate Based Application Configuration Id Trusted Certificate Authority ID
func (id DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId) ID() string {
	fmtString := "/directory/certificateAuthorities/certificateBasedApplicationConfigurations/%s/trustedCertificateAuthorities/%s"
	return fmt.Sprintf(fmtString, id.CertificateBasedApplicationConfigurationId, id.CertificateAuthorityAsEntityId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Certificate Authority Certificate Based Application Configuration Id Trusted Certificate Authority ID
func (id DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("certificateAuthorities", "certificateAuthorities", "certificateAuthorities"),
		resourceids.StaticSegment("certificateBasedApplicationConfigurations", "certificateBasedApplicationConfigurations", "certificateBasedApplicationConfigurations"),
		resourceids.UserSpecifiedSegment("certificateBasedApplicationConfigurationId", "certificateBasedApplicationConfigurationId"),
		resourceids.StaticSegment("trustedCertificateAuthorities", "trustedCertificateAuthorities", "trustedCertificateAuthorities"),
		resourceids.UserSpecifiedSegment("certificateAuthorityAsEntityId", "certificateAuthorityAsEntityId"),
	}
}

// String returns a human-readable description of this Directory Certificate Authority Certificate Based Application Configuration Id Trusted Certificate Authority ID
func (id DirectoryCertificateAuthorityCertificateBasedApplicationConfigurationIdTrustedCertificateAuthorityId) String() string {
	components := []string{
		fmt.Sprintf("Certificate Based Application Configuration: %q", id.CertificateBasedApplicationConfigurationId),
		fmt.Sprintf("Certificate Authority As Entity: %q", id.CertificateAuthorityAsEntityId),
	}
	return fmt.Sprintf("Directory Certificate Authority Certificate Based Application Configuration Id Trusted Certificate Authority (%s)", strings.Join(components, "\n"))
}
