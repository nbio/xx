package xx

import (
	"encoding/xml"
	"errors"
	"strings"
)

// ScanFunc is a callback that accepts an xml.StartElement, an
// xml.CharData, and an optional interface{} value for private use.
//
// The function is called for two XML tokens: xml.StartElement
// and xml.CharData. The xml.StartElement will always be the last
// StartElement parsed. CharData may be nil.
type ScanFunc func(ctx *Context) error

// Scanner scans XML from an xml.Decoder, looking for specific paths.
type Scanner struct {
	tree map[xml.Name]*Scanner
	se   ScanFunc
	cd   ScanFunc
}

// NewScanner returns an initialized Scanner, ready to use.
func NewScanner() *Scanner {
	return &Scanner{tree: make(map[xml.Name]*Scanner)}
}

// MustHandleStartElement initializes an XML StartElement handler.
// It panics if it cannot create a path handler.
func (s *Scanner) MustHandleStartElement(path string, f ScanFunc) {
	must(s.HandleStartElement(path, f))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}

// HandleStartElement initializes an XML StartElement handler.
//
// Paths must be in the form of "foo>bar" or "uri foo>uri bar".
// The path is delimited by > characters, and individual path
// elements are split on whitespace into a namespace and local
// tag name.
//
// HandleStartElement returns ErrInvalidPath if the specified path
// is malformed.
func (s *Scanner) HandleStartElement(path string, f ScanFunc) error {
	s2, err := s.makePath(path)
	if err != nil {
		return err
	}
	s2.se = f
	return nil
}

// MustHandleCharData initializes an XML CharData handler.
// It panics if it cannot create a path handler.
func (s *Scanner) MustHandleCharData(path string, f ScanFunc) {
	must(s.HandleCharData(path, f))
}

// HandleCharData initializes an XML CharData handler.
func (s *Scanner) HandleCharData(path string, f ScanFunc) error {
	s2, err := s.makePath(path)
	if err != nil {
		return err
	}
	s2.cd = f
	return nil
}

// makePath finds or creates a tree of Scanners.
// It returns the leaf node Scanner for the path or an error.
func (s *Scanner) makePath(path string) (*Scanner, error) {
	names := strings.SplitN(path, ">", 2)
	fields := strings.Fields(names[0])
	var name xml.Name
	switch len(fields) {
	case 0:
		if len(names) > 1 {
			return nil, ErrInvalidPath
		}
		return s, nil
	case 1:
		name.Local = fields[0]
	case 2:
		name.Space = fields[0]
		name.Local = fields[1]
	default:
		return nil, ErrInvalidPath
	}
	s2, ok := s.tree[name]
	if !ok {
		s2 = NewScanner()
		s.tree[name] = s2
	}
	if len(names) == 1 {
		return s2.makePath("")
	}
	return s2.makePath(names[1])
}

// ErrInvalidPath describes an invalid Scanner path
// returned by Scanner.Handle.
var ErrInvalidPath = errors.New("invalid scan path")

// Scan reads xml.Tokens from Decoder d, passing matching
// xml.StartElement or xml.CharData tokens and an optional
// value from the caller to the matching ScanFuncs in s.
// It returns any errors encountered.
// It will return if it encounters an xml.EndElement that matches
// the corresponding xml.StartElement it attempted to scan.
func (s *Scanner) Scan(d *xml.Decoder, ctxValue interface{}) error {
	ctx := Context{
		Decoder: d,
		Value:   ctxValue,
	}
	return s.scan(&ctx)
}

func (s *Scanner) scan(ctx *Context) error {
	for {
		t, err := ctx.Decoder.Token()
		if err != nil {
			return err
		}
		switch node := t.(type) {
		case xml.StartElement:
			s2, ok := s.tree[node.Name]
			if !ok {
				s2, ok = s.tree[xml.Name{"", node.Name.Local}]
				if !ok {
					err = ctx.Decoder.Skip()
					break
				}
			}
			child := *ctx
			child.Parent = ctx
			child.StartElement = node
			if s2.se != nil {
				err = s2.se(&child)
				if err != nil {
					return err
				}
			}
			err = s2.scan(&child)

		case xml.EndElement:
			return nil

		case xml.CharData:
			if s.cd != nil {
				ctx.CharData = node
				err = s.cd(ctx)
				ctx.CharData = nil
			}
		}
		if err != nil {
			return err
		}
	}
	return nil
}
