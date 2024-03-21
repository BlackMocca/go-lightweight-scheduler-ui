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

func (list ConnectionLists) FindById(id *uuid.UUID) int {
	var index = -1
	if len(list) > 0 {
		for index, item := range list {
			if item.Id.String() == id.String() {
				return index
			}
		}
	}
	return index
}

func (list ConnectionLists) Remove(index int) []*ConnectionList {
	if index == 0 && len(list) == 1 {
		return make([]*ConnectionList, 0)
	}
	newLists := append(list[:index], list[index+1:]...)

	return newLists
}
