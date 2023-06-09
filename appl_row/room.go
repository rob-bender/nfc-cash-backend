package appl_row

type Room struct {
	Id          int      `json:"id"`
	UidRoom     string   `json:"uid_room"`
	MemberCount int      `json:"member_count"`
	IsWorks     bool     `json:"is_works"`
	Members     []string `json:"members"`
	MemberFixed string   `json:"member_fixed"`
}

type LeaveRoom struct {
	UidRoom string `json:"uidRoom"`
	UidUser string `json:"uidUser"`
}
