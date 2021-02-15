package tables

import "time"

type AppUsers struct {
	Id        uint64
	Handle    string
	Password  string
	Name      string
	Birthday  time.Time
	Profile   string
	Image     string
	IsPrivate bool
	Disabled  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Tokens struct {
	Id     uint64
	UserId uint64
	Token  string
}

type ToFollows struct {
	Id         uint64
	ToUser     uint64
	ByUser     uint64
	Permission uint8
	CreatedAt  time.Time
	UpdatedAt  time.Time
}

type Posts struct {
	Id        uint64
	UserId    uint64
	Body      string
	Disabled  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type ToPosts struct {
	ByPost uint64
	ToPost uint64
}

type Praises struct {
	Id        uint64
	UserId    uint64
	PostId    uint64
	Disabled  bool
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Diffusions struct {
	UserId    uint64
	PostId    uint64
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Notifications struct {
	Id        uint64
	UserId    uint64
	Type      uint8
	Status    uint8
	CreatedAt time.Time
	UpdatedAt time.Time
}

type NotificationToFollows struct {
	NotificationId uint64
	ToFollowId     uint64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type NotificationPraises struct {
	NotificationId uint64
	PraiseId       uint64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type NotificationDiffusions struct {
	NotificationId uint64
	DiffusionId    uint64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type NotificationMentions struct {
	NotificationId uint64
	PostId         uint64
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type UserGroups struct {
	Id   uint64
	Name string
}

type GroupToUsers struct {
	GroupId uint64
	UserId  uint64
}

type InviteUserToGroups struct {
	GroupId uint64
	UserId  uint64
}

type MessageLogs struct {
	Id      uint64
	UserId  uint64
	IsGroup bool
	Body    string
}

type LogToUsers struct {
	UserId      uint64
	LogId       uint64
	IsConfirmed bool
}

type LogToGroups struct {
	GroupId uint64
	LogId   uint64
}
