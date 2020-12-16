package model

import "github.com/nlopes/slack"

type UploadInfo struct {
	ImageName string
	Service   string
	Board     string
	User      string
}

type SlackClient struct {
	Client            *slack.Client
	VerificationToken string
	ChannelID         string
}

type SlackMsg struct {
	Text       string
	Ts         string
	Channel    string
	Reaction   string
	Translated string
	Source     string
	Target     string
}
