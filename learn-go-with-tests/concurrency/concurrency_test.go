package concurrency

import (
	"reflect"
	"testing"
	"time"
)

func mockWebsiteChecker(url string) bool {
	if url == "google.com" {
		return false
	}
	return true
}

func TestCheckWebsites(t *testing.T) {
	websites := []string{
		"google.com",
		"xxx.com",
		"yyy.com",
	}

	want := map[string]bool{
		"google.com": false,
		"xxx.com":    true,
		"yyy.com":    true,
	}

	got := CheckWebsites(mockWebsiteChecker, websites)

	if !reflect.DeepEqual(want, got) {
		t.Fatalf("wanted %v, got %v", want, got)
	}
}

func slowStubWebsiteChecker(_ string) bool {
	time.Sleep(20 * time.Millisecond)
	return true
}

func BenchmarkCheckWebsites(b *testing.B) {

	urls := make([]string, 20)

	for i := 0; i < len(urls); i++ {
		urls[i] = "a url"
	}

	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		CheckWebsites(slowStubWebsiteChecker, urls)
	}
}
