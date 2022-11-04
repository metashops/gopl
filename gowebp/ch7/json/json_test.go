package main

import (
	`testing`
	`time`
)

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Wrong id, was expecting 1 but got", post.Id)
	}
	if post.Content != "Hello World!" {
		t.Error("Wrong Content, Hello World! but got", post.Content)
	}
}

func TestEncode(t *testing.T) {
	if testing.Short() {
		t.Skip("Skip encoding for now")
	}
	time.Sleep(time.Second * 10)
}
