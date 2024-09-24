package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId{}

// IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId is a struct representing the Resource ID for a Identity Governance Role Management Alert Alert Id Alert Incident
type IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId struct {
	UnifiedRoleManagementAlertId         string
	UnifiedRoleManagementAlertIncidentId string
}

// NewIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentID returns a new IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId struct
func NewIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentID(unifiedRoleManagementAlertId string, unifiedRoleManagementAlertIncidentId string) IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId {
	return IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId{
		UnifiedRoleManagementAlertId:         unifiedRoleManagementAlertId,
		UnifiedRoleManagementAlertIncidentId: unifiedRoleManagementAlertIncidentId,
	}
}

// ParseIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentID parses 'input' into a IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId
func ParseIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentID(input string) (*IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentIDInsensitively parses 'input' case-insensitively into a IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId
// note: this method should only be used for API response data and not user input
func ParseIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentIDInsensitively(input string) (*IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId, error) {
	parser := resourceids.NewParserFromResourceIdType(&IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.UnifiedRoleManagementAlertId, ok = input.Parsed["unifiedRoleManagementAlertId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementAlertId", input)
	}

	if id.UnifiedRoleManagementAlertIncidentId, ok = input.Parsed["unifiedRoleManagementAlertIncidentId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "unifiedRoleManagementAlertIncidentId", input)
	}

	return nil
}

// ValidateIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentID checks that 'input' can be parsed as a Identity Governance Role Management Alert Alert Id Alert Incident ID
func ValidateIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseIdentityGovernanceRoleManagementAlertAlertIdAlertIncidentID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Identity Governance Role Management Alert Alert Id Alert Incident ID
func (id IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId) ID() string {
	fmtString := "/identityGovernance/roleManagementAlerts/alerts/%s/alertIncidents/%s"
	return fmt.Sprintf(fmtString, id.UnifiedRoleManagementAlertId, id.UnifiedRoleManagementAlertIncidentId)
}

// Segments returns a slice of Resource ID Segments which comprise this Identity Governance Role Management Alert Alert Id Alert Incident ID
func (id IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("identityGovernance", "identityGovernance", "identityGovernance"),
		resourceids.StaticSegment("roleManagementAlerts", "roleManagementAlerts", "roleManagementAlerts"),
		resourceids.StaticSegment("alerts", "alerts", "alerts"),
		resourceids.UserSpecifiedSegment("unifiedRoleManagementAlertId", "unifiedRoleManagementAlertId"),
		resourceids.StaticSegment("alertIncidents", "alertIncidents", "alertIncidents"),
		resourceids.UserSpecifiedSegment("unifiedRoleManagementAlertIncidentId", "unifiedRoleManagementAlertIncidentId"),
	}
}

// String returns a human-readable description of this Identity Governance Role Management Alert Alert Id Alert Incident ID
func (id IdentityGovernanceRoleManagementAlertAlertIdAlertIncidentId) String() string {
	components := []string{
		fmt.Sprintf("Unified Role Management Alert: %q", id.UnifiedRoleManagementAlertId),
		fmt.Sprintf("Unified Role Management Alert Incident: %q", id.UnifiedRoleManagementAlertIncidentId),
	}
	return fmt.Sprintf("Identity Governance Role Management Alert Alert Id Alert Incident (%s)", strings.Join(components, "\n"))
}
