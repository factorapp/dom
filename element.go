package dom

import (
	"fmt"
	sjs "syscall/js"

	"github.com/factorapp/dom/js"
)

var _ DomNode = (*Element)(nil)

func AsElement(v js.Value) *Element {
	if !v.Valid() {
		return nil
	}
	return &Element{&Node{v: v}}
}

func AsNodeList(v js.Value) NodeList {
	if !v.Valid() {
		return nil
	}
	arr := make(NodeList, v.Length())
	for i := range arr {
		arr[i] = AsElement(v.Index(i))
	}
	return arr
}

var _ DomNode = (*Element)(nil)

type Element struct {
	*Node
}

func (e *Element) SetInnerHTML(s string) {
	e.v.Set("innerHTML", s)
}

func (e *Element) SetAttribute(k string, v interface{}) {
	e.v.Call("setAttribute", k, fmt.Sprint(v))
}

func (e *Element) GetAttribute(k string) js.Value {
	return e.v.Call("getAttribute", k)
}

func (e *Element) Style() *Style {
	return &Style{v: e.v.Get("style")}
}

func (e *Element) GetBoundingClientRect() Rect {
	rv := e.v.Call("getBoundingClientRect")
	x, y := rv.Get("x").Int(), rv.Get("y").Int()
	w, h := rv.Get("width").Int(), rv.Get("height").Int()
	return Rect{Min: Point{x, y}, Max: Point{x + w, y + h}}
}

func (e *Element) onMouseEvent(typ string, flags int, h MouseEventHandler) {
	e.AddEventListenerFlags(typ, flags, func(e Event) {
		h(e.(*MouseEvent))
	})
}

func (e *Element) OnClick(h MouseEventHandler) {
	e.onMouseEvent("click", int(sjs.StopPropagation), h)
}

func (e *Element) OnMouseDown(h MouseEventHandler) {
	e.onMouseEvent("mousedown", int(sjs.StopPropagation), h)
}

func (e *Element) OnMouseMove(h MouseEventHandler) {
	e.onMouseEvent("mousemove", 0, h)
}

func (e *Element) OnMouseUp(h MouseEventHandler) {
	e.onMouseEvent("mouseup", int(sjs.StopPropagation), h)
}
