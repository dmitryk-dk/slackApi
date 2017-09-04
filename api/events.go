package api

import "github.com/dmitryk-dk/slackbot/models"

type Event struct {
	Type string `json:"type"`
}

type Hello struct{ Event }

type Message struct {
	Event

	Channel string `json:"channel"`
	User    string `json:"user"`
	Text    string `json:"text"`
	TS      string `json:"ts"`

	Subtype string `json:"subtype,omitempty"`
	Hidden  bool   `json:"hidden,omitempty"`

	DeletedTS string `json:"deleted_ts,omitempty"`
	EventTS   string `json:"event_ts,omitempty"`

	Edited struct {
		User string `json:"user"`
		TS   string `json:"ts"`
	} `json:"edited,omitempty"`
}

type ChannelMarked struct {
	Event
	Channel string `json:"channel"`
	TS      string `json:"ts"`
}

type ChannelCreated struct {
	Event
	Channel struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Created int    `json:"created"`
		Creator string `json:"creator"`
	} `json:"channel"`
}

type ChannelJoined struct {
	Event
	Channel models.Channel `json:"channel"`
}

type ChannelLeft struct {
	Event
	Channel string `json:"channel"`
}

type ChannelDeleted struct {
	Event
	Channel string `json:"channel"`
}

type ChannelRename struct {
	Event
	Channel struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Created int    `json:"created"`
	}
}

type ChannelArchive struct {
	Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}

type ChannelUnarchive struct {
	Event
	Channel string `json:"channel"`
	User    string `json:"user"`
}

type ChannelHistoryChanged struct {
	Event
	Latest  string `json:"latest"`
	TS      string `json:"ts"`
	EventTS string `json:"event_ts"`
}

type IMCreated struct {
	Event
	User    string        `json:"user"`
	Channel models.Channel `json:"channel"`
}

type IMOpen struct {
	Event
	User    string `json:"user"`
	Channel string `json:"channel"`
}

type IMClose struct {
	Event
	User    string `json:"user"`
	Channel string `json:"channel"`
}

type IMMarked struct {
	Event
	Channel string `json:"channel"`
	TS      string `json:"ts"`
}

type IMHistoryChanged struct {
	Event
	Latest  string `json:"latest"`
	TS      string `json:"ts"`
	EventTS string `json:"event_ts"`
}



