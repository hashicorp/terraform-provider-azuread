package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityExportOptions string

const (
	SecurityExportOptions_CondensePaths          SecurityExportOptions = "condensePaths"
	SecurityExportOptions_FileInfo               SecurityExportOptions = "fileInfo"
	SecurityExportOptions_FriendlyName           SecurityExportOptions = "friendlyName"
	SecurityExportOptions_IncludeFolderAndPath   SecurityExportOptions = "includeFolderAndPath"
	SecurityExportOptions_OptimizedPartitionSize SecurityExportOptions = "optimizedPartitionSize"
	SecurityExportOptions_OriginalFiles          SecurityExportOptions = "originalFiles"
	SecurityExportOptions_PdfReplacement         SecurityExportOptions = "pdfReplacement"
	SecurityExportOptions_SplitSource            SecurityExportOptions = "splitSource"
	SecurityExportOptions_Tags                   SecurityExportOptions = "tags"
	SecurityExportOptions_Text                   SecurityExportOptions = "text"
)

func PossibleValuesForSecurityExportOptions() []string {
	return []string{
		string(SecurityExportOptions_CondensePaths),
		string(SecurityExportOptions_FileInfo),
		string(SecurityExportOptions_FriendlyName),
		string(SecurityExportOptions_IncludeFolderAndPath),
		string(SecurityExportOptions_OptimizedPartitionSize),
		string(SecurityExportOptions_OriginalFiles),
		string(SecurityExportOptions_PdfReplacement),
		string(SecurityExportOptions_SplitSource),
		string(SecurityExportOptions_Tags),
		string(SecurityExportOptions_Text),
	}
}

func (s *SecurityExportOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseSecurityExportOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseSecurityExportOptions(input string) (*SecurityExportOptions, error) {
	vals := map[string]SecurityExportOptions{
		"condensepaths":          SecurityExportOptions_CondensePaths,
		"fileinfo":               SecurityExportOptions_FileInfo,
		"friendlyname":           SecurityExportOptions_FriendlyName,
		"includefolderandpath":   SecurityExportOptions_IncludeFolderAndPath,
		"optimizedpartitionsize": SecurityExportOptions_OptimizedPartitionSize,
		"originalfiles":          SecurityExportOptions_OriginalFiles,
		"pdfreplacement":         SecurityExportOptions_PdfReplacement,
		"splitsource":            SecurityExportOptions_SplitSource,
		"tags":                   SecurityExportOptions_Tags,
		"text":                   SecurityExportOptions_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := SecurityExportOptions(input)
	return &out, nil
}
