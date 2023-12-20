package models

import (
	"encoding/hex"
)

type ConnectionList struct {
	Version    string `json:"version"`
	Favourites string `json:"favourites"`
	Host       string `json:"host"`
	Username   string `json:"username"`
	Password   string `json:"password"`
}

type ConnectionLists []*ConnectionList

func NewConnectionList(version, favourites, host, username, password string) *ConnectionList {
	return &ConnectionList{
		Version:    version,
		Favourites: favourites,
		Host:       host,
		Username:   username,
		Password:   password,
	}
}

func (c *ConnectionList) GetEncodePassword() string {
	return hex.EncodeToString([]byte(c.Password))
}
