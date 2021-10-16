package xx

import (
	"bytes"
	"fmt"
	"io"
	"testing"

	"github.com/nbio/xml"

	"github.com/nbio/st"
)

func decoder(s string) *xml.Decoder {
	return xml.NewDecoder(bytes.NewBufferString(s))
}

func TestScanner(t *testing.T) {
	s := NewScanner()
	s.MustHandleStartElement("epp", func(ctx *Context) error {
		st.Reject(t, ctx.Parent, nil)
		st.Expect(t, ctx.Parent.StartElement.Name, xml.Name{})
		return nil
	})
	s.MustHandleStartElement("epp>response>result", func(ctx *Context) error {
		st.Expect(t, ctx.path(), "epp>response>result")
		return nil
	})
	s.MustHandleCharData("epp>response>result>msg", func(ctx *Context) error {
		st.Expect(t, ctx.path(), "epp>response>result>msg")
		return nil
	})

	x := `<?xml version="1.0" encoding="utf-8"?>
<epp xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance" xsi:schemaLocation="urn:ietf:params:xml:ns:epp-1.0 epp-1.0.xsd" xmlns="urn:ietf:params:xml:ns:epp-1.0">
  <response>
    <result code="1000">
      <msg>Command completed successfully</msg>
    </result>
    <trID>
      <clTRID>0000000000000001</clTRID>
      <svTRID>3577a51b-5a4b-4025-8556-0a3e392b4097:1</svTRID>
    </trID>
  </response>
</epp>`

	d := decoder(x)
	var v struct{}
	err := s.Scan(d, &v)
	st.Expect(t, err, io.EOF)
}

func TestScannerInvalidXML(t *testing.T) {
	s := NewScanner()
	d := decoder(`<foo><bar/><baz/`)
	err := s.Scan(d, nil)
	_, ok := err.(*xml.SyntaxError)
	st.Expect(t, ok, true)
}

func ExampleScanner_Scan() {
	type Part struct {
		Name string
	}

	type Example struct {
		Size    int
		Enabled bool
		Parts   []Part
	}

	s := NewScanner()
	s.MustHandleStartElement("example", func(ctx *Context) error {
		v := ctx.Value.(*Example)
		v.Size = ctx.AttrInt("", "size")
		v.Enabled = ctx.AttrBool("", "enabled")
		return nil
	})
	s.MustHandleStartElement("example>part", func(ctx *Context) error {
		v := ctx.Value.(*Example)
		v.Parts = append(v.Parts, Part{})
		return nil
	})
	s.MustHandleCharData("example>part", func(ctx *Context) error {
		v := ctx.Value.(*Example)
		v.Parts[len(v.Parts)-1].Name = string(ctx.CharData)
		return nil
	})

	x := `<?xml version="1.0" encoding="utf-8"?>
<example size="1" enabled="true">
	<part>Foo</part>
	<part>Bar</part>
	<part>Baz</part>
</example>`

	d := xml.NewDecoder(bytes.NewBufferString(x))
	var v Example
	s.Scan(d, &v)
	fmt.Printf("Example size=%d enabled=%t parts=%s,%s,%s\n",
		v.Size, v.Enabled, v.Parts[0].Name, v.Parts[1].Name, v.Parts[2].Name)
	// Output: Example size=1 enabled=true parts=Foo,Bar,Baz
}
