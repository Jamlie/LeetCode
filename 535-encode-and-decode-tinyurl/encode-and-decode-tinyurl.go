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
	endpointLength := 6

	endpoint := generateRandomString(endpointLength)

	for {
		if _, ok := c.urls[endpoint]; ok {
			endpoint = generateRandomString(endpointLength)
		} else {
			break
		}
	}

	var (
		urlBase = "http://tinyurl.com/"
		newUrl  = strings.Builder{}
	)

	newUrl.Grow(len(urlBase) + len(endpoint))

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


/**
 * Your Codec object will be instantiated and called as such:
 * obj := Constructor();
 * url := obj.encode(longUrl);
 * ans := obj.decode(url);
 */
