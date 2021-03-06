// Code generated by protoc-gen-go.
// source: missle.proto
// DO NOT EDIT!

/*
Package pb is a generated protocol buffer package.

It is generated from these files:
	missle.proto

It has these top-level messages:
	Point
	Player
	CAuth
	EAuth
	CMatchEnter
	CMatchExit
	EMatcInit
	EMatchEnd
	EMatchTurn
	CPlayerMove
	EPlayerMove
	CPlayerFire
	EPlayerFire
	CPlayerHit
	EPlayerHit
*/
package pb

import proto "code.google.com/p/goprotobuf/proto"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = math.Inf

type Code int32

const (
	Code_C_AUTH          Code = 11
	Code_E_AUTH          Code = 12
	Code_C_MATCH_ENTER   Code = 21
	Code_E_MATCH_INIT    Code = 22
	Code_E_MATCH_TURN    Code = 23
	Code_E_MATCH_END     Code = 24
	Code_C_MATCH_EXIT    Code = 25
	Code_C_PLAYER_MOVE   Code = 31
	Code_E_PLAYER_MOVE   Code = 32
	Code_C_PLAYER_FIRE   Code = 33
	Code_E_PLAYER_FIRE   Code = 34
	Code_C_PLAYER_HIT    Code = 35
	Code_E_PLAYER_HIT    Code = 36
	Code_C_PLAYER_HEALTH Code = 37
)

var Code_name = map[int32]string{
	11: "C_AUTH",
	12: "E_AUTH",
	21: "C_MATCH_ENTER",
	22: "E_MATCH_INIT",
	23: "E_MATCH_TURN",
	24: "E_MATCH_END",
	25: "C_MATCH_EXIT",
	31: "C_PLAYER_MOVE",
	32: "E_PLAYER_MOVE",
	33: "C_PLAYER_FIRE",
	34: "E_PLAYER_FIRE",
	35: "C_PLAYER_HIT",
	36: "E_PLAYER_HIT",
	37: "C_PLAYER_HEALTH",
}
var Code_value = map[string]int32{
	"C_AUTH":          11,
	"E_AUTH":          12,
	"C_MATCH_ENTER":   21,
	"E_MATCH_INIT":    22,
	"E_MATCH_TURN":    23,
	"E_MATCH_END":     24,
	"C_MATCH_EXIT":    25,
	"C_PLAYER_MOVE":   31,
	"E_PLAYER_MOVE":   32,
	"C_PLAYER_FIRE":   33,
	"E_PLAYER_FIRE":   34,
	"C_PLAYER_HIT":    35,
	"E_PLAYER_HIT":    36,
	"C_PLAYER_HEALTH": 37,
}

func (x Code) Enum() *Code {
	p := new(Code)
	*p = x
	return p
}
func (x Code) String() string {
	return proto.EnumName(Code_name, int32(x))
}
func (x *Code) UnmarshalJSON(data []byte) error {
	value, err := proto.UnmarshalJSONEnum(Code_value, data, "Code")
	if err != nil {
		return err
	}
	*x = Code(value)
	return nil
}

type Point struct {
	X                *float32 `protobuf:"fixed32,1,req,name=x" json:"x,omitempty"`
	Y                *float32 `protobuf:"fixed32,2,req,name=y" json:"y,omitempty"`
	XXX_unrecognized []byte   `json:"-"`
}

func (m *Point) Reset()         { *m = Point{} }
func (m *Point) String() string { return proto.CompactTextString(m) }
func (*Point) ProtoMessage()    {}

func (m *Point) GetX() float32 {
	if m != nil && m.X != nil {
		return *m.X
	}
	return 0
}

func (m *Point) GetY() float32 {
	if m != nil && m.Y != nil {
		return *m.Y
	}
	return 0
}

