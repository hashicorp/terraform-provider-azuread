package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DefenderPromptForSampleSubmission string

const (
	DefenderPromptForSampleSubmission_AlwaysPrompt                    DefenderPromptForSampleSubmission = "alwaysPrompt"
	DefenderPromptForSampleSubmission_NeverSendData                   DefenderPromptForSampleSubmission = "neverSendData"
	DefenderPromptForSampleSubmission_PromptBeforeSendingPersonalData DefenderPromptForSampleSubmission = "promptBeforeSendingPersonalData"
	DefenderPromptForSampleSubmission_SendAllDataWithoutPrompting     DefenderPromptForSampleSubmission = "sendAllDataWithoutPrompting"
	DefenderPromptForSampleSubmission_UserDefined                     DefenderPromptForSampleSubmission = "userDefined"
)

func PossibleValuesForDefenderPromptForSampleSubmission() []string {
	return []string{
		string(DefenderPromptForSampleSubmission_AlwaysPrompt),
		string(DefenderPromptForSampleSubmission_NeverSendData),
		string(DefenderPromptForSampleSubmission_PromptBeforeSendingPersonalData),
		string(DefenderPromptForSampleSubmission_SendAllDataWithoutPrompting),
		string(DefenderPromptForSampleSubmission_UserDefined),
	}
}

func (s *DefenderPromptForSampleSubmission) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseDefenderPromptForSampleSubmission(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseDefenderPromptForSampleSubmission(input string) (*DefenderPromptForSampleSubmission, error) {
	vals := map[string]DefenderPromptForSampleSubmission{
		"alwaysprompt":                    DefenderPromptForSampleSubmission_AlwaysPrompt,
		"neversenddata":                   DefenderPromptForSampleSubmission_NeverSendData,
		"promptbeforesendingpersonaldata": DefenderPromptForSampleSubmission_PromptBeforeSendingPersonalData,
		"sendalldatawithoutprompting":     DefenderPromptForSampleSubmission_SendAllDataWithoutPrompting,
		"userdefined":                     DefenderPromptForSampleSubmission_UserDefined,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := DefenderPromptForSampleSubmission(input)
	return &out, nil
}
