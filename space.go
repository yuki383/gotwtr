package gotwtr

// User objects can found and expanded in the user resource.
// These objects are available for expansion by adding at least one of
// host_ids, creator_id, speaker_ids, mentioned_user_ids
// to the expansions query parameter.

// FYI
// https://developer.twitter.com/en/docs/twitter-api/data-dictionary/object-model/space

type SpaceField string

const (
	SpaceFieldHostIDs          SpaceField = "host_ids"
	SpaceFieldCreatedAt        SpaceField = "created_at"
	SpaceFieldCreatorID        SpaceField = "creator_id"
	SpaceFieldID               SpaceField = "id"
	SpaceFieldLanguage         SpaceField = "lang"
	SpaceFieldInvittedUserIDs  SpaceField = "invited_user_ids"
	SpaceFieldParticipantCount SpaceField = "participant_count"
	SpaceFieldSpeakerIDs       SpaceField = "speaker_ids"
	SpaceFieldStartedAt        SpaceField = "started_at"
	SpaceFieldState            SpaceField = "state"
	SpaceFieldTitle            SpaceField = "title"
	SpaceFieldUpdatedAt        SpaceField = "updated_at"
	SpaceFieldScheduledStart   SpaceField = "scheduled_start"
	SpaceFieldIsTicketed       SpaceField = "is_ticketed"
)

func spaceFieldsToString(sfs []SpaceField) []string {
	slice := make([]string, len(sfs))
	for i, sf := range sfs {
		slice[i] = string(sf)
	}
	return slice
}

type Space struct {
	ID               string   `json:"id"`
	State            string   `json:"state"`
	CreatedAt        string   `json:"created_at,omitempty"`
	HostIDs          []string `json:"host_ids,omitempty"`
	Lang             string   `json:"lang,omitempty"`
	IsTicketed       bool     `json:"is_ticketed,omitempty"`
	InvitedUserIDs   []string `json:"invited_user_ids,omitempty"`
	ParticipantCount int      `json:"participant_count,omitempty"`
	ScheduledStart   string   `json:"scheduled_start,omitempty"`
	SpeakerIDs       []string `json:"speaker_ids,omitempty"`
	StartedAt        string   `json:"started_at,omitempty"`
	Title            string   `json:"title,omitempty"`
	UpdatedAt        string   `json:"updated_at,omitempty"`
	CreatorID        string   `json:"creator_id,omitempty"`
}

type SearchSpacesResponse struct {
	Spaces   []*Space            `json:"data,omitempty"`
	Includes *SpaceIncludes      `json:"includes,omitempty"`
	Meta     *SearchSpacesMeta   `json:"meta"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
}

type SearchSpacesMeta struct {
	ResultCount int `json:"result_count"`
}

type SpaceIncludes struct {
	Users []*User
}

type SpaceLookUpByIDResponse struct {
	Space    *Space              `json:"data,omitempty"`
	Includes *SpaceIncludes      `json:"includes,omitempty"`
	Errors   []*APIResponseError `json:"errors,omitempty"`
}
