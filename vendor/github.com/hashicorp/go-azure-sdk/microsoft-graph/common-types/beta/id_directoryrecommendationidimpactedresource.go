package beta

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &DirectoryRecommendationIdImpactedResourceId{}

// DirectoryRecommendationIdImpactedResourceId is a struct representing the Resource ID for a Directory Recommendation Id Impacted Resource
type DirectoryRecommendationIdImpactedResourceId struct {
	RecommendationId   string
	ImpactedResourceId string
}

// NewDirectoryRecommendationIdImpactedResourceID returns a new DirectoryRecommendationIdImpactedResourceId struct
func NewDirectoryRecommendationIdImpactedResourceID(recommendationId string, impactedResourceId string) DirectoryRecommendationIdImpactedResourceId {
	return DirectoryRecommendationIdImpactedResourceId{
		RecommendationId:   recommendationId,
		ImpactedResourceId: impactedResourceId,
	}
}

// ParseDirectoryRecommendationIdImpactedResourceID parses 'input' into a DirectoryRecommendationIdImpactedResourceId
func ParseDirectoryRecommendationIdImpactedResourceID(input string) (*DirectoryRecommendationIdImpactedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryRecommendationIdImpactedResourceId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryRecommendationIdImpactedResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseDirectoryRecommendationIdImpactedResourceIDInsensitively parses 'input' case-insensitively into a DirectoryRecommendationIdImpactedResourceId
// note: this method should only be used for API response data and not user input
func ParseDirectoryRecommendationIdImpactedResourceIDInsensitively(input string) (*DirectoryRecommendationIdImpactedResourceId, error) {
	parser := resourceids.NewParserFromResourceIdType(&DirectoryRecommendationIdImpactedResourceId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := DirectoryRecommendationIdImpactedResourceId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *DirectoryRecommendationIdImpactedResourceId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.RecommendationId, ok = input.Parsed["recommendationId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "recommendationId", input)
	}

	if id.ImpactedResourceId, ok = input.Parsed["impactedResourceId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "impactedResourceId", input)
	}

	return nil
}

// ValidateDirectoryRecommendationIdImpactedResourceID checks that 'input' can be parsed as a Directory Recommendation Id Impacted Resource ID
func ValidateDirectoryRecommendationIdImpactedResourceID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseDirectoryRecommendationIdImpactedResourceID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Directory Recommendation Id Impacted Resource ID
func (id DirectoryRecommendationIdImpactedResourceId) ID() string {
	fmtString := "/directory/recommendations/%s/impactedResources/%s"
	return fmt.Sprintf(fmtString, id.RecommendationId, id.ImpactedResourceId)
}

// Segments returns a slice of Resource ID Segments which comprise this Directory Recommendation Id Impacted Resource ID
func (id DirectoryRecommendationIdImpactedResourceId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("directory", "directory", "directory"),
		resourceids.StaticSegment("recommendations", "recommendations", "recommendations"),
		resourceids.UserSpecifiedSegment("recommendationId", "recommendationId"),
		resourceids.StaticSegment("impactedResources", "impactedResources", "impactedResources"),
		resourceids.UserSpecifiedSegment("impactedResourceId", "impactedResourceId"),
	}
}

// String returns a human-readable description of this Directory Recommendation Id Impacted Resource ID
func (id DirectoryRecommendationIdImpactedResourceId) String() string {
	components := []string{
		fmt.Sprintf("Recommendation: %q", id.RecommendationId),
		fmt.Sprintf("Impacted Resource: %q", id.ImpactedResourceId),
	}
	return fmt.Sprintf("Directory Recommendation Id Impacted Resource (%s)", strings.Join(components, "\n"))
}
