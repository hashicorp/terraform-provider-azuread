package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RetrievalDataSource string

const (
	RetrievalDataSource_Calendar           RetrievalDataSource = "calendar"
	RetrievalDataSource_ExternalItems      RetrievalDataSource = "externalItems"
	RetrievalDataSource_Mail               RetrievalDataSource = "mail"
	RetrievalDataSource_OneDriveBusiness   RetrievalDataSource = "oneDriveBusiness"
	RetrievalDataSource_People             RetrievalDataSource = "people"
	RetrievalDataSource_SharePoint         RetrievalDataSource = "sharePoint"
	RetrievalDataSource_SharePointEmbedded RetrievalDataSource = "sharePointEmbedded"
	RetrievalDataSource_Teams              RetrievalDataSource = "teams"
)

func PossibleValuesForRetrievalDataSource() []string {
	return []string{
		string(RetrievalDataSource_Calendar),
		string(RetrievalDataSource_ExternalItems),
		string(RetrievalDataSource_Mail),
		string(RetrievalDataSource_OneDriveBusiness),
		string(RetrievalDataSource_People),
		string(RetrievalDataSource_SharePoint),
		string(RetrievalDataSource_SharePointEmbedded),
		string(RetrievalDataSource_Teams),
	}
}

func (s *RetrievalDataSource) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseRetrievalDataSource(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseRetrievalDataSource(input string) (*RetrievalDataSource, error) {
	vals := map[string]RetrievalDataSource{
		"calendar":           RetrievalDataSource_Calendar,
		"externalitems":      RetrievalDataSource_ExternalItems,
		"mail":               RetrievalDataSource_Mail,
		"onedrivebusiness":   RetrievalDataSource_OneDriveBusiness,
		"people":             RetrievalDataSource_People,
		"sharepoint":         RetrievalDataSource_SharePoint,
		"sharepointembedded": RetrievalDataSource_SharePointEmbedded,
		"teams":              RetrievalDataSource_Teams,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := RetrievalDataSource(input)
	return &out, nil
}
