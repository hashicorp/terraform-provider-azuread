package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PrinterBase = Printer{}

type Printer struct {
	AcceptingJobs nullable.Type[bool] `json:"acceptingJobs,omitempty"`

	// The connectors that are associated with the printer.
	Connectors *[]PrintConnector `json:"connectors,omitempty"`

	// True if the printer has a physical device for printing. Read-only.
	HasPhysicalDevice *bool `json:"hasPhysicalDevice,omitempty"`

	// True if the printer is shared; false otherwise. Read-only.
	IsShared *bool `json:"isShared,omitempty"`

	// The most recent dateTimeOffset when a printer interacted with Universal Print. Read-only.
	LastSeenDateTime nullable.Type[string] `json:"lastSeenDateTime,omitempty"`

	// The DateTimeOffset when the printer was registered. Read-only.
	RegisteredDateTime *string `json:"registeredDateTime,omitempty"`

	Share *PrinterShare `json:"share,omitempty"`

	// The list of printerShares that are associated with the printer. Currently, only one printerShare can be associated
	// with the printer. Read-only. Nullable.
	Shares *[]PrinterShare `json:"shares,omitempty"`

	// A list of task triggers that are associated with the printer.
	TaskTriggers *[]PrintTaskTrigger `json:"taskTriggers,omitempty"`

	// Fields inherited from PrinterBase

	// The capabilities of the printer/printerShare.
	Capabilities *PrinterCapabilities `json:"capabilities,omitempty"`

	// The default print settings of printer/printerShare.
	Defaults *PrinterDefaults `json:"defaults,omitempty"`

	// The name of the printer/printerShare.
	DisplayName *string `json:"displayName,omitempty"`

	// Specifies whether the printer/printerShare is currently accepting new print jobs.
	IsAcceptingJobs nullable.Type[bool] `json:"isAcceptingJobs,omitempty"`

	// The list of jobs that are queued for printing by the printer/printerShare.
	Jobs *[]PrintJob `json:"jobs,omitempty"`

	// The physical and/or organizational location of the printer/printerShare.
	Location *PrinterLocation `json:"location,omitempty"`

	// The manufacturer of the printer/printerShare.
	Manufacturer nullable.Type[string] `json:"manufacturer,omitempty"`

	// The model name of the printer/printerShare.
	Model nullable.Type[string] `json:"model,omitempty"`

	Name   nullable.Type[string] `json:"name,omitempty"`
	Status *PrinterStatus        `json:"status,omitempty"`

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

func (s Printer) PrinterBase() BasePrinterBaseImpl {
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

func (s Printer) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Printer{}

func (s Printer) MarshalJSON() ([]byte, error) {
	type wrapper Printer
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Printer: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Printer: %+v", err)
	}

	delete(decoded, "hasPhysicalDevice")
	delete(decoded, "isShared")
	delete(decoded, "lastSeenDateTime")
	delete(decoded, "registeredDateTime")
	delete(decoded, "shares")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.printer"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Printer: %+v", err)
	}

	return encoded, nil
}
