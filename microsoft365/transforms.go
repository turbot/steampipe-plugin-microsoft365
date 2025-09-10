package microsoft365

import (
	"time"

	"github.com/microsoftgraph/msgraph-sdk-go/models"
)

type Microsoft365CalendarInfo struct {
	models.Calendarable
	CalendarGroupID string
	UserID          string
}

type Microsoft365CalendarEventInfo struct {
	models.Eventable
	UserID string
}

type Microsoft365CalendarGroupInfo struct {
	models.CalendarGroupable
	UserID string
}
type Microsoft365ContactInfo struct {
	models.Contactable
	UserID string
}

type Microsoft365DriveInfo struct {
	models.Driveable
	UserID string
}

type Microsoft365DriveItemInfo struct {
	models.DriveItemable
	DriveID string
	UserID  string
}

type Microsoft365MailMessageInfo struct {
	models.Messageable
	UserID string
}

type Microsoft365OrgContactInfo struct {
	models.OrgContactable
}

type Microsoft365GroupInfo struct {
	models.Groupable
}

type Microsoft365TeamsSettingsInfo struct {
	TeamsCount int    `json:"teams_count"`
	Note       string `json:"note"`
}

type Microsoft365PlannerSettingsInfo struct {
	PlansCount int    `json:"plans_count"`
	Note       string `json:"note"`
}

type Microsoft365TeamworkSettingsInfo struct {
	WorkforceIntegrationsCount int    `json:"workforce_integrations_count"`
	Note                       string `json:"note"`
}

type Microsoft365SecurityInfo struct {
	models.Securityable
}

type Microsoft365TeamInfo struct {
	models.Teamable
	ID string
}

type Microsoft365TeamChannelInfo struct {
	models.Channelable
	TeamID string
	UserID string
}

type Microsoft365TeamMemberInfo struct {
	TeamID   string
	MemberID string
}

func (calendar *Microsoft365CalendarInfo) CalendarColor() string {
	if calendar.GetColor() == nil {
		return ""
	}
	return calendar.GetColor().String()
}

func (calendar *Microsoft365CalendarInfo) CalendarDefaultOnlineMeetingProvider() string {
	if calendar.GetDefaultOnlineMeetingProvider() == nil {
		return ""
	}
	return calendar.GetDefaultOnlineMeetingProvider().String()
}

func (calendar *Microsoft365CalendarInfo) CalendarMultiValueExtendedProperties() []map[string]interface{} {
	if calendar.GetMultiValueExtendedProperties() == nil {
		return nil
	}

	var multiValueExtendedProperties []map[string]interface{}
	for _, i := range calendar.GetMultiValueExtendedProperties() {
		multiValueExtendedProperties = append(multiValueExtendedProperties, map[string]interface{}{
			"value": i.GetValue(),
		})
	}
	return multiValueExtendedProperties
}

func (calendar *Microsoft365CalendarInfo) GetAllowedOnlineMeetingProviders() []string {
	if calendar.Calendarable.GetAllowedOnlineMeetingProviders() == nil {
		return nil
	}

	var providers []string
	for _, provider := range calendar.Calendarable.GetAllowedOnlineMeetingProviders() {
		// Convert enum values to string values
		providers = append(providers, provider.String())
	}
	return providers
}

// ConvertCalendarPermissionAllowedRoles converts numeric allowedRoles array to string array
func ConvertCalendarPermissionAllowedRoles(allowedRoles []models.CalendarRoleType) []string {
	if allowedRoles == nil {
		return nil
	}

	var roles []string
	for _, role := range allowedRoles {
		// Convert numeric enum values to string values
		roles = append(roles, role.String())

	}
	return roles
}

func (calendar *Microsoft365CalendarInfo) CalendarOwner() map[string]interface{} {
	if calendar.GetOwner() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if calendar.GetOwner().GetAddress() != nil {
		data["address"] = *calendar.GetOwner().GetAddress()
	}
	if calendar.GetOwner().GetName() != nil {
		data["name"] = *calendar.GetOwner().GetName()
	}
	// Always include @odata_type for backward compatibility
	if calendar.GetOwner().GetOdataType() != nil {
		data["@odata_type"] = *calendar.GetOwner().GetOdataType()
	} else {
		// Default to microsoft.graph.emailAddress if not provided
		data["@odata_type"] = "#microsoft.graph.emailAddress"
	}
	return data
}

func (contact *Microsoft365ContactInfo) ContactBusinessAddress() map[string]interface{} {
	if contact.GetBusinessAddress() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if contact.GetBusinessAddress().GetCity() != nil {
		data["city"] = *contact.GetBusinessAddress().GetCity()
	}
	if contact.GetBusinessAddress().GetCountryOrRegion() != nil {
		data["countryOrRegion"] = *contact.GetBusinessAddress().GetCountryOrRegion()
	}
	if contact.GetBusinessAddress().GetPostalCode() != nil {
		data["postalCode"] = *contact.GetBusinessAddress().GetPostalCode()
	}
	if contact.GetBusinessAddress().GetState() != nil {
		data["state"] = *contact.GetBusinessAddress().GetState()
	}
	if contact.GetBusinessAddress().GetStreet() != nil {
		data["street"] = *contact.GetBusinessAddress().GetStreet()
	}

	return data
}

func (contact *Microsoft365ContactInfo) ContactHomeAddress() map[string]interface{} {
	if contact.GetHomeAddress() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if contact.GetHomeAddress().GetCity() != nil {
		data["city"] = *contact.GetHomeAddress().GetCity()
	}
	if contact.GetHomeAddress().GetCountryOrRegion() != nil {
		data["countryOrRegion"] = *contact.GetHomeAddress().GetCountryOrRegion()
	}
	if contact.GetHomeAddress().GetPostalCode() != nil {
		data["postalCode"] = *contact.GetHomeAddress().GetPostalCode()
	}
	if contact.GetHomeAddress().GetState() != nil {
		data["state"] = *contact.GetHomeAddress().GetState()
	}
	if contact.GetHomeAddress().GetStreet() != nil {
		data["street"] = *contact.GetHomeAddress().GetStreet()
	}

	return data
}

func (contact *Microsoft365ContactInfo) ContactOtherAddress() map[string]interface{} {
	if contact.GetOtherAddress() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if contact.GetOtherAddress().GetCity() != nil {
		data["city"] = *contact.GetOtherAddress().GetCity()
	}
	if contact.GetOtherAddress().GetCountryOrRegion() != nil {
		data["countryOrRegion"] = *contact.GetOtherAddress().GetCountryOrRegion()
	}
	if contact.GetOtherAddress().GetPostalCode() != nil {
		data["postalCode"] = *contact.GetOtherAddress().GetPostalCode()
	}
	if contact.GetOtherAddress().GetState() != nil {
		data["state"] = *contact.GetOtherAddress().GetState()
	}
	if contact.GetOtherAddress().GetStreet() != nil {
		data["street"] = *contact.GetOtherAddress().GetStreet()
	}

	return data
}

func (contact *Microsoft365ContactInfo) ContactEmailAddresses() []map[string]interface{} {
	if contact.GetEmailAddresses() == nil {
		return nil
	}

	contacts := []map[string]interface{}{}
	for _, c := range contact.GetEmailAddresses() {
		data := map[string]interface{}{}
		if c.GetAddress() != nil {
			data["address"] = *c.GetAddress()
		}
		if c.GetName() != nil {
			data["name"] = *c.GetName()
		}
		contacts = append(contacts, data)
	}

	return contacts
}

func (drive *Microsoft365DriveInfo) DriveCreatedBy() map[string]interface{} {
	if drive.GetCreatedBy() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if drive.GetCreatedBy().GetApplication() != nil {
		applicationData := map[string]interface{}{}
		if drive.GetCreatedBy().GetApplication().GetDisplayName() != nil {
			applicationData["displayName"] = *drive.GetCreatedBy().GetApplication().GetDisplayName()
		}
		if drive.GetCreatedBy().GetApplication().GetId() != nil {
			applicationData["id"] = *drive.GetCreatedBy().GetApplication().GetId()
		}
		data["application"] = applicationData
	}
	if drive.GetCreatedBy().GetUser() != nil {
		userData := map[string]interface{}{}
		if drive.GetCreatedBy().GetUser().GetDisplayName() != nil {
			userData["displayName"] = *drive.GetCreatedBy().GetUser().GetDisplayName()
		}
		if drive.GetCreatedBy().GetUser().GetId() != nil {
			userData["id"] = *drive.GetCreatedBy().GetUser().GetId()
		}
		data["user"] = userData
	}
	if drive.GetCreatedBy().GetDevice() != nil {
		deviceData := map[string]interface{}{}
		if drive.GetCreatedBy().GetDevice().GetDisplayName() != nil {
			deviceData["displayName"] = *drive.GetCreatedBy().GetDevice().GetDisplayName()
		}
		if drive.GetCreatedBy().GetDevice().GetId() != nil {
			deviceData["id"] = *drive.GetCreatedBy().GetDevice().GetId()
		}
		data["device"] = deviceData
	}

	return data
}

func (drive *Microsoft365DriveInfo) DriveLastModifiedBy() map[string]interface{} {
	if drive.GetLastModifiedBy() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if drive.GetLastModifiedBy().GetApplication() != nil {
		applicationData := map[string]interface{}{}
		if drive.GetLastModifiedBy().GetApplication().GetDisplayName() != nil {
			applicationData["displayName"] = *drive.GetLastModifiedBy().GetApplication().GetDisplayName()
		}
		if drive.GetLastModifiedBy().GetApplication().GetId() != nil {
			applicationData["id"] = *drive.GetLastModifiedBy().GetApplication().GetId()
		}
		data["application"] = applicationData
	}
	if drive.GetLastModifiedBy().GetUser() != nil {
		userData := map[string]interface{}{}
		if drive.GetLastModifiedBy().GetUser().GetDisplayName() != nil {
			userData["displayName"] = *drive.GetLastModifiedBy().GetUser().GetDisplayName()
		}
		if drive.GetLastModifiedBy().GetUser().GetId() != nil {
			userData["id"] = *drive.GetLastModifiedBy().GetUser().GetId()
		}
		data["user"] = userData
	}
	if drive.GetLastModifiedBy().GetDevice() != nil {
		deviceData := map[string]interface{}{}
		if drive.GetLastModifiedBy().GetDevice().GetDisplayName() != nil {
			deviceData["displayName"] = *drive.GetLastModifiedBy().GetDevice().GetDisplayName()
		}
		if drive.GetLastModifiedBy().GetDevice().GetId() != nil {
			deviceData["id"] = *drive.GetLastModifiedBy().GetDevice().GetId()
		}
		data["device"] = deviceData
	}

	return data
}

func (drive *Microsoft365DriveInfo) DriveParentReference() map[string]interface{} {
	if drive.GetParentReference() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if drive.GetParentReference().GetDriveId() != nil {
		data["driveId"] = *drive.GetParentReference().GetDriveId()
	}
	if drive.GetParentReference().GetDriveType() != nil {
		data["driveType"] = *drive.GetParentReference().GetDriveType()
	}
	if drive.GetParentReference().GetId() != nil {
		data["id"] = *drive.GetParentReference().GetId()
	}
	if drive.GetParentReference().GetName() != nil {
		data["name"] = *drive.GetParentReference().GetName()
	}
	if drive.GetParentReference().GetPath() != nil {
		data["path"] = *drive.GetParentReference().GetPath()
	}
	if drive.GetParentReference().GetShareId() != nil {
		data["shareId"] = *drive.GetParentReference().GetShareId()
	}
	if drive.GetParentReference().GetSiteId() != nil {
		data["siteId"] = *drive.GetParentReference().GetSiteId()
	}
	if drive.GetParentReference().GetSharepointIds() != nil {
		sharePointData := map[string]interface{}{}
		if drive.GetParentReference().GetSharepointIds().GetListId() != nil {
			sharePointData["listId"] = *drive.GetParentReference().GetSharepointIds().GetListId()
		}
		if drive.GetParentReference().GetSharepointIds().GetListItemId() != nil {
			sharePointData["listItemId"] = *drive.GetParentReference().GetSharepointIds().GetListItemId()
		}
		if drive.GetParentReference().GetSharepointIds().GetListItemUniqueId() != nil {
			sharePointData["listItemUniqueId"] = *drive.GetParentReference().GetSharepointIds().GetListItemUniqueId()
		}
		if drive.GetParentReference().GetSharepointIds().GetSiteId() != nil {
			sharePointData["siteId"] = *drive.GetParentReference().GetSharepointIds().GetSiteId()
		}
		if drive.GetParentReference().GetSharepointIds().GetSiteUrl() != nil {
			sharePointData["siteUrl"] = *drive.GetParentReference().GetSharepointIds().GetSiteUrl()
		}
		if drive.GetParentReference().GetSharepointIds().GetTenantId() != nil {
			sharePointData["tenantId"] = *drive.GetParentReference().GetSharepointIds().GetTenantId()
		}
		if drive.GetParentReference().GetSharepointIds().GetWebId() != nil {
			sharePointData["webId"] = *drive.GetParentReference().GetSharepointIds().GetWebId()
		}
		data["sharePointIds"] = sharePointData
	}

	return data
}

func (driveItem *Microsoft365DriveItemInfo) DriveItemCreatedBy() map[string]interface{} {
	if driveItem.GetCreatedBy() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if driveItem.GetCreatedBy().GetApplication() != nil {
		applicationData := map[string]interface{}{}
		if driveItem.GetCreatedBy().GetApplication().GetDisplayName() != nil {
			applicationData["displayName"] = *driveItem.GetCreatedBy().GetApplication().GetDisplayName()
		}
		if driveItem.GetCreatedBy().GetApplication().GetId() != nil {
			applicationData["id"] = *driveItem.GetCreatedBy().GetApplication().GetId()
		}
		data["application"] = applicationData
	}
	if driveItem.GetCreatedBy().GetUser() != nil {
		userData := map[string]interface{}{}
		if driveItem.GetCreatedBy().GetUser().GetDisplayName() != nil {
			userData["displayName"] = *driveItem.GetCreatedBy().GetUser().GetDisplayName()
		}
		if driveItem.GetCreatedBy().GetUser().GetId() != nil {
			userData["id"] = *driveItem.GetCreatedBy().GetUser().GetId()
		}
		data["user"] = userData
	}
	if driveItem.GetCreatedBy().GetDevice() != nil {
		deviceData := map[string]interface{}{}
		if driveItem.GetCreatedBy().GetDevice().GetDisplayName() != nil {
			deviceData["displayName"] = *driveItem.GetCreatedBy().GetDevice().GetDisplayName()
		}
		if driveItem.GetCreatedBy().GetDevice().GetId() != nil {
			deviceData["id"] = *driveItem.GetCreatedBy().GetDevice().GetId()
		}
		data["device"] = deviceData
	}

	return data
}

func (driveItem *Microsoft365DriveItemInfo) DriveItemFile() map[string]interface{} {
	if driveItem.GetFile() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if driveItem.GetFile().GetMimeType() != nil {
		data["mimeType"] = *driveItem.GetFile().GetMimeType()
	}
	if driveItem.GetFile().GetProcessingMetadata() != nil {
		data["processingMetadata"] = *driveItem.GetFile().GetProcessingMetadata()
	}
	if driveItem.GetFile().GetHashes() != nil {
		hashData := map[string]interface{}{}
		if driveItem.GetFile().GetHashes().GetCrc32Hash() != nil {
			hashData["crc32Hash"] = *driveItem.GetFile().GetHashes().GetCrc32Hash()
		}
		if driveItem.GetFile().GetHashes().GetQuickXorHash() != nil {
			hashData["quickXorHash"] = *driveItem.GetFile().GetHashes().GetQuickXorHash()
		}
		if driveItem.GetFile().GetHashes().GetSha1Hash() != nil {
			hashData["sha1Hash"] = *driveItem.GetFile().GetHashes().GetSha1Hash()
		}
		if driveItem.GetFile().GetHashes().GetSha256Hash() != nil {
			hashData["sha256Hash"] = *driveItem.GetFile().GetHashes().GetSha256Hash()
		}
		data["hashes"] = hashData
	}

	return data
}

func (driveItem *Microsoft365DriveItemInfo) DriveItemFilePath() string {
	if driveItem.GetParentReference() != nil && driveItem.GetParentReference().GetPath() != nil {
		return *driveItem.GetParentReference().GetPath()
	}
	return ""
}

func (driveItem *Microsoft365DriveItemInfo) DriveItemFolder() map[string]interface{} {
	if driveItem.GetFolder() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if driveItem.GetFolder().GetChildCount() != nil {
		data["childCount"] = *driveItem.GetFolder().GetChildCount()
	}
	return data
}

func (driveItem *Microsoft365DriveItemInfo) DriveItemLastModifiedBy() map[string]interface{} {
	if driveItem.GetLastModifiedBy() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if driveItem.GetLastModifiedBy().GetApplication() != nil {
		applicationData := map[string]interface{}{}
		if driveItem.GetLastModifiedBy().GetApplication().GetDisplayName() != nil {
			applicationData["displayName"] = *driveItem.GetLastModifiedBy().GetApplication().GetDisplayName()
		}
		if driveItem.GetLastModifiedBy().GetApplication().GetId() != nil {
			applicationData["id"] = *driveItem.GetLastModifiedBy().GetApplication().GetId()
		}
		data["application"] = applicationData
	}
	if driveItem.GetLastModifiedBy().GetUser() != nil {
		userData := map[string]interface{}{}
		if driveItem.GetLastModifiedBy().GetUser().GetDisplayName() != nil {
			userData["displayName"] = *driveItem.GetLastModifiedBy().GetUser().GetDisplayName()
		}
		if driveItem.GetLastModifiedBy().GetUser().GetId() != nil {
			userData["id"] = *driveItem.GetLastModifiedBy().GetUser().GetId()
		}
		data["user"] = userData
	}
	if driveItem.GetLastModifiedBy().GetDevice() != nil {
		deviceData := map[string]interface{}{}
		if driveItem.GetLastModifiedBy().GetDevice().GetDisplayName() != nil {
			deviceData["displayName"] = *driveItem.GetLastModifiedBy().GetDevice().GetDisplayName()
		}
		if driveItem.GetLastModifiedBy().GetDevice().GetId() != nil {
			deviceData["id"] = *driveItem.GetLastModifiedBy().GetDevice().GetId()
		}
		data["device"] = deviceData
	}

	return data
}

