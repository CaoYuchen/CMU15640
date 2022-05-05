package tribserver

import (
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/rpc"

	"github.com/cmu440/tribbler/libstore"

	"github.com/cmu440/tribbler/rpc/tribrpc"
)

type tribServer struct {
	// Reference to the tribble server's libstore
	mLib libstore.Libstore
}

// NewTribServer creates, starts and returns a new TribServer. masterServerHostPort
// is the master storage server's host:port and myHostPort is the host:port on which
// the TribServer should listen. acceptRPC determines whether this server should listen
// for RPCs from the TribClient. A non-nil error should be returned if the TribServer
// could not be started.
//
// (FALL 2021) You should not have to modify this function. Doing so might break test cases.
func NewTribServer(masterServerHostPort, myHostPort string, mode libstore.LeaseMode, acceptRPC bool) (TribServer, error) {
	ts := new(tribServer)

	lib, err := libstore.NewLibstore(masterServerHostPort, myHostPort, mode)
	if err != nil {
		return nil, err
	}
	ts.mLib = lib

	// Setup the HTTP handler that will serve incoming RPCs.

	if acceptRPC {

		// Register the TribServer to receive RPCs from the TribClient.
		// See the tribclient package for examples illustrating how the
		// RPCs should be invoked.
		if err := rpc.RegisterName("TribServer", tribrpc.Wrap(ts)); err != nil {
			return nil, err
		}
		rpc.HandleHTTP()

		// Begin listening on the specified port.
		l, err := net.Listen("tcp", myHostPort)
		fmt.Println("Listening on", myHostPort)
		if err != nil {
			return nil, err
		}

		go http.Serve(l, nil)
	}

	return ts, nil
}

func (ts *tribServer) CreateUser(args *tribrpc.CreateUserArgs, reply *tribrpc.CreateUserReply) error {
	return errors.New("not implemented")
}

func (ts *tribServer) AddSubscription(args *tribrpc.SubscriptionArgs, reply *tribrpc.SubscriptionReply) error {
	return errors.New("not implemented")
}

func (ts *tribServer) RemoveSubscription(args *tribrpc.SubscriptionArgs, reply *tribrpc.SubscriptionReply) error {
	return errors.New("not implemented")
}

func (ts *tribServer) GetFriends(args *tribrpc.GetFriendsArgs, reply *tribrpc.GetFriendsReply) error {
	return errors.New("not implemented")
}

func (ts *tribServer) PostTribble(args *tribrpc.PostTribbleArgs, reply *tribrpc.PostTribbleReply) error {
	return errors.New("not implemented")
}

func (ts *tribServer) ModifyTribble(args *tribrpc.ModifyTribbleArgs, reply *tribrpc.ModifyTribbleReply) error {
	return errors.New("not implemented")
}

func (ts *tribServer) DeleteTribble(args *tribrpc.DeleteTribbleArgs, reply *tribrpc.DeleteTribbleReply) error {
	return errors.New("not implemented")
}

func (ts *tribServer) GetTribbles(args *tribrpc.GetTribblesArgs, reply *tribrpc.GetTribblesReply) error {
	return errors.New("not implemented")
}

func (ts *tribServer) GetTribblesBySubscription(args *tribrpc.GetTribblesArgs, reply *tribrpc.GetTribblesReply) error {
	return errors.New("not implemented")
}
