package beta

// Copyright IBM Corp. 2021, 2025 All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSPrivacyAccessControlItem struct {
	// Possible values of a property
	Accessibility *Enablement `json:"accessibility,omitempty"`

	// Possible values of a property
	AddressBook *Enablement `json:"addressBook,omitempty"`

	// Allow or deny the app or process to send a restricted Apple event to another app or process. You will need to know
	// the identifier, identifier type, and code requirement of the receiving app or process. This collection can contain a
	// maximum of 500 elements.
	AppleEventsAllowedReceivers *[]MacOSAppleEventReceiver `json:"appleEventsAllowedReceivers,omitempty"`

	// Block access to camera app.
	BlockCamera *bool `json:"blockCamera,omitempty"`

	// Block the app or process from listening to events from input devices such as mouse, keyboard, and trackpad.Requires
	// macOS 10.15 or later.
	BlockListenEvent *bool `json:"blockListenEvent,omitempty"`

	// Block access to microphone.
	BlockMicrophone *bool `json:"blockMicrophone,omitempty"`

	// Block app from capturing contents of system display. Requires macOS 10.15 or later.
	BlockScreenCapture *bool `json:"blockScreenCapture,omitempty"`

	// Possible values of a property
	Calendar *Enablement `json:"calendar,omitempty"`

	// Enter the code requirement, which can be obtained with the command 'codesign –display -r –' in the Terminal app.
	// Include everything after '=>'.
	CodeRequirement *string `json:"codeRequirement,omitempty"`

	// The display name of the app, process, or executable.
	DisplayName *string `json:"displayName,omitempty"`

	// Possible values of a property
	FileProviderPresence *Enablement `json:"fileProviderPresence,omitempty"`

	// The bundle ID or path of the app, process, or executable.
	Identifier *string `json:"identifier,omitempty"`

	// Process identifier types for MacOS Privacy Preferences
	IdentifierType *MacOSProcessIdentifierType `json:"identifierType,omitempty"`

	// Possible values of a property
	MediaLibrary *Enablement `json:"mediaLibrary,omitempty"`

	// The OData ID of this entity
	ODataId *string `json:"@odata.id,omitempty"`

	// The OData Type of this entity
	ODataType *string `json:"@odata.type,omitempty"`

	// Possible values of a property
	Photos *Enablement `json:"photos,omitempty"`

	// Possible values of a property
	PostEvent *Enablement `json:"postEvent,omitempty"`

	// Possible values of a property
	Reminders *Enablement `json:"reminders,omitempty"`

	// Possible values of a property
	SpeechRecognition *Enablement `json:"speechRecognition,omitempty"`

	// Statically validates the code requirement. Use this setting if the process invalidates its dynamic code signature.
	StaticCodeValidation *bool `json:"staticCodeValidation,omitempty"`

	// Possible values of a property
	SystemPolicyAllFiles *Enablement `json:"systemPolicyAllFiles,omitempty"`

	// Possible values of a property
	SystemPolicyDesktopFolder *Enablement `json:"systemPolicyDesktopFolder,omitempty"`

	// Possible values of a property
	SystemPolicyDocumentsFolder *Enablement `json:"systemPolicyDocumentsFolder,omitempty"`

	// Possible values of a property
	SystemPolicyDownloadsFolder *Enablement `json:"systemPolicyDownloadsFolder,omitempty"`

	// Possible values of a property
	SystemPolicyNetworkVolumes *Enablement `json:"systemPolicyNetworkVolumes,omitempty"`

	// Possible values of a property
	SystemPolicyRemovableVolumes *Enablement `json:"systemPolicyRemovableVolumes,omitempty"`

	// Possible values of a property
	SystemPolicySystemAdminFiles *Enablement `json:"systemPolicySystemAdminFiles,omitempty"`
}