func (driveItem *Microsoft365DriveItemInfo) DriveItemParentReference() map[string]interface{} {
	if driveItem.GetParentReference() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if driveItem.GetParentReference().GetDriveId() != nil {
		data["driveId"] = *driveItem.GetParentReference().GetDriveId()
	}
	if driveItem.GetParentReference().GetDriveType() != nil {
		data["driveType"] = *driveItem.GetParentReference().GetDriveType()
	}
	if driveItem.GetParentReference().GetId() != nil {
		data["id"] = *driveItem.GetParentReference().GetId()
	}
	if driveItem.GetParentReference().GetName() != nil {
		data["name"] = *driveItem.GetParentReference().GetName()
	}
	if driveItem.GetParentReference().GetPath() != nil {
		data["path"] = *driveItem.GetParentReference().GetPath()
	}
	if driveItem.GetParentReference().GetShareId() != nil {
		data["shareId"] = *driveItem.GetParentReference().GetShareId()
	}
	if driveItem.GetParentReference().GetSiteId() != nil {
		data["siteId"] = *driveItem.GetParentReference().GetSiteId()
	}
	if driveItem.GetParentReference().GetSharepointIds() != nil {
		sharePointData := map[string]interface{}{}
		if driveItem.GetParentReference().GetSharepointIds().GetListId() != nil {
			sharePointData["listId"] = *driveItem.GetParentReference().GetSharepointIds().GetListId()
		}
		if driveItem.GetParentReference().GetSharepointIds().GetListItemId() != nil {
			sharePointData["listItemId"] = *driveItem.GetParentReference().GetSharepointIds().GetListItemId()
		}
		if driveItem.GetParentReference().GetSharepointIds().GetListItemUniqueId() != nil {
			sharePointData["listItemUniqueId"] = *driveItem.GetParentReference().GetSharepointIds().GetListItemUniqueId()
		}
		if driveItem.GetParentReference().GetSharepointIds().GetSiteId() != nil {
			sharePointData["siteId"] = *driveItem.GetParentReference().GetSharepointIds().GetSiteId()
		}
		if driveItem.GetParentReference().GetSharepointIds().GetSiteUrl() != nil {
			sharePointData["siteUrl"] = *driveItem.GetParentReference().GetSharepointIds().GetSiteUrl()
		}
		if driveItem.GetParentReference().GetSharepointIds().GetTenantId() != nil {
			sharePointData["tenantId"] = *driveItem.GetParentReference().GetSharepointIds().GetTenantId()
		}
		if driveItem.GetParentReference().GetSharepointIds().GetWebId() != nil {
			sharePointData["webId"] = *driveItem.GetParentReference().GetSharepointIds().GetWebId()
		}
		data["sharePointIds"] = sharePointData
	}

	return data
}

func (event *Microsoft365CalendarEventInfo) EventAttendees() []map[string]interface{} {
	if event.GetAttendees() == nil {
		return nil
	}

	attendees := []map[string]interface{}{}
	for _, a := range event.GetAttendees() {
		attendeeInfo := map[string]interface{}{}
		if a.GetStatus() != nil {
			responseStatusInfo := map[string]interface{}{}
			if a.GetStatus().GetResponse() != nil {
				responseStatusInfo["response"] = a.GetStatus().GetResponse().String()
			}
			if a.GetStatus().GetTime() != nil {
				responseStatusInfo["time"] = *a.GetStatus().GetTime()
			}
			attendeeInfo["status"] = responseStatusInfo
		}
		if a.GetProposedNewTime() != nil {
			data := map[string]interface{}{}
			if a.GetProposedNewTime().GetEnd() != nil {
				endTimeInfo := map[string]interface{}{}
				if event.GetEnd().GetDateTime() != nil {
					endTimeInfo["dateTime"] = *event.GetEnd().GetDateTime()
				}
				if event.GetEnd().GetTimeZone() != nil {
					endTimeInfo["timeZone"] = *event.GetEnd().GetTimeZone()
				}
				data["end"] = endTimeInfo
			}
			if a.GetProposedNewTime().GetStart() != nil {
				startTimeInfo := map[string]interface{}{}
				if event.GetStart().GetDateTime() != nil {
					startTimeInfo["dateTime"] = *event.GetStart().GetDateTime()
				}
				if event.GetStart().GetTimeZone() != nil {
					startTimeInfo["timeZone"] = *event.GetStart().GetTimeZone()
				}
				data["start"] = startTimeInfo
			}
			attendeeInfo["proposedNewTime"] = data
		}

		attendees = append(attendees, attendeeInfo)
	}
	return attendees
}

func (event *Microsoft365CalendarEventInfo) EventBody() map[string]interface{} {
	if event.GetBody() == nil {
		return nil
	}

	bodyInfo := map[string]interface{}{}
	if event.GetBody().GetContent() != nil {
		bodyInfo["content"] = *event.GetBody().GetContent()
	}
	if event.GetBody().GetContentType() != nil {
		bodyInfo["contentType"] = *event.GetBody().GetContentType()
	}
	return bodyInfo
}

func (event *Microsoft365CalendarEventInfo) EventEnd() map[string]interface{} {
	if event.GetEnd() == nil {
		return nil
	}

	endTimeInfo := map[string]interface{}{}
	if event.GetEnd().GetDateTime() != nil {
		endTimeInfo["dateTime"] = *event.GetEnd().GetDateTime()
	}
	if event.GetEnd().GetDateTime() != nil {
		endTimeInfo["timeZone"] = *event.GetEnd().GetTimeZone()
	}
	return endTimeInfo
}

func (event *Microsoft365CalendarEventInfo) EventLocation() map[string]interface{} {
	if event.GetLocation() == nil {
		return nil
	}

	locationInfo := map[string]interface{}{}
	if event.GetLocation().GetDisplayName() != nil {
		locationInfo["displayName"] = *event.GetLocation().GetDisplayName()
	}
	if event.GetLocation().GetDisplayName() != nil {
		locationInfo["displayName"] = *event.GetLocation().GetDisplayName()
	}
	if event.GetLocation().GetLocationEmailAddress() != nil {
		locationInfo["locationEmailAddress"] = *event.GetLocation().GetLocationEmailAddress()
	}
	if event.GetLocation().GetLocationType() != nil {
		locationInfo["locationType"] = event.GetLocation().GetLocationType().String()
	}
	if event.GetLocation().GetLocationUri() != nil {
		locationInfo["locationUri"] = *event.GetLocation().GetLocationUri()
	}
	if event.GetLocation().GetOdataType() != nil {
		locationInfo["@odata_type"] = *event.GetLocation().GetOdataType()
	}
	if event.GetLocation().GetUniqueId() != nil {
		locationInfo["uniqueId"] = *event.GetLocation().GetUniqueId()
	}
	if event.GetLocation().GetUniqueIdType() != nil {
		locationInfo["uniqueIdType"] = event.GetLocation().GetUniqueIdType().String()
	}

	addressInfo := map[string]interface{}{}
	if event.GetLocation().GetAddress() != nil {
		if event.GetLocation().GetAddress().GetCity() != nil {
			addressInfo["city"] = *event.GetLocation().GetAddress().GetCity()
		}
		if event.GetLocation().GetAddress().GetCountryOrRegion() != nil {
			addressInfo["countryOrRegion"] = *event.GetLocation().GetAddress().GetCountryOrRegion()
		}
		if event.GetLocation().GetAddress().GetPostalCode() != nil {
			addressInfo["postalCode"] = *event.GetLocation().GetAddress().GetPostalCode()
		}
		if event.GetLocation().GetAddress().GetState() != nil {
			addressInfo["state"] = *event.GetLocation().GetAddress().GetState()
		}
		if event.GetLocation().GetAddress().GetStreet() != nil {
			addressInfo["street"] = *event.GetLocation().GetAddress().GetStreet()
		}
	}
	locationInfo["address"] = addressInfo

	coordinatesInfo := map[string]interface{}{}
	if event.GetLocation().GetCoordinates() != nil {
		if event.GetLocation().GetCoordinates().GetAccuracy() != nil {
			coordinatesInfo["accuracy"] = *event.GetLocation().GetCoordinates().GetAccuracy()
		}
		if event.GetLocation().GetCoordinates().GetAltitude() != nil {
			coordinatesInfo["altitude"] = *event.GetLocation().GetCoordinates().GetAltitude()
		}
		if event.GetLocation().GetCoordinates().GetAltitudeAccuracy() != nil {
			coordinatesInfo["altitudeAccuracy"] = *event.GetLocation().GetCoordinates().GetAltitudeAccuracy()
		}
		if event.GetLocation().GetCoordinates().GetLatitude() != nil {
			coordinatesInfo["latitude"] = *event.GetLocation().GetCoordinates().GetLatitude()
		}
		if event.GetLocation().GetCoordinates().GetLongitude() != nil {
			coordinatesInfo["longitude"] = *event.GetLocation().GetCoordinates().GetLongitude()
		}
	}
	locationInfo["coordinates"] = coordinatesInfo

	return locationInfo
}

func (event *Microsoft365CalendarEventInfo) EventLocations() []map[string]interface{} {
	if event.GetLocations() == nil {
		return nil
	}

	locInfo := []map[string]interface{}{}
	for _, l := range event.GetLocations() {
		locationInfo := map[string]interface{}{}
		if l.GetDisplayName() != nil {
			locationInfo["displayName"] = *l.GetDisplayName()
		}
		if l.GetDisplayName() != nil {
			locationInfo["displayName"] = *l.GetDisplayName()
		}
		if l.GetLocationEmailAddress() != nil {
			locationInfo["locationEmailAddress"] = *l.GetLocationEmailAddress()
		}
		if l.GetLocationType() != nil {
			locationInfo["locationType"] = l.GetLocationType().String()
		}
		if l.GetLocationUri() != nil {
			locationInfo["locationUri"] = *l.GetLocationUri()
		}
		if l.GetOdataType() != nil {
			locationInfo["@odata_type"] = *l.GetOdataType()
		}
		if l.GetUniqueId() != nil {
			locationInfo["uniqueId"] = *l.GetUniqueId()
		}
		if l.GetUniqueIdType() != nil {
			locationInfo["uniqueIdType"] = l.GetUniqueIdType().String()
		}

		addressInfo := map[string]interface{}{}
		if l.GetAddress() != nil {
			if l.GetAddress().GetCity() != nil {
				addressInfo["city"] = *l.GetAddress().GetCity()
			}
			if l.GetAddress().GetCountryOrRegion() != nil {
				addressInfo["countryOrRegion"] = *l.GetAddress().GetCountryOrRegion()
			}
			if l.GetAddress().GetPostalCode() != nil {
				addressInfo["postalCode"] = *l.GetAddress().GetPostalCode()
			}
			if l.GetAddress().GetState() != nil {
				addressInfo["state"] = *l.GetAddress().GetState()
			}
			if l.GetAddress().GetStreet() != nil {
				addressInfo["street"] = *l.GetAddress().GetStreet()
			}
		}
		locationInfo["address"] = addressInfo

		coordinatesInfo := map[string]interface{}{}
		if l.GetCoordinates() != nil {
			if l.GetCoordinates().GetAccuracy() != nil {
				coordinatesInfo["accuracy"] = *l.GetCoordinates().GetAccuracy()
			}
			if l.GetCoordinates().GetAltitude() != nil {
				coordinatesInfo["altitude"] = *l.GetCoordinates().GetAltitude()
			}
			if l.GetCoordinates().GetAltitudeAccuracy() != nil {
				coordinatesInfo["altitudeAccuracy"] = *l.GetCoordinates().GetAltitudeAccuracy()
			}
			if l.GetCoordinates().GetLatitude() != nil {
				coordinatesInfo["latitude"] = *l.GetCoordinates().GetLatitude()
			}
			if l.GetCoordinates().GetLongitude() != nil {
				coordinatesInfo["longitude"] = *l.GetCoordinates().GetLongitude()
			}
		}
		locationInfo["coordinates"] = coordinatesInfo

		locInfo = append(locInfo, locationInfo)
	}
	return locInfo
}

func (event *Microsoft365CalendarEventInfo) EventOnlineMeeting() map[string]interface{} {
	if event.GetOnlineMeeting() == nil {
		return nil
	}

	onlineMeetingInfo := map[string]interface{}{
		"tollFreeNumbers": event.GetOnlineMeeting().GetTollFreeNumbers(),
	}
	if event.GetOnlineMeeting().GetConferenceId() != nil {
		onlineMeetingInfo["conferenceId"] = *event.GetOnlineMeeting().GetConferenceId()
	}
	if event.GetOnlineMeeting().GetJoinUrl() != nil {
		onlineMeetingInfo["joinUrl"] = *event.GetOnlineMeeting().GetJoinUrl()
	}
	if event.GetOnlineMeeting().GetQuickDial() != nil {
		onlineMeetingInfo["quickDial"] = *event.GetOnlineMeeting().GetQuickDial()
	}
	if event.GetOnlineMeeting().GetTollNumber() != nil {
		onlineMeetingInfo["tollNumber"] = *event.GetOnlineMeeting().GetTollNumber()
	}
	if event.GetOnlineMeeting().GetPhones() != nil {
		phones := []map[string]interface{}{}
		for _, p := range event.GetOnlineMeeting().GetPhones() {
			data := map[string]interface{}{}
			if p.GetLanguage() != nil {
				data["language"] = *p.GetLanguage()
			}
			if p.GetNumber() != nil {
				data["number"] = *p.GetNumber()
			}
			if p.GetRegion() != nil {
				data["region"] = *p.GetRegion()
			}
			if p.GetTypeEscaped() != nil {
				data["type"] = p.GetTypeEscaped().String()
			}
			phones = append(phones, data)
		}
		onlineMeetingInfo["phones"] = phones
	}
	return onlineMeetingInfo
}

func (event *Microsoft365CalendarEventInfo) EventOrganizer() map[string]interface{} {
	if event.GetOrganizer() == nil {
		return nil
	}

	organizerInfo := map[string]interface{}{}
	if event.GetOrganizer().GetOdataType() != nil {
		organizerInfo["@odata_type"] = *event.GetOrganizer().GetOdataType()
	}

	if event.GetOrganizer().GetEmailAddress() != nil {
		emailAddressInfo := map[string]interface{}{}
		if event.GetOrganizer().GetEmailAddress().GetAddress() != nil {
			emailAddressInfo["address"] = *event.GetOrganizer().GetEmailAddress().GetAddress()
		}
		if event.GetOrganizer().GetEmailAddress().GetName() != nil {
			emailAddressInfo["name"] = *event.GetOrganizer().GetEmailAddress().GetName()
		}
		organizerInfo["emailAddress"] = emailAddressInfo
	}
	return organizerInfo
}

func (event *Microsoft365CalendarEventInfo) EventRecurrence() map[string]interface{} {
	if event.GetRecurrence() == nil {
		return nil
	}

	recurrenceInfo := map[string]interface{}{}

	if event.GetRecurrence().GetPattern() != nil {
		patternInfo := map[string]interface{}{
			"daysOfWeek": event.GetRecurrence().GetPattern().GetDaysOfWeek(),
		}
		if event.GetRecurrence().GetPattern().GetDayOfMonth() != nil {
			patternInfo["dayOfMonth"] = *event.GetRecurrence().GetPattern().GetDayOfMonth()
		}
		if event.GetRecurrence().GetPattern().GetFirstDayOfWeek() != nil {
			patternInfo["firstDayOfWeek"] = event.GetRecurrence().GetPattern().GetFirstDayOfWeek().String()
		}
		if event.GetRecurrence().GetPattern().GetIndex() != nil {
			patternInfo["index"] = event.GetRecurrence().GetPattern().GetIndex().String()
		}
		if event.GetRecurrence().GetPattern().GetInterval() != nil {
			patternInfo["interval"] = *event.GetRecurrence().GetPattern().GetInterval()
		}
		if event.GetRecurrence().GetPattern().GetMonth() != nil {
			patternInfo["month"] = *event.GetRecurrence().GetPattern().GetMonth()
		}
		if event.GetRecurrence().GetPattern().GetTypeEscaped() != nil {
			patternInfo["type"] = event.GetRecurrence().GetPattern().GetTypeEscaped().String()
		}
		recurrenceInfo["pattern"] = patternInfo
	}
	if event.GetRecurrence().GetRangeEscaped() != nil {
		recurrenceRange := event.GetRecurrence().GetRangeEscaped()
		rangeInfo := map[string]interface{}{}
		if recurrenceRange.GetTypeEscaped() != nil {
			rangeInfo["type"] = recurrenceRange.GetTypeEscaped().String()
		}
		if recurrenceRange.GetEndDate() != nil {
			rangeInfo["endDate"] = *recurrenceRange.GetEndDate()
		}
		if recurrenceRange.GetNumberOfOccurrences() != nil {
			rangeInfo["numberOfOccurrences"] = *recurrenceRange.GetNumberOfOccurrences()
		}
		if recurrenceRange.GetRecurrenceTimeZone() != nil {
			rangeInfo["recurrenceTimeZone"] = recurrenceRange.GetRecurrenceTimeZone()
		}
		if recurrenceRange.GetStartDate() != nil {
			rangeInfo["startDate"] = *recurrenceRange.GetStartDate()
		}
		recurrenceInfo["range"] = rangeInfo
	}

	return recurrenceInfo
}

func (event *Microsoft365CalendarEventInfo) EventResponseStatus() map[string]interface{} {
	if event.GetResponseStatus() == nil {
		return nil
	}

	responseStatusInfo := map[string]interface{}{}
	if event.GetResponseStatus().GetResponse() != nil {
		responseStatusInfo["response"] = event.GetResponseStatus().GetResponse().String()
	}

	if event.GetResponseStatus().GetTime() != nil {
		responseStatusInfo["time"] = *event.GetResponseStatus().GetTime()
	}
	return responseStatusInfo
}

func (event *Microsoft365CalendarEventInfo) EventStart() map[string]interface{} {
	if event.GetStart() == nil {
		return nil
	}

	startTimeInfo := map[string]interface{}{}
	if event.GetStart().GetDateTime() != nil {
		startTimeInfo["dateTime"] = *event.GetStart().GetDateTime()
	}
	if event.GetStart().GetDateTime() != nil {
		startTimeInfo["timeZone"] = *event.GetStart().GetTimeZone()
	}
	return startTimeInfo
}

