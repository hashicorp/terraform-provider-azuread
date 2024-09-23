package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAdditionalOptions string

const (
	SecurityAdditionalOptions_AllDocumentVersions         SecurityAdditionalOptions = "allDocumentVersions"
	SecurityAdditionalOptions_CloudAttachments            SecurityAdditionalOptions = "cloudAttachments"
	SecurityAdditionalOptions_ListAttachments             SecurityAdditionalOptions = "listAttachments"
	SecurityAdditionalOptions_None                        SecurityAdditionalOptions = "none"
	SecurityAdditionalOptions_SubfolderContents           SecurityAdditionalOptions = "subfolderContents"
	SecurityAdditionalOptions_TeamsAndYammerConversations SecurityAdditionalOptions = "teamsAndYammerConversations"
)

func PossibleValuesForSecurityAdditionalOptions() []string {
	return []string{
		string(SecurityAdditionalOptions_AllDocumentVersions),
		string(SecurityAdditionalOptions_CloudAttachments),
		string(SecurityAdditionalOptions_ListAttachments),
		string(SecurityAdditionalOptions_None),
		string(SecurityAdditionalOptions_SubfolderContents),
		string(SecurityAdditionalOptions_TeamsAndYammerConversations),
	}
}

func (s *SecurityAdditionalOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityAdditionalOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityAdditionalOptions(input string) (*SecurityAdditionalOptions, error) {
	vals := map[string]SecurityAdditionalOptions{
		"alldocumentversions":         SecurityAdditionalOptions_AllDocumentVersions,
		"cloudattachments":            SecurityAdditionalOptions_CloudAttachments,
		"listattachments":             SecurityAdditionalOptions_ListAttachments,
		"none":                        SecurityAdditionalOptions_None,
		"subfoldercontents":           SecurityAdditionalOptions_SubfolderContents,
		"teamsandyammerconversations": SecurityAdditionalOptions_TeamsAndYammerConversations,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityAdditionalOptions(input)
	return &out, nil
}
