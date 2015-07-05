package xx

import (
	"encoding/xml"
	"testing"

	"github.com/nbio/st"
)

func TestContextAttr(t *testing.T) {
	ctx := Context{
		StartElement: xml.StartElement{
			Name: xml.Name{"urn:ietf:params:xml:ns:epp-1.0", "example"},
			Attr: []xml.Attr{
				{xml.Name{"urn:ietf:params:xml:ns:epp-1.0", "avail"}, "1"},
				{xml.Name{"", "alpha"}, "TRUE"},
				{xml.Name{"", "beta"}, "1"},
				{xml.Name{"other", "gamma"}, "false"},
				{xml.Name{"other", "delta"}, "FALSE"},
				{xml.Name{"", "omega"}, "hammertime"},
				{xml.Name{"", "number"}, "42"},
			},
		},
	}

	st.Expect(t, ctx.Attr("", "avail"), "1")
	st.Expect(t, ctx.AttrBool("", "avail"), true)
	st.Expect(t, ctx.AttrInt("", "avail"), 1)
	st.Expect(t, ctx.Attr("other", "avail"), "")
	st.Expect(t, ctx.AttrBool("other", "avail"), false)
	st.Expect(t, ctx.AttrInt("other", "avail"), 0)
	st.Expect(t, ctx.AttrBool("", "alpha"), true)
	st.Expect(t, ctx.AttrBool("", "beta"), true)
	st.Expect(t, ctx.AttrBool("", "gamma"), false)
	st.Expect(t, ctx.AttrBool("", "delta"), false)
	st.Expect(t, ctx.Attr("other", "omega"), "")
	st.Expect(t, ctx.AttrBool("other", "omega"), false)
	st.Expect(t, ctx.Attr("", "omega"), "hammertime")
	st.Expect(t, ctx.AttrBool("", "omega"), true)
	st.Expect(t, ctx.AttrInt("", "number"), 42)
}