func (message *Microsoft365MailMessageInfo) MessageAttachments() []map[string]interface{} {
	if message.GetAttachments() == nil {
		return nil
	}

	data := []map[string]interface{}{}
	for _, i := range message.GetAttachments() {
		attachmentInfo := map[string]interface{}{
			"lastModifiedDateTime": i.GetLastModifiedDateTime(),
		}
		if i.GetName() != nil {
			attachmentInfo["name"] = *i.GetName()
		}
		if i.GetContentType() != nil {
			attachmentInfo["contentType"] = *i.GetContentType()
		}
		if i.GetIsInline() != nil {
			attachmentInfo["isInline"] = *i.GetIsInline()
		}
		if i.GetSize() != nil {
			attachmentInfo["size"] = *i.GetSize()
		}
		data = append(data, attachmentInfo)
	}
	return data
}

func (message *Microsoft365MailMessageInfo) MessageBccRecipients() []map[string]interface{} {
	if message.GetBccRecipients() == nil {
		return nil
	}

	bccRecipients := []map[string]interface{}{}
	for _, i := range message.GetBccRecipients() {
		recipientInfo := map[string]interface{}{}
		if i.GetEmailAddress() != nil {
			addressInfo := map[string]interface{}{}
			if i.GetEmailAddress().GetAddress() != nil {
				addressInfo["address"] = i.GetEmailAddress().GetAddress()
			}
			if i.GetEmailAddress().GetName() != nil {
				addressInfo["name"] = i.GetEmailAddress().GetName()
			}
			recipientInfo["emailAddress"] = addressInfo
		}
		bccRecipients = append(bccRecipients, recipientInfo)
	}
	return bccRecipients
}

func (message *Microsoft365MailMessageInfo) MessageBody() map[string]interface{} {
	if message.GetBody() == nil {
		return nil
	}

	bodyInfo := map[string]interface{}{}
	if message.GetBody().GetContent() != nil {
		bodyInfo["content"] = *message.GetBody().GetContent()
	}
	if message.GetBody().GetContentType() != nil {
		bodyInfo["contentType"] = message.GetBody().GetContentType().String()
	}
	return bodyInfo
}

func (message *Microsoft365MailMessageInfo) MessageCcRecipients() []map[string]interface{} {
	if message.GetCcRecipients() == nil {
		return nil
	}

	ccRecipients := []map[string]interface{}{}
	for _, i := range message.GetCcRecipients() {
		recipientInfo := map[string]interface{}{}
		if i.GetEmailAddress() != nil {
			addressInfo := map[string]interface{}{}
			if i.GetEmailAddress().GetAddress() != nil {
				addressInfo["address"] = i.GetEmailAddress().GetAddress()
			}
			if i.GetEmailAddress().GetName() != nil {
				addressInfo["name"] = i.GetEmailAddress().GetName()
			}
			recipientInfo["emailAddress"] = addressInfo
		}
		ccRecipients = append(ccRecipients, recipientInfo)
	}
	return ccRecipients
}

func (message *Microsoft365MailMessageInfo) MessageFrom() map[string]interface{} {
	if message.GetFrom() == nil {
		return nil
	}
	fromInfo := map[string]interface{}{}
	if message.GetFrom().GetEmailAddress() != nil {
		addressInfo := map[string]interface{}{}
		if message.GetFrom().GetEmailAddress().GetAddress() != nil {
			addressInfo["address"] = message.GetFrom().GetEmailAddress().GetAddress()
		}
		if message.GetFrom().GetEmailAddress().GetName() != nil {
			addressInfo["name"] = message.GetFrom().GetEmailAddress().GetName()
		}
		fromInfo["emailAddress"] = addressInfo
	}
	return fromInfo
}

func (message *Microsoft365MailMessageInfo) MessageImportance() interface{} {
	if message.GetImportance() == nil {
		return nil
	}
	return message.GetImportance().String()
}

func (message *Microsoft365MailMessageInfo) MessageInferenceClassification() interface{} {
	if message.GetInferenceClassification() == nil {
		return nil
	}
	return message.GetInferenceClassification().String()
}

func (message *Microsoft365MailMessageInfo) MessageReplyTo() []map[string]interface{} {
	if message.GetReplyTo() == nil {
		return nil
	}

	data := []map[string]interface{}{}
	for _, i := range message.GetReplyTo() {
		replyToInfo := map[string]interface{}{}
		if i.GetEmailAddress() != nil {
			addressInfo := map[string]interface{}{}
			if i.GetEmailAddress().GetAddress() != nil {
				addressInfo["address"] = i.GetEmailAddress().GetAddress()
			}
			if i.GetEmailAddress().GetName() != nil {
				addressInfo["name"] = i.GetEmailAddress().GetName()
			}
			replyToInfo["emailAddress"] = addressInfo
		}
		data = append(data, replyToInfo)
	}
	return data
}

func (message *Microsoft365MailMessageInfo) MessageSender() map[string]interface{} {
	if message.GetSender() == nil {
		return nil
	}
	senderInfo := map[string]interface{}{}
	if message.GetSender().GetEmailAddress() != nil {
		addressInfo := map[string]interface{}{}
		if message.GetSender().GetEmailAddress().GetAddress() != nil {
			addressInfo["address"] = message.GetSender().GetEmailAddress().GetAddress()
		}
		if message.GetSender().GetEmailAddress().GetName() != nil {
			addressInfo["name"] = message.GetSender().GetEmailAddress().GetName()
		}
		senderInfo["emailAddress"] = addressInfo
	}
	return senderInfo
}

func (message *Microsoft365MailMessageInfo) MessageToRecipients() []map[string]interface{} {
	if message.GetToRecipients() == nil {
		return nil
	}

	recipients := []map[string]interface{}{}
	for _, i := range message.GetToRecipients() {
		recipientInfo := map[string]interface{}{}
		if i.GetEmailAddress() != nil {
			addressInfo := map[string]interface{}{}
			if i.GetEmailAddress().GetAddress() != nil {
				addressInfo["address"] = i.GetEmailAddress().GetAddress()
			}
			if i.GetEmailAddress().GetName() != nil {
				addressInfo["name"] = i.GetEmailAddress().GetName()
			}
			recipientInfo["emailAddress"] = addressInfo
		}
		recipients = append(recipients, recipientInfo)
	}
	return recipients
}

func (orgContact *Microsoft365OrgContactInfo) OrgContactAddresses() []map[string]interface{} {
	if orgContact.GetAddresses() == nil {
		return nil
	}

	addresses := []map[string]interface{}{}
	for _, i := range orgContact.GetAddresses() {
		data := map[string]interface{}{}
		if i.GetCity() != nil {
			data["city"] = *i.GetCity()
		}
		if i.GetCountryOrRegion() != nil {
			data["countryOrRegion"] = *i.GetCountryOrRegion()
		}
		if i.GetOfficeLocation() != nil {
			data["officeLocation"] = *i.GetOfficeLocation()
		}
		if i.GetPostalCode() != nil {
			data["postalCode"] = *i.GetPostalCode()
		}
		if i.GetState() != nil {
			data["state"] = *i.GetState()
		}
		if i.GetStreet() != nil {
			data["street"] = *i.GetStreet()
		}
		addresses = append(addresses, data)
	}
	return addresses
}

func (orgContact *Microsoft365OrgContactInfo) OrgContactDirectReports() []map[string]interface{} {
	if orgContact.GetDirectReports() == nil {
		return nil
	}

	directReports := []map[string]interface{}{}
	for _, i := range orgContact.GetDirectReports() {
		data := map[string]interface{}{}
		if i.GetId() != nil {
			data["id"] = *i.GetId()
		}
		if i.GetDeletedDateTime() != nil {
			data["deletedDateTime"] = *i.GetDeletedDateTime()
		}
		directReports = append(directReports, data)
	}
	return directReports
}

func (orgContact *Microsoft365OrgContactInfo) OrgContactManager() map[string]interface{} {
	if orgContact.GetManager() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if orgContact.GetManager().GetId() != nil {
		data["id"] = *orgContact.GetManager().GetId()
	}
	if orgContact.GetManager().GetDeletedDateTime() != nil {
		data["deletedDateTime"] = *orgContact.GetManager().GetDeletedDateTime()
	}
	return data
}

func (orgContact *Microsoft365OrgContactInfo) OrgContactMemberOf() []map[string]interface{} {
	if orgContact.GetMemberOf() == nil {
		return nil
	}

	memberOf := []map[string]interface{}{}
	for _, i := range orgContact.GetMemberOf() {
		data := map[string]interface{}{}
		if i.GetId() != nil {
			data["id"] = *i.GetId()
		}
		if i.GetDeletedDateTime() != nil {
			data["deletedDateTime"] = *i.GetDeletedDateTime()
		}
		memberOf = append(memberOf, data)
	}
	return memberOf
}

func (orgContact *Microsoft365OrgContactInfo) OrgContactOnPremisesProvisioningErrors() []map[string]interface{} {
	if orgContact.GetOnPremisesProvisioningErrors() == nil {
		return nil
	}

	provisioningErrors := []map[string]interface{}{}
	for _, i := range orgContact.GetOnPremisesProvisioningErrors() {
		data := map[string]interface{}{}
		if i.GetCategory() != nil {
			data["category"] = *i.GetCategory()
		}
		if i.GetOccurredDateTime() != nil {
			data["occurredDateTime"] = *i.GetOccurredDateTime()
		}
		if i.GetPropertyCausingError() != nil {
			data["propertyCausingError"] = *i.GetPropertyCausingError()
		}
		if i.GetValue() != nil {
			data["value"] = *i.GetValue()
		}
		provisioningErrors = append(provisioningErrors, data)
	}
	return provisioningErrors
}

func (orgContact *Microsoft365OrgContactInfo) OrgContactPhones() []map[string]interface{} {
	if orgContact.GetPhones() == nil {
		return nil
	}

	phones := []map[string]interface{}{}
	for _, i := range orgContact.GetPhones() {
		data := map[string]interface{}{}
		if i.GetLanguage() != nil {
			data["language"] = *i.GetLanguage()
		}
		if i.GetNumber() != nil {
			data["number"] = *i.GetNumber()
		}
		if i.GetRegion() != nil {
			data["region"] = *i.GetRegion()
		}
		if i.GetTypeEscaped() != nil {
			data["type"] = i.GetTypeEscaped().String()
		}
		phones = append(phones, data)
	}
	return phones
}

func (orgContact *Microsoft365OrgContactInfo) OrgContactTransitiveMemberOf() []map[string]interface{} {
	if orgContact.GetTransitiveMemberOf() == nil {
		return nil
	}

	memberOf := []map[string]interface{}{}
	for _, i := range orgContact.GetTransitiveMemberOf() {
		data := map[string]interface{}{}
		if i.GetId() != nil {
			data["id"] = *i.GetId()
		}
		if i.GetDeletedDateTime() != nil {
			data["deletedDateTime"] = *i.GetDeletedDateTime()
		}
		memberOf = append(memberOf, data)
	}
	return memberOf
}

func (team *Microsoft365TeamInfo) TeamMembers() interface{} {
	if team.GetSpecialization() == nil {
		return nil
	}
	return team.GetSpecialization().String()
}

func (team *Microsoft365TeamInfo) TeamSpecialization() interface{} {
	if team.GetSpecialization() == nil {
		return nil
	}
	return team.GetSpecialization().String()
}

func (team *Microsoft365TeamInfo) TeamSummary() map[string]interface{} {
	if team.GetSummary() == nil {
		return nil
	}

	summary := map[string]interface{}{}
	if team.GetSummary().GetGuestsCount() != nil {
		summary["guests_count"] = *team.GetSummary().GetGuestsCount()
	}
	if team.GetSummary().GetMembersCount() != nil {
		summary["members_count"] = *team.GetSummary().GetMembersCount()
	}
	if team.GetSummary().GetOwnersCount() != nil {
		summary["owners_count"] = *team.GetSummary().GetOwnersCount()
	}

	return summary
}

func (team *Microsoft365TeamInfo) TeamTemplate() map[string]interface{} {
	if team.GetTemplate() == nil {
		return nil
	}

	template := map[string]interface{}{}
	if team.GetTemplate().GetId() != nil {
		template["id"] = *team.GetTemplate().GetId()
	}
	if team.GetTemplate().GetOdataType() != nil {
		template["@odata_type"] = *team.GetTemplate().GetOdataType()
	}

	return template
}

func (team *Microsoft365TeamInfo) TeamVisibility() interface{} {
	if team.GetVisibility() == nil {
		return nil
	}
	return team.GetVisibility().String()
}

func (team *Microsoft365TeamChannelInfo) TeamChannelMembershipType() interface{} {
	if team.GetMembershipType() == nil {
		return nil
	}
	return team.GetMembershipType().String()
}

// Site transform methods
func (site *Microsoft365SiteInfo) SiteAnalytics() map[string]interface{} {
	if site.GetAnalytics() == nil {
		return nil
	}

	analytics := map[string]interface{}{}
	if site.GetAnalytics().GetAllTime() != nil {
		analytics["all_time"] = site.GetAnalytics().GetAllTime()
	}
	if site.GetAnalytics().GetItemActivityStats() != nil {
		analytics["item_activity_stats"] = site.GetAnalytics().GetItemActivityStats()
	}
	if site.GetAnalytics().GetLastSevenDays() != nil {
		analytics["last_seven_days"] = site.GetAnalytics().GetLastSevenDays()
	}

	return analytics
}

func (site *Microsoft365SiteInfo) SiteColumns() []map[string]interface{} {
	if site.GetColumns() == nil {
		return nil
	}

	var columns []map[string]interface{}
	for _, column := range site.GetColumns() {
		columnData := map[string]interface{}{}
		if column.GetId() != nil {
			columnData["id"] = *column.GetId()
		}
		if column.GetName() != nil {
			columnData["name"] = *column.GetName()
		}
		if column.GetColumnGroup() != nil {
			columnData["column_group"] = *column.GetColumnGroup()
		}
		if column.GetDescription() != nil {
			columnData["description"] = *column.GetDescription()
		}
		if column.GetDisplayName() != nil {
			columnData["display_name"] = *column.GetDisplayName()
		}
		if column.GetEnforceUniqueValues() != nil {
			columnData["enforce_unique_values"] = *column.GetEnforceUniqueValues()
		}
		if column.GetHidden() != nil {
			columnData["hidden"] = *column.GetHidden()
		}
		if column.GetIndexed() != nil {
			columnData["indexed"] = *column.GetIndexed()
		}
		if column.GetIsDeletable() != nil {
			columnData["is_deletable"] = *column.GetIsDeletable()
		}
		if column.GetIsReorderable() != nil {
			columnData["is_reorderable"] = *column.GetIsReorderable()
		}
		if column.GetIsSealed() != nil {
			columnData["is_sealed"] = *column.GetIsSealed()
		}
		if column.GetReadOnly() != nil {
			columnData["read_only"] = *column.GetReadOnly()
		}
		if column.GetRequired() != nil {
			columnData["required"] = *column.GetRequired()
		}
		columns = append(columns, columnData)
	}

	return columns
}

func (site *Microsoft365SiteInfo) SiteContentTypes() []map[string]interface{} {
	if site.GetContentTypes() == nil {
		return nil
	}

	var contentTypes []map[string]interface{}
	for _, contentType := range site.GetContentTypes() {
		contentTypeData := map[string]interface{}{}
		if contentType.GetId() != nil {
			contentTypeData["id"] = *contentType.GetId()
		}
		if contentType.GetName() != nil {
			contentTypeData["name"] = *contentType.GetName()
		}
		if contentType.GetDescription() != nil {
			contentTypeData["description"] = *contentType.GetDescription()
		}
		if contentType.GetGroup() != nil {
			contentTypeData["group"] = *contentType.GetGroup()
		}
		if contentType.GetHidden() != nil {
			contentTypeData["hidden"] = *contentType.GetHidden()
		}
		if contentType.GetInheritedFrom() != nil {
			contentTypeData["inherited_from"] = contentType.GetInheritedFrom()
		}
		if contentType.GetIsBuiltIn() != nil {
			contentTypeData["is_built_in"] = *contentType.GetIsBuiltIn()
		}
		contentTypes = append(contentTypes, contentTypeData)
	}

	return contentTypes
}

func (site *Microsoft365SiteInfo) SiteDrives() []map[string]interface{} {
	if site.GetDrives() == nil {
		return nil
	}

	var drives []map[string]interface{}
	for _, drive := range site.GetDrives() {
		driveData := map[string]interface{}{}
		if drive.GetId() != nil {
			driveData["id"] = *drive.GetId()
		}
		if drive.GetName() != nil {
			driveData["name"] = *drive.GetName()
		}
		if drive.GetDriveType() != nil {
			driveData["drive_type"] = *drive.GetDriveType()
		}
		if drive.GetDescription() != nil {
			driveData["description"] = *drive.GetDescription()
		}
		if drive.GetCreatedDateTime() != nil {
			driveData["created_date_time"] = *drive.GetCreatedDateTime()
		}
		if drive.GetLastModifiedDateTime() != nil {
			driveData["last_modified_date_time"] = *drive.GetLastModifiedDateTime()
		}
		if drive.GetWebUrl() != nil {
			driveData["web_url"] = *drive.GetWebUrl()
		}
		drives = append(drives, driveData)
	}

	return drives
}

func (site *Microsoft365SiteInfo) SiteLists() []map[string]interface{} {
	if site.GetLists() == nil {
		return nil
	}

	var lists []map[string]interface{}
	for _, list := range site.GetLists() {
		listData := map[string]interface{}{}
		if list.GetId() != nil {
			listData["id"] = *list.GetId()
		}
		if list.GetName() != nil {
			listData["name"] = *list.GetName()
		}
		if list.GetDisplayName() != nil {
			listData["display_name"] = *list.GetDisplayName()
		}
		if list.GetDescription() != nil {
			listData["description"] = *list.GetDescription()
		}
		if list.GetCreatedDateTime() != nil {
			listData["created_date_time"] = *list.GetCreatedDateTime()
		}
		if list.GetLastModifiedDateTime() != nil {
			listData["last_modified_date_time"] = *list.GetLastModifiedDateTime()
		}
		if list.GetWebUrl() != nil {
			listData["web_url"] = *list.GetWebUrl()
		}
		lists = append(lists, listData)
	}

	return lists
}

