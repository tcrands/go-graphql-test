package main

var albums []Album = []Album{
	Album{
		ID:     "0",
		Artist: "1",
		Title:  "Fearless",
		Year:   "2008",
		Type:   "album",
	},
	Album{
		ID:     "1",
		Artist: "123123",
		Title:  "Fearle123123ss",
		Year:   "2008",
		Type:   "album",
	},
}

var artists []Artist = []Artist{
	Artist{
		ID:   "1",
		Name: "Taylor Swift",
		Type: "artist",
	},
}

var songs []Song = []Song{
	Song{
		ID:       "1",
		Album:    "ts-fearless",
		Title:    "Fearless",
		Duration: "4:01",
		Type:     "song",
	},
	Song{
		ID:       "2",
		Album:    "fearless",
		Title:    "Fifteen",
		Duration: "4:54",
		Type:     "song",
	},
}
