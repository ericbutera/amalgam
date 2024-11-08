package rss

import (
	"crypto/sha256"
	"encoding/binary"
	"encoding/xml"
	"fmt"
	"math"
	"math/big"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/google/uuid"
)

const (
	Version      = "2.0"
	ItemCount    = 25
	LinkTemplate = "https://faker:8080/feed/%s"
)

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
	var seed int64

	// note: this is a test helper so it doesn't matter. in production, this would be a bad idea.
	attempt := binary.BigEndian.Uint64(hash[:8])
	if attempt > math.MaxInt64 {
		bigAttempt := new(big.Int).SetUint64(attempt)
		bigMaxInt64 := new(big.Int).SetInt64(math.MaxInt64)
		seed = bigAttempt.Mod(bigAttempt, bigMaxInt64).Int64()
	} else {
		seed = int64(attempt)
	}

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
