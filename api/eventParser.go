package api

import "encoding/json"

func EventParser (msg []byte) (interface{}, error){
	var event = Event{}

	if err := json.Unmarshal(msg, &event); err != nil {
		return nil, err
	}

	switch event.Type {
		case "channel_archive":
			var event = ChannelArchive{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
		case "channel_created":
			var event = ChannelCreated{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
		case "channel_deleted":
			var event = ChannelDeleted{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
		case "channel_history_changed":
			var event = ChannelHistoryChanged{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
		case "channel_joined":
			var event = ChannelJoined{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
		case "channel_left":
			var event = ChannelLeft{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
		case "channel_marked":
			var event = ChannelMarked{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
		case "channel_rename":
			var event = ChannelRename{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
		case "channel_unarchive":
			var event = ChannelUnarchive{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
		case "hello":
			var event = Hello{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
		case "im_close":
			var event = IMClose{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
		case "im_created":
			var event = IMCreated{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
		case "im_history_changed":
			var event = IMHistoryChanged{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
		case "im_marked":
			var event = IMMarked{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
		case "im_open":
			var event = IMOpen{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
		case "message":
			var event = Message{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
		default:
			var event = Event{}
			if err := json.Unmarshal(msg, &event); err != nil {
				return nil, err
			}
			return event, nil
	}
}
