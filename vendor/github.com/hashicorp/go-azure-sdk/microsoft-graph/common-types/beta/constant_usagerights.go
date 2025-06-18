package beta

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UsageRights string

const (
	UsageRights_AccessDenied                                   UsageRights = "accessDenied"
	UsageRights_Comment                                        UsageRights = "comment"
	UsageRights_DocEdit                                        UsageRights = "docEdit"
	UsageRights_Edit                                           UsageRights = "edit"
	UsageRights_EditRightsData                                 UsageRights = "editRightsData"
	UsageRights_EncryptedProtectionTypeNotSupportedException   UsageRights = "encryptedProtectionTypeNotSupportedException"
	UsageRights_Exception                                      UsageRights = "exception"
	UsageRights_Export                                         UsageRights = "export"
	UsageRights_Extract                                        UsageRights = "extract"
	UsageRights_Forward                                        UsageRights = "forward"
	UsageRights_ObjModel                                       UsageRights = "objModel"
	UsageRights_Owner                                          UsageRights = "owner"
	UsageRights_Print                                          UsageRights = "print"
	UsageRights_PurviewClaimsChallengeNotSupportedException    UsageRights = "purviewClaimsChallengeNotSupportedException"
	UsageRights_Reply                                          UsageRights = "reply"
	UsageRights_ReplyAll                                       UsageRights = "replyAll"
	UsageRights_Unknown                                        UsageRights = "unknown"
	UsageRights_UserDefinedProtectionTypeNotSupportedException UsageRights = "userDefinedProtectionTypeNotSupportedException"
	UsageRights_View                                           UsageRights = "view"
	UsageRights_ViewRightsData                                 UsageRights = "viewRightsData"
)

func PossibleValuesForUsageRights() []string {
	return []string{
		string(UsageRights_AccessDenied),
		string(UsageRights_Comment),
		string(UsageRights_DocEdit),
		string(UsageRights_Edit),
		string(UsageRights_EditRightsData),
		string(UsageRights_EncryptedProtectionTypeNotSupportedException),
		string(UsageRights_Exception),
		string(UsageRights_Export),
		string(UsageRights_Extract),
		string(UsageRights_Forward),
		string(UsageRights_ObjModel),
		string(UsageRights_Owner),
		string(UsageRights_Print),
		string(UsageRights_PurviewClaimsChallengeNotSupportedException),
		string(UsageRights_Reply),
		string(UsageRights_ReplyAll),
		string(UsageRights_Unknown),
		string(UsageRights_UserDefinedProtectionTypeNotSupportedException),
		string(UsageRights_View),
		string(UsageRights_ViewRightsData),
	}
}

func (s *UsageRights) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parseUsageRights(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parseUsageRights(input string) (*UsageRights, error) {
	vals := map[string]UsageRights{
		"accessdenied":   UsageRights_AccessDenied,
		"comment":        UsageRights_Comment,
		"docedit":        UsageRights_DocEdit,
		"edit":           UsageRights_Edit,
		"editrightsdata": UsageRights_EditRightsData,
		"encryptedprotectiontypenotsupportedexception": UsageRights_EncryptedProtectionTypeNotSupportedException,
		"exception": UsageRights_Exception,
		"export":    UsageRights_Export,
		"extract":   UsageRights_Extract,
		"forward":   UsageRights_Forward,
		"objmodel":  UsageRights_ObjModel,
		"owner":     UsageRights_Owner,
		"print":     UsageRights_Print,
		"purviewclaimschallengenotsupportedexception": UsageRights_PurviewClaimsChallengeNotSupportedException,
		"reply":    UsageRights_Reply,
		"replyall": UsageRights_ReplyAll,
		"unknown":  UsageRights_Unknown,
		"userdefinedprotectiontypenotsupportedexception": UsageRights_UserDefinedProtectionTypeNotSupportedException,
		"view":           UsageRights_View,
		"viewrightsdata": UsageRights_ViewRightsData,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := UsageRights(input)
	return &out, nil
}
