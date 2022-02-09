package image

import "path"

type Image struct {
	Id         string   `json:"id"`
	Name       string   `json:"name"`
	Tag        string   `json:"tag"`
	Size       string   `json:"size"`
	Counts     int      `json:"counts"`
	Work       string   `json:"work"`
	Hostname   string   `json:"hostname"`
	Entrypoint []string `json:"entrypoint"`
	Command    []string `json:"command"`
	Envs       []string `json:"envs"`
	CreateTime string   `json:"create_time"`
}

func (image Image) RootDir() string {
	return path.Join("", image.Id)
}
