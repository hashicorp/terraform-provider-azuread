package stable

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PayloadTheme string

const (
	PayloadTheme_AccountActivation     PayloadTheme = "accountActivation"
	PayloadTheme_AccountVerification   PayloadTheme = "accountVerification"
	PayloadTheme_Advertisement         PayloadTheme = "advertisement"
	PayloadTheme_Billing               PayloadTheme = "billing"
	PayloadTheme_CleanUpMail           PayloadTheme = "cleanUpMail"
	PayloadTheme_Controversial         PayloadTheme = "controversial"
	PayloadTheme_DocumentReceived      PayloadTheme = "documentReceived"
	PayloadTheme_EmployeeEngagement    PayloadTheme = "employeeEngagement"
	PayloadTheme_Expense               PayloadTheme = "expense"
	PayloadTheme_Fax                   PayloadTheme = "fax"
	PayloadTheme_FinanceReport         PayloadTheme = "financeReport"
	PayloadTheme_IncomingMessages      PayloadTheme = "incomingMessages"
	PayloadTheme_Invoice               PayloadTheme = "invoice"
	PayloadTheme_ItemReceived          PayloadTheme = "itemReceived"
	PayloadTheme_LoginAlert            PayloadTheme = "loginAlert"
	PayloadTheme_MailReceived          PayloadTheme = "mailReceived"
	PayloadTheme_Other                 PayloadTheme = "other"
	PayloadTheme_Password              PayloadTheme = "password"
	PayloadTheme_Payment               PayloadTheme = "payment"
	PayloadTheme_Payroll               PayloadTheme = "payroll"
	PayloadTheme_PersonalizedOffer     PayloadTheme = "personalizedOffer"
	PayloadTheme_Quarantine            PayloadTheme = "quarantine"
	PayloadTheme_RemoteWork            PayloadTheme = "remoteWork"
	PayloadTheme_ReviewMessage         PayloadTheme = "reviewMessage"
	PayloadTheme_SecurityUpdate        PayloadTheme = "securityUpdate"
	PayloadTheme_ServiceSuspended      PayloadTheme = "serviceSuspended"
	PayloadTheme_SignatureRequired     PayloadTheme = "signatureRequired"
	PayloadTheme_Unknown               PayloadTheme = "unknown"
	PayloadTheme_UpgradeMailboxStorage PayloadTheme = "upgradeMailboxStorage"
	PayloadTheme_VerifyMailbox         PayloadTheme = "verifyMailbox"
	PayloadTheme_Voicemail             PayloadTheme = "voicemail"
)

func PossibleValuesForPayloadTheme() []string {
	return []string{
		string(PayloadTheme_AccountActivation),
		string(PayloadTheme_AccountVerification),
		string(PayloadTheme_Advertisement),
		string(PayloadTheme_Billing),
		string(PayloadTheme_CleanUpMail),
		string(PayloadTheme_Controversial),
		string(PayloadTheme_DocumentReceived),
		string(PayloadTheme_EmployeeEngagement),
		string(PayloadTheme_Expense),
		string(PayloadTheme_Fax),
		string(PayloadTheme_FinanceReport),
		string(PayloadTheme_IncomingMessages),
		string(PayloadTheme_Invoice),
		string(PayloadTheme_ItemReceived),
		string(PayloadTheme_LoginAlert),
		string(PayloadTheme_MailReceived),
		string(PayloadTheme_Other),
		string(PayloadTheme_Password),
		string(PayloadTheme_Payment),
		string(PayloadTheme_Payroll),
		string(PayloadTheme_PersonalizedOffer),
		string(PayloadTheme_Quarantine),
		string(PayloadTheme_RemoteWork),
		string(PayloadTheme_ReviewMessage),
		string(PayloadTheme_SecurityUpdate),
		string(PayloadTheme_ServiceSuspended),
		string(PayloadTheme_SignatureRequired),
		string(PayloadTheme_Unknown),
		string(PayloadTheme_UpgradeMailboxStorage),
		string(PayloadTheme_VerifyMailbox),
		string(PayloadTheme_Voicemail),
	}
}

func (s *PayloadTheme) UnmarshalJSON(bytes []byte) error {
	var decoded string
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}
	out, err := parsePayloadTheme(decoded)
	if err != nil {
		return fmt.Errorf("parsing %q: %+v", decoded, err)
	}
	*s = *out
	return nil
}

func parsePayloadTheme(input string) (*PayloadTheme, error) {
	vals := map[string]PayloadTheme{
		"accountactivation":     PayloadTheme_AccountActivation,
		"accountverification":   PayloadTheme_AccountVerification,
		"advertisement":         PayloadTheme_Advertisement,
		"billing":               PayloadTheme_Billing,
		"cleanupmail":           PayloadTheme_CleanUpMail,
		"controversial":         PayloadTheme_Controversial,
		"documentreceived":      PayloadTheme_DocumentReceived,
		"employeeengagement":    PayloadTheme_EmployeeEngagement,
		"expense":               PayloadTheme_Expense,
		"fax":                   PayloadTheme_Fax,
		"financereport":         PayloadTheme_FinanceReport,
		"incomingmessages":      PayloadTheme_IncomingMessages,
		"invoice":               PayloadTheme_Invoice,
		"itemreceived":          PayloadTheme_ItemReceived,
		"loginalert":            PayloadTheme_LoginAlert,
		"mailreceived":          PayloadTheme_MailReceived,
		"other":                 PayloadTheme_Other,
		"password":              PayloadTheme_Password,
		"payment":               PayloadTheme_Payment,
		"payroll":               PayloadTheme_Payroll,
		"personalizedoffer":     PayloadTheme_PersonalizedOffer,
		"quarantine":            PayloadTheme_Quarantine,
		"remotework":            PayloadTheme_RemoteWork,
		"reviewmessage":         PayloadTheme_ReviewMessage,
		"securityupdate":        PayloadTheme_SecurityUpdate,
		"servicesuspended":      PayloadTheme_ServiceSuspended,
		"signaturerequired":     PayloadTheme_SignatureRequired,
		"unknown":               PayloadTheme_Unknown,
		"upgrademailboxstorage": PayloadTheme_UpgradeMailboxStorage,
		"verifymailbox":         PayloadTheme_VerifyMailbox,
		"voicemail":             PayloadTheme_Voicemail,
	}
	if v, ok := vals[strings.ToLower(input)]; ok {
		return &v, nil
	}

	// otherwise presume it's an undefined value and best-effort it
	out := PayloadTheme(input)
	return &out, nil
}
