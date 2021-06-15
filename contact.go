package intercom

import "fmt"

// ContactService handles interactions with the API through a ContactRepository.
type ContactService struct {
	Repository ContactRepository
}

// ContactList holds a list of Contacts and paging information
type ContactList struct {
	Contacts   []Contact  `json:"data"`
	Pages      PageParams `json:"pages"`
	TotalCount int64      `json:"total_count"`
}

// Contact represents a Contact within Intercom.
// Not all of the fields are writeable to the API, non-writeable fields are
// stripped out from the request. Please see the API documentation for details.
type Contact struct {
	// The type of object - contact.
	Type string `json:"type,omitempty"`
	// The unique identifier for the contact which is given by Intercom.
	ID string `json:"id,omitempty"`
	// The id of the workspace which the contact belongs to.
	WorkspaceID string `json:"workspace_id,omitempty"`
	// A unique identifier for the contact which is given to Intercom.
	UserID string `json:"external_id,omitempty"`
	// The role of the contact - user or lead.
	Role string `json:"role,omitempty"`
	// The contacts email.
	Email string `json:"email,omitempty"`
	// The contacts phone.
	Phone string `json:"phone,omitempty"`
	// The contacts name.
	Name string `json:"name,omitempty"`
	// An image URL containing the avatar of a contact.
	Avatar string `json:"avatar,omitempty"`
	// The id of an admin that has been assigned account ownership of the contact.
	OwnerID int64 `json:"owner_id,omitempty"`
	// A list of social profiles associated to the contact.
	SocialProfiles *SocialProfileList `json:"social_profiles,omitempty"`
	// Whether the contact has had an email sent to them hard bounce.
	HasHardBounced bool `json:"has_hard_bounced"`
	// Whether the contact has marked an email sent to them as spam.
	MarkedEmailAsSpam bool `json:"marked_email_as_spam"`
	// Whether the contact is unsubscribed from emails.
	UnsubscribedFromEmails bool `json:"unsubscribed_from_emails"`
	// The time when the contact was created.
	CreatedAt int64 `json:"created_at,omitempty"`
	// The time when the contact was last updated.
	UpdatedAt int64 `json:"updated_at,omitempty"`
	// The time specified for when a contact signed up.
	SignedUpAt int64 `json:"signed_up_at,omitempty"`
	// The time when the contact was last seen (either where the Intercom Messenger was installed or when specified manually).
	LastSeenAt int64 `json:"last_seen_at,omitempty"`
	// The time when the contact last messaged in.
	LastRepliedAt int64 `json:"last_replied_at,omitempty"`
	// The time when the contact was last messaged.
	LastContactedAt int64 `json:"last_contacted_at,omitempty"`
	// The time when the contact last opened an email.
	LastEmailOpenedAt int64 `json:"last_email_opened_at,omitempty"`
	// The time when the contact last clicked a link in an email.
	LastEmailClickedAt int64 `json:"last_email_clicked_at,omitempty"`
	// A preferred language setting for the contact, used by the Intercom Messenger even if their browser settings change.
	LanguageOverride string `json:"language_override,omitempty"`
	// The name of the browser which the contact is using.
	Browser string `json:"browser,omitempty"`
	// The version of the browser which the contact is using.
	BrowserVersion string `json:"browser_version,omitempty"`
	// The language set by the browser which the contact is using.
	BrowserLanguage string `json:"browser_language,omitempty"`
	// The operating system which the contact is using.
	Os string `json:"os,omitempty"`
	// An object showing location details of the contact.
	Location *Location `json:"location,omitempty"`
	// The name of the Android app which the contact is using.
	AndroidAppName string `json:"android_app_name,omitempty"`
	// The version of the Android app which the contact is using.
	AndroidAppVersion string `json:"android_app_version,omitempty"`
	// The Android device which the contact is using.
	AndroidDevice string `json:"android_device,omitempty"`
	// The version of the Android OS which the contact is using.
	AndroidOsVersion string `json:"android_os_version,omitempty"`
	// The version of the Android SDK which the contact is using.
	AndroidSdkVersion string `json:"android_sdk_version,omitempty"`
	// The last time the contact used the Android app.
	AndroidLastSeenAt int64 `json:"android_last_seen_at,omitempty"`
	// The name of the iOS app which the contact is using.
	IOSAppName string `json:"ios_app_name,omitempty"`
	// The version of the iOS app which the contact is using.
	IOSAppVersion string `json:"ios_app_version,omitempty"`
	// The iOS device which the contact is using.
	IOSDevice string `json:"ios_device,omitempty"`
	// The version of iOS which the contact is using.
	IOSOsVersion string `json:"ios_os_version,omitempty"`
	// The version of the iOS SDK which the contact is using.
	IOSSdkVersion string `json:"ios_sdk_version,omitempty"`
	// The last time the contact used the iOS app.
	IOSLastSeenAt int64 `json:"ios_last_seen_at,omitempty"`
	// The custom attributes which are set for the contact.
	CustomAttributes map[string]interface{} `json:"custom_attributes,omitempty"`
	// The tags which have been added to the contact.
	Tags *AddressableList `json:"tags,omitempty"`
	// The notes which have been added to the contact.
	Notes *AddressableList `json:"notes,omitempty"`
	// The companies which the contact belongs to.
	Companies *AddressableList `json:"companies,omitempty"`
}

