package stable

import (
	"fmt"
	"strings"

	"github.com/hashicorp/go-azure-helpers/resourcemanager/resourceids"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ resourceids.ResourceId = &MeInferenceClassificationOverrideId{}

// MeInferenceClassificationOverrideId is a struct representing the Resource ID for a Me Inference Classification Override
type MeInferenceClassificationOverrideId struct {
	InferenceClassificationOverrideId string
}

// NewMeInferenceClassificationOverrideID returns a new MeInferenceClassificationOverrideId struct
func NewMeInferenceClassificationOverrideID(inferenceClassificationOverrideId string) MeInferenceClassificationOverrideId {
	return MeInferenceClassificationOverrideId{
		InferenceClassificationOverrideId: inferenceClassificationOverrideId,
	}
}

// ParseMeInferenceClassificationOverrideID parses 'input' into a MeInferenceClassificationOverrideId
func ParseMeInferenceClassificationOverrideID(input string) (*MeInferenceClassificationOverrideId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInferenceClassificationOverrideId{})
	parsed, err := parser.Parse(input, false)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInferenceClassificationOverrideId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

// ParseMeInferenceClassificationOverrideIDInsensitively parses 'input' case-insensitively into a MeInferenceClassificationOverrideId
// note: this method should only be used for API response data and not user input
func ParseMeInferenceClassificationOverrideIDInsensitively(input string) (*MeInferenceClassificationOverrideId, error) {
	parser := resourceids.NewParserFromResourceIdType(&MeInferenceClassificationOverrideId{})
	parsed, err := parser.Parse(input, true)
	if err != nil {
		return nil, fmt.Errorf("parsing %q: %+v", input, err)
	}

	id := MeInferenceClassificationOverrideId{}
	if err = id.FromParseResult(*parsed); err != nil {
		return nil, err
	}

	return &id, nil
}

func (id *MeInferenceClassificationOverrideId) FromParseResult(input resourceids.ParseResult) error {
	var ok bool

	if id.InferenceClassificationOverrideId, ok = input.Parsed["inferenceClassificationOverrideId"]; !ok {
		return resourceids.NewSegmentNotSpecifiedError(id, "inferenceClassificationOverrideId", input)
	}

	return nil
}

// ValidateMeInferenceClassificationOverrideID checks that 'input' can be parsed as a Me Inference Classification Override ID
func ValidateMeInferenceClassificationOverrideID(input interface{}, key string) (warnings []string, errors []error) {
	v, ok := input.(string)
	if !ok {
		errors = append(errors, fmt.Errorf("expected %q to be a string", key))
		return
	}

	if _, err := ParseMeInferenceClassificationOverrideID(v); err != nil {
		errors = append(errors, err)
	}

	return
}

// ID returns the formatted Me Inference Classification Override ID
func (id MeInferenceClassificationOverrideId) ID() string {
	fmtString := "/me/inferenceClassification/overrides/%s"
	return fmt.Sprintf(fmtString, id.InferenceClassificationOverrideId)
}

// Segments returns a slice of Resource ID Segments which comprise this Me Inference Classification Override ID
func (id MeInferenceClassificationOverrideId) Segments() []resourceids.Segment {
	return []resourceids.Segment{
		resourceids.StaticSegment("me", "me", "me"),
		resourceids.StaticSegment("inferenceClassification", "inferenceClassification", "inferenceClassification"),
		resourceids.StaticSegment("overrides", "overrides", "overrides"),
		resourceids.UserSpecifiedSegment("inferenceClassificationOverrideId", "inferenceClassificationOverrideId"),
	}
}

// String returns a human-readable description of this Me Inference Classification Override ID
func (id MeInferenceClassificationOverrideId) String() string {
	components := []string{
		fmt.Sprintf("Inference Classification Override: %q", id.InferenceClassificationOverrideId),
	}
	return fmt.Sprintf("Me Inference Classification Override (%s)", strings.Join(components, "\n"))
}
