package dom

import "github.com/factorapp/dom/js"

func GetDocument() *Document {
	doc := js.Get("document")
	if !doc.Valid() {
		return nil
	}
	return &Document{&Node{v: doc}}
}

var _ DomNode = (*Document)(nil)

type Document struct {
	*Node
}

func (d *Document) CreateElement(tag string) *Element {
	v := d.v.Call("createElement", tag)
	return &Element{
		&Node{
			Type: ElementNode,
			Data: tag,
			v:    v,
		},
	}
}

func (d *Document) CreateElementNS(ns string, tag string) *Element {
	v := d.v.Call("createElementNS", ns, tag)
	return AsElement(v)
}
func (d *Document) GetElementsByTagName(tag string) NodeList {
	v := d.v.Call("getElementsByTagName", tag)
	return AsNodeList(v)
}
func (d *Document) QuerySelector(qu string) *Element {
	v := d.v.Call("querySelector", qu)
	return AsElement(v)
}
func (d *Document) QuerySelectorAll(qu string) NodeList {
	v := d.v.Call("querySelectorAll", qu)
	return AsNodeList(v)
}
