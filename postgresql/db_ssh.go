package postgresql

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
	"net"
	"time"

	"connectQuerySQL/config"

	"github.com/jmoiron/sqlx"
	"github.com/lib/pq"
	"golang.org/x/crypto/ssh"

	internalSSH "connectQuerySQL/ssh"
)

const (
	postgresSSHDriverName = "postgres+ssh"
)

type ViaSSHDialer struct {
	client *ssh.Client
}

func (self *ViaSSHDialer) Open(s string) (_ driver.Conn, err error) {
	return pq.DialOpen(self, s)
}

func (self *ViaSSHDialer) Dial(network, address string) (net.Conn, error) {
	return self.client.Dial(network, address)
}

func (self *ViaSSHDialer) DialTimeout(network, address string, timeout time.Duration) (net.Conn, error) {
	return self.client.Dial(network, address)
}

// OpenSQLViaSSH - opens a SQL database via SSH and returns its handler
// ensure to use close function and set as defer after using this method
func OpenSQLViaSSH(dbCfg config.DBConfig, sshCfg config.SSHConfig) (*sql.DB, func(), error) {
	sshConn, err := internalSSH.OpenSSHConn(sshCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open ssh connection for db - err: %v", err)
	}

	// Now we register the ViaSSHDialer with the ssh connection as a parameter
	sql.Register(postgresSSHDriverName, &ViaSSHDialer{sshConn})

	db, err := openSQL(dbCfg, true)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open database - err: %v", err)
	}

	close := func() {
		sshConn.Close()
		db.Close()
	}

	return db, close, nil
}

// OpenSQLXViaSSH - opens a SQLX database via SSH and returns its handler
// ensure to use close function and set as defer after using this method
func OpenSQLXViaSSH(dbCfg config.DBConfig, sshCfg config.SSHConfig) (*sqlx.DB, func(), error) {
	sshConn, err := internalSSH.OpenSSHConn(sshCfg)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open ssh connection for db - err: %v", err)
	}

	// Now we register the ViaSSHDialer with the ssh connection as a parameter
	sql.Register(postgresSSHDriverName, &ViaSSHDialer{sshConn})

	db, err := openSQLX(dbCfg, true)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to open database - err: %v", err)
	}

	close := func() {
		sshConn.Close()
		db.Close()
	}

	return db, close, nil
}
