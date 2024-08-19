var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func generateRandomString(length int) string {
	randomString := make([]byte, length)

	for i := range randomString {
		randomString[i] = chars[rand.IntN(len(chars))]
	}

	return string(randomString)
}

type Codec struct {
	urls map[string]string
}

func Constructor() Codec {
	return Codec{
		urls: make(map[string]string),
	}
}

// Encodes a URL to a shortened URL.
func (c *Codec) encode(longUrl string) string {
	const (
		urlBase        = "http://tinyurl.com/"
		endpointLength = 6
	)

	endpoint := generateRandomString(endpointLength)

	for {
		if _, ok := c.urls[endpoint]; ok {
			endpoint = generateRandomString(endpointLength)
		} else {
			break
		}
	}

	newUrl := strings.Builder{}

	newUrl.Grow(len(urlBase) + len(endpoint))

	newUrl.WriteString(urlBase)
	newUrl.WriteString(endpoint)

	c.urls[newUrl.String()] = longUrl

	return newUrl.String()
}

// Decodes a shortened URL to its original URL.
func (c *Codec) decode(shortUrl string) string {
	if longUrl, ok := c.urls[shortUrl]; ok {
		return longUrl
	}
	return ""
}