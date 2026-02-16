package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAlertTemplate struct {
	// Category assigned to the alert triggered by the custom detection rule.
	Category *string `json:"category,omitempty"`

	// Description of the alert triggered by the custom detection rule.
	Description *string `json:"description,omitempty"`

	// Which asset or assets were impacted based on the alert triggered by the custom detection rule.
	ImpactedAssets *[]SecurityImpactedAsset `json:"impactedAssets,omitempty"`

	// MITRE technique assigned to the alert triggered by the custom detection rule.
	MitreTechniques *[]string `json:"mitreTechniques,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Recommended actions to mitigate the threat related to the alert triggered by the custom detection rule.
	RecommendedActions nullable.Type[string] `json:"recommendedActions,omitempty"`

	Severity *SecurityAlertSeverity `json:"severity,omitempty"`

	// Name of the alert triggered by the custom detection rule.
	Title *string `json:"title,omitempty"`
}

var _ json.Unmarshaler = &SecurityAlertTemplate{}

func (s *SecurityAlertTemplate) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Category           *string                `json:"category,omitempty"`
		Description        *string                `json:"description,omitempty"`
		MitreTechniques    *[]string              `json:"mitreTechniques,omitempty"`
		ODataId            *string                `json:"@odata.id,omitempty"`
		ODataType          *string                `json:"@odata.type,omitempty"`
		RecommendedActions nullable.Type[string]  `json:"recommendedActions,omitempty"`
		Severity           *SecurityAlertSeverity `json:"severity,omitempty"`
		Title              *string                `json:"title,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Category = decoded.Category
	s.Description = decoded.Description
	s.MitreTechniques = decoded.MitreTechniques
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RecommendedActions = decoded.RecommendedActions
	s.Severity = decoded.Severity
	s.Title = decoded.Title

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SecurityAlertTemplate into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["impactedAssets"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ImpactedAssets into list []json.RawMessage: %+v", err)
		}

		output := make([]SecurityImpactedAsset, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalSecurityImpactedAssetImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ImpactedAssets' for 'SecurityAlertTemplate': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ImpactedAssets = &output
	}

	return nil
}
