package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &UserIdInformationProtectionThreatAssessmentRequestId{}

// UserIdInformationProtectionThreatAssessmentRequestId is a struct representing the Resource ID for a User Id Information Protection Threat Assessment Request
type UserIdInformationProtectionThreatAssessmentRequestId struct {
	UserId                    string
	ThreatAssessmentRequestId string
}

// NewUserIdInformationProtectionThreatAssessmentRequestID returns a new UserIdInformationProtectionThreatAssessmentRequestId struct
func NewUserIdInformationProtectionThreatAssessmentRequestID(userId string, threatAssessmentRequestId string) UserIdInformationProtectionThreatAssessmentRequestId {
	return UserIdInformationProtectionThreatAssessmentRequestId{
		UserId:                    userId,
		ThreatAssessmentRequestId: threatAssessmentRequestId,
	}
}

// ParseUserIdInformationProtectionThreatAssessmentRequestID parses 'input' into a UserIdInformationProtectionThreatAssessmentRequestId
func ParseUserIdInformationProtectionThreatAssessmentRequestID(input string) (*UserIdInformationProtectionThreatAssessmentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInformationProtectionThreatAssessmentRequestId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInformationProtectionThreatAssessmentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseUserIdInformationProtectionThreatAssessmentRequestIDInsensitively parses 'input' case-insensitively into a UserIdInformationProtectionThreatAssessmentRequestId
// note: this method should only be used for API response data and not user input
func ParseUserIdInformationProtectionThreatAssessmentRequestIDInsensitively(input string) (*UserIdInformationProtectionThreatAssessmentRequestId, error) {
	parser := resourceids.NewParserFromResourceIdType(&UserIdInformationProtectionThreatAssessmentRequestId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := UserIdInformationProtectionThreatAssessmentRequestId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *UserIdInformationProtectionThreatAssessmentRequestId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UserId, ok = input.Parsed["userId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "userId", input)
	}

	if id.ThreatAssessmentRequestId, ok = input.Parsed["threatAssessmentRequestId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "threatAssessmentRequestId", input)
	}

	return nil
}

// ValidateUserIdInformationProtectionThreatAssessmentRequestID checks that 'input' can be parsed as a User Id Information Protection Threat Assessment Request ID
func ValidateUserIdInformationProtectionThreatAssessmentRequestID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseUserIdInformationProtectionThreatAssessmentRequestID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted User Id Information Protection Threat Assessment Request ID
func (id UserIdInformationProtectionThreatAssessmentRequestId) ID() string {
	fmtString := "/users/%s/informationProtection/threatAssessmentRequests/%s"
	return fmt.Sprintf(fmtString, id.UserId, id.ThreatAssessmentRequestId)
}

// Segments returns a slice of Resource ID Segments which comprise this User Id Information Protection Threat Assessment Request ID
func (id UserIdInformationProtectionThreatAssessmentRequestId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("users", "users", "users"),
		resourceids.UserSpecifiedSegment("userId", "userId"),
		resourceids.StaticSegment("informationProtection", "informationProtection", "informationProtection"),
		resourceids.StaticSegment("threatAssessmentRequests", "threatAssessmentRequests", "threatAssessmentRequests"),
		resourceids.UserSpecifiedSegment("threatAssessmentRequestId", "threatAssessmentRequestId"),
	}
}

// String returns a human-readable description of this User Id Information Protection Threat Assessment Request ID
func (id UserIdInformationProtectionThreatAssessmentRequestId) String() string {
	components := []string{
		fmt.Sprintf("User: %q", id.UserId),
		fmt.Sprintf("Threat Assessment Request: %q", id.ThreatAssessmentRequestId),
	}
	return fmt.Sprintf("User Id Information Protection Threat Assessment Request (%s)", strings.Join(components, "\n"))
}
