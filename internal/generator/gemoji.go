package main

import (
	"encoding/json"
)

const (
	gemojiURL             = "https://raw.githubusercontent.com/github/gemoji/master/db/emoji.json"
	gemojiLatestCommitURL = "https://api.github.com/repos/github/gemoji/commits?path=/db/emoji.json&per_page=1"
)

type gemoji struct {
	Emoji   string   `json:"emoji"`
	Aliases []string `json:"aliases"`
}

func fetchGemojis() (map[string]string, string, error) {
	b, err := fetchData(gemojiURL)
	if err != nil {
		return nil, "", err
	}

	v, err := fetchData(gemojiLatestCommitURL)
	if err != nil {
		return nil, "", err
	}
	var commits []struct {
		SHA string `json:"sha"`
	}
	if err = json.Unmarshal(v, &commits); err != nil {
		return nil, "", err
	}

	var gemojis []gemoji
	r := make(map[string]string)

	if err = json.Unmarshal(b, &gemojis); err != nil {
		return nil, "", err
	}

	for _, gemoji := range gemojis {
		for _, alias := range gemoji.Aliases {
			if len(alias) == 0 || len(gemoji.Emoji) == 0 {
				continue
			}

			r[makeAlias(alias)] = gemoji.Emoji
		}
	}

	return r, commits[0].SHA, nil
}
