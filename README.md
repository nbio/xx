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
// Put something here
```

## Author

© 2015 nb.io, LLC
