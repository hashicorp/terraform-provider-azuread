package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &RoleManagementExchangeCustomAppScopeId{}

// RoleManagementExchangeCustomAppScopeId is a struct representing the Resource ID for a Role Management Exchange Custom App Scope
type RoleManagementExchangeCustomAppScopeId struct {
	CustomAppScopeId string
}

// NewRoleManagementExchangeCustomAppScopeID returns a new RoleManagementExchangeCustomAppScopeId struct
func NewRoleManagementExchangeCustomAppScopeID(customAppScopeId string) RoleManagementExchangeCustomAppScopeId {
	return RoleManagementExchangeCustomAppScopeId{
		CustomAppScopeId: customAppScopeId,
	}
}

// ParseRoleManagementExchangeCustomAppScopeID parses 'input' into a RoleManagementExchangeCustomAppScopeId
func ParseRoleManagementExchangeCustomAppScopeID(input string) (*RoleManagementExchangeCustomAppScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementExchangeCustomAppScopeId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementExchangeCustomAppScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseRoleManagementExchangeCustomAppScopeIDInsensitively parses 'input' case-insensitively into a RoleManagementExchangeCustomAppScopeId
// note: this method should only be used for API response data and not user input
func ParseRoleManagementExchangeCustomAppScopeIDInsensitively(input string) (*RoleManagementExchangeCustomAppScopeId, error) {
	parser := resourceids.NewParserFromResourceIdType(&RoleManagementExchangeCustomAppScopeId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := RoleManagementExchangeCustomAppScopeId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *RoleManagementExchangeCustomAppScopeId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.CustomAppScopeId, ok = input.Parsed["customAppScopeId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "customAppScopeId", input)
	}

	return nil
}

// ValidateRoleManagementExchangeCustomAppScopeID checks that 'input' can be parsed as a Role Management Exchange Custom App Scope ID
func ValidateRoleManagementExchangeCustomAppScopeID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseRoleManagementExchangeCustomAppScopeID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Role Management Exchange Custom App Scope ID
func (id RoleManagementExchangeCustomAppScopeId) ID() string {
	fmtString := "/roleManagement/exchange/customAppScopes/%s"
	return fmt.Sprintf(fmtString, id.CustomAppScopeId)
}

// Segments returns a slice of Resource ID Segments which comprise this Role Management Exchange Custom App Scope ID
func (id RoleManagementExchangeCustomAppScopeId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("roleManagement", "roleManagement", "roleManagement"),
		resourceids.StaticSegment("exchange", "exchange", "exchange"),
		resourceids.StaticSegment("customAppScopes", "customAppScopes", "customAppScopes"),
		resourceids.UserSpecifiedSegment("customAppScopeId", "customAppScopeId"),
	}
}

// String returns a human-readable description of this Role Management Exchange Custom App Scope ID
func (id RoleManagementExchangeCustomAppScopeId) String() string {
	components := []string{
		fmt.Sprintf("Custom App Scope: %q", id.CustomAppScopeId),
	}
	return fmt.Sprintf("Role Management Exchange Custom App Scope (%s)", strings.Join(components, "\n"))
}