func (site *Microsoft365SiteInfo) SiteOperations() []map[string]interface{} {
	if site.GetOperations() == nil {
		return nil
	}

	var operations []map[string]interface{}
	for _, operation := range site.GetOperations() {
		operationData := map[string]interface{}{}
		if operation.GetId() != nil {
			operationData["id"] = *operation.GetId()
		}
		if operation.GetStatus() != nil {
			operationData["status"] = operation.GetStatus().String()
		}
		if operation.GetCreatedDateTime() != nil {
			operationData["created_date_time"] = *operation.GetCreatedDateTime()
		}
		if operation.GetLastActionDateTime() != nil {
			operationData["last_action_date_time"] = *operation.GetLastActionDateTime()
		}
		if operation.GetPercentageComplete() != nil {
			operationData["percentage_complete"] = *operation.GetPercentageComplete()
		}
		if operation.GetResourceId() != nil {
			operationData["resource_id"] = *operation.GetResourceId()
		}
		if operation.GetResourceLocation() != nil {
			operationData["resource_location"] = *operation.GetResourceLocation()
		}
		if operation.GetTypeEscaped() != nil {
			operationData["type"] = *operation.GetTypeEscaped()
		}
		operations = append(operations, operationData)
	}

	return operations
}

func (site *Microsoft365SiteInfo) SitePages() []map[string]interface{} {
	if site.GetPages() == nil {
		return nil
	}

	var pages []map[string]interface{}
	for _, page := range site.GetPages() {
		pageData := map[string]interface{}{}
		if page.GetId() != nil {
			pageData["id"] = *page.GetId()
		}
		if page.GetName() != nil {
			pageData["name"] = *page.GetName()
		}
		if page.GetTitle() != nil {
			pageData["title"] = *page.GetTitle()
		}
		if page.GetWebUrl() != nil {
			pageData["web_url"] = *page.GetWebUrl()
		}
		if page.GetCreatedDateTime() != nil {
			pageData["created_date_time"] = *page.GetCreatedDateTime()
		}
		if page.GetLastModifiedDateTime() != nil {
			pageData["last_modified_date_time"] = *page.GetLastModifiedDateTime()
		}
		pages = append(pages, pageData)
	}

	return pages
}

func (site *Microsoft365SiteInfo) SitePermissions() []map[string]interface{} {
	if site.GetPermissions() == nil {
		return nil
	}

	var permissions []map[string]interface{}
	for _, permission := range site.GetPermissions() {
		permissionData := map[string]interface{}{}
		if permission.GetId() != nil {
			permissionData["id"] = *permission.GetId()
		}
		if permission.GetRoles() != nil {
			permissionData["roles"] = permission.GetRoles()
		}
		if permission.GetGrantedTo() != nil {
			permissionData["granted_to"] = permission.GetGrantedTo()
		}
		if permission.GetGrantedToIdentities() != nil {
			permissionData["granted_to_identities"] = permission.GetGrantedToIdentities()
		}
		if permission.GetInheritedFrom() != nil {
			permissionData["inherited_from"] = permission.GetInheritedFrom()
		}
		if permission.GetLink() != nil {
			permissionData["link"] = permission.GetLink()
		}
		if permission.GetShareId() != nil {
			permissionData["share_id"] = *permission.GetShareId()
		}
		permissions = append(permissions, permissionData)
	}

	return permissions
}

func (site *Microsoft365SiteInfo) SiteSites() []map[string]interface{} {
	if site.GetSites() == nil {
		return nil
	}

	var sites []map[string]interface{}
	for _, subSite := range site.GetSites() {
		siteData := map[string]interface{}{}
		if subSite.GetId() != nil {
			siteData["id"] = *subSite.GetId()
		}
		if subSite.GetName() != nil {
			siteData["name"] = *subSite.GetName()
		}
		if subSite.GetDisplayName() != nil {
			siteData["display_name"] = *subSite.GetDisplayName()
		}
		if subSite.GetDescription() != nil {
			siteData["description"] = *subSite.GetDescription()
		}
		if subSite.GetWebUrl() != nil {
			siteData["web_url"] = *subSite.GetWebUrl()
		}
		if subSite.GetCreatedDateTime() != nil {
			siteData["created_date_time"] = *subSite.GetCreatedDateTime()
		}
		if subSite.GetLastModifiedDateTime() != nil {
			siteData["last_modified_date_time"] = *subSite.GetLastModifiedDateTime()
		}
		if subSite.GetIsPersonalSite() != nil {
			siteData["is_personal_site"] = *subSite.GetIsPersonalSite()
		}
		sites = append(sites, siteData)
	}

	return sites
}

// Additional site transform methods
func (site *Microsoft365SiteInfo) SiteCreatedBy() map[string]interface{} {
	if site.GetCreatedBy() == nil {
		return nil
	}

	createdBy := map[string]interface{}{}
	if site.GetCreatedBy().GetUser() != nil {
		user := map[string]interface{}{}
		if site.GetCreatedBy().GetUser().GetId() != nil {
			user["id"] = *site.GetCreatedBy().GetUser().GetId()
		}
		if site.GetCreatedBy().GetUser().GetDisplayName() != nil {
			user["display_name"] = *site.GetCreatedBy().GetUser().GetDisplayName()
		}
		// Note: Identityable doesn't have GetUserPrincipalName method
		createdBy["user"] = user
	}
	if site.GetCreatedBy().GetApplication() != nil {
		app := map[string]interface{}{}
		if site.GetCreatedBy().GetApplication().GetId() != nil {
			app["id"] = *site.GetCreatedBy().GetApplication().GetId()
		}
		if site.GetCreatedBy().GetApplication().GetDisplayName() != nil {
			app["display_name"] = *site.GetCreatedBy().GetApplication().GetDisplayName()
		}
		createdBy["application"] = app
	}
	if site.GetCreatedBy().GetDevice() != nil {
		device := map[string]interface{}{}
		if site.GetCreatedBy().GetDevice().GetId() != nil {
			device["id"] = *site.GetCreatedBy().GetDevice().GetId()
		}
		if site.GetCreatedBy().GetDevice().GetDisplayName() != nil {
			device["display_name"] = *site.GetCreatedBy().GetDevice().GetDisplayName()
		}
		createdBy["device"] = device
	}

	return createdBy
}

func (site *Microsoft365SiteInfo) SiteLastModifiedBy() map[string]interface{} {
	if site.GetLastModifiedBy() == nil {
		return nil
	}

	lastModifiedBy := map[string]interface{}{}
	if site.GetLastModifiedBy().GetUser() != nil {
		user := map[string]interface{}{}
		if site.GetLastModifiedBy().GetUser().GetId() != nil {
			user["id"] = *site.GetLastModifiedBy().GetUser().GetId()
		}
		if site.GetLastModifiedBy().GetUser().GetDisplayName() != nil {
			user["display_name"] = *site.GetLastModifiedBy().GetUser().GetDisplayName()
		}
		// Note: Identityable doesn't have GetUserPrincipalName method
		lastModifiedBy["user"] = user
	}
	if site.GetLastModifiedBy().GetApplication() != nil {
		app := map[string]interface{}{}
		if site.GetLastModifiedBy().GetApplication().GetId() != nil {
			app["id"] = *site.GetLastModifiedBy().GetApplication().GetId()
		}
		if site.GetLastModifiedBy().GetApplication().GetDisplayName() != nil {
			app["display_name"] = *site.GetLastModifiedBy().GetApplication().GetDisplayName()
		}
		lastModifiedBy["application"] = app
	}
	if site.GetLastModifiedBy().GetDevice() != nil {
		device := map[string]interface{}{}
		if site.GetLastModifiedBy().GetDevice().GetId() != nil {
			device["id"] = *site.GetLastModifiedBy().GetDevice().GetId()
		}
		if site.GetLastModifiedBy().GetDevice().GetDisplayName() != nil {
			device["display_name"] = *site.GetLastModifiedBy().GetDevice().GetDisplayName()
		}
		lastModifiedBy["device"] = device
	}

	return lastModifiedBy
}

func (site *Microsoft365SiteInfo) SiteParentReference() map[string]interface{} {
	if site.GetParentReference() == nil {
		return nil
	}

	parentRef := map[string]interface{}{}
	if site.GetParentReference().GetDriveId() != nil {
		parentRef["drive_id"] = *site.GetParentReference().GetDriveId()
	}
	if site.GetParentReference().GetDriveType() != nil {
		parentRef["drive_type"] = *site.GetParentReference().GetDriveType()
	}
	if site.GetParentReference().GetId() != nil {
		parentRef["id"] = *site.GetParentReference().GetId()
	}
	if site.GetParentReference().GetName() != nil {
		parentRef["name"] = *site.GetParentReference().GetName()
	}
	if site.GetParentReference().GetPath() != nil {
		parentRef["path"] = *site.GetParentReference().GetPath()
	}
	if site.GetParentReference().GetSharepointIds() != nil {
		parentRef["sharepoint_ids"] = site.GetParentReference().GetSharepointIds()
	}
	if site.GetParentReference().GetSiteId() != nil {
		parentRef["site_id"] = *site.GetParentReference().GetSiteId()
	}

	return parentRef
}

func (site *Microsoft365SiteInfo) SiteSharepointIds() map[string]interface{} {
	if site.GetSharepointIds() == nil {
		return nil
	}

	sharepointIds := map[string]interface{}{}
	if site.GetSharepointIds().GetListId() != nil {
		sharepointIds["list_id"] = *site.GetSharepointIds().GetListId()
	}
	if site.GetSharepointIds().GetListItemId() != nil {
		sharepointIds["list_item_id"] = *site.GetSharepointIds().GetListItemId()
	}
	if site.GetSharepointIds().GetListItemUniqueId() != nil {
		sharepointIds["list_item_unique_id"] = *site.GetSharepointIds().GetListItemUniqueId()
	}
	if site.GetSharepointIds().GetSiteId() != nil {
		sharepointIds["site_id"] = *site.GetSharepointIds().GetSiteId()
	}
	if site.GetSharepointIds().GetSiteUrl() != nil {
		sharepointIds["site_url"] = *site.GetSharepointIds().GetSiteUrl()
	}
	if site.GetSharepointIds().GetTenantId() != nil {
		sharepointIds["tenant_id"] = *site.GetSharepointIds().GetTenantId()
	}
	if site.GetSharepointIds().GetWebId() != nil {
		sharepointIds["web_id"] = *site.GetSharepointIds().GetWebId()
	}

	return sharepointIds
}

func (site *Microsoft365SiteInfo) SiteSiteCollection() map[string]interface{} {
	if site.GetSiteCollection() == nil {
		return nil
	}

	siteCollection := map[string]interface{}{}
	if site.GetSiteCollection().GetHostname() != nil {
		siteCollection["hostname"] = *site.GetSiteCollection().GetHostname()
	}
	if site.GetSiteCollection().GetRoot() != nil {
		siteCollection["root"] = site.GetSiteCollection().GetRoot()
	}

	return siteCollection
}

func (site *Microsoft365SiteInfo) SiteRoot() map[string]interface{} {
	if site.GetRoot() == nil {
		return nil
	}

	// Note: Rootable interface only has GetOdataType method
	// Return basic structure for now
	root := map[string]interface{}{
		"odata_type": site.GetRoot().GetOdataType(),
	}

	return root
}

type Microsoft365ListInfo struct {
	models.Listable
	SiteID string
}

type Microsoft365OrganizationInfo struct {
	models.Organizationable
}

type Microsoft365MailSettingsInfo struct {
	models.MailboxSettingsable
	UserID string
}

type Microsoft365SecuritySettingsInfo struct {
	PolicyData interface{} // Can be various policy types
	PolicyType string
}

type Microsoft365CalendarSettingsInfo struct {
	models.MailboxSettingsable
	UserID    string
	Calendars models.CalendarCollectionResponseable
}

type Microsoft365SharePointSettingsInfo struct {
	models.SharepointSettingsable
}

type Microsoft365AuthenticationSettingsInfo struct {
	models.AuthenticationMethodsPolicyable
}

type Microsoft365SecurityDefaultsSettingsInfo struct {
	models.IdentitySecurityDefaultsEnforcementPolicyable
}

type Microsoft365SiteInfo struct {
	models.Siteable
}

type Microsoft365UserInfo struct {
	models.Userable
	MailboxSettings models.MailboxSettingsable
}

// List transform methods
func (list *Microsoft365ListInfo) ListCreatedBy() map[string]interface{} {
	if list.GetCreatedBy() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if list.GetCreatedBy().GetUser() != nil {
		data["user"] = list.GetCreatedBy().GetUser()
	}
	if list.GetCreatedBy().GetApplication() != nil {
		data["application"] = list.GetCreatedBy().GetApplication()
	}
	if list.GetCreatedBy().GetDevice() != nil {
		data["device"] = list.GetCreatedBy().GetDevice()
	}

	return data
}

func (list *Microsoft365ListInfo) ListLastModifiedBy() map[string]interface{} {
	if list.GetLastModifiedBy() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if list.GetLastModifiedBy().GetUser() != nil {
		data["user"] = list.GetLastModifiedBy().GetUser()
	}
	if list.GetLastModifiedBy().GetApplication() != nil {
		data["application"] = list.GetLastModifiedBy().GetApplication()
	}
	if list.GetLastModifiedBy().GetDevice() != nil {
		data["device"] = list.GetLastModifiedBy().GetDevice()
	}

	return data
}

func (list *Microsoft365ListInfo) ListParentReference() map[string]interface{} {
	if list.GetParentReference() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if list.GetParentReference().GetId() != nil {
		data["id"] = *list.GetParentReference().GetId()
	}
	if list.GetParentReference().GetPath() != nil {
		data["path"] = *list.GetParentReference().GetPath()
	}
	if list.GetParentReference().GetDriveId() != nil {
		data["drive_id"] = *list.GetParentReference().GetDriveId()
	}
	if list.GetParentReference().GetDriveType() != nil {
		data["drive_type"] = *list.GetParentReference().GetDriveType()
	}
	if list.GetParentReference().GetSharepointIds() != nil {
		data["sharepoint_ids"] = list.GetParentReference().GetSharepointIds()
	}

	return data
}

func (list *Microsoft365ListInfo) ListSharepointIds() map[string]interface{} {
	if list.GetSharepointIds() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if list.GetSharepointIds().GetListId() != nil {
		data["list_id"] = *list.GetSharepointIds().GetListId()
	}
	if list.GetSharepointIds().GetListItemId() != nil {
		data["list_item_id"] = *list.GetSharepointIds().GetListItemId()
	}
	if list.GetSharepointIds().GetListItemUniqueId() != nil {
		data["list_item_unique_id"] = *list.GetSharepointIds().GetListItemUniqueId()
	}
	if list.GetSharepointIds().GetSiteId() != nil {
		data["site_id"] = *list.GetSharepointIds().GetSiteId()
	}
	if list.GetSharepointIds().GetSiteUrl() != nil {
		data["site_url"] = *list.GetSharepointIds().GetSiteUrl()
	}
	if list.GetSharepointIds().GetTenantId() != nil {
		data["tenant_id"] = *list.GetSharepointIds().GetTenantId()
	}
	if list.GetSharepointIds().GetWebId() != nil {
		data["web_id"] = *list.GetSharepointIds().GetWebId()
	}

	return data
}

func (list *Microsoft365ListInfo) ListColumns() []map[string]interface{} {
	if list.GetColumns() == nil {
		return nil
	}

	var columns []map[string]interface{}
	for _, column := range list.GetColumns() {
		columnData := map[string]interface{}{}
		if column.GetId() != nil {
			columnData["id"] = *column.GetId()
		}
		if column.GetName() != nil {
			columnData["name"] = *column.GetName()
		}
		if column.GetColumnGroup() != nil {
			columnData["column_group"] = *column.GetColumnGroup()
		}
		if column.GetDescription() != nil {
			columnData["description"] = *column.GetDescription()
		}
		if column.GetDisplayName() != nil {
			columnData["display_name"] = *column.GetDisplayName()
		}
		if column.GetEnforceUniqueValues() != nil {
			columnData["enforce_unique_values"] = *column.GetEnforceUniqueValues()
		}
		if column.GetHidden() != nil {
			columnData["hidden"] = *column.GetHidden()
		}
		if column.GetIndexed() != nil {
			columnData["indexed"] = *column.GetIndexed()
		}
		if column.GetIsDeletable() != nil {
			columnData["is_deletable"] = *column.GetIsDeletable()
		}
		if column.GetIsReorderable() != nil {
			columnData["is_reorderable"] = *column.GetIsReorderable()
		}
		if column.GetIsSealed() != nil {
			columnData["is_sealed"] = *column.GetIsSealed()
		}
		if column.GetReadOnly() != nil {
			columnData["read_only"] = *column.GetReadOnly()
		}
		if column.GetRequired() != nil {
			columnData["required"] = *column.GetRequired()
		}
		columns = append(columns, columnData)
	}

	return columns
}

func (list *Microsoft365ListInfo) ListContentTypes() []map[string]interface{} {
	if list.GetContentTypes() == nil {
		return nil
	}

	var contentTypes []map[string]interface{}
	for _, contentType := range list.GetContentTypes() {
		contentTypeData := map[string]interface{}{}
		if contentType.GetId() != nil {
			contentTypeData["id"] = *contentType.GetId()
		}
		if contentType.GetName() != nil {
			contentTypeData["name"] = *contentType.GetName()
		}
		if contentType.GetDescription() != nil {
			contentTypeData["description"] = *contentType.GetDescription()
		}
		if contentType.GetGroup() != nil {
			contentTypeData["group"] = *contentType.GetGroup()
		}
		if contentType.GetHidden() != nil {
			contentTypeData["hidden"] = *contentType.GetHidden()
		}
		if contentType.GetInheritedFrom() != nil {
			contentTypeData["inherited_from"] = contentType.GetInheritedFrom()
		}
		if contentType.GetIsBuiltIn() != nil {
			contentTypeData["is_built_in"] = *contentType.GetIsBuiltIn()
		}
		contentTypes = append(contentTypes, contentTypeData)
	}

	return contentTypes
}

