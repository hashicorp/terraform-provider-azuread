package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId{}

// DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId is a struct representing the Resource ID for a Directory Public Key Infrastructure Certificate Based Auth Configuration Id Certificate Authority
type DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId struct {
	CertificateBasedAuthPkiId    string
	CertificateAuthorityDetailId string
}

// NewDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityID returns a new DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId struct
func NewDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityID(certificateBasedAuthPkiId string, certificateAuthorityDetailId string) DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId {
	return DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId{
		CertificateBasedAuthPkiId:    certificateBasedAuthPkiId,
		CertificateAuthorityDetailId: certificateAuthorityDetailId,
	}
}

// ParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityID parses 'input' into a DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId
func ParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityID(input string) (*DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityIDInsensitively parses 'input' case-insensitively into a DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId
// note: this method should only be used for API response data and not user input
func ParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityIDInsensitively(input string) (*DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CertificateBasedAuthPkiId, ok = input.Parsed["certificateBasedAuthPkiId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "certificateBasedAuthPkiId", input)
	}

	if id.CertificateAuthorityDetailId, ok = input.Parsed["certificateAuthorityDetailId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "certificateAuthorityDetailId", input)
	}

	return nil
}

// ValidateDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityID checks that 'input' can be parsed as a Directory Public Key Infrastructure Certificate Based Auth Configuration Id Certificate Authority ID
func ValidateDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Public Key Infrastructure Certificate Based Auth Configuration Id Certificate Authority ID
func (id DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId) ID() string {
	fmtString := "/directory/publicKeyInfrastructure/certificateBasedAuthConfigurations/%s/certificateAuthorities/%s"
	return fmt.Sprintf(fmtString, id.CertificateBasedAuthPkiId, id.CertificateAuthorityDetailId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Public Key Infrastructure Certificate Based Auth Configuration Id Certificate Authority ID
func (id DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("publicKeyInfrastructure", "publicKeyInfrastructure", "publicKeyInfrastructure"),
		resourceids.StaticSegment("certificateBasedAuthConfigurations", "certificateBasedAuthConfigurations", "certificateBasedAuthConfigurations"),
		resourceids.UserSpecifiedSegment("certificateBasedAuthPkiId", "certificateBasedAuthPkiId"),
		resourceids.StaticSegment("certificateAuthorities", "certificateAuthorities", "certificateAuthorities"),
		resourceids.UserSpecifiedSegment("certificateAuthorityDetailId", "certificateAuthorityDetailId"),
	}
}

// String returns a human-readable description of this Directory Public Key Infrastructure Certificate Based Auth Configuration Id Certificate Authority ID
func (id DirectoryPublicKeyInfrastructureCertificateBasedAuthConfigurationIdCertificateAuthorityId) String() string {
	components := []string{
		fmt.Sprintf("Certificate Based Auth Pki: %q", id.CertificateBasedAuthPkiId),
		fmt.Sprintf("Certificate Authority Detail: %q", id.CertificateAuthorityDetailId),
	}
	return fmt.Sprintf("Directory Public Key Infrastructure Certificate Based Auth Configuration Id Certificate Authority (%s)", strings.Join(components, "\n"))
}
