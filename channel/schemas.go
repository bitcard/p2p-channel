package channel

import (
	"encoding/json"
	"github.com/cyber-rhizome/p2p-channel/msglist"
)

/* constant */
const (
	chatMESSAGE   string = "chat"
	syncMESSAGE   string = "sync"
	stateMESSAGE  string = "state"
	peerJOIN      string = "join"
	peerCONNECTED string = "connected"
	peerLEAVE     string = "leave"
	addPEER       string = "add"
	unixMinute    int64  = 60
)

///////////////////
/* channel level */
///////////////////
// one per channel, recore channel meta data
type channelData struct {
	ChannelName string `storm:"id"`
	ChannelType string `storm:"index"`
	PeerList    map[string]struct{}
}

// many chat to one user, recore one chat message
type chatConent string

// state: INPUTTING, INSERT, DELETE
// comes with sync, For rope structure
type docContent struct {
	Action   string
	Position int
	Length   int
	Edit     string
}

////////////////////////////////////
/* communication between channels */
////////////////////////////////////

// data sent between users in the channel
// message types chatMESSAGE, syncMESSAGE, stateMESSAGE
type pubsubRawMessage struct {
	MsgType string
	Data    json.RawMessage
}

type pubsubMessage struct {
	MsgType string
	Sender  string
	Data    json.RawMessage
}

// sync every 5 minute, or when user log in, send my message latest id list
type syncRawMessage struct {
	TargerID     string
	TargetName   string
	P2pAddrs     []string
	LastSyncTime int64
	LatestMsgID  int
}

type syncMessage struct {
	SenderID     string
	TargerID     string
	TargetName   string
	P2pAddrs     []string
	LastSyncTime int64
	LatestMsgID  int
}

type syncRawState struct {
	P2pAddrs []string
	State    string
}

type syncState struct {
	SenderID string
	P2pAddrs []string
	State    string
}

type chatRawMessage struct {
	ID        int
	Content   string
	CreatedAt int64
}

type chatMessage struct {
	ID        int
	AuthorID  string
	Content   string
	CreatedAt int64
}

///////////////////////////
/* sync between channels */
///////////////////////////

// ask for missing messages
type syncRequest struct {
	ChannelName   string
	TargerID      string
	MissingMsgIDs []int
}

// response for asking
type syncResponse struct {
	ChannelName string
	TargerID    string
	MissingMsg  msglist.Message
}
