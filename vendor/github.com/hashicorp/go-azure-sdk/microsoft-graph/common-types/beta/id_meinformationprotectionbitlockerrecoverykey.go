package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInformationProtectionBitlockerRecoveryKeyId{}

// MeInformationProtectionBitlockerRecoveryKeyId is a struct representing the Resource ID for a Me Information Protection Bitlocker Recovery Key
type MeInformationProtectionBitlockerRecoveryKeyId struct {
	BitlockerRecoveryKeyId string
}

// NewMeInformationProtectionBitlockerRecoveryKeyID returns a new MeInformationProtectionBitlockerRecoveryKeyId struct
func NewMeInformationProtectionBitlockerRecoveryKeyID(bitlockerRecoveryKeyId string) MeInformationProtectionBitlockerRecoveryKeyId {
	return MeInformationProtectionBitlockerRecoveryKeyId{
		BitlockerRecoveryKeyId: bitlockerRecoveryKeyId,
	}
}

// ParseMeInformationProtectionBitlockerRecoveryKeyID parses 'input' into a MeInformationProtectionBitlockerRecoveryKeyId
func ParseMeInformationProtectionBitlockerRecoveryKeyID(input string) (*MeInformationProtectionBitlockerRecoveryKeyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInformationProtectionBitlockerRecoveryKeyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInformationProtectionBitlockerRecoveryKeyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeInformationProtectionBitlockerRecoveryKeyIDInsensitively parses 'input' case-insensitively into a MeInformationProtectionBitlockerRecoveryKeyId
// note: this method should only be used for API response data and not user input
func ParseMeInformationProtectionBitlockerRecoveryKeyIDInsensitively(input string) (*MeInformationProtectionBitlockerRecoveryKeyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInformationProtectionBitlockerRecoveryKeyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInformationProtectionBitlockerRecoveryKeyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeInformationProtectionBitlockerRecoveryKeyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.BitlockerRecoveryKeyId, ok = input.Parsed["bitlockerRecoveryKeyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "bitlockerRecoveryKeyId", input)
	}

	return nil
}

// ValidateMeInformationProtectionBitlockerRecoveryKeyID checks that 'input' can be parsed as a Me Information Protection Bitlocker Recovery Key ID
func ValidateMeInformationProtectionBitlockerRecoveryKeyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeInformationProtectionBitlockerRecoveryKeyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Information Protection Bitlocker Recovery Key ID
func (id MeInformationProtectionBitlockerRecoveryKeyId) ID() string {
	fmtString := "/me/informationProtection/bitlocker/recoveryKeys/%s"
	return fmt.Sprintf(fmtString, id.BitlockerRecoveryKeyId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Information Protection Bitlocker Recovery Key ID
func (id MeInformationProtectionBitlockerRecoveryKeyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("bitlocker", "bitlocker", "bitlocker"),
		resourceids.StaticSegment("recoveryKeys", "recoveryKeys", "recoveryKeys"),
		resourceids.UserSpecifiedSegment("bitlockerRecoveryKeyId", "bitlockerRecoveryKeyId"),
	}
}

// String returns a human-readable description of this Me Information Protection Bitlocker Recovery Key ID
func (id MeInformationProtectionBitlockerRecoveryKeyId) String() string {
	components := []string{
		fmt.Sprintf("Bitlocker Recovery Key: %q", id.BitlockerRecoveryKeyId),
	}
	return fmt.Sprintf("Me Information Protection Bitlocker Recovery Key (%s)", strings.Join(components, "\n"))
}
