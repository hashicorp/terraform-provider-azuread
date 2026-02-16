package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ExternalConnectorsLabel string

const (
	ExternalConnectorsLabel_Authors              ExternalConnectorsLabel = "authors"
	ExternalConnectorsLabel_ContainerName        ExternalConnectorsLabel = "containerName"
	ExternalConnectorsLabel_ContainerUrl         ExternalConnectorsLabel = "containerUrl"
	ExternalConnectorsLabel_CreatedBy            ExternalConnectorsLabel = "createdBy"
	ExternalConnectorsLabel_CreatedDateTime      ExternalConnectorsLabel = "createdDateTime"
	ExternalConnectorsLabel_FileExtension        ExternalConnectorsLabel = "fileExtension"
	ExternalConnectorsLabel_FileName             ExternalConnectorsLabel = "fileName"
	ExternalConnectorsLabel_IconUrl              ExternalConnectorsLabel = "iconUrl"
	ExternalConnectorsLabel_LastModifiedBy       ExternalConnectorsLabel = "lastModifiedBy"
	ExternalConnectorsLabel_LastModifiedDateTime ExternalConnectorsLabel = "lastModifiedDateTime"
	ExternalConnectorsLabel_Title                ExternalConnectorsLabel = "title"
	ExternalConnectorsLabel_Url                  ExternalConnectorsLabel = "url"
)

func PossibleValuesForExternalConnectorsLabel() []string {
	return []string{
		string(ExternalConnectorsLabel_Authors),
		string(ExternalConnectorsLabel_ContainerName),
		string(ExternalConnectorsLabel_ContainerUrl),
		string(ExternalConnectorsLabel_CreatedBy),
		string(ExternalConnectorsLabel_CreatedDateTime),
		string(ExternalConnectorsLabel_FileExtension),
		string(ExternalConnectorsLabel_FileName),
		string(ExternalConnectorsLabel_IconUrl),
		string(ExternalConnectorsLabel_LastModifiedBy),
		string(ExternalConnectorsLabel_LastModifiedDateTime),
		string(ExternalConnectorsLabel_Title),
		string(ExternalConnectorsLabel_Url),
	}
}

func (s *ExternalConnectorsLabel) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseExternalConnectorsLabel(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseExternalConnectorsLabel(input string) (*ExternalConnectorsLabel, error) {
	vals := map[string]ExternalConnectorsLabel{
		"authors":              ExternalConnectorsLabel_Authors,
		"containername":        ExternalConnectorsLabel_ContainerName,
		"containerurl":         ExternalConnectorsLabel_ContainerUrl,
		"createdby":            ExternalConnectorsLabel_CreatedBy,
		"createddatetime":      ExternalConnectorsLabel_CreatedDateTime,
		"fileextension":        ExternalConnectorsLabel_FileExtension,
		"filename":             ExternalConnectorsLabel_FileName,
		"iconurl":              ExternalConnectorsLabel_IconUrl,
		"lastmodifiedby":       ExternalConnectorsLabel_LastModifiedBy,
		"lastmodifieddatetime": ExternalConnectorsLabel_LastModifiedDateTime,
		"title":                ExternalConnectorsLabel_Title,
		"url":                  ExternalConnectorsLabel_Url,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := ExternalConnectorsLabel(input)
	return &out, nil
}
