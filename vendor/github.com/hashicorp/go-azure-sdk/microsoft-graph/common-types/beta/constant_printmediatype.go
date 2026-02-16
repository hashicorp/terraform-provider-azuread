package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrintMediaType string

const (
	PrintMediaType_Continuous      PrintMediaType = "continuous"
	PrintMediaType_ContinuousLong  PrintMediaType = "continuousLong"
	PrintMediaType_ContinuousShort PrintMediaType = "continuousShort"
	PrintMediaType_Envelope        PrintMediaType = "envelope"
	PrintMediaType_EnvelopePlain   PrintMediaType = "envelopePlain"
	PrintMediaType_EnvelopeWindow  PrintMediaType = "envelopeWindow"
	PrintMediaType_Labels          PrintMediaType = "labels"
	PrintMediaType_MultiLayer      PrintMediaType = "multiLayer"
	PrintMediaType_MultiPartForm   PrintMediaType = "multiPartForm"
	PrintMediaType_Screen          PrintMediaType = "screen"
	PrintMediaType_ScreenPaged     PrintMediaType = "screenPaged"
	PrintMediaType_Stationery      PrintMediaType = "stationery"
	PrintMediaType_Transparency    PrintMediaType = "transparency"
)

func PossibleValuesForPrintMediaType() []string {
	return []string{
		string(PrintMediaType_Continuous),
		string(PrintMediaType_ContinuousLong),
		string(PrintMediaType_ContinuousShort),
		string(PrintMediaType_Envelope),
		string(PrintMediaType_EnvelopePlain),
		string(PrintMediaType_EnvelopeWindow),
		string(PrintMediaType_Labels),
		string(PrintMediaType_MultiLayer),
		string(PrintMediaType_MultiPartForm),
		string(PrintMediaType_Screen),
		string(PrintMediaType_ScreenPaged),
		string(PrintMediaType_Stationery),
		string(PrintMediaType_Transparency),
	}
}

func (s *PrintMediaType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePrintMediaType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePrintMediaType(input string) (*PrintMediaType, error) {
	vals := map[string]PrintMediaType{
		"continuous":      PrintMediaType_Continuous,
		"continuouslong":  PrintMediaType_ContinuousLong,
		"continuousshort": PrintMediaType_ContinuousShort,
		"envelope":        PrintMediaType_Envelope,
		"envelopeplain":   PrintMediaType_EnvelopePlain,
		"envelopewindow":  PrintMediaType_EnvelopeWindow,
		"labels":          PrintMediaType_Labels,
		"multilayer":      PrintMediaType_MultiLayer,
		"multipartform":   PrintMediaType_MultiPartForm,
		"screen":          PrintMediaType_Screen,
		"screenpaged":     PrintMediaType_ScreenPaged,
		"stationery":      PrintMediaType_Stationery,
		"transparency":    PrintMediaType_Transparency,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PrintMediaType(input)
	return &out, nil
}
