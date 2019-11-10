package main

import "time"

type Task struct {
	Title       string
	Created     time.Time
	Description string
	Due         time.Time
	Modified    time.Time
	Project     string
	Tags        []string
	UUID        string
	VirtualTags []string
}
