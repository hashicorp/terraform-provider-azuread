package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInformationProtectionBitlockerRecoveryKeyId{}

// UserIdInformationProtectionBitlockerRecoveryKeyId is a struct representing the Resource ID for a User Id Information Protection Bitlocker Recovery Key
type UserIdInformationProtectionBitlockerRecoveryKeyId struct {
	UserId                 string
	BitlockerRecoveryKeyId string
}

// NewUserIdInformationProtectionBitlockerRecoveryKeyID returns a new UserIdInformationProtectionBitlockerRecoveryKeyId struct
func NewUserIdInformationProtectionBitlockerRecoveryKeyID(userId string, bitlockerRecoveryKeyId string) UserIdInformationProtectionBitlockerRecoveryKeyId {
	return UserIdInformationProtectionBitlockerRecoveryKeyId{
		UserId:                 userId,
		BitlockerRecoveryKeyId: bitlockerRecoveryKeyId,
	}
}

// ParseUserIdInformationProtectionBitlockerRecoveryKeyID parses 'input' into a UserIdInformationProtectionBitlockerRecoveryKeyId
func ParseUserIdInformationProtectionBitlockerRecoveryKeyID(input string) (*UserIdInformationProtectionBitlockerRecoveryKeyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInformationProtectionBitlockerRecoveryKeyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInformationProtectionBitlockerRecoveryKeyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdInformationProtectionBitlockerRecoveryKeyIDInsensitively parses 'input' case-insensitively into a UserIdInformationProtectionBitlockerRecoveryKeyId
// note: this method should only be used for API response data and not user input
func ParseUserIdInformationProtectionBitlockerRecoveryKeyIDInsensitively(input string) (*UserIdInformationProtectionBitlockerRecoveryKeyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInformationProtectionBitlockerRecoveryKeyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInformationProtectionBitlockerRecoveryKeyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdInformationProtectionBitlockerRecoveryKeyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.BitlockerRecoveryKeyId, ok = input.Parsed["bitlockerRecoveryKeyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "bitlockerRecoveryKeyId", input)
	}

	return nil
}

// ValidateUserIdInformationProtectionBitlockerRecoveryKeyID checks that 'input' can be parsed as a User Id Information Protection Bitlocker Recovery Key ID
func ValidateUserIdInformationProtectionBitlockerRecoveryKeyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdInformationProtectionBitlockerRecoveryKeyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Information Protection Bitlocker Recovery Key ID
func (id UserIdInformationProtectionBitlockerRecoveryKeyId) ID() string {
	fmtString := "/users/%s/informationProtection/bitlocker/recoveryKeys/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.BitlockerRecoveryKeyId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Information Protection Bitlocker Recovery Key ID
func (id UserIdInformationProtectionBitlockerRecoveryKeyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("bitlocker", "bitlocker", "bitlocker"),
		resourceids.StaticSegment("recoveryKeys", "recoveryKeys", "recoveryKeys"),
		resourceids.UserSpecifiedSegment("bitlockerRecoveryKeyId", "bitlockerRecoveryKeyId"),
	}
}

// String returns a human-readable description of this User Id Information Protection Bitlocker Recovery Key ID
func (id UserIdInformationProtectionBitlockerRecoveryKeyId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Bitlocker Recovery Key: %q", id.BitlockerRecoveryKeyId),
	}
	return fmt.Sprintf("User Id Information Protection Bitlocker Recovery Key (%s)", strings.Join(components, "\n"))
}