func (list *Microsoft365ListInfo) ListDrive() map[string]interface{} {
	if list.GetDrive() == nil {
		return nil
	}

	driveData := map[string]interface{}{}
	if list.GetDrive().GetId() != nil {
		driveData["id"] = *list.GetDrive().GetId()
	}
	if list.GetDrive().GetName() != nil {
		driveData["name"] = *list.GetDrive().GetName()
	}
	if list.GetDrive().GetDriveType() != nil {
		driveData["drive_type"] = *list.GetDrive().GetDriveType()
	}
	if list.GetDrive().GetDescription() != nil {
		driveData["description"] = *list.GetDrive().GetDescription()
	}
	if list.GetDrive().GetCreatedDateTime() != nil {
		driveData["created_date_time"] = *list.GetDrive().GetCreatedDateTime()
	}
	if list.GetDrive().GetLastModifiedDateTime() != nil {
		driveData["last_modified_date_time"] = *list.GetDrive().GetLastModifiedDateTime()
	}
	if list.GetDrive().GetWebUrl() != nil {
		driveData["web_url"] = *list.GetDrive().GetWebUrl()
	}

	return driveData
}

func (list *Microsoft365ListInfo) ListItems() []map[string]interface{} {
	if list.GetItems() == nil {
		return nil
	}

	var items []map[string]interface{}
	for _, item := range list.GetItems() {
		itemData := map[string]interface{}{}
		if item.GetId() != nil {
			itemData["id"] = *item.GetId()
		}
		if item.GetName() != nil {
			itemData["name"] = *item.GetName()
		}
		// Note: ListItemable doesn't have GetTitle method, using name instead
		if item.GetName() != nil {
			itemData["title"] = *item.GetName()
		}
		if item.GetWebUrl() != nil {
			itemData["web_url"] = *item.GetWebUrl()
		}
		if item.GetCreatedDateTime() != nil {
			itemData["created_date_time"] = *item.GetCreatedDateTime()
		}
		if item.GetLastModifiedDateTime() != nil {
			itemData["last_modified_date_time"] = *item.GetLastModifiedDateTime()
		}
		if item.GetCreatedBy() != nil {
			itemData["created_by"] = item.GetCreatedBy()
		}
		if item.GetLastModifiedBy() != nil {
			itemData["last_modified_by"] = item.GetLastModifiedBy()
		}
		items = append(items, itemData)
	}

	return items
}

func (list *Microsoft365ListInfo) ListInfo() map[string]interface{} {
	if list.GetList() == nil {
		return nil
	}

	info := map[string]interface{}{}
	if list.GetList().GetTemplate() != nil {
		info["template"] = *list.GetList().GetTemplate()
	}

	return info
}

func (list *Microsoft365ListInfo) ListOperations() []map[string]interface{} {
	if list.GetOperations() == nil {
		return nil
	}

	var operations []map[string]interface{}
	for _, operation := range list.GetOperations() {
		operationData := map[string]interface{}{}
		if operation.GetId() != nil {
			operationData["id"] = *operation.GetId()
		}
		if operation.GetStatus() != nil {
			operationData["status"] = operation.GetStatus().String()
		}
		if operation.GetCreatedDateTime() != nil {
			operationData["created_date_time"] = *operation.GetCreatedDateTime()
		}
		if operation.GetLastActionDateTime() != nil {
			operationData["last_action_date_time"] = *operation.GetLastActionDateTime()
		}
		if operation.GetPercentageComplete() != nil {
			operationData["percentage_complete"] = *operation.GetPercentageComplete()
		}
		if operation.GetResourceId() != nil {
			operationData["resource_id"] = *operation.GetResourceId()
		}
		if operation.GetResourceLocation() != nil {
			operationData["resource_location"] = *operation.GetResourceLocation()
		}
		if operation.GetTypeEscaped() != nil {
			operationData["type"] = *operation.GetTypeEscaped()
		}
		operations = append(operations, operationData)
	}

	return operations
}

func (list *Microsoft365ListInfo) ListSubscriptions() []map[string]interface{} {
	if list.GetSubscriptions() == nil {
		return nil
	}

	var subscriptions []map[string]interface{}
	for _, subscription := range list.GetSubscriptions() {
		subscriptionData := map[string]interface{}{}
		if subscription.GetId() != nil {
			subscriptionData["id"] = *subscription.GetId()
		}
		if subscription.GetApplicationId() != nil {
			subscriptionData["application_id"] = *subscription.GetApplicationId()
		}
		if subscription.GetChangeType() != nil {
			subscriptionData["change_type"] = *subscription.GetChangeType()
		}
		if subscription.GetClientState() != nil {
			subscriptionData["client_state"] = *subscription.GetClientState()
		}
		if subscription.GetCreatorId() != nil {
			subscriptionData["creator_id"] = *subscription.GetCreatorId()
		}
		if subscription.GetEncryptionCertificate() != nil {
			subscriptionData["encryption_certificate"] = *subscription.GetEncryptionCertificate()
		}
		if subscription.GetEncryptionCertificateId() != nil {
			subscriptionData["encryption_certificate_id"] = *subscription.GetEncryptionCertificateId()
		}
		if subscription.GetExpirationDateTime() != nil {
			subscriptionData["expiration_date_time"] = *subscription.GetExpirationDateTime()
		}
		if subscription.GetIncludeResourceData() != nil {
			subscriptionData["include_resource_data"] = *subscription.GetIncludeResourceData()
		}
		if subscription.GetLatestSupportedTlsVersion() != nil {
			subscriptionData["latest_supported_tls_version"] = *subscription.GetLatestSupportedTlsVersion()
		}
		if subscription.GetLifecycleNotificationUrl() != nil {
			subscriptionData["lifecycle_notification_url"] = *subscription.GetLifecycleNotificationUrl()
		}
		if subscription.GetNotificationUrl() != nil {
			subscriptionData["notification_url"] = *subscription.GetNotificationUrl()
		}
		if subscription.GetResource() != nil {
			subscriptionData["resource"] = *subscription.GetResource()
		}
		subscriptions = append(subscriptions, subscriptionData)
	}

	return subscriptions
}

func (list *Microsoft365ListInfo) ListSystem() map[string]interface{} {
	if list.GetSystem() == nil {
		return nil
	}

	system := map[string]interface{}{}
	if list.GetSystem().GetOdataType() != nil {
		system["@odata_type"] = *list.GetSystem().GetOdataType()
	}

	return system
}

// Drive transform methods for SharePoint-specific fields
func (drive *Microsoft365DriveInfo) DriveSharepointIds() map[string]interface{} {
	if drive.GetSharePointIds() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if drive.GetSharePointIds().GetListId() != nil {
		data["list_id"] = *drive.GetSharePointIds().GetListId()
	}
	if drive.GetSharePointIds().GetListItemId() != nil {
		data["list_item_id"] = *drive.GetSharePointIds().GetListItemId()
	}
	if drive.GetSharePointIds().GetListItemUniqueId() != nil {
		data["list_item_unique_id"] = *drive.GetSharePointIds().GetListItemUniqueId()
	}
	if drive.GetSharePointIds().GetSiteId() != nil {
		data["site_id"] = *drive.GetSharePointIds().GetSiteId()
	}
	if drive.GetSharePointIds().GetSiteUrl() != nil {
		data["site_url"] = *drive.GetSharePointIds().GetSiteUrl()
	}
	if drive.GetSharePointIds().GetTenantId() != nil {
		data["tenant_id"] = *drive.GetSharePointIds().GetTenantId()
	}
	if drive.GetSharePointIds().GetWebId() != nil {
		data["web_id"] = *drive.GetSharePointIds().GetWebId()
	}

	return data
}

func (drive *Microsoft365DriveInfo) DriveList() map[string]interface{} {
	if drive.GetList() == nil {
		return nil
	}

	listData := map[string]interface{}{}
	if drive.GetList().GetId() != nil {
		listData["id"] = *drive.GetList().GetId()
	}
	if drive.GetList().GetName() != nil {
		listData["name"] = *drive.GetList().GetName()
	}
	if drive.GetList().GetDisplayName() != nil {
		listData["display_name"] = *drive.GetList().GetDisplayName()
	}
	if drive.GetList().GetDescription() != nil {
		listData["description"] = *drive.GetList().GetDescription()
	}
	if drive.GetList().GetCreatedDateTime() != nil {
		listData["created_date_time"] = *drive.GetList().GetCreatedDateTime()
	}
	if drive.GetList().GetLastModifiedDateTime() != nil {
		listData["last_modified_date_time"] = *drive.GetList().GetLastModifiedDateTime()
	}
	if drive.GetList().GetWebUrl() != nil {
		listData["web_url"] = *drive.GetList().GetWebUrl()
	}

	return listData
}

func (drive *Microsoft365DriveInfo) DriveOwner() map[string]interface{} {
	if drive.GetOwner() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if drive.GetOwner().GetUser() != nil {
		data["user"] = drive.GetOwner().GetUser()
	}
	if drive.GetOwner().GetApplication() != nil {
		data["application"] = drive.GetOwner().GetApplication()
	}
	if drive.GetOwner().GetDevice() != nil {
		data["device"] = drive.GetOwner().GetDevice()
	}

	return data
}

func (drive *Microsoft365DriveInfo) DriveQuota() map[string]interface{} {
	if drive.GetQuota() == nil {
		return nil
	}

	data := map[string]interface{}{}
	if drive.GetQuota().GetDeleted() != nil {
		data["deleted"] = *drive.GetQuota().GetDeleted()
	}
	if drive.GetQuota().GetRemaining() != nil {
		data["remaining"] = *drive.GetQuota().GetRemaining()
	}
	if drive.GetQuota().GetState() != nil {
		data["state"] = *drive.GetQuota().GetState()
	}
	if drive.GetQuota().GetStoragePlanInformation() != nil {
		data["storage_plan_information"] = drive.GetQuota().GetStoragePlanInformation()
	}
	if drive.GetQuota().GetTotal() != nil {
		data["total"] = *drive.GetQuota().GetTotal()
	}
	if drive.GetQuota().GetUsed() != nil {
		data["used"] = *drive.GetQuota().GetUsed()
	}

	return data
}

func (drive *Microsoft365DriveInfo) DriveSystem() map[string]interface{} {
	if drive.GetSystem() == nil {
		return nil
	}

	system := map[string]interface{}{}
	if drive.GetSystem().GetOdataType() != nil {
		system["@odata_type"] = *drive.GetSystem().GetOdataType()
	}

	return system
}

// Organization transform methods
func (org *Microsoft365OrganizationInfo) OrganizationAssignedPlans() []map[string]interface{} {
	if org.GetAssignedPlans() == nil {
		return nil
	}

	var plans []map[string]interface{}
	for _, plan := range org.GetAssignedPlans() {
		planData := map[string]interface{}{}
		if plan.GetAssignedDateTime() != nil {
			planData["assigned_date_time"] = *plan.GetAssignedDateTime()
		}
		if plan.GetCapabilityStatus() != nil {
			planData["capability_status"] = *plan.GetCapabilityStatus()
		}
		if plan.GetService() != nil {
			planData["service"] = *plan.GetService()
		}
		if plan.GetServicePlanId() != nil {
			planData["service_plan_id"] = *plan.GetServicePlanId()
		}
		plans = append(plans, planData)
	}

	return plans
}

func (org *Microsoft365OrganizationInfo) OrganizationBranding() map[string]interface{} {
	if org.GetBranding() == nil {
		return nil
	}

	branding := map[string]interface{}{}
	if org.GetBranding().GetBackgroundColor() != nil {
		branding["background_color"] = *org.GetBranding().GetBackgroundColor()
	}
	if org.GetBranding().GetBackgroundImage() != nil {
		branding["background_image"] = org.GetBranding().GetBackgroundImage()
	}
	if org.GetBranding().GetBannerLogo() != nil {
		branding["banner_logo"] = org.GetBranding().GetBannerLogo()
	}
	if org.GetBranding().GetSignInPageText() != nil {
		branding["sign_in_page_text"] = *org.GetBranding().GetSignInPageText()
	}
	if org.GetBranding().GetSquareLogo() != nil {
		branding["square_logo"] = org.GetBranding().GetSquareLogo()
	}
	if org.GetBranding().GetUsernameHintText() != nil {
		branding["username_hint_text"] = *org.GetBranding().GetUsernameHintText()
	}

	return branding
}

func (org *Microsoft365OrganizationInfo) OrganizationBusinessPhones() []string {
	if org.GetBusinessPhones() == nil {
		return nil
	}
	return org.GetBusinessPhones()
}

func (org *Microsoft365OrganizationInfo) OrganizationCertificateBasedAuthConfiguration() []map[string]interface{} {
	if org.GetCertificateBasedAuthConfiguration() == nil {
		return nil
	}

	var configs []map[string]interface{}
	for _, config := range org.GetCertificateBasedAuthConfiguration() {
		configData := map[string]interface{}{}
		if config.GetId() != nil {
			configData["id"] = *config.GetId()
		}
		if config.GetCertificateAuthorities() != nil {
			configData["certificate_authorities"] = config.GetCertificateAuthorities()
		}
		configs = append(configs, configData)
	}

	return configs
}

func (org *Microsoft365OrganizationInfo) OrganizationExtensions() []map[string]interface{} {
	if org.GetExtensions() == nil {
		return nil
	}

	var extensions []map[string]interface{}
	for _, extension := range org.GetExtensions() {
		extensionData := map[string]interface{}{}
		if extension.GetId() != nil {
			extensionData["id"] = *extension.GetId()
		}
		if extension.GetOdataType() != nil {
			extensionData["@odata_type"] = *extension.GetOdataType()
		}
		extensions = append(extensions, extensionData)
	}

	return extensions
}

func (org *Microsoft365OrganizationInfo) OrganizationMarketingNotificationEmails() []string {
	if org.GetMarketingNotificationEmails() == nil {
		return nil
	}
	return org.GetMarketingNotificationEmails()
}

func (org *Microsoft365OrganizationInfo) OrganizationMobileDeviceManagementAuthority() interface{} {
	if org.GetMobileDeviceManagementAuthority() == nil {
		return nil
	}
	return org.GetMobileDeviceManagementAuthority().String()
}

func (org *Microsoft365OrganizationInfo) OrganizationPartnerTenantType() interface{} {
	if org.GetPartnerTenantType() == nil {
		return nil
	}
	return org.GetPartnerTenantType().String()
}

func (org *Microsoft365OrganizationInfo) OrganizationPrivacyProfile() map[string]interface{} {
	if org.GetPrivacyProfile() == nil {
		return nil
	}

	profile := map[string]interface{}{}
	if org.GetPrivacyProfile().GetContactEmail() != nil {
		profile["contact_email"] = *org.GetPrivacyProfile().GetContactEmail()
	}
	if org.GetPrivacyProfile().GetStatementUrl() != nil {
		profile["statement_url"] = *org.GetPrivacyProfile().GetStatementUrl()
	}
	// Note: PrivacyProfile doesn't have GetUrl method, only GetStatementUrl

	return profile
}

func (org *Microsoft365OrganizationInfo) OrganizationProvisionedPlans() []map[string]interface{} {
	if org.GetProvisionedPlans() == nil {
		return nil
	}

	var plans []map[string]interface{}
	for _, plan := range org.GetProvisionedPlans() {
		planData := map[string]interface{}{}
		if plan.GetCapabilityStatus() != nil {
			planData["capability_status"] = *plan.GetCapabilityStatus()
		}
		if plan.GetProvisioningStatus() != nil {
			planData["provisioning_status"] = *plan.GetProvisioningStatus()
		}
		if plan.GetService() != nil {
			planData["service"] = *plan.GetService()
		}
		plans = append(plans, planData)
	}

	return plans
}

func (org *Microsoft365OrganizationInfo) OrganizationSecurityComplianceNotificationMails() []string {
	if org.GetSecurityComplianceNotificationMails() == nil {
		return nil
	}
	return org.GetSecurityComplianceNotificationMails()
}

func (org *Microsoft365OrganizationInfo) OrganizationSecurityComplianceNotificationPhones() []string {
	if org.GetSecurityComplianceNotificationPhones() == nil {
		return nil
	}
	return org.GetSecurityComplianceNotificationPhones()
}

func (org *Microsoft365OrganizationInfo) OrganizationTechnicalNotificationMails() []string {
	if org.GetTechnicalNotificationMails() == nil {
		return nil
	}
	return org.GetTechnicalNotificationMails()
}

func (org *Microsoft365OrganizationInfo) OrganizationVerifiedDomains() []map[string]interface{} {
	if org.GetVerifiedDomains() == nil {
		return nil
	}

	var domains []map[string]interface{}
	for _, domain := range org.GetVerifiedDomains() {
		domainData := map[string]interface{}{}
		if domain.GetCapabilities() != nil {
			domainData["capabilities"] = *domain.GetCapabilities()
		}
		if domain.GetIsDefault() != nil {
			domainData["is_default"] = *domain.GetIsDefault()
		}
		if domain.GetIsInitial() != nil {
			domainData["is_initial"] = *domain.GetIsInitial()
		}
		if domain.GetName() != nil {
			domainData["name"] = *domain.GetName()
		}
		if domain.GetTypeEscaped() != nil {
			domainData["type"] = *domain.GetTypeEscaped()
		}
		domains = append(domains, domainData)
	}

	return domains
}

