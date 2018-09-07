package models

import (
	"net/url"
	"testing"
)

func makeUnderTest(urlStr string, path string, t *testing.T) URL {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		t.FailNow()
	}
	url := URL{
		url:  parsedURL,
		path: path,
	}

	return url
}

func TestToShortEmpty(t *testing.T) {
	url := makeUnderTest("http://localhost/test/testing", "test/testing", t)

	result := url.ToShort()
	if result != "xqc123" {
		t.Logf("Result should equal xqc123, was '%s' instead", result)
		t.Fail()
	}
}

func TestToShortDifferentPathsShouldBeDifferent(t *testing.T) {
	url1 := makeUnderTest("http://localhost/testing", "testing", t)
	url2 := makeUnderTest("http://localhost/test", "test", t)

	result1 := url1.ToShort()
	result2 := url2.ToShort()

	if result1 == result2 {
		t.Logf("Different URL paths, test and testing, should yield different short URLs (generally), got '%s' and '%s'", result1, result2)
		t.Fail()
	}
}
