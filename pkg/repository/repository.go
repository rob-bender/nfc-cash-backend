package repository

import (
	"time"

	"github.com/jmoiron/sqlx"
	"github.com/rob-bender/nfc-cash-backend/appl_row"
)

type TodoAuth interface {
	CreateUser(userForm appl_row.CreateUser) (string, int, error)
	GetUser(username string, password string) ([]appl_row.User, error)
	CheckEmailExist(userForm appl_row.CheckEmailExist) (bool, int, error)
	CheckUsernameExist(userForm appl_row.CheckUsernameExist) (bool, int, error)
	CheckConfirmAccount(userForm appl_row.CheckConfirmAccount) (bool, int, error)
	AddRefreshToken(id int, refreshToken string, expiresAt time.Time) (int, error)
	GetUserIdByRefreshToken(refreshToken string) (int, int, error)
}

type TodoVerify interface {
	CheckEmailVerify(uid string) (bool, int, error)
	EmailVerify(uid string) (bool, int, error)
}

type TodoRecovery interface {
	GetUserUidByEmail(userForm appl_row.RecoveryPasswordSendMessage) (string, int, error)
	CheckRecoveryPassword(uid string) (bool, int, error)
	LaunchRecoveryPassword(uid string) (int, error)
	CompleteRecoveryPassword(uid string) (int, error)
	RecoveryPasswordCompare(uid string, password string) (bool, int, error)
	RecoveryPassword(userForm appl_row.RecoveryPassword) (bool, int, error)
}

type TodoRoom interface {
	CreateRoom() (string, int, error)
	JoinRoom(uidRoom string, uidUser string) (string, int, error)
	LeaveRoom(uidRoom string, uidUser string) (int, error)
	GetRoom(uidRoom string) ([]appl_row.Room, int, error)
}

type TodoMessage interface {
	CreateMessage(messageForm appl_row.CreateMessage) (bool, int, error)
	GetRoomMessages(uidRoom string) ([]appl_row.Message, int, error)
}

type TodoOrder interface {
	OrderCreate(orderForm appl_row.OrderCreate) (bool, int, error)
	GetOrder(uidOrder string) ([]appl_row.Order, int, error)
	GetOrders() ([]appl_row.Orders, int, error)
}

type TodoIp interface {
	BlockIp(ipAddress string) (bool, int, error)
	CheckIpBlock(ipAddress string) (bool, int, error)
}

type TodoUser interface {
	GetUserProfile(id int) ([]appl_row.UserProfile, int, error)
	CheckIsAdmin(id int) (bool, int, error)
}

type TodoAdmin interface {
	GetUsersConfirm(id int) ([]appl_row.GetUsersConfirm, int, error)
	GetUsersUnConfirm(id int) ([]appl_row.GetUsersUnConfirm, int, error)
	GetUserProfile(id int) ([]appl_row.UserProfile, int, error)
	UserConfirmAccount(id int, userForm appl_row.UserConfirmAccount) (bool, int, error)
	ChangeUser(id int, userForm appl_row.ChangeUser) (bool, int, error)
}

type TodoTelegram interface {
	BotCreate(id int, botForm appl_row.BotCreate) (bool, int, error)
	BotDelete(id int, botForm appl_row.BotDelete) (bool, int, error)
	TurnOnBot(id int, token string) (bool, int, error)
	SwitchOffBot(id int, token string) (bool, int, error)
	GetBots(id int) ([]appl_row.Bot, int, error)
}

type Repository struct {
	TodoAuth
	TodoVerify
	TodoRecovery
	TodoRoom
	TodoMessage
	TodoOrder
	TodoIp
	TodoUser
	TodoAdmin
	TodoTelegram
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		TodoAuth:     NewAuthPostgres(db),
		TodoVerify:   NewVerifyPostgres(db),
		TodoRecovery: NewRecoveryPostgres(db),
		TodoRoom:     NewRoomPostgres(db),
		TodoMessage:  NewMessagePostgres(db),
		TodoOrder:    NewOrderPostgres(db),
		TodoIp:       NewIpPostgres(db),
		TodoUser:     NewUserPostgres(db),
		TodoAdmin:    NewAdminPostgres(db),
		TodoTelegram: NewTelegramPostgres(db),
	}
}
