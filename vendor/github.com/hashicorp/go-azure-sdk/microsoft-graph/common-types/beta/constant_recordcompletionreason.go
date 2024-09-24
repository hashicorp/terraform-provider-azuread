package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RecordCompletionReason string

const (
	RecordCompletionReason_InitialSilenceTimeout    RecordCompletionReason = "initialSilenceTimeout"
	RecordCompletionReason_MaxRecordDurationReached RecordCompletionReason = "maxRecordDurationReached"
	RecordCompletionReason_MaxSilenceTimeout        RecordCompletionReason = "maxSilenceTimeout"
	RecordCompletionReason_MediaReceiveTimeout      RecordCompletionReason = "mediaReceiveTimeout"
	RecordCompletionReason_OperationCanceled        RecordCompletionReason = "operationCanceled"
	RecordCompletionReason_PlayBeepFailed           RecordCompletionReason = "playBeepFailed"
	RecordCompletionReason_PlayPromptFailed         RecordCompletionReason = "playPromptFailed"
	RecordCompletionReason_StopToneDetected         RecordCompletionReason = "stopToneDetected"
	RecordCompletionReason_UnspecifiedError         RecordCompletionReason = "unspecifiedError"
)

func PossibleValuesForRecordCompletionReason() []string {
	return []string{
		string(RecordCompletionReason_InitialSilenceTimeout),
		string(RecordCompletionReason_MaxRecordDurationReached),
		string(RecordCompletionReason_MaxSilenceTimeout),
		string(RecordCompletionReason_MediaReceiveTimeout),
		string(RecordCompletionReason_OperationCanceled),
		string(RecordCompletionReason_PlayBeepFailed),
		string(RecordCompletionReason_PlayPromptFailed),
		string(RecordCompletionReason_StopToneDetected),
		string(RecordCompletionReason_UnspecifiedError),
	}
}

func (s *RecordCompletionReason) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRecordCompletionReason(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRecordCompletionReason(input string) (*RecordCompletionReason, error) {
	vals := map[string]RecordCompletionReason{
		"initialsilencetimeout":    RecordCompletionReason_InitialSilenceTimeout,
		"maxrecorddurationreached": RecordCompletionReason_MaxRecordDurationReached,
		"maxsilencetimeout":        RecordCompletionReason_MaxSilenceTimeout,
		"mediareceivetimeout":      RecordCompletionReason_MediaReceiveTimeout,
		"operationcanceled":        RecordCompletionReason_OperationCanceled,
		"playbeepfailed":           RecordCompletionReason_PlayBeepFailed,
		"playpromptfailed":         RecordCompletionReason_PlayPromptFailed,
		"stoptonedetected":         RecordCompletionReason_StopToneDetected,
		"unspecifiederror":         RecordCompletionReason_UnspecifiedError,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RecordCompletionReason(input)
	return &out, nil
}
