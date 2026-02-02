package stable

import (
	"encoding/json"
	"fmt"

	"github.com/hashicorp/go-azure-sdk/sdk/nullable"
)

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Entity = Invitation{}

type Invitation struct {
	// The URL the user can use to redeem their invitation. Read-only.
	InviteRedeemUrl nullable.Type[string] `json:"inviteRedeemUrl,omitempty"`

	// The URL the user should be redirected to after the invitation is redeemed. Required.
	InviteRedirectUrl string `json:"inviteRedirectUrl"`

	// The user created as part of the invitation creation. Read-only. The id property is required in the request body to
	// reset a redemption status.
	InvitedUser *User `json:"invitedUser,omitempty"`

	// The display name of the user being invited.
	InvitedUserDisplayName nullable.Type[string] `json:"invitedUserDisplayName,omitempty"`

	// The email address of the user being invited. Required. The following special characters aren't permitted in the email
	// address:Tilde (~)Exclamation point (!)Number sign (#)Dollar sign ($)Percent (%)Circumflex (^)Ampersand (&)Asterisk
	// (*)Parentheses (( ))Plus sign (+)Equal sign (=)Brackets ([ ])Braces ({ })Backslash (/)Slash mark (/)Pipe
	// (/|)Semicolon (;)Colon (:)Quotation marks (')Angle brackets (< >)Question mark (?)Comma (,)However, the following
	// exceptions apply:A period (.) or a hyphen (-) is permitted anywhere in the user name, except at the beginning or end
	// of the name.An underscore (_) is permitted anywhere in the user name, including at the beginning or end of the name.
	InvitedUserEmailAddress string `json:"invitedUserEmailAddress"`

	// Contains configuration for the message being sent to the invited user, including customizing message text, language,
	// and cc recipient list.
	InvitedUserMessageInfo *InvitedUserMessageInfo `json:"invitedUserMessageInfo,omitempty"`

	// The users or groups who are sponsors of the invited user. Sponsors are users and groups that are responsible for
	// guest users' privileges in the tenant and for keeping the guest users' information and access up to date.
	InvitedUserSponsors *[]DirectoryObject `json:"invitedUserSponsors,omitempty"`

	// List of OData IDs for `InvitedUserSponsors` to bind to this entity
	InvitedUserSponsors_ODataBind *[]string `json:"invitedUserSponsors@odata.bind,omitempty"`

	// The userType of the user being invited. By default, this is Guest. You can invite as Member if you're a company
	// administrator.
	InvitedUserType nullable.Type[string] `json:"invitedUserType,omitempty"`

	// Reset the user's redemption status and reinvite a user while retaining their user identifier, group memberships, and
	// app assignments. This property allows you to enable a user to sign-in using a different email address from the one in
	// the previous invitation. When true, the invitedUser/id relationship is required. For more information about using
	// this property, see Reset redemption status for a guest user.
	ResetRedemption nullable.Type[bool] `json:"resetRedemption,omitempty"`

	// Indicates whether an email should be sent to the user being invited. The default is false.
	SendInvitationMessage nullable.Type[bool] `json:"sendInvitationMessage,omitempty"`

	// The status of the invitation. Possible values are: PendingAcceptance, Completed, InProgress, and Error.
	Status nullable.Type[string] `json:"status,omitempty"`

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

func (s Invitation) Entity() BaseEntityImpl {
	return BaseEntityImpl{
		Id:        s.Id,
		ODataId:   s.ODataId,
		ODataType: s.ODataType,
	}
}

var _ json.Marshaler = Invitation{}

func (s Invitation) MarshalJSON() ([]byte, error) {
	type wrapper Invitation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling Invitation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling Invitation: %+v", err)
	}

	delete(decoded, "inviteRedeemUrl")
	delete(decoded, "invitedUser")

	if !s.OmitDiscriminatedValue {
		decoded["@odata.type"] = "#microsoft.graph.invitation"
	}

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling Invitation: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &Invitation{}

func (s *Invitation) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		InviteRedeemUrl               nullable.Type[string]   `json:"inviteRedeemUrl,omitempty"`
		InviteRedirectUrl             string                  `json:"inviteRedirectUrl"`
		InvitedUser                   *User                   `json:"invitedUser,omitempty"`
		InvitedUserDisplayName        nullable.Type[string]   `json:"invitedUserDisplayName,omitempty"`
		InvitedUserEmailAddress       string                  `json:"invitedUserEmailAddress"`
		InvitedUserMessageInfo        *InvitedUserMessageInfo `json:"invitedUserMessageInfo,omitempty"`
		InvitedUserSponsors_ODataBind *[]string               `json:"invitedUserSponsors@odata.bind,omitempty"`
		InvitedUserType               nullable.Type[string]   `json:"invitedUserType,omitempty"`
		ResetRedemption               nullable.Type[bool]     `json:"resetRedemption,omitempty"`
		SendInvitationMessage         nullable.Type[bool]     `json:"sendInvitationMessage,omitempty"`
		Status                        nullable.Type[string]   `json:"status,omitempty"`
		Id                            *string                 `json:"id,omitempty"`
		ODataId                       *string                 `json:"@odata.id,omitempty"`
		ODataType                     *string                 `json:"@odata.type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.InviteRedeemUrl = decoded.InviteRedeemUrl
	s.InviteRedirectUrl = decoded.InviteRedirectUrl
	s.InvitedUser = decoded.InvitedUser
	s.InvitedUserDisplayName = decoded.InvitedUserDisplayName
	s.InvitedUserEmailAddress = decoded.InvitedUserEmailAddress
	s.InvitedUserMessageInfo = decoded.InvitedUserMessageInfo
	s.InvitedUserSponsors_ODataBind = decoded.InvitedUserSponsors_ODataBind
	s.InvitedUserType = decoded.InvitedUserType
	s.ResetRedemption = decoded.ResetRedemption
	s.SendInvitationMessage = decoded.SendInvitationMessage
	s.Status = decoded.Status
	s.Id = decoded.Id
	s.ODataId = decoded.ODataId
	s.ODataType = decoded.ODataType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Invitation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["invitedUserSponsors"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling InvitedUserSponsors into list []json.RawMessage: %+v", err)
		}

		output := make([]DirectoryObject, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalDirectoryObjectImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'InvitedUserSponsors' for 'Invitation': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.InvitedUserSponsors = &output
	}

	return nil
}
