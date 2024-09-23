package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInformationProtectionDataLossPreventionPolicyId{}

// UserIdInformationProtectionDataLossPreventionPolicyId is a struct representing the Resource ID for a User Id Information Protection Data Loss Prevention Policy
type UserIdInformationProtectionDataLossPreventionPolicyId struct {
	UserId                     string
	DataLossPreventionPolicyId string
}

// NewUserIdInformationProtectionDataLossPreventionPolicyID returns a new UserIdInformationProtectionDataLossPreventionPolicyId struct
func NewUserIdInformationProtectionDataLossPreventionPolicyID(userId string, dataLossPreventionPolicyId string) UserIdInformationProtectionDataLossPreventionPolicyId {
	return UserIdInformationProtectionDataLossPreventionPolicyId{
		UserId:                     userId,
		DataLossPreventionPolicyId: dataLossPreventionPolicyId,
	}
}

// ParseUserIdInformationProtectionDataLossPreventionPolicyID parses 'input' into a UserIdInformationProtectionDataLossPreventionPolicyId
func ParseUserIdInformationProtectionDataLossPreventionPolicyID(input string) (*UserIdInformationProtectionDataLossPreventionPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInformationProtectionDataLossPreventionPolicyId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInformationProtectionDataLossPreventionPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdInformationProtectionDataLossPreventionPolicyIDInsensitively parses 'input' case-insensitively into a UserIdInformationProtectionDataLossPreventionPolicyId
// note: this method should only be used for API response data and not user input
func ParseUserIdInformationProtectionDataLossPreventionPolicyIDInsensitively(input string) (*UserIdInformationProtectionDataLossPreventionPolicyId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInformationProtectionDataLossPreventionPolicyId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInformationProtectionDataLossPreventionPolicyId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdInformationProtectionDataLossPreventionPolicyId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.DataLossPreventionPolicyId, ok = input.Parsed["dataLossPreventionPolicyId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "dataLossPreventionPolicyId", input)
	}

	return nil
}

// ValidateUserIdInformationProtectionDataLossPreventionPolicyID checks that 'input' can be parsed as a User Id Information Protection Data Loss Prevention Policy ID
func ValidateUserIdInformationProtectionDataLossPreventionPolicyID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdInformationProtectionDataLossPreventionPolicyID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Information Protection Data Loss Prevention Policy ID
func (id UserIdInformationProtectionDataLossPreventionPolicyId) ID() string {
	fmtString := "/users/%s/informationProtection/dataLossPreventionPolicies/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.DataLossPreventionPolicyId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Information Protection Data Loss Prevention Policy ID
func (id UserIdInformationProtectionDataLossPreventionPolicyId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("dataLossPreventionPolicies", "dataLossPreventionPolicies", "dataLossPreventionPolicies"),
		resourceids.UserSpecifiedSegment("dataLossPreventionPolicyId", "dataLossPreventionPolicyId"),
	}
}

// String returns a human-readable description of this User Id Information Protection Data Loss Prevention Policy ID
func (id UserIdInformationProtectionDataLossPreventionPolicyId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Data Loss Prevention Policy: %q", id.DataLossPreventionPolicyId),
	}
	return fmt.Sprintf("User Id Information Protection Data Loss Prevention Policy (%s)", strings.Join(components, "\n"))
}
