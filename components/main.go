package main

import (
	"syscall/js"

	"github.com/PaulRosset/go-hacknews"

	"fmt"
)

func topStories() []hacknews.Post {
	init := hacknews.Initializer{Story: "topstories", NbPosts: 30}
	codes, err := init.GetCodesStory()
	if err != nil {
		panic(err)
	}
	posts, err := init.GetPostStory(codes)
	if err != nil {
		panic(err)
	}
	return posts
}

var document = js.Global().Get("document")

func renderPost(post hacknews.Post, parent js.Value) {
	str1 := post.Score
	str2 := "points,"
	str3 := "by"
	str4 := post.By
	scoreAndBy := fmt.Sprintf("%d %s %s %s", str1, str2, str3, str4)

	li := document.Call("createElement", "li")
	a := document.Call("createElement", "a")
	div := document.Call("createElement", "div")

	text := document.Call("createTextNode", post.Title)
	score := document.Call("createTextNode", scoreAndBy)

	div.Call("appendChild", score)
	a.Set("href", post.Url)
	a.Call("appendChild", text)
	li.Call("appendChild", a)
	li.Call("appendChild", div)
	parent.Call("appendChild", li)
}

func renderPosts(posts []hacknews.Post, parent js.Value) {
	ul := document.Call("createElement", "ul")
	parent.Call("appendChild", ul)
	for _, post := range posts {
		renderPost(post, ul)
	}
}

func main() {
	posts := topStories()
	body := document.Get("body")
	renderPosts(posts, body)
}
