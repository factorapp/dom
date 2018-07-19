package dom

import (
	"github.com/factorapp/dom/js"
)

type DomNode interface {
	EventTarget

	// properties

	BaseURI() string
	NodeName() string
	ChildNodes() NodeList
	ParentNode() *Node
	ParentElement() *Element
	TextContent() string
	SetTextContent(s string)

	// methods

	AppendChild(n *Node)
	Contains(n *Node) bool
	IsEqualNode(n *Node) bool
	IsSameNode(n *Node) bool
	RemoveChild(n *Node)
	ReplaceChild(n, old *Node) *Node
}

type NodeList []*Element

func (e *Node) JSRef() js.Ref {
	return e.v.JSRef()
}

func (e *Node) Remove() {
	e.ParentNode().RemoveChild(e)
	for _, c := range e.callbacks {
		c.Release()
	}
	e.callbacks = nil
}

func (e *Node) AddEventListenerFlags(typ string, flags int, h EventHandler) {
	cb := js.NewEventCallbackFlags(flags, func(v js.Value) {
		h(convertEvent(v))
	})
	e.callbacks = append(e.callbacks, cb)
	e.v.Call("addEventListener", typ, cb)
}
func (e *Node) AddEventListener(typ string, h EventHandler) {
	e.AddEventListenerFlags(typ, 0, h)
}

func (e *Node) BaseURI() string {
	return e.v.Get("baseURI").String()
}

func (e *Node) NodeName() string {
	return e.v.Get("nodeName").String()
}

func (e *Node) ChildNodes() NodeList {
	return AsNodeList(e.v.Get("childNodes"))
}

func (e *Node) ParentNode() *Node {
	return e.Parent
}

func (e *Node) ParentElement() *Element {

	return &Element{
		&Node{
			Type: ElementNode,
			v:    e.v.Get("parentElement"),
		},
	}
}

func (e *Node) TextContent() string {
	return e.v.Get("textContent").String()
}

func (e *Node) SetTextContent(s string) {
	e.v.Set("textContent", s)
}

func (e *Node) Contains(n *Node) bool {
	return e.v.Call("contains", n.JSRef()).Bool()
}

func (e *Node) IsEqualNode(n *Node) bool {
	return e.v.Call("isEqualNode", n.JSRef()).Bool()
}

func (e *Node) IsSameNode(n *Node) bool {
	return e.v.Call("isSameNode", n.JSRef()).Bool()
}

func (e *Node) ReplaceChild(n, old *Node) *Node {
	//return AsElement(e.v.Call("replaceChild", n.JSRef(), old.JSRef()))
	return &Node{} // TODO
}