type contactListParams struct {
	PageParams
	SegmentID string `url:"segment_id,omitempty"`
	TagID     string `url:"tag_id,omitempty"`
	Email     string `url:"email,omitempty"`
}

// FindByID looks up a Contact by their Intercom ID.
func (c *ContactService) FindByID(id string) (Contact, error) {
	return c.findWithIdentifiers(UserIdentifiers{ID: id})
}

// FindByUserID looks up a Contact by their UserID (automatically generated server side).
func (c *ContactService) FindByUserID(userID string) (Contact, error) {
	return c.findWithIdentifiers(UserIdentifiers{UserID: userID})
}

func (c *ContactService) findWithIdentifiers(identifiers UserIdentifiers) (Contact, error) {
	return c.Repository.find(identifiers)
}

// List all Contacts for App.
func (c *ContactService) List(params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params})
}

// List all Contacts for App via Scroll API
func (c *ContactService) Scroll(scrollParam string) (ContactList, error) {
	return c.Repository.scroll(scrollParam)
}

// ListByEmail looks up a list of Contacts by their Email.
func (c *ContactService) ListByEmail(email string, params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params, Email: email})
}

// List Contacts by Segment.
func (c *ContactService) ListBySegment(segmentID string, params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params, SegmentID: segmentID})
}

// List Contacts By Tag.
func (c *ContactService) ListByTag(tagID string, params PageParams) (ContactList, error) {
	return c.Repository.list(contactListParams{PageParams: params, TagID: tagID})
}

// Create Contact
func (c *ContactService) Create(contact *Contact) (Contact, error) {
	return c.Repository.create(contact)
}

// Update Contact
func (c *ContactService) Update(contact *Contact) (Contact, error) {
	return c.Repository.update(contact)
}

// Convert Contact to User
func (c *ContactService) Convert(contact *Contact, user *User) (User, error) {
	return c.Repository.convert(contact, user)
}

// Delete Contact
func (c *ContactService) Delete(contact *Contact) (Contact, error) {
	return c.Repository.delete(contact.ID)
}

// MessageAddress gets the address for a Contact in order to message them
func (c Contact) MessageAddress() MessageAddress {
	return MessageAddress{
		Type:   "contact",
		ID:     c.ID,
		Email:  c.Email,
		UserID: c.UserID,
	}
}

func (c Contact) String() string {
	return fmt.Sprintf("[intercom] contact { id: %s name: %s, user_id: %s, email: %s }", c.ID, c.Name, c.UserID, c.Email)
}
