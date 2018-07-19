// Copyright 2012 The Go Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// This example demonstrates parsing HTML data and walking the resulting tree.
package dom_test

import (
	"fmt"
	"log"
	"strings"

	"github.com/factorapp/dom"
)

func ExampleParse() {
	s := `<p>Links:</p><ul><li><a href="foo">Foo</a><li><a href="/bar/baz">BarBaz</a></ul>`
	doc, err := dom.Parse(strings.NewReader(s))
	if err != nil {
		log.Fatal(err)
	}
	var f func(*dom.Node)
	f = func(n *dom.Node) {
		if n.Type == dom.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key == "href" {
					fmt.Println(a.Val)
					break
				}
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(doc)
	// Output:
	// foo
	// /bar/baz
}
