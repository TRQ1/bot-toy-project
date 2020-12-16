package tools

import (
	"fmt"
	"os"

	model "github.com/TRQ1/devops-toy-project/ssBot/core/model"
	"github.com/nlopes/slack"
)

func UploadImage(u model.UploadInfo) error {
	t := os.Getenv("SlackToken")
	c := os.Getenv("SlackChannel")
	env := os.Getenv("Environment")
	a := slack.New(t)

	filePath := ""
	if env == "local" {
		filePath = u.ImageName
	} else {
		filePath = "/tmp/" + u.ImageName
	}

	params := slack.FileUploadParameters{
		Title:          u.User + " 님이 요청하신 " + u.ImageName,
		File:           filePath,
		InitialComment: u.Service,
		Channels:       []string{c},
	}

	file, err := a.UploadFile(params)
	if err != nil {
		fmt.Printf("%s\n", err)
		return err
	}
	fmt.Printf("Name: %s is sent.\n", file.Name)

	return nil
}
