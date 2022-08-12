package office365

import "github.com/microsoftgraph/msgraph-sdk-go/models"

type Office365CalendarEventInfo struct {
	models.Eventable
}

type Office365DriveInfo struct {
	models.Driveable
}

type Office365DriveItemInfo struct {
	models.DriveItemable
	DriveID string
}

type Office365MailMessageInfo struct {
	models.Messageable
}

type Office365TeamInfo struct {
	models.Teamable
	UserIdentifier string
}

type Office365TeamChannelInfo struct {
	models.Channelable
	TeamID string
}

func (driveItem *Office365DriveItemInfo) DriveItemCreatedBy() map[string]interface{} {
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

func (driveItem *Office365DriveItemInfo) DriveItemFile() map[string]interface{} {
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

func (driveItem *Office365DriveItemInfo) DriveItemLastModifiedBy() map[string]interface{} {
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

func (driveItem *Office365DriveItemInfo) DriveItemParentReference() map[string]interface{} {
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

func (event *Office365CalendarEventInfo) EventAttendees() []map[string]interface{} {
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
			attendeeInfo["status"] = attendeeInfo
		}
		if a.GetProposedNewTime() != nil {
			data := map[string]interface{}{}
			if a.GetProposedNewTime().GetEnd() != nil {
				endTimeInfo := map[string]interface{}{}
				if event.GetEnd().GetDateTime() != nil {
					endTimeInfo["dateTime"] = *event.GetEnd().GetDateTime()
				}
				if event.GetEnd().GetDateTime() != nil {
					endTimeInfo["timeZone"] = *event.GetEnd().GetTimeZone()
				}
				data["end"] = endTimeInfo
			}
			if a.GetProposedNewTime().GetStart() != nil {
				startTimeInfo := map[string]interface{}{}
				if event.GetStart().GetDateTime() != nil {
					startTimeInfo["dateTime"] = *event.GetStart().GetDateTime()
				}
				if event.GetStart().GetDateTime() != nil {
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

func (event *Office365CalendarEventInfo) EventBody() map[string]interface{} {
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

func (event *Office365CalendarEventInfo) EventEnd() map[string]interface{} {
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

func (event *Office365CalendarEventInfo) EventLocation() map[string]interface{} {
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
	if event.GetLocation().GetType() != nil {
		locationInfo["type"] = *event.GetLocation().GetType()
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

func (event *Office365CalendarEventInfo) EventLocations() []map[string]interface{} {
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
		if l.GetType() != nil {
			locationInfo["type"] = *l.GetType()
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

func (event *Office365CalendarEventInfo) EventOnlineMeeting() map[string]interface{} {
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
			if p.GetType() != nil {
				data["type"] = p.GetType().String()
			}
			phones = append(phones, data)
		}
		onlineMeetingInfo["phones"] = phones
	}
	return onlineMeetingInfo
}

func (event *Office365CalendarEventInfo) EventOrganizer() map[string]interface{} {
	if event.GetOrganizer() == nil {
		return nil
	}

	organizerInfo := map[string]interface{}{}
	if event.GetOrganizer().GetType() != nil {
		organizerInfo["type"] = *event.GetOrganizer().GetType()
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

func (event *Office365CalendarEventInfo) EventRecurrence() map[string]interface{} {
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
		if event.GetRecurrence().GetPattern().GetType() != nil {
			patternInfo["type"] = *event.GetRecurrence().GetPattern().GetType()
		}
		recurrenceInfo["pattern"] = patternInfo
	}
	if event.GetRecurrence().GetRange() != nil {
		rangeInfo := map[string]interface{}{
			"type": event.GetRecurrence().GetRange().GetType().String(),
		}
		if event.GetRecurrence().GetRange().GetEndDate() != nil {
			rangeInfo["endDate"] = *event.GetRecurrence().GetRange().GetEndDate()
		}
		if event.GetRecurrence().GetRange().GetNumberOfOccurrences() != nil {
			rangeInfo["numberOfOccurrences"] = *event.GetRecurrence().GetRange().GetNumberOfOccurrences()
		}
		if event.GetRecurrence().GetRange().GetRecurrenceTimeZone() != nil {
			rangeInfo["recurrenceTimeZone"] = event.GetRecurrence().GetRange().GetRecurrenceTimeZone()
		}
		if event.GetRecurrence().GetRange().GetStartDate() != nil {
			rangeInfo["startDate"] = *event.GetRecurrence().GetRange().GetStartDate()
		}
		recurrenceInfo["range"] = rangeInfo
	}

	return recurrenceInfo
}

func (event *Office365CalendarEventInfo) EventResponseStatus() map[string]interface{} {
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

func (event *Office365CalendarEventInfo) EventStart() map[string]interface{} {
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

func (message *Office365MailMessageInfo) MessageAttachments() []map[string]interface{} {
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

func (message *Office365MailMessageInfo) MessageBccRecipients() []map[string]interface{} {
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

func (message *Office365MailMessageInfo) MessageBody() map[string]interface{} {
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

func (message *Office365MailMessageInfo) MessageCcRecipients() []map[string]interface{} {
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

func (message *Office365MailMessageInfo) MessageFrom() map[string]interface{} {
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

func (message *Office365MailMessageInfo) MessageImportance() interface{} {
	if message.GetImportance() == nil {
		return nil
	}
	return message.GetImportance().String()
}

func (message *Office365MailMessageInfo) MessageInferenceClassification() interface{} {
	if message.GetInferenceClassification() == nil {
		return nil
	}
	return message.GetInferenceClassification().String()
}

func (message *Office365MailMessageInfo) MessageReplyTo() []map[string]interface{} {
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

func (message *Office365MailMessageInfo) MessageSender() map[string]interface{} {
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

func (message *Office365MailMessageInfo) MessageToRecipients() []map[string]interface{} {
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

func (team *Office365TeamInfo) TeamMembers() interface{} {
	if team.GetSpecialization() == nil {
		return nil
	}
	return team.GetSpecialization().String()
}

func (team *Office365TeamInfo) TeamSpecialization() interface{} {
	if team.GetSpecialization() == nil {
		return nil
	}
	return team.GetSpecialization().String()
}

func (team *Office365TeamInfo) TeamSummary() map[string]interface{} {
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

func (team *Office365TeamInfo) TeamTemplate() map[string]interface{} {
	if team.GetTemplate() == nil {
		return nil
	}

	template := map[string]interface{}{}
	if team.GetTemplate().GetId() != nil {
		template["id"] = *team.GetTemplate().GetId()
	}
	if team.GetSummary().GetMembersCount() != nil {
		template["type"] = *team.GetTemplate().GetType()
	}

	return template
}

func (team *Office365TeamInfo) TeamVisibility() interface{} {
	if team.GetVisibility() == nil {
		return nil
	}
	return team.GetVisibility().String()
}

func (team *Office365TeamChannelInfo) TeamChannelMembershipType() interface{} {
	if team.GetMembershipType() == nil {
		return nil
	}
	return team.GetMembershipType().String()
}

