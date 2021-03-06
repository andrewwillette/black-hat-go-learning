package main

import (
	"fmt"

	"github.com/spf13/viper"
	"github.com/stacktitan/ldapauth"
)

type User struct {
	Username  string
	Firstname string
	IsAdmin   bool
}

func (u *User) Attributes() []string {
	return []string{"givenName"}
}

func main() {
	viper.AutomaticEnv()
	viper.SetConfigName("ldap")
	viper.addConfigPath(".")

	ldp := &ldapauth.LDAP{
		Address: viper.GetString("ldap.address"),
		UID:     viper.GetString("ldap.uid"),
	}
	fmt.Println("hello world")
}
