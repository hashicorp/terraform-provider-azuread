package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyOperationType string

const (
	GroupPolicyOperationType_AddLanguageFiles    GroupPolicyOperationType = "addLanguageFiles"
	GroupPolicyOperationType_None                GroupPolicyOperationType = "none"
	GroupPolicyOperationType_Remove              GroupPolicyOperationType = "remove"
	GroupPolicyOperationType_RemoveLanguageFiles GroupPolicyOperationType = "removeLanguageFiles"
	GroupPolicyOperationType_UpdateLanguageFiles GroupPolicyOperationType = "updateLanguageFiles"
	GroupPolicyOperationType_Upload              GroupPolicyOperationType = "upload"
	GroupPolicyOperationType_UploadNewVersion    GroupPolicyOperationType = "uploadNewVersion"
)

func PossibleValuesForGroupPolicyOperationType() []string {
	return []string{
		string(GroupPolicyOperationType_AddLanguageFiles),
		string(GroupPolicyOperationType_None),
		string(GroupPolicyOperationType_Remove),
		string(GroupPolicyOperationType_RemoveLanguageFiles),
		string(GroupPolicyOperationType_UpdateLanguageFiles),
		string(GroupPolicyOperationType_Upload),
		string(GroupPolicyOperationType_UploadNewVersion),
	}
}

func (s *GroupPolicyOperationType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseGroupPolicyOperationType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseGroupPolicyOperationType(input string) (*GroupPolicyOperationType, error) {
	vals := map[string]GroupPolicyOperationType{
		"addlanguagefiles":    GroupPolicyOperationType_AddLanguageFiles,
		"none":                GroupPolicyOperationType_None,
		"remove":              GroupPolicyOperationType_Remove,
		"removelanguagefiles": GroupPolicyOperationType_RemoveLanguageFiles,
		"updatelanguagefiles": GroupPolicyOperationType_UpdateLanguageFiles,
		"upload":              GroupPolicyOperationType_Upload,
		"uploadnewversion":    GroupPolicyOperationType_UploadNewVersion,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := GroupPolicyOperationType(input)
	return &out, nil
}
