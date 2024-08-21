package struct_interface

import (
	"crypto/sha1"
	"compress/gzip"
	"fmt"
	"io"
	"net/http"
)

// shalContent calculates the SHA1 hash of the content at the given URL.
// It handles both plain text and gzipped content.
// It returns the hash as a hex-encoded string.

func main() {
	url := "http://httpbin.org/gzip"
	hash, err := shalContent(url)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("SHA1 hash of", url, "is", hash)
}

func shalContent(url string) (string, error) {
	resp, err := http. Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	var r io.Reader = resp.Body
	if resp.Header.Get("Content-Encoding") == "gzip" {
		var err error
		r, err = gzip.NewReader(r)

		if err != nil {
			return "", err
		}
	}

	// Calculate the SHA1 hash
	h := sha1.New()
	if _, err := io.Copy(h, r); err != nil {
		return "", err
	}

	return fmt.Sprintf("%x", h.Sum(nil)), nil
}

