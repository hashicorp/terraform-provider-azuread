package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAdditionalOptions string

const (
	SecurityAdditionalOptions_AdvancedIndexing            SecurityAdditionalOptions = "advancedIndexing"
	SecurityAdditionalOptions_AllDocumentVersions         SecurityAdditionalOptions = "allDocumentVersions"
	SecurityAdditionalOptions_AllItemsInFolder            SecurityAdditionalOptions = "allItemsInFolder"
	SecurityAdditionalOptions_CloudAttachments            SecurityAdditionalOptions = "cloudAttachments"
	SecurityAdditionalOptions_CondensePaths               SecurityAdditionalOptions = "condensePaths"
	SecurityAdditionalOptions_FriendlyName                SecurityAdditionalOptions = "friendlyName"
	SecurityAdditionalOptions_HtmlTranscripts             SecurityAdditionalOptions = "htmlTranscripts"
	SecurityAdditionalOptions_IncludeFolderAndPath        SecurityAdditionalOptions = "includeFolderAndPath"
	SecurityAdditionalOptions_IncludeReport               SecurityAdditionalOptions = "includeReport"
	SecurityAdditionalOptions_ListAttachments             SecurityAdditionalOptions = "listAttachments"
	SecurityAdditionalOptions_None                        SecurityAdditionalOptions = "none"
	SecurityAdditionalOptions_SplitSource                 SecurityAdditionalOptions = "splitSource"
	SecurityAdditionalOptions_SubfolderContents           SecurityAdditionalOptions = "subfolderContents"
	SecurityAdditionalOptions_TeamsAndYammerConversations SecurityAdditionalOptions = "teamsAndYammerConversations"
)

func PossibleValuesForSecurityAdditionalOptions() []string {
	return []string{
		string(SecurityAdditionalOptions_AdvancedIndexing),
		string(SecurityAdditionalOptions_AllDocumentVersions),
		string(SecurityAdditionalOptions_AllItemsInFolder),
		string(SecurityAdditionalOptions_CloudAttachments),
		string(SecurityAdditionalOptions_CondensePaths),
		string(SecurityAdditionalOptions_FriendlyName),
		string(SecurityAdditionalOptions_HtmlTranscripts),
		string(SecurityAdditionalOptions_IncludeFolderAndPath),
		string(SecurityAdditionalOptions_IncludeReport),
		string(SecurityAdditionalOptions_ListAttachments),
		string(SecurityAdditionalOptions_None),
		string(SecurityAdditionalOptions_SplitSource),
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
		"advancedindexing":            SecurityAdditionalOptions_AdvancedIndexing,
		"alldocumentversions":         SecurityAdditionalOptions_AllDocumentVersions,
		"allitemsinfolder":            SecurityAdditionalOptions_AllItemsInFolder,
		"cloudattachments":            SecurityAdditionalOptions_CloudAttachments,
		"condensepaths":               SecurityAdditionalOptions_CondensePaths,
		"friendlyname":                SecurityAdditionalOptions_FriendlyName,
		"htmltranscripts":             SecurityAdditionalOptions_HtmlTranscripts,
		"includefolderandpath":        SecurityAdditionalOptions_IncludeFolderAndPath,
		"includereport":               SecurityAdditionalOptions_IncludeReport,
		"listattachments":             SecurityAdditionalOptions_ListAttachments,
		"none":                        SecurityAdditionalOptions_None,
		"splitsource":                 SecurityAdditionalOptions_SplitSource,
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
