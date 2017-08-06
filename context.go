package xx

import (
	"encoding/xml"
	"strconv"
)

// Context holds XML scanning context.
type Context struct {
	Decoder      *xml.Decoder
	Value        interface{}
	StartElement xml.StartElement
	CharData     xml.CharData
	Parent       *Context
}

// Attr returns the string value for the XML attributed named space:local.
// Pass an empty string in space to match any namespace.
func (ctx *Context) Attr(space, local string) string {
	for _, a := range ctx.StartElement.Attr {
		if (space == "" || a.Name.Space == space) && a.Name.Local == local {
			return a.Value
		}
	}
	return ""
}

// AttrBool returns a boolean value for the XML attributed named space:local.
// Pass an empty string in space to match any namespace.
func (ctx *Context) AttrBool(space, local string) bool {
	v := ctx.Attr(space, local)
	if len(v) == 0 || v == "0" || v[0] == 'f' || v[0] == 'F' {
		return false
	}
	return true
}

// AttrInt returns a integer value for the XML attributed named space:local.
// Pass an empty string in space to match any namespace.
func (ctx *Context) AttrInt(space, local string) int {
	v := ctx.Attr(space, local)
	n, _ := strconv.Atoi(v)
	return n
}

// path returns a > delimited path for nested Context(s)
func (ctx *Context) path() string {
	p := ctx.StartElement.Name.Local
	if ctx.Parent != nil {
		p2 := ctx.Parent.path()
		if p2 != "" {
			p = p2 + ">" + p
		}
	}
	return p
}
