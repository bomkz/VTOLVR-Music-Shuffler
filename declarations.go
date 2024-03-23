package main

type vdflib struct {
	Path                     string
	Label                    string
	ContentID                string
	TotalSize                string
	UpdateCleanBytesTally    string
	TimeLastUpdateCorruption string
	Apps                     []libapp
}

type libapp struct {
	AppID   string
	BuildID string
}

var tick = make(chan bool)
var logLines []string
