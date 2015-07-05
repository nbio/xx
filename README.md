# xx

[![build status](https://img.shields.io/circleci/project/nbio/xx/master.svg)](https://circleci.com/gh/nbio/xx)
[![godoc](http://img.shields.io/badge/docs-GoDoc-blue.svg)](https://godoc.org/github.com/nbio/xx)

Minimal [SAX](https://en.wikipedia.org/wiki/Simple_API_for_XML)-ish XML scanner for [Go](https://golang./org). Extracted from and used in production at [domainr.com](https://domainr.com).

## What do I get?

Two things: `xml.StartElement` and `xml.CharData`. No processing instructions, comments, or end tags. Need them? PRs accepted.

## Why?

1. So you can parse XML without [reflect](https://godoc.org/reflect).
2. XX is shorter and [sounds better](http://thexx.info) than [XML](https://godoc.org/encoding/xml).
3. [Because](http://www.theatlantic.com/technology/archive/2013/11/english-has-a-new-preposition-because-internet/281601/) [XML](http://harmful.cat-v.org/software/xml/).


## Install

`go get github.com/nbio/xx`

## Usage

```go
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
```

## Author

© 2015 nb.io, LLC
