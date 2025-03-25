package main

import (
	"io"
	"net/http"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
)

type BeefSummary struct {
	Beef map[string]int `json:"beef"`
}

// HTTPClient interface for better testability
type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

// DefaultHTTPClient implements HTTPClient
type DefaultHTTPClient struct{}

func (c *DefaultHTTPClient) Get(url string) (*http.Response, error) {
	return http.Get(url)
}

var httpClient HTTPClient = &DefaultHTTPClient{}

// List of valid meat types
var validMeats = map[string]bool{
	"t-bone": true, "fatback": true, "pastrami": true, "pork": true,
	"meatloaf": true, "jowl": true, "bresaola": true, "bacon": true,
	"beef": true, "chicken": true, "turkey": true, "ham": true,
	"salami": true, "prosciutto": true, "ribs": true, "steak": true,
	"sirloin": true, "tenderloin": true, "brisket": true, "short": true,
	"loin": true, "shank": true, "flank": true, "ribeye": true,
	"strip": true, "ground": true, "round": true, "chuck": true,
	"rump": true, "tongue": true, "tri-tip": true, "turducken": true,
	"venison": true, "buffalo": true, "kielbasa": true, "andouille": true,
	"biltong": true, "boudin": true, "capicola": true, "chislic": true,
	"corned": true, "doner": true, "drumstick": true, "frankfurter": true,
	"hamburger": true, "jerky": true, "landjaeger": true, "leberkas": true,
	"meatball": true, "pancetta": true, "picanha": true, "porchetta": true,
	"sausage": true, "shankle": true, "spare": true, "swine": true,
}

func fetchMeatText() (string, error) {
	resp, err := httpClient.Get("https://baconipsum.com/api/?type=meat-and-filler&paras=99&format=text")
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	return string(body), nil
}

func countMeat(text string) map[string]int {
	// Remove punctuation and convert to lowercase
	re := regexp.MustCompile(`[.,]`)
	text = re.ReplaceAllString(text, " ")
	text = strings.ToLower(text)

	// Split into words and count only valid meat types
	words := strings.Fields(text)
	counts := make(map[string]int)

	for _, word := range words {
		if word != "" && validMeats[word] {
			counts[word]++
		}
	}

	return counts
}

func getBeefSummary(c *fiber.Ctx) error {
	text, err := fetchMeatText()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"error": "Failed to fetch meat text",
		})
	}

	counts := countMeat(text)
	summary := BeefSummary{
		Beef: counts,
	}

	return c.JSON(summary)
}

func startMeatAPI() {
	app := fiber.New()
	app.Get("/beef/summary", getBeefSummary)
	app.Listen(":8080")
}
