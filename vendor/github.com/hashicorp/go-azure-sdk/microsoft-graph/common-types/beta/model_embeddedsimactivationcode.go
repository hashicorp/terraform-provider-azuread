package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EmbeddedSIMActivationCode struct {
	// The Integrated Circuit Card Identifier (ICCID) for this embedded SIM activation code as provided by the mobile
	// operator.
	IntegratedCircuitCardIdentifier *string `json:"integratedCircuitCardIdentifier,omitempty"`

	// The MatchingIdentifier (MatchingID) as specified in the GSMA Association SGP.22 RSP Technical Specification section
	// 4.1.
	MatchingIdentifier *string `json:"matchingIdentifier,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The fully qualified domain name of the SM-DP+ server as specified in the GSM Association SPG .22 RSP Technical
	// Specification.
	SmdpPlusServerAddress *string `json:"smdpPlusServerAddress,omitempty"`
}
