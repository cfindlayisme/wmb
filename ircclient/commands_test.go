package ircclient_test

import (
	"fmt"
	"net"
	"testing"

	"github.com/cfindlayisme/wmb/ircclient"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

type MockConn struct {
	mock.Mock
	net.Conn
}

func (m *MockConn) Write(b []byte) (n int, err error) {
	args := m.Called(b)
	return args.Int(0), args.Error(1)
}

func TestSetNick(t *testing.T) {
	conn := new(MockConn)
	conn.On("Write", []byte("NICK nick\r\n")).Return(10, nil)

	err := ircclient.SetNick(conn, "nick")
	require.NoError(t, err)

	conn.AssertExpectations(t)
}

func TestJoinChannel(t *testing.T) {
	conn := new(MockConn)
	conn.On("Write", []byte("JOIN #channel\r\n")).Return(12, nil)

	err := ircclient.JoinChannel(conn, "#channel")
	require.NoError(t, err)

	conn.AssertExpectations(t)
}

func TestPartChannel(t *testing.T) {
	conn := new(MockConn)
	conn.On("Write", []byte("PART #channel\r\n")).Return(12, nil)

	err := ircclient.PartChannel(conn, "#channel")
	require.NoError(t, err)

	conn.AssertExpectations(t)
}

func TestSetMode(t *testing.T) {
	conn := new(MockConn)
	conn.On("Write", []byte("MODE #channel +o\r\n")).Return(15, nil)

	err := ircclient.SetMode(conn, "#channel", "+o")
	require.NoError(t, err)

	conn.AssertExpectations(t)
}
func TestSetTopic(t *testing.T) {
	conn := new(MockConn)
	conn.On("Write", []byte("TOPIC #channel topic\r\n")).Return(20, nil) // Remove the `:` before `topic`

	err := ircclient.SetTopic(conn, "#channel", "topic")
	require.NoError(t, err)

	conn.AssertExpectations(t)
}

func TestInviteUser(t *testing.T) {
	conn := new(MockConn)
	conn.On("Write", []byte("INVITE nick #channel\r\n")).Return(20, nil)

	err := ircclient.InviteUser(conn, "nick", "#channel")
	require.NoError(t, err)

	conn.AssertExpectations(t)
}

func TestKickUser(t *testing.T) {
	conn := new(MockConn)
	conn.On("Write", []byte("KICK #channel nick :Kicked\r\n")).Return(26, nil)

	err := ircclient.KickUser(conn, "nick", "#channel", "")
	require.NoError(t, err)

	conn.AssertExpectations(t)
}

func TestQuote(t *testing.T) {
	conn := new(MockConn)
	conn.On("Write", []byte("command\r\n")).Return(9, nil) // Remove the `QUOTE ` before `command`

	err := ircclient.Quote(conn, "command")
	require.NoError(t, err)

	conn.AssertExpectations(t)
}

func TestSendMessage(t *testing.T) {
	conn := new(MockConn)
	conn.On("Write", []byte("PRIVMSG #channel :message\r\n")).Return(26, nil)

	err := ircclient.SendMessage(conn, "#channel", "message")
	require.NoError(t, err)

	conn.AssertExpectations(t)
}

func TestSendNotice(t *testing.T) {
	conn := new(MockConn)
	conn.On("Write", []byte("NOTICE #channel :message\r\n")).Return(24, nil)

	err := ircclient.SendNotice(conn, "#channel", "message")
	require.NoError(t, err)

	conn.AssertExpectations(t)
}

func TestSetUser(t *testing.T) {
	conn := new(MockConn)
	conn.On("Write", []byte("USER wmb 0 * :Webhook message bot\r\r\n")).Return(37, nil)

	err := ircclient.SetUser(conn)
	require.NoError(t, err)

	conn.AssertExpectations(t)
}

func TestSendQuit(t *testing.T) {
	conn := new(MockConn)
	quitMessage := "Client Quit"
	conn.On("Write", []byte(fmt.Sprintf("QUIT :%s\r\n", quitMessage))).Return(len(quitMessage)+7, nil)

	err := ircclient.SendQuit(conn, quitMessage)
	require.NoError(t, err)

	conn.AssertExpectations(t)
}
