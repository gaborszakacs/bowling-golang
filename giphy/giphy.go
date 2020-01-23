package giphy

import (
	libgiphy "github.com/sanzaru/go-giphy"
)

type Client struct {
	APIkey string
}

func (c Client) Random(topic string) (string, error) {
	giphy := libgiphy.NewGiphy(c.APIkey)

	dataRandom, err := giphy.GetRandom(topic)
	if err != nil {
		return "", err
	}
	return dataRandom.Data.Url, nil
}
