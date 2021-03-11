package newsfeed

import "testing"

func TestAdd(t *testing.T) {
	feed := New()
	feed.Add(Item{
		ID:    1,
		Title: "Gin is pretty trash, ngl",
		Post:  "anf",
	})

	if len(feed.Items) == 0 {
		t.Errorf("Item was not added to feed")
	}
}

func TestGetAll(t *testing.T) {
	feed := New()
	feed.Add(Item{
		ID:    1,
		Title: "Gin is pretty trash, ngl",
		Post:  "anf",
	})

	if len(feed.GetAll()) != 1 {
		t.Errorf("Something is wrong")
	}
}
