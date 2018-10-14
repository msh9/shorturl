package shorturl

import (
	"encoding/ascii85"
	"encoding/binary"
	"hash/crc32"
	"net/url"
)

// ShortenableURL defines a method to explicitly request a URL be made shorter
type ShortenableURL interface {
	Shorten() *url.URL
}

// HashedURL is a net/url which may be explicitly shortened
type HashedURL url.URL

// ToShort renders a non-unique, but deterministically repeatable short string version of the URL
func (hashedUrl *HashedURL) ToShort(base string) string {
	return hashedUrl.Shorten(base).String()
}

// String strings a HashURL type
func (hashedUrl *HashedURL) String() string {
	thisURL := url.URL(*hashedUrl)

	return thisURL.String()
}

// Shorten renders a non-unique, but deterministically repeatable short version of the URL
func (hashedUrl *HashedURL) Shorten(base string) *url.URL {
	completeString := hashedUrl.String()

	checksum := crc32.Checksum([]byte(completeString), crc32.IEEETable)
	checkSumBuf := make([]byte, 4)
	binary.LittleEndian.PutUint32(checkSumBuf, checksum)
	encodedShort := make([]byte, ascii85.MaxEncodedLen(len(checkSumBuf)))
	ascii85.Encode(encodedShort, []byte(checkSumBuf))
	urlPath := string(encodedShort)

	url, err := url.Parse(base + "/" + urlPath)
	if err != nil {

	}
}
