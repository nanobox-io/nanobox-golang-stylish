package stylish

import "testing"

func TestShortenProgress(t *testing.T) {
	if shortenProgress(20, ProcessStart("this is a progress")) != "+ this is a progress ---------------- >\n" {
		t.Errorf("shortenProgress didnt do it right but instead returned %s", shortenProgress(20, ProcessStart("this is a progress")))
	}
}