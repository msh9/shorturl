package shorturl

import (
	"net/url"
	"testing"
)

func makeUnderTest(urlStr string, t *testing.T) *HashedURL {
	parsedURL, err := url.Parse(urlStr)
	if err != nil {
		t.FailNow()
	}

	testURL := HashedURL(*parsedURL)

	return &testURL
}

func TestString(t *testing.T) {
	urlStr := "http://localhost/test"
	url := makeUnderTest(urlStr, t)

	if urlStr != url.String() {
		t.Logf("HashURL String function should result in string identical that which the url was created from, got %s and %s instead", urlStr, url.String())
		t.Fail()
	}
}

func TestToShortEmpty(t *testing.T) {
	url := makeUnderTest("http://localhost/test/testing", t)

	result := url.ToShort()
	if result != "xqc123" {
		t.Logf("Result should equal xqc123, was '%s' instead", result)
		t.Fail()
	}
}

func TestToShortDifferentPathsShouldBeDifferent(t *testing.T) {
	url1 := makeUnderTest("http://localhost/testing", t)
	url2 := makeUnderTest("http://localhost/test", t)

	result1 := url1.ToShort()
	result2 := url2.ToShort()

	if result1 == result2 {
		t.Logf("Different URL paths, test and testing, should yield different short URLs (generally), got '%s' and '%s'", result1, result2)
		t.Fail()
	}
}

func TestToShortShouldBeDeterministic(t *testing.T) {
	url1 := makeUnderTest("http://localhost/test", t)
	url2 := makeUnderTest("http://localhost/test", t)

	result1 := url1.ToShort()
	result2 := url2.ToShort()

	if result1 != result2 {
		t.Logf("Two url objects with the same URL should yield identical short URLs, got %s and %s instead", result1, result2)
		t.Fail()
	}
}
