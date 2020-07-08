package newsfeeder

import "testing"

var feeds = New()

func TestAdd(t *testing.T) {
	item := Item{Title: "title 1", Post: "post 1"}
	feeds.Add(item)
	if len(feeds.Items) == 0 {
		t.Errorf("test case failed")
	}
}

func TestGetall(t *testing.T) {
	feeds.GetAll()
	if len(feeds.Items) == 0 {
		t.Errorf("test case failed")
	}
}
