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

type User struct {
	ID string `json:"id"`

	TeamID string `json:"team_id"`
	Name   string `json:"name"`
	Color  string `json:"color"`

	Profile Profile `json:"profile"`

	Deleted           bool `json:"deleted"`
	IsBot             bool `json:"is_bot"`
	IsAdmin           bool `json:"is_admin"`
	IsOwner           bool `json:"is_owner"`
	IsPrimaryOwner    bool `json:"is_primary_owner"`
	IsRestricted      bool `json:"is_restricted"`
	IsUltraRestricted bool `json:"is_ultra_restricted"`

	Has2FA   bool `json:"has_2fa"`
	HasFiles bool `json:"has_files"`
}

type Profile struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	BotID     string `json:"bot_id"`

	RealName           string `json:"real_name"`
	RealNameNormalized string `json:"real_name_normalized"`

	Title string `json:"title"`

	Skype string `json:"skype"`
	Phone string `json:"phone"`

	Image24       string `json:"image_24"`
	Image32       string `json:"image_32"`
	Image48       string `json:"image_48"`
	Image72       string `json:"image_72"`
	Image192      string `json:"image_192"`
	Image512      string `json:"image_512"`
	Image1024     string `json:"image_1024"`
	ImageOriginal string `json:"image_original"`
}
