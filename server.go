package socks

import (
	"bytes"
	"context"
	"crypto/tls"
	"io"
	"net"
	"os"

	"github.com/lucas-clemente/quic-go"
	"github.com/pkg/errors"
)

type Server struct {
	password []byte
	pwdLen   int
	listener quic.Listener
}

func NewServer(address string, tlsConfig *tls.Config, password string) (*Server, error) {
	err := os.Setenv("GODEBUG", "bbr=1")
	if err != nil {
		return nil, err
	}
	tlsConfig.NextProtos = append(tlsConfig.NextProtos, "h2")
	listener, err := quic.ListenAddr(address, tlsConfig, nil)
	if err != nil {
		return nil, errors.WithStack(err)
	}
	l := len(password)
	if l > 32 {
		return nil, errors.New("password size > 32")
	}
	return &Server{
		password: []byte(password),
		pwdLen:   l,
		listener: listener,
	}, nil
}

func (s *Server) ListenAndServe() error {
	for {
		session, err := s.listener.Accept(context.Background())
		if err != nil {
			return err
		}
		go s.handleSession(session)
	}
}

func (s *Server) handleSession(session quic.Session) {
	defer func() { _ = session.Close() }()
	for {
		stream, err := session.AcceptStream(context.Background())
		if err != nil {
			return
		}
		go s.handleStream(stream)
	}
}

func (s *Server) handleStream(stream quic.Stream) {
	defer func() {
		recover()
		_ = stream.Close()
	}()
	str := &deadlineStream{Stream: stream}
	// read password
	pwd := make([]byte, s.pwdLen)
	_, err := io.ReadFull(str, pwd)
	if err != nil {
		return
	}

	if !bytes.Equal(pwd, s.password) {
		// invalid password
		_, _ = str.Write([]byte{respInvalidPWD})
		return
	}

	// get connect host
	address, err := unpackHostData(str)
	if err != nil {
		_, _ = str.Write([]byte{respInvalidHost})
		return
	}

	// connect
	conn, err := net.Dial("tcp", address)
	if err != nil {
		_, _ = str.Write([]byte{respConnectFailed})
		return
	}
	defer func() { _ = conn.Close() }()
	_, _ = str.Write([]byte{respOK})

	// copy
	go func() { _, _ = io.Copy(stream, conn) }()
	_, _ = io.Copy(conn, stream)
}

func (s *Server) Close() {
	_ = s.listener.Close()
}
