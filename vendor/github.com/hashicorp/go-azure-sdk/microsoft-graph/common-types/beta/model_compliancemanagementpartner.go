package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = ComplianceManagementPartner{}

type ComplianceManagementPartner struct {
	// User groups which enroll Android devices through partner.
	AndroidEnrollmentAssignments *[]ComplianceManagementPartnerAssignment `json:"androidEnrollmentAssignments,omitempty"`

	// Partner onboarded for Android devices.
	AndroidOnboarded *bool `json:"androidOnboarded,omitempty"`

	// Partner display name
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// User groups which enroll ios devices through partner.
	IosEnrollmentAssignments *[]ComplianceManagementPartnerAssignment `json:"iosEnrollmentAssignments,omitempty"`

	// Partner onboarded for ios devices.
	IosOnboarded *bool `json:"iosOnboarded,omitempty"`

	// Timestamp of last heartbeat after admin onboarded to the compliance management partner
	LastHeartbeatDateTime *string `json:"lastHeartbeatDateTime,omitempty"`

	// User groups which enroll Linux devices through partner.
	LinuxEnrollmentAssignments *[]ComplianceManagementPartnerAssignment `json:"linuxEnrollmentAssignments,omitempty"`

	// Partner onboarded for Linux devices.
	LinuxOnboarded *bool `json:"linuxOnboarded,omitempty"`

	// User groups which enroll Mac devices through partner.
	MacOsEnrollmentAssignments *[]ComplianceManagementPartnerAssignment `json:"macOsEnrollmentAssignments,omitempty"`

	// Partner onboarded for Mac devices.
	MacOsOnboarded *bool `json:"macOsOnboarded,omitempty"`

	// Partner state of this tenant.
	PartnerState *DeviceManagementPartnerTenantState `json:"partnerState,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s ComplianceManagementPartner) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = ComplianceManagementPartner{}

func (s ComplianceManagementPartner) MarshalJSON() ([]byte, error) {
	type wrapper ComplianceManagementPartner
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ComplianceManagementPartner: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ComplianceManagementPartner: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.complianceManagementPartner"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ComplianceManagementPartner: %+v", err)
	}

	return encoded, nil
}
