package model

type TaskGroup struct {
	Id     int    `json:"id"`
	Name   string `json:"name"`
	Path   string `json:"path"`
	Status string `json:"status"`
}

type TaskGroups struct {
	TaskGroups []TaskGroup `json:"taskgroups"`
}
