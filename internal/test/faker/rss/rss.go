package rss

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

const Version = "2.0"
const ItemCount = 25
const LinkTemplate = "https://faker:8080/feed/%s"

type RSS struct {
	XMLName xml.Name `xml:"rss"`
	Version string   `xml:"version,attr"`
	Channel Channel  `xml:"channel"`
}

type Channel struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	Items       []Item `xml:"item"`
}

type Item struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func uuidToSeed(id string) (int64, error) {
	uuid, err := uuid.Parse(id)
	if err != nil {
		return 0, err
	}
	hash := sha256.Sum256(uuid[:])
	seed := int64(binary.BigEndian.Uint64(hash[:8]))
	return seed, nil
}

func generateChannel(seed int64, id string) Channel {
	gofakeit.Seed(seed)
	channel := Channel{
		Title:       gofakeit.BuzzWord(),
		Link:        fmt.Sprintf(LinkTemplate, id),
		Description: gofakeit.Sentence(10),
		Items:       generateItems(),
	}
	return channel
}

func generateItems() []Item {
	gofakeit.Seed(0)
	items := make([]Item, ItemCount)
	for i := range items {
		items[i] = Item{
			Title:       gofakeit.Sentence(5),
			Link:        gofakeit.URL(),
			Description: gofakeit.Paragraph(1, 2, 5, " "),
			PubDate:     time.Now().Add(time.Duration(-i) * time.Hour).Format(time.RFC1123),
		}
	}
	return items
}

// Generate a deterministic RSS channel based on a UUID.
// Items within the channel are random.
func Generate(id string) (*RSS, error) {
	seed, err := uuidToSeed(id)
	if err != nil {
		return nil, err
	}

	return &RSS{
		Version: Version,
		Channel: generateChannel(seed, id),
	}, nil
}
