package hashTable

import (
	"crypto/md5"
	"encoding/hex"
	"hash"
)

// TinyURL is a URL shortening service where you enter a URL such as
// https://leetcode.com/problems/design-tinyurl and it returns a short URL such
// as http://tinyurl.com/4e9iAk.
// Design the encode and decode methods for the TinyURL service. There is no
// restriction on how your encode/decode algorithm should work. You just need to
// ensure that a URL can be encoded to a tiny URL and the tiny URL can be decoded
// to the original URL.

type Codec struct {
	h       hash.Hash
	tinyMap map[string]string
	prefix  string
}

func Constructor535() Codec {
	return Codec{
		h:       md5.New(),
		tinyMap: make(map[string]string),
		prefix:  "http://tinyurl.com/",
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	this.h.Write([]byte(longUrl))
	tinyStr := hex.EncodeToString(this.h.Sum(nil))[0:8]
	this.tinyMap[tinyStr] = longUrl
	return this.prefix + tinyStr
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	tinyStr := shortUrl[len(shortUrl)-8:]
	return this.tinyMap[tinyStr]
}
