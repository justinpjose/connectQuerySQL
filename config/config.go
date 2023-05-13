package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

const (
	unsetCMD = `db_unset`
)

// DBConfig represents DB configuration
// Get it from 1Password
// OR to get the Password, get the DBPasswordKey from service env vars
// and then get the value of that from AWS Secret Manager
type DBConfig struct {
	Name     string `envconfig:"DB_NAME" required:"true"`
	Host     string `envconfig:"DB_HOST" required:"true"`
	Password string `envconfig:"DB_PASSWORD" required:"true"`
	Port     int    `envconfig:"DB_PORT" required:"true"`
	Username string `envconfig:"DB_USERNAME" required:"true"`
}

// SSHConfig represents SSH configuration
// Go to DBeaver and get the information through 'Edit connection' and then SSH tab
type SSHConfig struct {
	Host               string `envconfig:"SSH_HOST" required:"true"`
	Port               int    `envconfig:"SSH_PORT" required:"true"`
	User               string `envconfig:"SSH_USER" required:"true"`
	PrivateKeyPath     string `envconfig:"SSH_PRIVATE_KEY_PATH" required:"true"`
	PrivateKeyPassword string `envconfig:"SSH_PRIVATE_KEY_PASSWORD" required:"true"`
}

func Get() (*DBConfig, *SSHConfig, error) {
	var dbCfg DBConfig
	if err := envconfig.Process("", &dbCfg); err != nil {
		return nil, nil, fmt.Errorf("failed to get db config - err: %w", err)
	}

	var sshCfg SSHConfig
	if err := envconfig.Process("", &sshCfg); err != nil {
		return nil, nil, fmt.Errorf("failed to get ssh config - err: %w", err)
	}

	if dbCfg.Host == "" || sshCfg.Host == "" {
		return nil, nil, fmt.Errorf("set config environment variables for db and ssh - db ?")
	}

	return &dbCfg, &sshCfg, nil
}

func Unset() {
	log.Print()
	log.Print("ALERT: REMEMBER TO UNSET ENV_VARS - db_unset")
}
