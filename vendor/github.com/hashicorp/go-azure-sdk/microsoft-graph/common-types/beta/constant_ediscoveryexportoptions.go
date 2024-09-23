package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EdiscoveryExportOptions string

const (
	EdiscoveryExportOptions_FileInfo       EdiscoveryExportOptions = "fileInfo"
	EdiscoveryExportOptions_OriginalFiles  EdiscoveryExportOptions = "originalFiles"
	EdiscoveryExportOptions_PdfReplacement EdiscoveryExportOptions = "pdfReplacement"
	EdiscoveryExportOptions_Tags           EdiscoveryExportOptions = "tags"
	EdiscoveryExportOptions_Text           EdiscoveryExportOptions = "text"
)

func PossibleValuesForEdiscoveryExportOptions() []string {
	return []string{
		string(EdiscoveryExportOptions_FileInfo),
		string(EdiscoveryExportOptions_OriginalFiles),
		string(EdiscoveryExportOptions_PdfReplacement),
		string(EdiscoveryExportOptions_Tags),
		string(EdiscoveryExportOptions_Text),
	}
}

func (s *EdiscoveryExportOptions) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseEdiscoveryExportOptions(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseEdiscoveryExportOptions(input string) (*EdiscoveryExportOptions, error) {
	vals := map[string]EdiscoveryExportOptions{
		"fileinfo":       EdiscoveryExportOptions_FileInfo,
		"originalfiles":  EdiscoveryExportOptions_OriginalFiles,
		"pdfreplacement": EdiscoveryExportOptions_PdfReplacement,
		"tags":           EdiscoveryExportOptions_Tags,
		"text":           EdiscoveryExportOptions_Text,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := EdiscoveryExportOptions(input)
	return &out, nil
}
