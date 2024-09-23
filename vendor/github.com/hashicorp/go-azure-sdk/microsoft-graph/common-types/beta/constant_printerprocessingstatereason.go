package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrinterProcessingStateReason string

const (
	PrinterProcessingStateReason_ConnectingToDevice                 PrinterProcessingStateReason = "connectingToDevice"
	PrinterProcessingStateReason_CoverOpen                          PrinterProcessingStateReason = "coverOpen"
	PrinterProcessingStateReason_DeveloperEmpty                     PrinterProcessingStateReason = "developerEmpty"
	PrinterProcessingStateReason_DeveloperLow                       PrinterProcessingStateReason = "developerLow"
	PrinterProcessingStateReason_DoorOpen                           PrinterProcessingStateReason = "doorOpen"
	PrinterProcessingStateReason_FuserOverTemp                      PrinterProcessingStateReason = "fuserOverTemp"
	PrinterProcessingStateReason_FuserUnderTemp                     PrinterProcessingStateReason = "fuserUnderTemp"
	PrinterProcessingStateReason_InputTrayMissing                   PrinterProcessingStateReason = "inputTrayMissing"
	PrinterProcessingStateReason_InterlockOpen                      PrinterProcessingStateReason = "interlockOpen"
	PrinterProcessingStateReason_InterpreterResourceUnavailable     PrinterProcessingStateReason = "interpreterResourceUnavailable"
	PrinterProcessingStateReason_MarkerSupplyEmpty                  PrinterProcessingStateReason = "markerSupplyEmpty"
	PrinterProcessingStateReason_MarkerSupplyLow                    PrinterProcessingStateReason = "markerSupplyLow"
	PrinterProcessingStateReason_MarkerWasteAlmostFull              PrinterProcessingStateReason = "markerWasteAlmostFull"
	PrinterProcessingStateReason_MarkerWasteFull                    PrinterProcessingStateReason = "markerWasteFull"
	PrinterProcessingStateReason_MediaEmpty                         PrinterProcessingStateReason = "mediaEmpty"
	PrinterProcessingStateReason_MediaJam                           PrinterProcessingStateReason = "mediaJam"
	PrinterProcessingStateReason_MediaLow                           PrinterProcessingStateReason = "mediaLow"
	PrinterProcessingStateReason_MediaNeeded                        PrinterProcessingStateReason = "mediaNeeded"
	PrinterProcessingStateReason_MovingToPaused                     PrinterProcessingStateReason = "movingToPaused"
	PrinterProcessingStateReason_None                               PrinterProcessingStateReason = "none"
	PrinterProcessingStateReason_OpticalPhotoConductorLifeOver      PrinterProcessingStateReason = "opticalPhotoConductorLifeOver"
	PrinterProcessingStateReason_OpticalPhotoConductorNearEndOfLife PrinterProcessingStateReason = "opticalPhotoConductorNearEndOfLife"
	PrinterProcessingStateReason_Other                              PrinterProcessingStateReason = "other"
	PrinterProcessingStateReason_OutputAreaAlmostFull               PrinterProcessingStateReason = "outputAreaAlmostFull"
	PrinterProcessingStateReason_OutputAreaFull                     PrinterProcessingStateReason = "outputAreaFull"
	PrinterProcessingStateReason_OutputTrayMissing                  PrinterProcessingStateReason = "outputTrayMissing"
	PrinterProcessingStateReason_Paused                             PrinterProcessingStateReason = "paused"
	PrinterProcessingStateReason_Shutdown                           PrinterProcessingStateReason = "shutdown"
	PrinterProcessingStateReason_SpoolAreaFull                      PrinterProcessingStateReason = "spoolAreaFull"
	PrinterProcessingStateReason_StoppedPartially                   PrinterProcessingStateReason = "stoppedPartially"
	PrinterProcessingStateReason_Stopping                           PrinterProcessingStateReason = "stopping"
	PrinterProcessingStateReason_TimedOut                           PrinterProcessingStateReason = "timedOut"
	PrinterProcessingStateReason_TonerEmpty                         PrinterProcessingStateReason = "tonerEmpty"
	PrinterProcessingStateReason_TonerLow                           PrinterProcessingStateReason = "tonerLow"
)

