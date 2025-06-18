package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CloudPCDomainJoinConfiguration struct {
	// Specifies the method by which the provisioned Cloud PC joins Microsoft Entra ID. If you choose the hybridAzureADJoin
	// type, only provide a value for the onPremisesConnectionId property and leave the regionName property empty. If you
	// choose the azureADJoin type, provide a value for either the onPremisesConnectionId or the regionName property.
	// Possible values are: azureADJoin, hybridAzureADJoin, unknownFutureValue.
	DomainJoinType *CloudPCDomainJoinType `json:"domainJoinType,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// The Azure network connection ID that matches the virtual network IT admins want the provisioning policy to use when
	// they create Cloud PCs. You can use this property in both domain join types: Azure AD joined or Hybrid Microsoft Entra
	// joined. If you enter an onPremisesConnectionId, leave the regionName property empty.
	OnPremisesConnectionId nullable.Type[string] `json:"onPremisesConnectionId,omitempty"`

	// The logical geographic group this region belongs to. Multiple regions can belong to one region group. A customer can
	// select a regionGroup when they provision a Cloud PC, and the Cloud PC is put in one of the regions in the group based
	// on resource status. For example, the Europe region group contains the Northern Europe and Western Europe regions.
	// Possible values are: default, australia, canada, usCentral, usEast, usWest, france, germany, europeUnion,
	// unitedKingdom, japan, asia, india, southAmerica, euap, usGovernment, usGovernmentDOD, unknownFutureValue, norway,
	// switzerland, southKorea, middleEast, mexico, australasia, europe. Use the Prefer: include-unknown-enum-members
	// request header to get the following values in this evolvable enum: norway, switzerland, southKorea, middleEast,
	// mexico, australasia, europe. Read-only.
	RegionGroup *CloudPCRegionGroup `json:"regionGroup,omitempty"`

	// The supported Azure region where the IT admin wants the provisioning policy to create Cloud PCs. The underlying
	// virtual network is created and managed by the Windows 365 service. This can only be entered if the IT admin chooses
	// Microsoft Entra joined as the domain join type. If you enter a regionName, leave the onPremisesConnectionId property
	// empty.
	RegionName nullable.Type[string] `json:"regionName,omitempty"`

	// Specifies the method by which the provisioned Cloud PC joins Microsoft Entra ID. If you choose the hybridAzureADJoin
	// type, only provide a value for the onPremisesConnectionId property and leave regionName as empty. If you choose the
	// azureADJoin type, provide a value for either onPremisesConnectionId or regionName. The possible values are:
	// azureADJoin, hybridAzureADJoin, unknownFutureValue. The type property is deprecated and will stop returning data on
	// January 31, 2024. Going forward, use the domainJoinType property.
	Type *CloudPCDomainJoinType `json:"type,omitempty"`
}

var _ json.Marshaler = CloudPCDomainJoinConfiguration{}

func (s CloudPCDomainJoinConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper CloudPCDomainJoinConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPCDomainJoinConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPCDomainJoinConfiguration: %+v", err)
	}

	delete(decoded, "regionGroup")

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPCDomainJoinConfiguration: %+v", err)
	}

	return encoded, nil
}