// Mail Settings transform methods
func (mail *Microsoft365MailSettingsInfo) MailSettingsId() *string {
	return &mail.UserID
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsDisplayName() *string {
	return &mail.UserID // Use UserID as display name for now
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsCreatedDateTime() *time.Time {
	// MailboxSettings doesn't have created date, return nil
	return nil
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsLastModifiedDateTime() *time.Time {
	// MailboxSettings doesn't have last modified date, return nil
	return nil
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsAutomaticRepliesEnabled() *bool {
	if mail.GetAutomaticRepliesSetting() == nil {
		return nil
	}
	if mail.GetAutomaticRepliesSetting().GetStatus() == nil {
		return nil
	}
	// Convert AutomaticRepliesStatus to bool
	status := mail.GetAutomaticRepliesSetting().GetStatus()
	enabled := *status != models.DISABLED_AUTOMATICREPLIESSTATUS
	return &enabled
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsAutomaticRepliesStatus() interface{} {
	if mail.GetAutomaticRepliesSetting() == nil {
		return nil
	}
	if mail.GetAutomaticRepliesSetting().GetStatus() == nil {
		return nil
	}
	return mail.GetAutomaticRepliesSetting().GetStatus().String()
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsAutomaticRepliesExternalAudience() interface{} {
	if mail.GetAutomaticRepliesSetting() == nil {
		return nil
	}
	if mail.GetAutomaticRepliesSetting().GetExternalAudience() == nil {
		return nil
	}
	return mail.GetAutomaticRepliesSetting().GetExternalAudience().String()
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsAutomaticRepliesScheduledStartDateTime() *time.Time {
	if mail.GetAutomaticRepliesSetting() == nil {
		return nil
	}
	if mail.GetAutomaticRepliesSetting().GetScheduledStartDateTime() == nil {
		return nil
	}
	// DateTimeTimeZoneable needs to be converted to time.Time
	// For now, return nil as this requires more complex conversion
	return nil
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsAutomaticRepliesScheduledEndDateTime() *time.Time {
	if mail.GetAutomaticRepliesSetting() == nil {
		return nil
	}
	if mail.GetAutomaticRepliesSetting().GetScheduledEndDateTime() == nil {
		return nil
	}
	// DateTimeTimeZoneable needs to be converted to time.Time
	// For now, return nil as this requires more complex conversion
	return nil
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsAutomaticRepliesInternalReplyMessage() *string {
	if mail.GetAutomaticRepliesSetting() == nil {
		return nil
	}
	return mail.GetAutomaticRepliesSetting().GetInternalReplyMessage()
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsAutomaticRepliesExternalReplyMessage() *string {
	if mail.GetAutomaticRepliesSetting() == nil {
		return nil
	}
	return mail.GetAutomaticRepliesSetting().GetExternalReplyMessage()
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsArchiveFolder() *string {
	return mail.GetArchiveFolder()
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsTimeZone() *string {
	return mail.GetTimeZone()
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsLanguage() map[string]interface{} {
	if mail.GetLanguage() == nil {
		return nil
	}

	language := map[string]interface{}{}
	if mail.GetLanguage().GetDisplayName() != nil {
		language["display_name"] = *mail.GetLanguage().GetDisplayName()
	}
	if mail.GetLanguage().GetLocale() != nil {
		language["locale"] = *mail.GetLanguage().GetLocale()
	}

	return language
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsWorkingHours() map[string]interface{} {
	if mail.GetWorkingHours() == nil {
		return nil
	}

	workingHours := map[string]interface{}{}
	if mail.GetWorkingHours().GetDaysOfWeek() != nil {
		var days []string
		for _, day := range mail.GetWorkingHours().GetDaysOfWeek() {
			days = append(days, day.String())
		}
		workingHours["days_of_week"] = days
	}
	if mail.GetWorkingHours().GetStartTime() != nil {
		workingHours["start_time"] = *mail.GetWorkingHours().GetStartTime()
	}
	if mail.GetWorkingHours().GetEndTime() != nil {
		workingHours["end_time"] = *mail.GetWorkingHours().GetEndTime()
	}
	if mail.GetWorkingHours().GetTimeZone() != nil {
		workingHours["time_zone"] = mail.GetWorkingHours().GetTimeZone()
	}

	return workingHours
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsAutomaticRepliesSetting() map[string]interface{} {
	if mail.GetAutomaticRepliesSetting() == nil {
		return nil
	}

	setting := map[string]interface{}{}
	if mail.GetAutomaticRepliesSetting().GetStatus() != nil {
		setting["status"] = mail.GetAutomaticRepliesSetting().GetStatus().String()
	}
	if mail.GetAutomaticRepliesSetting().GetExternalAudience() != nil {
		setting["external_audience"] = mail.GetAutomaticRepliesSetting().GetExternalAudience().String()
	}
	if mail.GetAutomaticRepliesSetting().GetScheduledStartDateTime() != nil {
		setting["scheduled_start_date_time"] = mail.GetAutomaticRepliesSetting().GetScheduledStartDateTime()
	}
	if mail.GetAutomaticRepliesSetting().GetScheduledEndDateTime() != nil {
		setting["scheduled_end_date_time"] = mail.GetAutomaticRepliesSetting().GetScheduledEndDateTime()
	}
	if mail.GetAutomaticRepliesSetting().GetInternalReplyMessage() != nil {
		setting["internal_reply_message"] = *mail.GetAutomaticRepliesSetting().GetInternalReplyMessage()
	}
	if mail.GetAutomaticRepliesSetting().GetExternalReplyMessage() != nil {
		setting["external_reply_message"] = *mail.GetAutomaticRepliesSetting().GetExternalReplyMessage()
	}

	return setting
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsDateFormat() *string {
	return mail.GetDateFormat()
}

func (mail *Microsoft365MailSettingsInfo) MailSettingsTimeFormat() *string {
	return mail.GetTimeFormat()
}

// Security Settings transform methods
func (security *Microsoft365SecuritySettingsInfo) SecuritySettingsId() *string {
	// Extract ID based on policy type
	switch security.PolicyType {
	case "conditional_access_policy":
		if policy, ok := security.PolicyData.(models.ConditionalAccessPolicyable); ok {
			return policy.GetId()
		}
	case "identity_security_defaults_enforcement_policy":
		if policy, ok := security.PolicyData.(models.IdentitySecurityDefaultsEnforcementPolicyable); ok {
			return policy.GetId()
		}
	case "authentication_methods_policy":
		if policy, ok := security.PolicyData.(models.AuthenticationMethodsPolicyable); ok {
			return policy.GetId()
		}
	case "authorization_policy":
		if policy, ok := security.PolicyData.(models.AuthorizationPolicyable); ok {
			return policy.GetId()
		}
	}
	return nil
}

func (security *Microsoft365SecuritySettingsInfo) SecuritySettingsDisplayName() *string {
	// Extract display name based on policy type
	switch security.PolicyType {
	case "conditional_access_policy":
		if policy, ok := security.PolicyData.(models.ConditionalAccessPolicyable); ok {
			return policy.GetDisplayName()
		}
	case "identity_security_defaults_enforcement_policy":
		if policy, ok := security.PolicyData.(models.IdentitySecurityDefaultsEnforcementPolicyable); ok {
			return policy.GetDisplayName()
		}
	case "authentication_methods_policy":
		if policy, ok := security.PolicyData.(models.AuthenticationMethodsPolicyable); ok {
			return policy.GetDisplayName()
		}
	case "authorization_policy":
		if policy, ok := security.PolicyData.(models.AuthorizationPolicyable); ok {
			return policy.GetDisplayName()
		}
	}
	return nil
}

func (security *Microsoft365SecuritySettingsInfo) SecuritySettingsCreatedDateTime() *time.Time {
	// Extract created date based on policy type
	switch security.PolicyType {
	case "conditional_access_policy":
		if policy, ok := security.PolicyData.(models.ConditionalAccessPolicyable); ok {
			return policy.GetCreatedDateTime()
		}
	}
	return nil
}

func (security *Microsoft365SecuritySettingsInfo) SecuritySettingsLastModifiedDateTime() *time.Time {
	// Extract last modified date based on policy type
	switch security.PolicyType {
	case "conditional_access_policy":
		if policy, ok := security.PolicyData.(models.ConditionalAccessPolicyable); ok {
			return policy.GetModifiedDateTime()
		}
	case "authentication_methods_policy":
		if policy, ok := security.PolicyData.(models.AuthenticationMethodsPolicyable); ok {
			return policy.GetLastModifiedDateTime()
		}
	}
	return nil
}

func (security *Microsoft365SecuritySettingsInfo) SecuritySettingsConditionalAccessState() interface{} {
	if policy, ok := security.PolicyData.(models.ConditionalAccessPolicyable); ok {
		if policy.GetState() == nil {
			return nil
		}
		return policy.GetState().String()
	}
	return nil
}

func (security *Microsoft365SecuritySettingsInfo) SecuritySettingsConditionalAccessConditions() map[string]interface{} {
	if policy, ok := security.PolicyData.(models.ConditionalAccessPolicyable); ok {
		if policy.GetConditions() == nil {
			return nil
		}

		conditions := map[string]interface{}{}
		if policy.GetConditions().GetApplications() != nil {
			conditions["applications"] = policy.GetConditions().GetApplications()
		}
		if policy.GetConditions().GetUsers() != nil {
			conditions["users"] = policy.GetConditions().GetUsers()
		}
		if policy.GetConditions().GetClientApplications() != nil {
			conditions["client_applications"] = policy.GetConditions().GetClientApplications()
		}
		if policy.GetConditions().GetDevices() != nil {
			conditions["devices"] = policy.GetConditions().GetDevices()
		}
		if policy.GetConditions().GetLocations() != nil {
			conditions["locations"] = policy.GetConditions().GetLocations()
		}
		if policy.GetConditions().GetPlatforms() != nil {
			conditions["platforms"] = policy.GetConditions().GetPlatforms()
		}
		if policy.GetConditions().GetSignInRiskLevels() != nil {
			var riskLevels []string
			for _, level := range policy.GetConditions().GetSignInRiskLevels() {
				riskLevels = append(riskLevels, level.String())
			}
			conditions["sign_in_risk_levels"] = riskLevels
		}
		if policy.GetConditions().GetUserRiskLevels() != nil {
			var riskLevels []string
			for _, level := range policy.GetConditions().GetUserRiskLevels() {
				riskLevels = append(riskLevels, level.String())
			}
			conditions["user_risk_levels"] = riskLevels
		}

		return conditions
	}
	return nil
}

func (security *Microsoft365SecuritySettingsInfo) SecuritySettingsConditionalAccessGrantControls() map[string]interface{} {
	if policy, ok := security.PolicyData.(models.ConditionalAccessPolicyable); ok {
		if policy.GetGrantControls() == nil {
			return nil
		}

		grantControls := map[string]interface{}{}
		if policy.GetGrantControls().GetOperator() != nil {
			grantControls["operator"] = *policy.GetGrantControls().GetOperator()
		}
		if policy.GetGrantControls().GetBuiltInControls() != nil {
			var controls []string
			for _, control := range policy.GetGrantControls().GetBuiltInControls() {
				controls = append(controls, control.String())
			}
			grantControls["built_in_controls"] = controls
		}
		if policy.GetGrantControls().GetCustomAuthenticationFactors() != nil {
			var factors []string
			factors = append(factors, policy.GetGrantControls().GetCustomAuthenticationFactors()...)
			grantControls["custom_authentication_factors"] = factors
		}
		if policy.GetGrantControls().GetTermsOfUse() != nil {
			var terms []string
			terms = append(terms, policy.GetGrantControls().GetTermsOfUse()...)
			grantControls["terms_of_use"] = terms
		}

		return grantControls
	}
	return nil
}

func (security *Microsoft365SecuritySettingsInfo) SecuritySettingsConditionalAccessSessionControls() map[string]interface{} {
	if policy, ok := security.PolicyData.(models.ConditionalAccessPolicyable); ok {
		if policy.GetSessionControls() == nil {
			return nil
		}

		sessionControls := map[string]interface{}{}
		if policy.GetSessionControls().GetApplicationEnforcedRestrictions() != nil {
			sessionControls["application_enforced_restrictions"] = policy.GetSessionControls().GetApplicationEnforcedRestrictions()
		}
		if policy.GetSessionControls().GetCloudAppSecurity() != nil {
			sessionControls["cloud_app_security"] = policy.GetSessionControls().GetCloudAppSecurity()
		}
		if policy.GetSessionControls().GetDisableResilienceDefaults() != nil {
			sessionControls["disable_resilience_defaults"] = *policy.GetSessionControls().GetDisableResilienceDefaults()
		}
		if policy.GetSessionControls().GetPersistentBrowser() != nil {
			sessionControls["persistent_browser"] = policy.GetSessionControls().GetPersistentBrowser()
		}
		if policy.GetSessionControls().GetSignInFrequency() != nil {
			sessionControls["sign_in_frequency"] = policy.GetSessionControls().GetSignInFrequency()
		}

		return sessionControls
	}
	return nil
}

func (security *Microsoft365SecuritySettingsInfo) SecuritySettingsSecurityDefaultsEnabled() *bool {
	if policy, ok := security.PolicyData.(models.IdentitySecurityDefaultsEnforcementPolicyable); ok {
		return policy.GetIsEnabled()
	}
	return nil
}

func (security *Microsoft365SecuritySettingsInfo) SecuritySettingsSecurityDefaultsIsEnabled() *bool {
	if policy, ok := security.PolicyData.(models.IdentitySecurityDefaultsEnforcementPolicyable); ok {
		return policy.GetIsEnabled()
	}
	return nil
}

func (security *Microsoft365SecuritySettingsInfo) SecuritySettingsAuthenticationMethodsPolicy() map[string]interface{} {
	if policy, ok := security.PolicyData.(models.AuthenticationMethodsPolicyable); ok {
		authPolicy := map[string]interface{}{}
		if policy.GetId() != nil {
			authPolicy["id"] = *policy.GetId()
		}
		if policy.GetDescription() != nil {
			authPolicy["description"] = *policy.GetDescription()
		}
		if policy.GetDisplayName() != nil {
			authPolicy["display_name"] = *policy.GetDisplayName()
		}
		if policy.GetLastModifiedDateTime() != nil {
			authPolicy["last_modified_date_time"] = *policy.GetLastModifiedDateTime()
		}
		if policy.GetPolicyVersion() != nil {
			authPolicy["policy_version"] = *policy.GetPolicyVersion()
		}
		if policy.GetReconfirmationInDays() != nil {
			authPolicy["reconfirmation_in_days"] = *policy.GetReconfirmationInDays()
		}
		if policy.GetRegistrationEnforcement() != nil {
			authPolicy["registration_enforcement"] = policy.GetRegistrationEnforcement()
		}

		return authPolicy
	}
	return nil
}

func (security *Microsoft365SecuritySettingsInfo) SecuritySettingsAuthorizationPolicy() map[string]interface{} {
	if policy, ok := security.PolicyData.(models.AuthorizationPolicyable); ok {
		authPolicy := map[string]interface{}{}
		if policy.GetId() != nil {
			authPolicy["id"] = *policy.GetId()
		}
		if policy.GetDescription() != nil {
			authPolicy["description"] = *policy.GetDescription()
		}
		if policy.GetDisplayName() != nil {
			authPolicy["display_name"] = *policy.GetDisplayName()
		}
		// Note: AuthorizationPolicy doesn't have LastModifiedDateTime or Version methods
		if policy.GetAllowInvitesFrom() != nil {
			authPolicy["allow_invites_from"] = policy.GetAllowInvitesFrom().String()
		}
		if policy.GetBlockMsolPowerShell() != nil {
			authPolicy["block_msol_powershell"] = *policy.GetBlockMsolPowerShell()
		}
		if policy.GetDefaultUserRolePermissions() != nil {
			authPolicy["default_user_role_permissions"] = policy.GetDefaultUserRolePermissions()
		}
		// Note: AuthorizationPolicy doesn't have GetEnabledPreviewFeatures method
		if policy.GetGuestUserRoleId() != nil {
			authPolicy["guest_user_role_id"] = *policy.GetGuestUserRoleId()
		}
		// Note: AuthorizationPolicy doesn't have GetPermissionGrantPolicyIdsAssignedToDefaultUserRole method

		return authPolicy
	}
	return nil
}

func (security *Microsoft365SecuritySettingsInfo) SecuritySettingsIdentitySecurityDefaultsEnforcementPolicy() map[string]interface{} {
	if policy, ok := security.PolicyData.(models.IdentitySecurityDefaultsEnforcementPolicyable); ok {
		policyData := map[string]interface{}{}
		if policy.GetId() != nil {
			policyData["id"] = *policy.GetId()
		}
		if policy.GetDescription() != nil {
			policyData["description"] = *policy.GetDescription()
		}
		if policy.GetDisplayName() != nil {
			policyData["display_name"] = *policy.GetDisplayName()
		}
		// Note: IdentitySecurityDefaultsEnforcementPolicy doesn't have LastModifiedDateTime or Version methods
		if policy.GetIsEnabled() != nil {
			policyData["is_enabled"] = *policy.GetIsEnabled()
		}

		return policyData
	}
	return nil
}

func (security *Microsoft365SecuritySettingsInfo) SecuritySettingsPolicyDetails() map[string]interface{} {
	details := map[string]interface{}{
		"policy_type": security.PolicyType,
	}

	// Add type-specific details
	switch security.PolicyType {
	case "conditional_access_policy":
		if policy, ok := security.PolicyData.(models.ConditionalAccessPolicyable); ok {
			if policy.GetId() != nil {
				details["id"] = *policy.GetId()
			}
			if policy.GetDisplayName() != nil {
				details["display_name"] = *policy.GetDisplayName()
			}
			if policy.GetState() != nil {
				details["state"] = policy.GetState().String()
			}
			if policy.GetCreatedDateTime() != nil {
				details["created_date_time"] = *policy.GetCreatedDateTime()
			}
			if policy.GetModifiedDateTime() != nil {
				details["modified_date_time"] = *policy.GetModifiedDateTime()
			}
		}
	}

	return details
}

func (security *Microsoft365SecuritySettingsInfo) SecuritySettingsComplianceSettings() map[string]interface{} {
	// This would contain compliance-related settings
	// For now, return basic structure
	return map[string]interface{}{
		"policy_type":        security.PolicyType,
		"compliance_enabled": true, // Placeholder
	}
}

func (security *Microsoft365SecuritySettingsInfo) SecuritySettingsPrivacySettings() map[string]interface{} {
	// This would contain privacy-related settings
	// For now, return basic structure
	return map[string]interface{}{
		"policy_type":     security.PolicyType,
		"privacy_enabled": true, // Placeholder
	}
}

// Calendar Settings transform methods
func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsId() *string {
	return &calendar.UserID
}

func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsDisplayName() *string {
	return &calendar.UserID // Use UserID as display name for now
}

func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsCreatedDateTime() *time.Time {
	// Calendar settings don't have a specific created date, return nil
	return nil
}

func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsLastModifiedDateTime() *time.Time {
	// Calendar settings don't have a specific last modified date, return nil
	return nil
}

func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsCalendarPermissions() map[string]interface{} {
	// Extract calendar permissions from calendars
	if calendar.Calendars == nil || calendar.Calendars.GetValue() == nil {
		return nil
	}

	permissions := map[string]interface{}{
		"total_calendars": len(calendar.Calendars.GetValue()),
		"calendars":       []map[string]interface{}{},
	}

	var calendarList []map[string]interface{}
	for _, cal := range calendar.Calendars.GetValue() {
		calInfo := map[string]interface{}{}
		if cal.GetId() != nil {
			calInfo["id"] = *cal.GetId()
		}
		if cal.GetName() != nil {
			calInfo["name"] = *cal.GetName()
		}
		if cal.GetCanEdit() != nil {
			calInfo["can_edit"] = *cal.GetCanEdit()
		}
		if cal.GetCanShare() != nil {
			calInfo["can_share"] = *cal.GetCanShare()
		}
		if cal.GetCanViewPrivateItems() != nil {
			calInfo["can_view_private_items"] = *cal.GetCanViewPrivateItems()
		}
		if cal.GetIsDefaultCalendar() != nil {
			calInfo["is_default_calendar"] = *cal.GetIsDefaultCalendar()
		}
		if cal.GetIsRemovable() != nil {
			calInfo["is_removable"] = *cal.GetIsRemovable()
		}
		if cal.GetIsTallyingResponses() != nil {
			calInfo["is_tallying_responses"] = *cal.GetIsTallyingResponses()
		}
		if cal.GetColor() != nil {
			calInfo["color"] = *cal.GetColor()
		}
		if cal.GetHexColor() != nil {
			calInfo["hex_color"] = *cal.GetHexColor()
		}
		calendarList = append(calendarList, calInfo)
	}
	permissions["calendars"] = calendarList

	return permissions
}

func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsCalendarSharing() map[string]interface{} {
	// Calendar sharing settings
	sharing := map[string]interface{}{
		"sharing_enabled": true, // Placeholder
		"external_sharing": map[string]interface{}{
			"enabled": false, // Placeholder
		},
		"internal_sharing": map[string]interface{}{
			"enabled": true, // Placeholder
		},
	}
	return sharing
}

func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsWorkingHours() map[string]interface{} {
	// Extract working hours from mailbox settings
	if calendar.GetWorkingHours() == nil {
		return nil
	}

	workingHours := map[string]interface{}{}
	if calendar.GetWorkingHours().GetDaysOfWeek() != nil {
		var days []string
		for _, day := range calendar.GetWorkingHours().GetDaysOfWeek() {
			days = append(days, day.String())
		}
		workingHours["days_of_week"] = days
	}
	if calendar.GetWorkingHours().GetStartTime() != nil {
		workingHours["start_time"] = *calendar.GetWorkingHours().GetStartTime()
	}
	if calendar.GetWorkingHours().GetEndTime() != nil {
		workingHours["end_time"] = *calendar.GetWorkingHours().GetEndTime()
	}
	if calendar.GetWorkingHours().GetTimeZone() != nil {
		workingHours["time_zone"] = calendar.GetWorkingHours().GetTimeZone()
	}

	return workingHours
}

func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsReminderSettings() map[string]interface{} {
	// Calendar reminder settings
	reminders := map[string]interface{}{
		"default_reminder_minutes": 15,   // Placeholder
		"reminder_sound_enabled":   true, // Placeholder
		"email_reminders": map[string]interface{}{
			"enabled":         true,
			"advance_minutes": []int{15, 30, 60, 1440}, // 15 min, 30 min, 1 hour, 1 day
		},
		"popup_reminders": map[string]interface{}{
			"enabled":         true,
			"advance_minutes": []int{5, 10, 15},
		},
	}
	return reminders
}

func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsTimeZone() *string {
	return calendar.GetTimeZone()
}

func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsLanguage() map[string]interface{} {
	if calendar.GetLanguage() == nil {
		return nil
	}

	language := map[string]interface{}{}
	if calendar.GetLanguage().GetDisplayName() != nil {
		language["display_name"] = *calendar.GetLanguage().GetDisplayName()
	}
	if calendar.GetLanguage().GetLocale() != nil {
		language["locale"] = *calendar.GetLanguage().GetLocale()
	}

	return language
}

func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsExchangeOnlineSettings() map[string]interface{} {
	// Exchange Online specific settings
	exchangeSettings := map[string]interface{}{
		"exchange_online_enabled": true,
		"calendar_sync_enabled":   true,
		"free_busy_sharing": map[string]interface{}{
			"enabled": true,
			"level":   "limited", // limited, detailed, full
		},
		"meeting_requests": map[string]interface{}{
			"auto_accept":            false,
			"auto_decline_conflicts": false,
			"send_responses":         true,
		},
	}
	return exchangeSettings
}

func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsDelegateSettings() map[string]interface{} {
	// Delegate and assistant settings
	delegateSettings := map[string]interface{}{
		"delegates_enabled": true,
		"assistant_access": map[string]interface{}{
			"enabled":     false,
			"permissions": []string{"read", "write"}, // read, write, delete
		},
		"delegate_permissions": map[string]interface{}{
			"can_view_private_items": false,
			"can_edit_items":         false,
			"can_delete_items":       false,
		},
	}
	return delegateSettings
}

func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsMeetingSettings() map[string]interface{} {
	// Meeting and scheduling settings
	meetingSettings := map[string]interface{}{
		"meeting_requests": map[string]interface{}{
			"auto_accept":            false,
			"auto_decline_conflicts": false,
			"send_responses":         true,
		},
		"scheduling_assistant": map[string]interface{}{
			"enabled":        true,
			"show_free_busy": true,
		},
		"meeting_reminders": map[string]interface{}{
			"default_minutes": 15,
			"email_enabled":   true,
			"popup_enabled":   true,
		},
	}
	return meetingSettings
}

func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsResourceBookingSettings() map[string]interface{} {
	// Resource booking and room scheduling settings
	resourceSettings := map[string]interface{}{
		"resource_booking_enabled": true,
		"room_scheduling": map[string]interface{}{
			"enabled":             true,
			"auto_accept":         false,
			"conflict_resolution": "decline", // accept, decline, suggest_alternatives
		},
		"equipment_booking": map[string]interface{}{
			"enabled":           true,
			"approval_required": false,
		},
	}
	return resourceSettings
}

func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsOrganizationPolicies() map[string]interface{} {
	// Organization-wide calendar policies
	orgPolicies := map[string]interface{}{
		"calendar_policies": map[string]interface{}{
			"enabled":         true,
			"retention_days":  2555, // 7 years default
			"archive_enabled": true,
		},
		"sharing_policies": map[string]interface{}{
			"external_sharing": "blocked", // allowed, blocked, limited
			"internal_sharing": "allowed",
		},
		"compliance_policies": map[string]interface{}{
			"litigation_hold":  false,
			"retention_policy": "default",
		},
	}
	return orgPolicies
}

func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsComplianceSettings() map[string]interface{} {
	// Compliance and retention settings
	complianceSettings := map[string]interface{}{
		"retention": map[string]interface{}{
			"enabled":            true,
			"period_days":        2555, // 7 years
			"archive_after_days": 1095, // 3 years
		},
		"litigation_hold": map[string]interface{}{
			"enabled":       false,
			"hold_duration": "indefinite",
		},
		"e_discovery": map[string]interface{}{
			"enabled":        true,
			"search_enabled": true,
		},
	}
	return complianceSettings
}

func (calendar *Microsoft365CalendarSettingsInfo) CalendarSettingsSecuritySettings() map[string]interface{} {
	// Security-related calendar settings
	securitySettings := map[string]interface{}{
		"access_control": map[string]interface{}{
			"require_authentication": true,
			"allow_anonymous_access": false,
		},
		"encryption": map[string]interface{}{
			"at_rest":    true,
			"in_transit": true,
		},
		"audit_logging": map[string]interface{}{
			"enabled":        true,
			"retention_days": 90,
		},
	}
	return securitySettings
}

// User transform methods
func (user *Microsoft365UserInfo) UserAssignedLicenses() []map[string]interface{} {
	if user.GetAssignedLicenses() == nil {
		return nil
	}

	licenses := []map[string]interface{}{}
	for _, license := range user.GetAssignedLicenses() {
		licenseData := map[string]interface{}{}
		if license.GetDisabledPlans() != nil {
			licenseData["disabled_plans"] = license.GetDisabledPlans()
		}
		if license.GetSkuId() != nil {
			licenseData["sku_id"] = *license.GetSkuId()
		}
		licenses = append(licenses, licenseData)
	}
	return licenses
}

func (user *Microsoft365UserInfo) UserAssignedPlans() []map[string]interface{} {
	if user.GetAssignedPlans() == nil {
		return nil
	}

	plans := []map[string]interface{}{}
	for _, plan := range user.GetAssignedPlans() {
		planData := map[string]interface{}{}
		if plan.GetAssignedDateTime() != nil {
			planData["assigned_date_time"] = *plan.GetAssignedDateTime()
		}
		if plan.GetCapabilityStatus() != nil {
			planData["capability_status"] = *plan.GetCapabilityStatus()
		}
		if plan.GetService() != nil {
			planData["service"] = *plan.GetService()
		}
		if plan.GetServicePlanId() != nil {
			planData["service_plan_id"] = *plan.GetServicePlanId()
		}
		plans = append(plans, planData)
	}
	return plans
}

func (user *Microsoft365UserInfo) UserProvisionedPlans() []map[string]interface{} {
	if user.GetProvisionedPlans() == nil {
		return nil
	}

	plans := []map[string]interface{}{}
	for _, plan := range user.GetProvisionedPlans() {
		planData := map[string]interface{}{}
		if plan.GetCapabilityStatus() != nil {
			planData["capability_status"] = *plan.GetCapabilityStatus()
		}
		if plan.GetProvisioningStatus() != nil {
			planData["provisioning_status"] = *plan.GetProvisioningStatus()
		}
		if plan.GetService() != nil {
			planData["service"] = *plan.GetService()
		}
		plans = append(plans, planData)
	}
	return plans
}

func (user *Microsoft365UserInfo) UserPasswordProfile() map[string]interface{} {
	if user.GetPasswordProfile() == nil {
		return nil
	}

	profile := map[string]interface{}{}
	if user.GetPasswordProfile().GetForceChangePasswordNextSignIn() != nil {
		profile["force_change_password_next_sign_in"] = *user.GetPasswordProfile().GetForceChangePasswordNextSignIn()
	}
	if user.GetPasswordProfile().GetForceChangePasswordNextSignInWithMfa() != nil {
		profile["force_change_password_next_sign_in_with_mfa"] = *user.GetPasswordProfile().GetForceChangePasswordNextSignInWithMfa()
	}
	if user.GetPasswordProfile().GetPassword() != nil {
		profile["password"] = *user.GetPasswordProfile().GetPassword()
	}
	return profile
}

func (user *Microsoft365UserInfo) UserOnPremisesExtensionAttributes() map[string]interface{} {
	if user.GetOnPremisesExtensionAttributes() == nil {
		return nil
	}

	attrs := map[string]interface{}{}
	if user.GetOnPremisesExtensionAttributes().GetExtensionAttribute1() != nil {
		attrs["extension_attribute1"] = *user.GetOnPremisesExtensionAttributes().GetExtensionAttribute1()
	}
	if user.GetOnPremisesExtensionAttributes().GetExtensionAttribute2() != nil {
		attrs["extension_attribute2"] = *user.GetOnPremisesExtensionAttributes().GetExtensionAttribute2()
	}
	if user.GetOnPremisesExtensionAttributes().GetExtensionAttribute3() != nil {
		attrs["extension_attribute3"] = *user.GetOnPremisesExtensionAttributes().GetExtensionAttribute3()
	}
	if user.GetOnPremisesExtensionAttributes().GetExtensionAttribute4() != nil {
		attrs["extension_attribute4"] = *user.GetOnPremisesExtensionAttributes().GetExtensionAttribute4()
	}
	if user.GetOnPremisesExtensionAttributes().GetExtensionAttribute5() != nil {
		attrs["extension_attribute5"] = *user.GetOnPremisesExtensionAttributes().GetExtensionAttribute5()
	}
	if user.GetOnPremisesExtensionAttributes().GetExtensionAttribute6() != nil {
		attrs["extension_attribute6"] = *user.GetOnPremisesExtensionAttributes().GetExtensionAttribute6()
	}
	if user.GetOnPremisesExtensionAttributes().GetExtensionAttribute7() != nil {
		attrs["extension_attribute7"] = *user.GetOnPremisesExtensionAttributes().GetExtensionAttribute7()
	}
	if user.GetOnPremisesExtensionAttributes().GetExtensionAttribute8() != nil {
		attrs["extension_attribute8"] = *user.GetOnPremisesExtensionAttributes().GetExtensionAttribute8()
	}
	if user.GetOnPremisesExtensionAttributes().GetExtensionAttribute9() != nil {
		attrs["extension_attribute9"] = *user.GetOnPremisesExtensionAttributes().GetExtensionAttribute9()
	}
	if user.GetOnPremisesExtensionAttributes().GetExtensionAttribute10() != nil {
		attrs["extension_attribute10"] = *user.GetOnPremisesExtensionAttributes().GetExtensionAttribute10()
	}
	if user.GetOnPremisesExtensionAttributes().GetExtensionAttribute11() != nil {
		attrs["extension_attribute11"] = *user.GetOnPremisesExtensionAttributes().GetExtensionAttribute11()
	}
	if user.GetOnPremisesExtensionAttributes().GetExtensionAttribute12() != nil {
		attrs["extension_attribute12"] = *user.GetOnPremisesExtensionAttributes().GetExtensionAttribute12()
	}
	if user.GetOnPremisesExtensionAttributes().GetExtensionAttribute13() != nil {
		attrs["extension_attribute13"] = *user.GetOnPremisesExtensionAttributes().GetExtensionAttribute13()
	}
	if user.GetOnPremisesExtensionAttributes().GetExtensionAttribute14() != nil {
		attrs["extension_attribute14"] = *user.GetOnPremisesExtensionAttributes().GetExtensionAttribute14()
	}
	if user.GetOnPremisesExtensionAttributes().GetExtensionAttribute15() != nil {
		attrs["extension_attribute15"] = *user.GetOnPremisesExtensionAttributes().GetExtensionAttribute15()
	}
	return attrs
}

func (user *Microsoft365UserInfo) UserOnPremisesProvisioningErrors() []map[string]interface{} {
	if user.GetOnPremisesProvisioningErrors() == nil {
		return nil
	}

	errors := []map[string]interface{}{}
	for _, err := range user.GetOnPremisesProvisioningErrors() {
		errorData := map[string]interface{}{}
		if err.GetCategory() != nil {
			errorData["category"] = *err.GetCategory()
		}
		if err.GetOccurredDateTime() != nil {
			errorData["occurred_date_time"] = *err.GetOccurredDateTime()
		}
		if err.GetPropertyCausingError() != nil {
			errorData["property_causing_error"] = *err.GetPropertyCausingError()
		}
		if err.GetValue() != nil {
			errorData["value"] = *err.GetValue()
		}
		errors = append(errors, errorData)
	}
	return errors
}

func (user *Microsoft365UserInfo) UserServiceProvisioningErrors() []map[string]interface{} {
	if user.GetServiceProvisioningErrors() == nil {
		return nil
	}

	errors := []map[string]interface{}{}
	for _, err := range user.GetServiceProvisioningErrors() {
		errorData := map[string]interface{}{}
		if err.GetCreatedDateTime() != nil {
			errorData["created_date_time"] = *err.GetCreatedDateTime()
		}
		if err.GetIsResolved() != nil {
			errorData["is_resolved"] = *err.GetIsResolved()
		}
		if err.GetServiceInstance() != nil {
			errorData["service_instance"] = *err.GetServiceInstance()
		}
		errors = append(errors, errorData)
	}
	return errors
}

func (user *Microsoft365UserInfo) UserIdentities() []map[string]interface{} {
	if user.GetIdentities() == nil {
		return nil
	}

	identities := []map[string]interface{}{}
	for _, identity := range user.GetIdentities() {
		identityData := map[string]interface{}{}
		if identity.GetSignInType() != nil {
			identityData["sign_in_type"] = *identity.GetSignInType()
		}
		if identity.GetIssuer() != nil {
			identityData["issuer"] = *identity.GetIssuer()
		}
		if identity.GetIssuerAssignedId() != nil {
			identityData["issuer_assigned_id"] = *identity.GetIssuerAssignedId()
		}
		identities = append(identities, identityData)
	}
	return identities
}

func (user *Microsoft365UserInfo) UserLicenseAssignmentStates() []map[string]interface{} {
	if user.GetLicenseAssignmentStates() == nil {
		return nil
	}

	states := []map[string]interface{}{}
	for _, state := range user.GetLicenseAssignmentStates() {
		stateData := map[string]interface{}{}
		if state.GetAssignedByGroup() != nil {
			stateData["assigned_by_group"] = *state.GetAssignedByGroup()
		}
		if state.GetDisabledPlans() != nil {
			stateData["disabled_plans"] = state.GetDisabledPlans()
		}
		if state.GetError() != nil {
			stateData["error"] = *state.GetError()
		}
		if state.GetLastUpdatedDateTime() != nil {
			stateData["last_updated_date_time"] = *state.GetLastUpdatedDateTime()
		}
		if state.GetSkuId() != nil {
			stateData["sku_id"] = *state.GetSkuId()
		}
		if state.GetState() != nil {
			stateData["state"] = *state.GetState()
		}
		states = append(states, stateData)
	}
	return states
}

func (user *Microsoft365UserInfo) UserLicenseDetails() []map[string]interface{} {
	if user.GetLicenseDetails() == nil {
		return nil
	}

	details := []map[string]interface{}{}
	for _, detail := range user.GetLicenseDetails() {
		detailData := map[string]interface{}{}
		if detail.GetServicePlans() != nil {
			servicePlans := []map[string]interface{}{}
			for _, plan := range detail.GetServicePlans() {
				planData := map[string]interface{}{}
				if plan.GetServicePlanId() != nil {
					planData["service_plan_id"] = *plan.GetServicePlanId()
				}
				if plan.GetServicePlanName() != nil {
					planData["service_plan_name"] = *plan.GetServicePlanName()
				}
				if plan.GetProvisioningStatus() != nil {
					planData["provisioning_status"] = *plan.GetProvisioningStatus()
				}
				if plan.GetAppliesTo() != nil {
					planData["applies_to"] = *plan.GetAppliesTo()
				}
				servicePlans = append(servicePlans, planData)
			}
			detailData["service_plans"] = servicePlans
		}
		if detail.GetSkuId() != nil {
			detailData["sku_id"] = *detail.GetSkuId()
		}
		if detail.GetSkuPartNumber() != nil {
			detailData["sku_part_number"] = *detail.GetSkuPartNumber()
		}
		details = append(details, detailData)
	}
	return details
}

// User Mailbox Settings transform methods - these will be called on the mailbox settings from hydrate
func (user *Microsoft365UserInfo) GetUserPurpose() string {
	if user.MailboxSettings == nil || user.MailboxSettings.GetUserPurpose() == nil {
		return ""
	}
	return user.MailboxSettings.GetUserPurpose().String()
}

func (user *Microsoft365UserInfo) GetArchiveFolder() string {
	if user.MailboxSettings == nil || user.MailboxSettings.GetArchiveFolder() == nil {
		return ""
	}
	return *user.MailboxSettings.GetArchiveFolder()
}

func (user *Microsoft365UserInfo) GetAutomaticRepliesSetting() map[string]interface{} {
	if user.MailboxSettings == nil || user.MailboxSettings.GetAutomaticRepliesSetting() == nil {
		return nil
	}

	setting := map[string]interface{}{}
	if user.MailboxSettings.GetAutomaticRepliesSetting().GetStatus() != nil {
		setting["status"] = user.MailboxSettings.GetAutomaticRepliesSetting().GetStatus().String()
	}
	if user.MailboxSettings.GetAutomaticRepliesSetting().GetExternalAudience() != nil {
		setting["external_audience"] = user.MailboxSettings.GetAutomaticRepliesSetting().GetExternalAudience().String()
	}
	if user.MailboxSettings.GetAutomaticRepliesSetting().GetScheduledStartDateTime() != nil {
		setting["scheduled_start_date_time"] = user.MailboxSettings.GetAutomaticRepliesSetting().GetScheduledStartDateTime()
	}
	if user.MailboxSettings.GetAutomaticRepliesSetting().GetScheduledEndDateTime() != nil {
		setting["scheduled_end_date_time"] = user.MailboxSettings.GetAutomaticRepliesSetting().GetScheduledEndDateTime()
	}
	if user.MailboxSettings.GetAutomaticRepliesSetting().GetInternalReplyMessage() != nil {
		setting["internal_reply_message"] = *user.MailboxSettings.GetAutomaticRepliesSetting().GetInternalReplyMessage()
	}
	if user.MailboxSettings.GetAutomaticRepliesSetting().GetExternalReplyMessage() != nil {
		setting["external_reply_message"] = *user.MailboxSettings.GetAutomaticRepliesSetting().GetExternalReplyMessage()
	}
	return setting
}

func (user *Microsoft365UserInfo) GetDateFormat() string {
	if user.MailboxSettings == nil || user.MailboxSettings.GetDateFormat() == nil {
		return ""
	}
	return *user.MailboxSettings.GetDateFormat()
}

func (user *Microsoft365UserInfo) GetTimeFormat() string {
	if user.MailboxSettings == nil || user.MailboxSettings.GetTimeFormat() == nil {
		return ""
	}
	return *user.MailboxSettings.GetTimeFormat()
}

func (user *Microsoft365UserInfo) GetTimeZone() string {
	if user.MailboxSettings == nil || user.MailboxSettings.GetTimeZone() == nil {
		return ""
	}
	return *user.MailboxSettings.GetTimeZone()
}

func (user *Microsoft365UserInfo) GetLanguage() map[string]interface{} {
	if user.MailboxSettings == nil || user.MailboxSettings.GetLanguage() == nil {
		return nil
	}

	language := map[string]interface{}{}
	if user.MailboxSettings.GetLanguage().GetDisplayName() != nil {
		language["display_name"] = *user.MailboxSettings.GetLanguage().GetDisplayName()
	}
	if user.MailboxSettings.GetLanguage().GetLocale() != nil {
		language["locale"] = *user.MailboxSettings.GetLanguage().GetLocale()
	}
	return language
}

func (user *Microsoft365UserInfo) GetWorkingHours() map[string]interface{} {
	if user.MailboxSettings == nil || user.MailboxSettings.GetWorkingHours() == nil {
		return nil
	}

	workingHours := map[string]interface{}{}
	if user.MailboxSettings.GetWorkingHours().GetTimeZone() != nil {
		workingHours["time_zone"] = user.MailboxSettings.GetWorkingHours().GetTimeZone()
	}
	if user.MailboxSettings.GetWorkingHours().GetDaysOfWeek() != nil {
		days := []string{}
		for _, day := range user.MailboxSettings.GetWorkingHours().GetDaysOfWeek() {
			days = append(days, day.String())
		}
		workingHours["days_of_week"] = days
	}
	if user.MailboxSettings.GetWorkingHours().GetStartTime() != nil {
		workingHours["start_time"] = *user.MailboxSettings.GetWorkingHours().GetStartTime()
	}
	if user.MailboxSettings.GetWorkingHours().GetEndTime() != nil {
		workingHours["end_time"] = *user.MailboxSettings.GetWorkingHours().GetEndTime()
	}
	return workingHours
}

func (user *Microsoft365UserInfo) GetDelegateMeetingMessageDeliveryOptions() string {
	if user.MailboxSettings == nil || user.MailboxSettings.GetDelegateMeetingMessageDeliveryOptions() == nil {
		return ""
	}
	return user.MailboxSettings.GetDelegateMeetingMessageDeliveryOptions().String()
}

// SharePoint Settings transform methods
func (sp *Microsoft365SharePointSettingsInfo) SharePointSettingsDetails() map[string]interface{} {
	result := map[string]interface{}{}

	if sp.GetId() != nil {
		result["id"] = *sp.GetId()
	}
	if sp.GetOdataType() != nil {
		result["odata_type"] = *sp.GetOdataType()
	}
	if sp.GetIsCommentingOnSitePagesEnabled() != nil {
		result["is_commenting_on_site_pages_enabled"] = *sp.GetIsCommentingOnSitePagesEnabled()
	}
	if sp.GetIsFileActivityNotificationEnabled() != nil {
		result["is_file_activity_notification_enabled"] = *sp.GetIsFileActivityNotificationEnabled()
	}
	if sp.GetIsLegacyAuthProtocolsEnabled() != nil {
		result["is_legacy_auth_protocols_enabled"] = *sp.GetIsLegacyAuthProtocolsEnabled()
	}
	if sp.GetIsLoopEnabled() != nil {
		result["is_loop_enabled"] = *sp.GetIsLoopEnabled()
	}
	if sp.GetIsMacSyncAppEnabled() != nil {
		result["is_mac_sync_app_enabled"] = *sp.GetIsMacSyncAppEnabled()
	}
	if sp.GetIsRequireAcceptingUserToMatchInvitedUserEnabled() != nil {
		result["is_require_accepting_user_to_match_invited_user_enabled"] = *sp.GetIsRequireAcceptingUserToMatchInvitedUserEnabled()
	}
	if sp.GetIsResharingByExternalUsersEnabled() != nil {
		result["is_resharing_by_external_users_enabled"] = *sp.GetIsResharingByExternalUsersEnabled()
	}
	if sp.GetIsSharePointMobileNotificationEnabled() != nil {
		result["is_sharepoint_mobile_notification_enabled"] = *sp.GetIsSharePointMobileNotificationEnabled()
	}
	if sp.GetIsSharePointNewsfeedEnabled() != nil {
		result["is_sharepoint_newsfeed_enabled"] = *sp.GetIsSharePointNewsfeedEnabled()
	}
	if sp.GetIsSiteCreationEnabled() != nil {
		result["is_site_creation_enabled"] = *sp.GetIsSiteCreationEnabled()
	}
	if sp.GetIsSiteCreationUIEnabled() != nil {
		result["is_site_creation_ui_enabled"] = *sp.GetIsSiteCreationUIEnabled()
	}
	if sp.GetIsSitePagesCreationEnabled() != nil {
		result["is_site_pages_creation_enabled"] = *sp.GetIsSitePagesCreationEnabled()
	}
	if sp.GetIsSitesStorageLimitAutomatic() != nil {
		result["is_sites_storage_limit_automatic"] = *sp.GetIsSitesStorageLimitAutomatic()
	}
	if sp.GetIsSyncButtonHiddenOnPersonalSite() != nil {
		result["is_sync_button_hidden_on_personal_site"] = *sp.GetIsSyncButtonHiddenOnPersonalSite()
	}
	if sp.GetIsUnmanagedSyncAppForTenantRestricted() != nil {
		result["is_unmanaged_sync_app_for_tenant_restricted"] = *sp.GetIsUnmanagedSyncAppForTenantRestricted()
	}
	if sp.GetPersonalSiteDefaultStorageLimitInMB() != nil {
		result["personal_site_default_storage_limit_in_mb"] = *sp.GetPersonalSiteDefaultStorageLimitInMB()
	}
	if sp.GetSharingAllowedDomainList() != nil {
		result["sharing_allowed_domain_list"] = sp.GetSharingAllowedDomainList()
	}
	if sp.GetSharingBlockedDomainList() != nil {
		result["sharing_blocked_domain_list"] = sp.GetSharingBlockedDomainList()
	}
	if sp.GetSharingCapability() != nil {
		result["sharing_capability"] = sp.GetSharingCapability().String()
	}
	if sp.GetSharingDomainRestrictionMode() != nil {
		result["sharing_domain_restriction_mode"] = sp.GetSharingDomainRestrictionMode().String()
	}
	if sp.GetSiteCreationDefaultManagedPath() != nil {
		result["site_creation_default_managed_path"] = *sp.GetSiteCreationDefaultManagedPath()
	}
	if sp.GetSiteCreationDefaultStorageLimitInMB() != nil {
		result["site_creation_default_storage_limit_in_mb"] = *sp.GetSiteCreationDefaultStorageLimitInMB()
	}
	if sp.GetTenantDefaultTimezone() != nil {
		result["tenant_default_timezone"] = *sp.GetTenantDefaultTimezone()
	}
	if sp.GetAllowedDomainGuidsForSyncApp() != nil {
		result["allowed_domain_guids_for_sync_app"] = sp.GetAllowedDomainGuidsForSyncApp()
	}
	if sp.GetAvailableManagedPathsForSiteCreation() != nil {
		result["available_managed_paths_for_site_creation"] = sp.GetAvailableManagedPathsForSiteCreation()
	}
	if sp.GetDeletedUserPersonalSiteRetentionPeriodInDays() != nil {
		result["deleted_user_personal_site_retention_period_in_days"] = *sp.GetDeletedUserPersonalSiteRetentionPeriodInDays()
	}
	if sp.GetExcludedFileExtensionsForSyncApp() != nil {
		result["excluded_file_extensions_for_sync_app"] = sp.GetExcludedFileExtensionsForSyncApp()
	}
	if sp.GetImageTaggingOption() != nil {
		result["image_tagging_option"] = sp.GetImageTaggingOption().String()
	}
	if sp.GetIdleSessionSignOut() != nil {
		result["idle_session_sign_out"] = sp.GetIdleSessionSignOut()
	}

	return result
}

// Authentication Settings transform methods
func (auth *Microsoft365AuthenticationSettingsInfo) AuthenticationSettingsDetails() map[string]interface{} {
	result := map[string]interface{}{}

	if auth.GetId() != nil {
		result["id"] = *auth.GetId()
	}
	if auth.GetDescription() != nil {
		result["description"] = *auth.GetDescription()
	}
	if auth.GetDisplayName() != nil {
		result["display_name"] = *auth.GetDisplayName()
	}
	if auth.GetPolicyVersion() != nil {
		result["policy_version"] = *auth.GetPolicyVersion()
	}
	if auth.GetLastModifiedDateTime() != nil {
		result["last_modified_date_time"] = *auth.GetLastModifiedDateTime()
	}
	if auth.GetPolicyMigrationState() != nil {
		result["policy_migration_state"] = *auth.GetPolicyMigrationState()
	}
	if auth.GetRegistrationEnforcement() != nil {
		registrationEnforcement := map[string]interface{}{}
		regEnf := auth.GetRegistrationEnforcement()
		if regEnf.GetOdataType() != nil {
			registrationEnforcement["odata_type"] = *regEnf.GetOdataType()
		}
		if regEnf.GetAuthenticationMethodsRegistrationCampaign() != nil {
			registrationEnforcement["authentication_methods_registration_campaign"] = regEnf.GetAuthenticationMethodsRegistrationCampaign()
		}
		result["registration_enforcement"] = registrationEnforcement
	}
	if auth.GetAuthenticationMethodConfigurations() != nil {
		var authConfigs []map[string]interface{}
		for _, config := range auth.GetAuthenticationMethodConfigurations() {
			configData := map[string]interface{}{}
			if config.GetId() != nil {
				configData["id"] = *config.GetId()
			}
			if config.GetOdataType() != nil {
				configData["odata_type"] = *config.GetOdataType()
			}
			if config.GetState() != nil {
				configData["state"] = *config.GetState()
			}
			if config.GetExcludeTargets() != nil {
				configData["exclude_targets"] = config.GetExcludeTargets()
			}
			authConfigs = append(authConfigs, configData)
		}
		result["authentication_method_configurations"] = authConfigs
	}

	return result
}

// Security Defaults Settings transform methods
func (sec *Microsoft365SecurityDefaultsSettingsInfo) SecurityDefaultsSettingsDetails() map[string]interface{} {
	result := map[string]interface{}{}

	if sec.GetId() != nil {
		result["id"] = *sec.GetId()
	}
	if sec.GetDisplayName() != nil {
		result["display_name"] = *sec.GetDisplayName()
	}
	if sec.GetDescription() != nil {
		result["description"] = *sec.GetDescription()
	}
	if sec.GetIsEnabled() != nil {
		result["is_enabled"] = *sec.GetIsEnabled()
	}

	return result
}

// Group transform methods
func (g *Microsoft365GroupInfo) GroupAssignedLicenses() []map[string]interface{} {
	var result []map[string]interface{}
	if g.GetAssignedLicenses() != nil {
		for _, license := range g.GetAssignedLicenses() {
			licenseData := map[string]interface{}{}
			if license.GetDisabledPlans() != nil {
				licenseData["disabled_plans"] = license.GetDisabledPlans()
			}
			if license.GetSkuId() != nil {
				licenseData["sku_id"] = *license.GetSkuId()
			}
			result = append(result, licenseData)
		}
	}
	return result
}

func (g *Microsoft365GroupInfo) GroupAssignedLabels() []map[string]interface{} {
	var result []map[string]interface{}
	if g.GetAssignedLabels() != nil {
		for _, label := range g.GetAssignedLabels() {
			labelData := map[string]interface{}{}
			if label.GetLabelId() != nil {
				labelData["label_id"] = *label.GetLabelId()
			}
			if label.GetDisplayName() != nil {
				labelData["display_name"] = *label.GetDisplayName()
			}
			result = append(result, labelData)
		}
	}
	return result
}

func (g *Microsoft365GroupInfo) GroupOnPremisesProvisioningErrors() []map[string]interface{} {
	var result []map[string]interface{}
	if g.GetOnPremisesProvisioningErrors() != nil {
		for _, error := range g.GetOnPremisesProvisioningErrors() {
			errorData := map[string]interface{}{}
			if error.GetCategory() != nil {
				errorData["category"] = *error.GetCategory()
			}
			if error.GetPropertyCausingError() != nil {
				errorData["property_causing_error"] = *error.GetPropertyCausingError()
			}
			if error.GetValue() != nil {
				errorData["value"] = *error.GetValue()
			}
			if error.GetOccurredDateTime() != nil {
				errorData["occurred_date_time"] = *error.GetOccurredDateTime()
			}
			result = append(result, errorData)
		}
	}
	return result
}

func (g *Microsoft365GroupInfo) GroupServiceProvisioningErrors() []map[string]interface{} {
	var result []map[string]interface{}
	if g.GetServiceProvisioningErrors() != nil {
		for _, error := range g.GetServiceProvisioningErrors() {
			errorData := map[string]interface{}{}
			if error.GetCreatedDateTime() != nil {
				errorData["created_date_time"] = *error.GetCreatedDateTime()
			}
			if error.GetIsResolved() != nil {
				errorData["is_resolved"] = *error.GetIsResolved()
			}
			if error.GetServiceInstance() != nil {
				errorData["service_instance"] = *error.GetServiceInstance()
			}
			if error.GetBackingStore() != nil {
				errorData["backing_store"] = error.GetBackingStore()
			}
			result = append(result, errorData)
		}
	}
	return result
}

// Teams Settings transform methods
func (teams *Microsoft365TeamsSettingsInfo) TeamsSettingsDetails() map[string]interface{} {
	result := map[string]interface{}{}
	result["teams_count"] = teams.TeamsCount
	result["note"] = teams.Note
	return result
}

// Planner Settings transform methods
func (planner *Microsoft365PlannerSettingsInfo) PlannerSettingsDetails() map[string]interface{} {
	result := map[string]interface{}{}
	result["plans_count"] = planner.PlansCount
	result["note"] = planner.Note
	return result
}

// Teamwork Settings transform methods
func (teamwork *Microsoft365TeamworkSettingsInfo) TeamworkSettingsDetails() map[string]interface{} {
	result := map[string]interface{}{}
	result["workforce_integrations_count"] = teamwork.WorkforceIntegrationsCount
	result["note"] = teamwork.Note
	return result
}

// Security Settings transform methods
func (security *Microsoft365SecurityInfo) SecuritySettingsDetails() map[string]interface{} {
	result := map[string]interface{}{}

	// Get alerts count
	if security.GetAlerts() != nil {
		result["alerts_count"] = len(security.GetAlerts())
	} else {
		result["alerts_count"] = 0
	}

	// Get secure scores count
	if security.GetSecureScores() != nil {
		result["secure_scores_count"] = len(security.GetSecureScores())
	} else {
		result["secure_scores_count"] = 0
	}

	// Get secure score control profiles count
	if security.GetSecureScoreControlProfiles() != nil {
		result["secure_score_control_profiles_count"] = len(security.GetSecureScoreControlProfiles())
	} else {
		result["secure_score_control_profiles_count"] = 0
	}

	// Get subject rights requests count
	if security.GetSubjectRightsRequests() != nil {
		result["subject_rights_requests_count"] = len(security.GetSubjectRightsRequests())
	} else {
		result["subject_rights_requests_count"] = 0
	}

	// Check if attack simulation is available
	if security.GetAttackSimulation() != nil {
		result["attack_simulation_available"] = true
	} else {
		result["attack_simulation_available"] = false
	}

	// Check if data security and governance is available
	if security.GetDataSecurityAndGovernance() != nil {
		result["data_security_and_governance_available"] = true
	} else {
		result["data_security_and_governance_available"] = false
	}

	return result
}