func PossibleValuesForPrinterProcessingStateReason() []string {
	return []string{
		string(PrinterProcessingStateReason_ConnectingToDevice),
		string(PrinterProcessingStateReason_CoverOpen),
		string(PrinterProcessingStateReason_DeveloperEmpty),
		string(PrinterProcessingStateReason_DeveloperLow),
		string(PrinterProcessingStateReason_DoorOpen),
		string(PrinterProcessingStateReason_FuserOverTemp),
		string(PrinterProcessingStateReason_FuserUnderTemp),
		string(PrinterProcessingStateReason_InputTrayMissing),
		string(PrinterProcessingStateReason_InterlockOpen),
		string(PrinterProcessingStateReason_InterpreterResourceUnavailable),
		string(PrinterProcessingStateReason_MarkerSupplyEmpty),
		string(PrinterProcessingStateReason_MarkerSupplyLow),
		string(PrinterProcessingStateReason_MarkerWasteAlmostFull),
		string(PrinterProcessingStateReason_MarkerWasteFull),
		string(PrinterProcessingStateReason_MediaEmpty),
		string(PrinterProcessingStateReason_MediaJam),
		string(PrinterProcessingStateReason_MediaLow),
		string(PrinterProcessingStateReason_MediaNeeded),
		string(PrinterProcessingStateReason_MovingToPaused),
		string(PrinterProcessingStateReason_None),
		string(PrinterProcessingStateReason_OpticalPhotoConductorLifeOver),
		string(PrinterProcessingStateReason_OpticalPhotoConductorNearEndOfLife),
		string(PrinterProcessingStateReason_Other),
		string(PrinterProcessingStateReason_OutputAreaAlmostFull),
		string(PrinterProcessingStateReason_OutputAreaFull),
		string(PrinterProcessingStateReason_OutputTrayMissing),
		string(PrinterProcessingStateReason_Paused),
		string(PrinterProcessingStateReason_Shutdown),
		string(PrinterProcessingStateReason_SpoolAreaFull),
		string(PrinterProcessingStateReason_StoppedPartially),
		string(PrinterProcessingStateReason_Stopping),
		string(PrinterProcessingStateReason_TimedOut),
		string(PrinterProcessingStateReason_TonerEmpty),
		string(PrinterProcessingStateReason_TonerLow),
	}
}

func (s *PrinterProcessingStateReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrinterProcessingStateReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrinterProcessingStateReason(input string) (*PrinterProcessingStateReason, error) {
	vals := map[string]PrinterProcessingStateReason{
		"connectingtodevice":                 PrinterProcessingStateReason_ConnectingToDevice,
		"coveropen":                          PrinterProcessingStateReason_CoverOpen,
		"developerempty":                     PrinterProcessingStateReason_DeveloperEmpty,
		"developerlow":                       PrinterProcessingStateReason_DeveloperLow,
		"dooropen":                           PrinterProcessingStateReason_DoorOpen,
		"fuserovertemp":                      PrinterProcessingStateReason_FuserOverTemp,
		"fuserundertemp":                     PrinterProcessingStateReason_FuserUnderTemp,
		"inputtraymissing":                   PrinterProcessingStateReason_InputTrayMissing,
		"interlockopen":                      PrinterProcessingStateReason_InterlockOpen,
		"interpreterresourceunavailable":     PrinterProcessingStateReason_InterpreterResourceUnavailable,
		"markersupplyempty":                  PrinterProcessingStateReason_MarkerSupplyEmpty,
		"markersupplylow":                    PrinterProcessingStateReason_MarkerSupplyLow,
		"markerwastealmostfull":              PrinterProcessingStateReason_MarkerWasteAlmostFull,
		"markerwastefull":                    PrinterProcessingStateReason_MarkerWasteFull,
		"mediaempty":                         PrinterProcessingStateReason_MediaEmpty,
		"mediajam":                           PrinterProcessingStateReason_MediaJam,
		"medialow":                           PrinterProcessingStateReason_MediaLow,
		"medianeeded":                        PrinterProcessingStateReason_MediaNeeded,
		"movingtopaused":                     PrinterProcessingStateReason_MovingToPaused,
		"none":                               PrinterProcessingStateReason_None,
		"opticalphotoconductorlifeover":      PrinterProcessingStateReason_OpticalPhotoConductorLifeOver,
		"opticalphotoconductornearendoflife": PrinterProcessingStateReason_OpticalPhotoConductorNearEndOfLife,
		"other":                              PrinterProcessingStateReason_Other,
		"outputareaalmostfull":               PrinterProcessingStateReason_OutputAreaAlmostFull,
		"outputareafull":                     PrinterProcessingStateReason_OutputAreaFull,
		"outputtraymissing":                  PrinterProcessingStateReason_OutputTrayMissing,
		"paused":                             PrinterProcessingStateReason_Paused,
		"shutdown":                           PrinterProcessingStateReason_Shutdown,
		"spoolareafull":                      PrinterProcessingStateReason_SpoolAreaFull,
		"stoppedpartially":                   PrinterProcessingStateReason_StoppedPartially,
		"stopping":                           PrinterProcessingStateReason_Stopping,
		"timedout":                           PrinterProcessingStateReason_TimedOut,
		"tonerempty":                         PrinterProcessingStateReason_TonerEmpty,
		"tonerlow":                           PrinterProcessingStateReason_TonerLow,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrinterProcessingStateReason(input)
	return &out, nil
}