type Player struct {
	Id               *int64  `protobuf:"varint,1,req,name=id" json:"id,omitempty"`
	NickName         *string `protobuf:"bytes,2,req,name=nickName" json:"nickName,omitempty"`
	Avatar           *string `protobuf:"bytes,3,req,name=avatar" json:"avatar,omitempty"`
	IsLeft           *bool   `protobuf:"varint,4,req,name=isLeft" json:"isLeft,omitempty"`
	Position         *Point  `protobuf:"bytes,5,req,name=position" json:"position,omitempty"`
	Health           *int32  `protobuf:"varint,6,req,name=health" json:"health,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *Player) Reset()         { *m = Player{} }
func (m *Player) String() string { return proto.CompactTextString(m) }
func (*Player) ProtoMessage()    {}

func (m *Player) GetId() int64 {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return 0
}

func (m *Player) GetNickName() string {
	if m != nil && m.NickName != nil {
		return *m.NickName
	}
	return ""
}

func (m *Player) GetAvatar() string {
	if m != nil && m.Avatar != nil {
		return *m.Avatar
	}
	return ""
}

func (m *Player) GetIsLeft() bool {
	if m != nil && m.IsLeft != nil {
		return *m.IsLeft
	}
	return false
}

func (m *Player) GetPosition() *Point {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *Player) GetHealth() int32 {
	if m != nil && m.Health != nil {
		return *m.Health
	}
	return 0
}

type CAuth struct {
	UserName         *string `protobuf:"bytes,1,req,name=userName" json:"userName,omitempty"`
	Password         *string `protobuf:"bytes,2,req,name=password" json:"password,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *CAuth) Reset()         { *m = CAuth{} }
func (m *CAuth) String() string { return proto.CompactTextString(m) }
func (*CAuth) ProtoMessage()    {}

func (m *CAuth) GetUserName() string {
	if m != nil && m.UserName != nil {
		return *m.UserName
	}
	return ""
}

func (m *CAuth) GetPassword() string {
	if m != nil && m.Password != nil {
		return *m.Password
	}
	return ""
}

type EAuth struct {
	Code             *int32  `protobuf:"varint,1,req,name=code" json:"code,omitempty"`
	UserId           *int64  `protobuf:"varint,2,req,name=userId" json:"userId,omitempty"`
	Message          *string `protobuf:"bytes,3,opt,name=message" json:"message,omitempty"`
	XXX_unrecognized []byte  `json:"-"`
}

func (m *EAuth) Reset()         { *m = EAuth{} }
func (m *EAuth) String() string { return proto.CompactTextString(m) }
func (*EAuth) ProtoMessage()    {}

func (m *EAuth) GetCode() int32 {
	if m != nil && m.Code != nil {
		return *m.Code
	}
	return 0
}

func (m *EAuth) GetUserId() int64 {
	if m != nil && m.UserId != nil {
		return *m.UserId
	}
	return 0
}

func (m *EAuth) GetMessage() string {
	if m != nil && m.Message != nil {
		return *m.Message
	}
	return ""
}

type CMatchEnter struct {
	XXX_unrecognized []byte `json:"-"`
}

func (m *CMatchEnter) Reset()         { *m = CMatchEnter{} }
func (m *CMatchEnter) String() string { return proto.CompactTextString(m) }
func (*CMatchEnter) ProtoMessage()    {}

type CMatchExit struct {
	MatchId          *int64 `protobuf:"varint,1,req,name=matchId" json:"matchId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CMatchExit) Reset()         { *m = CMatchExit{} }
func (m *CMatchExit) String() string { return proto.CompactTextString(m) }
func (*CMatchExit) ProtoMessage()    {}

func (m *CMatchExit) GetMatchId() int64 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

type EMatcInit struct {
	MatchId          *int64    `protobuf:"varint,1,req,name=matchId" json:"matchId,omitempty"`
	Players          []*Player `protobuf:"bytes,2,rep,name=players" json:"players,omitempty"`
	Points           []*Point  `protobuf:"bytes,3,rep,name=points" json:"points,omitempty"`
	XXX_unrecognized []byte    `json:"-"`
}

func (m *EMatcInit) Reset()         { *m = EMatcInit{} }
func (m *EMatcInit) String() string { return proto.CompactTextString(m) }
func (*EMatcInit) ProtoMessage()    {}

func (m *EMatcInit) GetMatchId() int64 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *EMatcInit) GetPlayers() []*Player {
	if m != nil {
		return m.Players
	}
	return nil
}

func (m *EMatcInit) GetPoints() []*Point {
	if m != nil {
		return m.Points
	}
	return nil
}

type EMatchEnd struct {
	MatchId          *int64 `protobuf:"varint,1,req,name=matchId" json:"matchId,omitempty"`
	Points           *int32 `protobuf:"varint,2,req,name=points" json:"points,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EMatchEnd) Reset()         { *m = EMatchEnd{} }
func (m *EMatchEnd) String() string { return proto.CompactTextString(m) }
func (*EMatchEnd) ProtoMessage()    {}

func (m *EMatchEnd) GetMatchId() int64 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *EMatchEnd) GetPoints() int32 {
	if m != nil && m.Points != nil {
		return *m.Points
	}
	return 0
}

type EMatchTurn struct {
	MatchId          *int64 `protobuf:"varint,1,req,name=matchId" json:"matchId,omitempty"`
	PlayerId         *int64 `protobuf:"varint,2,req,name=playerId" json:"playerId,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EMatchTurn) Reset()         { *m = EMatchTurn{} }
func (m *EMatchTurn) String() string { return proto.CompactTextString(m) }
func (*EMatchTurn) ProtoMessage()    {}

func (m *EMatchTurn) GetMatchId() int64 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *EMatchTurn) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

type CPlayerMove struct {
	MatchId          *int64 `protobuf:"varint,1,req,name=matchId" json:"matchId,omitempty"`
	Position         *Point `protobuf:"bytes,2,req,name=position" json:"position,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CPlayerMove) Reset()         { *m = CPlayerMove{} }
func (m *CPlayerMove) String() string { return proto.CompactTextString(m) }
func (*CPlayerMove) ProtoMessage()    {}

func (m *CPlayerMove) GetMatchId() int64 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *CPlayerMove) GetPosition() *Point {
	if m != nil {
		return m.Position
	}
	return nil
}

type EPlayerMove struct {
	MatchId          *int64 `protobuf:"varint,1,req,name=matchId" json:"matchId,omitempty"`
	PlayerId         *int64 `protobuf:"varint,2,req,name=playerId" json:"playerId,omitempty"`
	Position         *Point `protobuf:"bytes,3,req,name=position" json:"position,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EPlayerMove) Reset()         { *m = EPlayerMove{} }
func (m *EPlayerMove) String() string { return proto.CompactTextString(m) }
func (*EPlayerMove) ProtoMessage()    {}

func (m *EPlayerMove) GetMatchId() int64 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *EPlayerMove) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *EPlayerMove) GetPosition() *Point {
	if m != nil {
		return m.Position
	}
	return nil
}

type CPlayerFire struct {
	MatchId          *int64 `protobuf:"varint,1,req,name=matchId" json:"matchId,omitempty"`
	PlayerId         *int64 `protobuf:"varint,2,req,name=playerId" json:"playerId,omitempty"`
	Position         *Point `protobuf:"bytes,3,req,name=position" json:"position,omitempty"`
	Velocity         *Point `protobuf:"bytes,4,req,name=velocity" json:"velocity,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CPlayerFire) Reset()         { *m = CPlayerFire{} }
func (m *CPlayerFire) String() string { return proto.CompactTextString(m) }
func (*CPlayerFire) ProtoMessage()    {}

func (m *CPlayerFire) GetMatchId() int64 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *CPlayerFire) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *CPlayerFire) GetPosition() *Point {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *CPlayerFire) GetVelocity() *Point {
	if m != nil {
		return m.Velocity
	}
	return nil
}

