package models

import "encoding/hex"

type ConnectionList struct {
	Favourites string `json:"favourites"`
	Host       string `json:"host"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

func NewConnectionList(favourites, host, username, password string) *ConnectionList {
	return &ConnectionList{
		Favourites: favourites,
		Host:       host,
		Username:   username,
		Password:   password,
	}
}

func (c *ConnectionList) GetEncodePassword() string {
	return hex.EncodeToString([]byte(c.Password))
}
