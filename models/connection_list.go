package models

import (
	"encoding/hex"

	"github.com/gofrs/uuid"
)

type ConnectionList struct {
	Id         *uuid.UUID `json:"id"`
	Version    string     `json:"version"`
	Favourites string     `json:"favourites"`
	Host       string     `json:"host"`
	Username   string     `json:"username"`
	Password   string     `json:"password"`
}

type ConnectionLists []*ConnectionList

func NewConnectionList(version, favourites, host, username, password string) *ConnectionList {
	id, _ := uuid.NewV4()
	return &ConnectionList{
		Id:         &id,
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

func (c *ConnectionList) GetDecodePassword() string {
	bu, _ := hex.DecodeString(c.Password)
	return string(bu)
}