type EPlayerFire struct {
	MatchId          *int64 `protobuf:"varint,1,req,name=matchId" json:"matchId,omitempty"`
	PlayerId         *int64 `protobuf:"varint,2,req,name=playerId" json:"playerId,omitempty"`
	Position         *Point `protobuf:"bytes,3,req,name=position" json:"position,omitempty"`
	Velocity         *Point `protobuf:"bytes,4,req,name=velocity" json:"velocity,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EPlayerFire) Reset()         { *m = EPlayerFire{} }
func (m *EPlayerFire) String() string { return proto.CompactTextString(m) }
func (*EPlayerFire) ProtoMessage()    {}

func (m *EPlayerFire) GetMatchId() int64 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *EPlayerFire) GetPlayerId() int64 {
	if m != nil && m.PlayerId != nil {
		return *m.PlayerId
	}
	return 0
}

func (m *EPlayerFire) GetPosition() *Point {
	if m != nil {
		return m.Position
	}
	return nil
}

func (m *EPlayerFire) GetVelocity() *Point {
	if m != nil {
		return m.Velocity
	}
	return nil
}

type CPlayerHit struct {
	MatchId          *int64 `protobuf:"varint,1,req,name=matchId" json:"matchId,omitempty"`
	P1               *int64 `protobuf:"varint,2,req,name=p1" json:"p1,omitempty"`
	P2               *int64 `protobuf:"varint,3,req,name=p2" json:"p2,omitempty"`
	Damage           *int32 `protobuf:"varint,4,req,name=damage" json:"damage,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *CPlayerHit) Reset()         { *m = CPlayerHit{} }
func (m *CPlayerHit) String() string { return proto.CompactTextString(m) }
func (*CPlayerHit) ProtoMessage()    {}

func (m *CPlayerHit) GetMatchId() int64 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *CPlayerHit) GetP1() int64 {
	if m != nil && m.P1 != nil {
		return *m.P1
	}
	return 0
}

func (m *CPlayerHit) GetP2() int64 {
	if m != nil && m.P2 != nil {
		return *m.P2
	}
	return 0
}

func (m *CPlayerHit) GetDamage() int32 {
	if m != nil && m.Damage != nil {
		return *m.Damage
	}
	return 0
}

type EPlayerHit struct {
	MatchId          *int64 `protobuf:"varint,1,req,name=matchId" json:"matchId,omitempty"`
	P1               *int64 `protobuf:"varint,2,req,name=p1" json:"p1,omitempty"`
	P2               *int64 `protobuf:"varint,3,req,name=p2" json:"p2,omitempty"`
	Damage           *int32 `protobuf:"varint,4,req,name=damage" json:"damage,omitempty"`
	XXX_unrecognized []byte `json:"-"`
}

func (m *EPlayerHit) Reset()         { *m = EPlayerHit{} }
func (m *EPlayerHit) String() string { return proto.CompactTextString(m) }
func (*EPlayerHit) ProtoMessage()    {}

func (m *EPlayerHit) GetMatchId() int64 {
	if m != nil && m.MatchId != nil {
		return *m.MatchId
	}
	return 0
}

func (m *EPlayerHit) GetP1() int64 {
	if m != nil && m.P1 != nil {
		return *m.P1
	}
	return 0
}

func (m *EPlayerHit) GetP2() int64 {
	if m != nil && m.P2 != nil {
		return *m.P2
	}
	return 0
}

func (m *EPlayerHit) GetDamage() int32 {
	if m != nil && m.Damage != nil {
		return *m.Damage
	}
	return 0
}

func init() {
	proto.RegisterEnum("pb.Code", Code_name, Code_value)
}
