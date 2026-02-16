package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Label string

const (
	Label_Authors              Label = "authors"
	Label_CreatedBy            Label = "createdBy"
	Label_CreatedDateTime      Label = "createdDateTime"
	Label_FileExtension        Label = "fileExtension"
	Label_FileName             Label = "fileName"
	Label_LastModifiedBy       Label = "lastModifiedBy"
	Label_LastModifiedDateTime Label = "lastModifiedDateTime"
	Label_Title                Label = "title"
	Label_Url                  Label = "url"
)

func PossibleValuesForLabel() []string {
	return []string{
		string(Label_Authors),
		string(Label_CreatedBy),
		string(Label_CreatedDateTime),
		string(Label_FileExtension),
		string(Label_FileName),
		string(Label_LastModifiedBy),
		string(Label_LastModifiedDateTime),
		string(Label_Title),
		string(Label_Url),
	}
}

func (s *Label) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseLabel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseLabel(input string) (*Label, error) {
	vals := map[string]Label{
		"authors":              Label_Authors,
		"createdby":            Label_CreatedBy,
		"createddatetime":      Label_CreatedDateTime,
		"fileextension":        Label_FileExtension,
		"filename":             Label_FileName,
		"lastmodifiedby":       Label_LastModifiedBy,
		"lastmodifieddatetime": Label_LastModifiedDateTime,
		"title":                Label_Title,
		"url":                  Label_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := Label(input)
	return &out, nil
}
