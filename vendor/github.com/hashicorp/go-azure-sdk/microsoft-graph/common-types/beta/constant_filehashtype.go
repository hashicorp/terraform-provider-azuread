package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type FileHashType string

const (
	FileHashType_AuthenticodeHash256 FileHashType = "authenticodeHash256"
	FileHashType_Ctph                FileHashType = "ctph"
	FileHashType_LsHash              FileHashType = "lsHash"
	FileHashType_Md5                 FileHashType = "md5"
	FileHashType_Sha1                FileHashType = "sha1"
	FileHashType_Sha256              FileHashType = "sha256"
	FileHashType_Unknown             FileHashType = "unknown"
)

func PossibleValuesForFileHashType() []string {
	return []string{
		string(FileHashType_AuthenticodeHash256),
		string(FileHashType_Ctph),
		string(FileHashType_LsHash),
		string(FileHashType_Md5),
		string(FileHashType_Sha1),
		string(FileHashType_Sha256),
		string(FileHashType_Unknown),
	}
}

func (s *FileHashType) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseFileHashType(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseFileHashType(input string) (*FileHashType, error) {
	vals := map[string]FileHashType{
		"authenticodehash256": FileHashType_AuthenticodeHash256,
		"ctph":                FileHashType_Ctph,
		"lshash":              FileHashType_LsHash,
		"md5":                 FileHashType_Md5,
		"sha1":                FileHashType_Sha1,
		"sha256":              FileHashType_Sha256,
		"unknown":             FileHashType_Unknown,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := FileHashType(input)
	return &out, nil
}
