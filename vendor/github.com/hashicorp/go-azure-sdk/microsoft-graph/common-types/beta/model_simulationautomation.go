package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = SimulationAutomation{}

type SimulationAutomation struct {
	// Identity of the user who created the attack simulation automation.
	CreatedBy *EmailIdentity `json:"createdBy,omitempty"`

	// Date and time when the attack simulation automation was created.
	CreatedDateTime nullable.Type[string] `json:"createdDateTime,omitempty"`

	// Description of the attack simulation automation.
	Description nullable.Type[string] `json:"description,omitempty"`

	// Display name of the attack simulation automation. Supports $filter and $orderby.
	DisplayName nullable.Type[string] `json:"displayName,omitempty"`

	// Identity of the user who most recently modified the attack simulation automation.
	LastModifiedBy *EmailIdentity `json:"lastModifiedBy,omitempty"`

	// Date and time when the attack simulation automation was most recently modified.
	LastModifiedDateTime nullable.Type[string] `json:"lastModifiedDateTime,omitempty"`

	// Date and time of the latest run of the attack simulation automation.
	LastRunDateTime nullable.Type[string] `json:"lastRunDateTime,omitempty"`

	// Date and time of the upcoming run of the attack simulation automation.
	NextRunDateTime nullable.Type[string] `json:"nextRunDateTime,omitempty"`

	// A collection of simulation automation runs.
	Runs *[]SimulationAutomationRun `json:"runs,omitempty"`

	// Status of the attack simulation automation. Supports $filter and $orderby. The possible values are: unknown, draft,
	// notRunning, running, completed, unknownFutureValue.
	Status *SimulationAutomationStatus `json:"status,omitempty"`

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

func (s SimulationAutomation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = SimulationAutomation{}

func (s SimulationAutomation) MarshalJSON() ([]byte, error) {
	type wrapper SimulationAutomation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SimulationAutomation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SimulationAutomation: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.simulationAutomation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SimulationAutomation: %+v", err)
	}

	return encoded, nil
}
