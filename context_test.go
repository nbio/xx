package xx

import (
	"encoding/xml"
	"testing"

	"github.com/nbio/st"
)

func TestContextAttr(t *testing.T) {
	ctx := Context{
		StartElement: xml.StartElement{
			Name: xml.Name{Space: "urn:ietf:params:xml:ns:epp-1.0", Local: "example"},
			Attr: []xml.Attr{
				{Name: xml.Name{Space: "urn:ietf:params:xml:ns:epp-1.0", Local: "avail"}, Value: "1"},
				{Name: xml.Name{Local: "alpha"}, Value: "TRUE"},
				{Name: xml.Name{Local: "beta"}, Value: "1"},
				{Name: xml.Name{Space: "other", Local: "gamma"}, Value: "false"},
				{Name: xml.Name{Space: "other", Local: "delta"}, Value: "FALSE"},
				{Name: xml.Name{Local: "omega"}, Value: "hammertime"},
				{Name: xml.Name{Local: "number"}, Value: "42"},
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
