package beta

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = TiIndicator{}

type TiIndicator struct {
	// The action to apply if the indicator is matched from within the targetProduct security tool. Possible values are:
	// unknown, allow, block, alert. Required.
	Action TiAction `json:"action"`

	// The cyber threat intelligence name(s) for the parties responsible for the malicious activity covered by the threat
	// indicator.
	ActivityGroupNames *[]string `json:"activityGroupNames,omitempty"`

	// A catchall area for extra data from the indicator that is not specifically covered by other tiIndicator properties.
	// The security tool specified by targetProduct typically does not utilize this data.
	AdditionalInformation nullable.Type[string] `json:"additionalInformation,omitempty"`

	// Stamped by the system when the indicator is ingested. The Microsoft Entra tenant id of submitting client. Required.
	AzureTenantId nullable.Type[string] `json:"azureTenantId,omitempty"`

	// An integer representing the confidence the data within the indicator accurately identifies malicious behavior.
	// Acceptable values are 0 – 100 with 100 being the highest.
	Confidence nullable.Type[int64] `json:"confidence,omitempty"`

	// Brief description (100 characters or less) of the threat represented by the indicator. Required.
	Description nullable.Type[string] `json:"description,omitempty"`

	// The area of the Diamond Model in which this indicator exists. Possible values are: unknown, adversary, capability,
	// infrastructure, victim.
	DiamondModel *DiamondModel `json:"diamondModel,omitempty"`

	DomainName           nullable.Type[string] `json:"domainName,omitempty"`
	EmailEncoding        nullable.Type[string] `json:"emailEncoding,omitempty"`
	EmailLanguage        nullable.Type[string] `json:"emailLanguage,omitempty"`
	EmailRecipient       nullable.Type[string] `json:"emailRecipient,omitempty"`
	EmailSenderAddress   nullable.Type[string] `json:"emailSenderAddress,omitempty"`
	EmailSenderName      nullable.Type[string] `json:"emailSenderName,omitempty"`
	EmailSourceDomain    nullable.Type[string] `json:"emailSourceDomain,omitempty"`
	EmailSourceIPAddress nullable.Type[string] `json:"emailSourceIpAddress,omitempty"`
	EmailSubject         nullable.Type[string] `json:"emailSubject,omitempty"`
	EmailXMailer         nullable.Type[string] `json:"emailXMailer,omitempty"`

	// DateTime string indicating when the Indicator expires. All indicators must have an expiration date to avoid stale
	// indicators persisting in the system. The Timestamp type represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z. Required.
	ExpirationDateTime nullable.Type[string] `json:"expirationDateTime,omitempty"`

	// An identification number that ties the indicator back to the indicator provider’s system (for example, a foreign
	// key).
	ExternalId nullable.Type[string] `json:"externalId,omitempty"`

	FileCompileDateTime nullable.Type[string] `json:"fileCompileDateTime,omitempty"`
	FileCreatedDateTime nullable.Type[string] `json:"fileCreatedDateTime,omitempty"`
	FileHashType        *FileHashType         `json:"fileHashType,omitempty"`
	FileHashValue       nullable.Type[string] `json:"fileHashValue,omitempty"`
	FileMutexName       nullable.Type[string] `json:"fileMutexName,omitempty"`
	FileName            nullable.Type[string] `json:"fileName,omitempty"`
	FilePacker          nullable.Type[string] `json:"filePacker,omitempty"`
	FilePath            nullable.Type[string] `json:"filePath,omitempty"`
	FileSize            nullable.Type[int64]  `json:"fileSize,omitempty"`
	FileType            nullable.Type[string] `json:"fileType,omitempty"`

	// Stamped by the system when the indicator is ingested. The Timestamp type represents date and time information using
	// ISO 8601 format and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	IngestedDateTime nullable.Type[string] `json:"ingestedDateTime,omitempty"`

	// Used to deactivate indicators within system. By default, any indicator submitted is set as active. However, providers
	// may submit existing indicators with this set to ‘False’ to deactivate indicators in the system.
	IsActive nullable.Type[bool] `json:"isActive,omitempty"`

	// A JSON array of strings that describes which point or points on the Kill Chain this indicator targets. See
	// ‘killChain values’ below for exact values.
	KillChain *[]string `json:"killChain,omitempty"`

	// Scenarios in which the indicator may cause false positives. This should be human-readable text.
	KnownFalsePositives nullable.Type[string] `json:"knownFalsePositives,omitempty"`

	// The last time the indicator was seen. The Timestamp type represents date and time information using ISO 8601 format
	// and is always in UTC time. For example, midnight UTC on Jan 1, 2014 is 2014-01-01T00:00:00Z
	LastReportedDateTime nullable.Type[string] `json:"lastReportedDateTime,omitempty"`

	// The malware family name associated with an indicator if it exists. Microsoft prefers the Microsoft malware family
	// name if at all possible that can be found via the Windows Defender Security Intelligence threat encyclopedia.
	MalwareFamilyNames *[]string `json:"malwareFamilyNames,omitempty"`

	NetworkCIDRBlock            nullable.Type[string] `json:"networkCidrBlock,omitempty"`
	NetworkDestinationAsn       nullable.Type[int64]  `json:"networkDestinationAsn,omitempty"`
	NetworkDestinationCIDRBlock nullable.Type[string] `json:"networkDestinationCidrBlock,omitempty"`
	NetworkDestinationIPv4      nullable.Type[string] `json:"networkDestinationIPv4,omitempty"`
	NetworkDestinationIPv6      nullable.Type[string] `json:"networkDestinationIPv6,omitempty"`
	NetworkDestinationPort      nullable.Type[int64]  `json:"networkDestinationPort,omitempty"`
	NetworkIPv4                 nullable.Type[string] `json:"networkIPv4,omitempty"`
	NetworkIPv6                 nullable.Type[string] `json:"networkIPv6,omitempty"`
	NetworkPort                 nullable.Type[int64]  `json:"networkPort,omitempty"`
	NetworkProtocol             nullable.Type[int64]  `json:"networkProtocol,omitempty"`
	NetworkSourceAsn            nullable.Type[int64]  `json:"networkSourceAsn,omitempty"`
	NetworkSourceCIDRBlock      nullable.Type[string] `json:"networkSourceCidrBlock,omitempty"`
	NetworkSourceIPv4           nullable.Type[string] `json:"networkSourceIPv4,omitempty"`
	NetworkSourceIPv6           nullable.Type[string] `json:"networkSourceIPv6,omitempty"`
	NetworkSourcePort           nullable.Type[int64]  `json:"networkSourcePort,omitempty"`

	// Determines if the indicator should trigger an event that is visible to an end-user. When set to ‘true,’ security
	// tools won't notify the end user that a ‘hit’ has occurred. This is most often treated as audit or silent mode by
	// security products where they'll simply log that a match occurred but won't perform the action. Default value is
	// false.
	PassiveOnly nullable.Type[bool] `json:"passiveOnly,omitempty"`

	// An integer representing the severity of the malicious behavior identified by the data within the indicator.
	// Acceptable values are 0 – 5 where 5 is the most severe and zero isn't severe at all. Default value is 3.
	Severity nullable.Type[int64] `json:"severity,omitempty"`

	// A JSON array of strings that stores arbitrary tags/keywords.
	Tags *[]string `json:"tags,omitempty"`

	// A string value representing a single security product to which the indicator should be applied. Acceptable values
	// are: Azure Sentinel, Microsoft Defender ATP. Required
	TargetProduct *string `json:"targetProduct,omitempty"`

	// Each indicator must have a valid Indicator Threat Type. Possible values are: Botnet, C2, CryptoMining, Darknet, DDoS,
	// MaliciousUrl, Malware, Phishing, Proxy, PUA, WatchList. Required.
	ThreatType nullable.Type[string] `json:"threatType,omitempty"`

	// Traffic Light Protocol value for the indicator. Possible values are: unknown, white, green, amber, red. Required.
	TlpLevel TlpLevel `json:"tlpLevel"`

	Url       nullable.Type[string] `json:"url,omitempty"`
	UserAgent nullable.Type[string] `json:"userAgent,omitempty"`

	// Fields inherited from Entity

	// The unique identifier for an entity. Read-only.
	Id *string `json:"id,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Model Behaviors
	OmitDiscriminatedValue bool `json:"-"`
}

func (s TiIndicator) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = TiIndicator{}

func (s TiIndicator) MarshalJSON() ([]byte, error) {
	type wrapper TiIndicator
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TiIndicator: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TiIndicator: %+v", err)
	}

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.tiIndicator"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TiIndicator: %+v", err)
	}

	return encoded, nil
}
