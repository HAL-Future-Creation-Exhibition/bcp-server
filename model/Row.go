package model

import "time"

type Row struct {
	CurrentPath string		`json:"current_path"`
	Name        string		`json:"name"`
	IsDir       bool   		`json:"isDir"`
	UpdateTime	time.Time	`json:"update_time"`
}

type Rows []Row

func (r Rows) Len() int {
	return len(r)
}

func (r Rows) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func (r Rows) Less(i, j int) bool {
	return r[i].UpdateTime.Unix() > r[j].UpdateTime.Unix()
}