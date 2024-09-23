package stable

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityProtectionRoot struct {
	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Risk detection in Microsoft Entra ID Protection and the associated information about the detection.
	RiskDetections *[]RiskDetection `json:"riskDetections,omitempty"`

	// Microsoft Entra service principals that are at risk.
	RiskyServicePrincipals *[]RiskyServicePrincipal `json:"riskyServicePrincipals,omitempty"`

	// Users that are flagged as at-risk by Microsoft Entra ID Protection.
	RiskyUsers *[]RiskyUser `json:"riskyUsers,omitempty"`

	// Represents information about detected at-risk service principals in a Microsoft Entra tenant.
	ServicePrincipalRiskDetections *[]ServicePrincipalRiskDetection `json:"servicePrincipalRiskDetections,omitempty"`
}

var _ json.Unmarshaler = &IdentityProtectionRoot{}

func (s *IdentityProtectionRoot) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ODataId                        *string                          `json:"@odata.id,omitempty"`
		ODataType                      *string                          `json:"@odata.type,omitempty"`
		RiskDetections                 *[]RiskDetection                 `json:"riskDetections,omitempty"`
		ServicePrincipalRiskDetections *[]ServicePrincipalRiskDetection `json:"servicePrincipalRiskDetections,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType
	s.RiskDetections = decoded.RiskDetections
	s.ServicePrincipalRiskDetections = decoded.ServicePrincipalRiskDetections

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling IdentityProtectionRoot into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["riskyServicePrincipals"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RiskyServicePrincipals into list []json.RawMessage: %+v", err)
		}

		output := make([]RiskyServicePrincipal, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRiskyServicePrincipalImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RiskyServicePrincipals' for 'IdentityProtectionRoot': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RiskyServicePrincipals = &output
	}

	if v, ok := temp["riskyUsers"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling RiskyUsers into list []json.RawMessage: %+v", err)
		}

		output := make([]RiskyUser, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalRiskyUserImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'RiskyUsers' for 'IdentityProtectionRoot': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.RiskyUsers = &output
	}

	return nil
}
