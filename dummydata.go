package main

import "time"

var (
	tasks = []Task{
		Task{
			Title:       "Clean the house",
			Created:     time.Now(),
			Description: "vacuum, mop, blablabla, bla bla blaaaaaaaaaa",
			Due:         time.Now().Add(7 * 24 * time.Hour),
			Modified:    time.Now(),
			Project:     "alltag",
			Tags:        []string{"duties"},
			UUID:        "????????",
			VirtualTags: []string{},
		},
		Task{
			Title:       "Win the lottery",
			Created:     time.Now(),
			Description: "play and win millions.. dudududmmd udm dudmdm",
			Due:         time.Now().Add(33333 * time.Hour),
			Modified:    time.Now(),
			Project:     "dreams",
			Tags:        []string{"gambling"},
			UUID:        "????????",
			VirtualTags: []string{},
		},
		Task{
			Title:       "Wash the car",
			Created:     time.Now(),
			Description: "water & soap, water & soap, water & soap, and leather",
			Due:         time.Now().Add(7 * 24 * time.Hour),
			Modified:    time.Now(),
			Project:     "alltag",
			Tags:        []string{"duties"},
			UUID:        "????????",
			VirtualTags: []string{},
		},
		Task{
			Title:       "Learn to play harmonica",
			Created:     time.Now(),
			Description: "practise playing, again and again and again and again",
			Due:         time.Now().Add(7 * 24 * time.Hour),
			Modified:    time.Now(),
			Project:     "alltag",
			Tags:        []string{"joy", "hobby", "music"},
			UUID:        "????????",
			VirtualTags: []string{},
		},
	}
)
