package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserPfxPaddingScheme string

const (
	UserPfxPaddingScheme_None       UserPfxPaddingScheme = "none"
	UserPfxPaddingScheme_OaepSha1   UserPfxPaddingScheme = "oaepSha1"
	UserPfxPaddingScheme_OaepSha256 UserPfxPaddingScheme = "oaepSha256"
	UserPfxPaddingScheme_OaepSha384 UserPfxPaddingScheme = "oaepSha384"
	UserPfxPaddingScheme_OaepSha512 UserPfxPaddingScheme = "oaepSha512"
	UserPfxPaddingScheme_Pkcs1      UserPfxPaddingScheme = "pkcs1"
)

func PossibleValuesForUserPfxPaddingScheme() []string {
	return []string{
		string(UserPfxPaddingScheme_None),
		string(UserPfxPaddingScheme_OaepSha1),
		string(UserPfxPaddingScheme_OaepSha256),
		string(UserPfxPaddingScheme_OaepSha384),
		string(UserPfxPaddingScheme_OaepSha512),
		string(UserPfxPaddingScheme_Pkcs1),
	}
}

func (s *UserPfxPaddingScheme) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUserPfxPaddingScheme(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUserPfxPaddingScheme(input string) (*UserPfxPaddingScheme, error) {
	vals := map[string]UserPfxPaddingScheme{
		"none":       UserPfxPaddingScheme_None,
		"oaepsha1":   UserPfxPaddingScheme_OaepSha1,
		"oaepsha256": UserPfxPaddingScheme_OaepSha256,
		"oaepsha384": UserPfxPaddingScheme_OaepSha384,
		"oaepsha512": UserPfxPaddingScheme_OaepSha512,
		"pkcs1":      UserPfxPaddingScheme_Pkcs1,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UserPfxPaddingScheme(input)
	return &out, nil
}
