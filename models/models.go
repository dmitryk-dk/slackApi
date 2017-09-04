package models

type Topic struct {
	Value   string `json:"value"`
	Creator string `json:"creator"`
	LastSet int    `json:"last_set"`
}

type Channel struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	IsChanel   bool     `json:"is_channel"`
	Created    int      `json:"created"`
	Creator    string   `json:"creator"`
	IsArchived bool     `json:"is_archived"`
	IsGeneral  bool     `json:"is_general"`
	Members    []string `json:"members"`
	Topic      Topic
	Purpose    Topic
	IsMember   bool `json:"is_member"`
}

type IM struct {
	Id            string `json:"id"`
	IsIM          bool   `json:"is_im"`
	User          string `json:"user"`
	Created       int    `json:"created"`
	IsUserDeleted bool   `json:"is_user_deleted"`
}
