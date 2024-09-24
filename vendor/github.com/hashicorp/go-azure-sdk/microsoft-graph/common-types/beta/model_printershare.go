package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PrinterBase = PrinterShare{}

type PrinterShare struct {
	// If true, all users and groups will be granted access to this printer share. This supersedes the allow lists defined
	// by the allowedUsers and allowedGroups navigation properties.
	AllowAllUsers *bool `json:"allowAllUsers,omitempty"`

	// The groups whose users have access to print using the printer.
	AllowedGroups *[]Group `json:"allowedGroups,omitempty"`

	// The users who have access to print using the printer.
	AllowedUsers *[]User `json:"allowedUsers,omitempty"`

	// The DateTimeOffset when the printer share was created. Read-only.
	CreatedDateTime *string `json:"createdDateTime,omitempty"`

	// The printer that this printer share is related to.
	Printer *Printer `json:"printer,omitempty"`

	// Additional data for a printer share as viewed by the signed-in user.
	ViewPoint *PrinterShareViewpoint `json:"viewPoint,omitempty"`

	// Fields inherited from PrinterBase

	Capabilities    *PrinterCapabilities  `json:"capabilities,omitempty"`
	Defaults        *PrinterDefaults      `json:"defaults,omitempty"`
	DisplayName     *string               `json:"displayName,omitempty"`
	IsAcceptingJobs nullable.Type[bool]   `json:"isAcceptingJobs,omitempty"`
	Jobs            *[]PrintJob           `json:"jobs,omitempty"`
	Location        *PrinterLocation      `json:"location,omitempty"`
	Manufacturer    nullable.Type[string] `json:"manufacturer,omitempty"`
	Model           nullable.Type[string] `json:"model,omitempty"`
	Name            nullable.Type[string] `json:"name,omitempty"`
	Status          *PrinterStatus        `json:"status,omitempty"`

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

func (s PrinterShare) PrinterBase() BasePrinterBaseImpl {
	return BasePrinterBaseImpl{
		Capabilities:    s.Capabilities,
		Defaults:        s.Defaults,
		DisplayName:     s.DisplayName,
		IsAcceptingJobs: s.IsAcceptingJobs,
		Jobs:            s.Jobs,
		Location:        s.Location,
		Manufacturer:    s.Manufacturer,
		Model:           s.Model,
		Name:            s.Name,
		Status:          s.Status,
		Id:              s.Id,
		ODataId:         s.ODataId,
		ODataType:       s.ODataType,
	}
}

func (s PrinterShare) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = PrinterShare{}

func (s PrinterShare) MarshalJSON() ([]byte, error) {
	type wrapper PrinterShare
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrinterShare: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrinterShare: %+v", err)
	}

	delete(decoded, "createdDateTime")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.printerShare"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrinterShare: %+v", err)
	}

	return encoded, nil
}
