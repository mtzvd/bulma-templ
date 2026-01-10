package elements

import (
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestBaseElement_RendersTagAndClasses(t *testing.T) {
	html := render(t,
		BaseElement(
			BaseElementProps{
				Tag:         "section",
				BaseClasses: []string{"section", "is-large"},
			},
			Items{},
		),
	)

	if !strings.HasPrefix(html, "<section") {
		t.Fatalf("expected <section>, got: %s", html)
	}

	if !strings.Contains(html, `class="section is-large"`) {
		t.Fatalf("expected classes applied, got: %s", html)
	}

	if !strings.HasSuffix(html, "</section>") {
		t.Fatalf("expected closing tag, got: %s", html)
	}
}

func TestBaseElement_DelegatesAttrMerging(t *testing.T) {
	html := render(t,
		BaseElement(
			BaseElementProps{
				Tag:         "div",
				BaseClasses: []string{"a"},
				Attr: templ.Attributes{
					"class": "b",
					"id":    "x",
				},
			},
			Items{},
		),
	)

	if !strings.Contains(html, `class="a b"`) {
		t.Fatalf("class not merged correctly: %s", html)
	}

	if !strings.Contains(html, `id="x"`) {
		t.Fatalf("attribute lost: %s", html)
	}

	if strings.Count(html, "class=") != 1 {
		t.Fatalf("class duplicated: %s", html)
	}
}

func TestBaseElement_WithContent(t *testing.T) {
	html := render(t,
		BaseElement(
			BaseElementProps{
				Tag: "div",
			},
			Items{
				Html("A"),
				Html("B"),
			},
		),
	)

	if !strings.Contains(html, ">AB<") {
		t.Fatalf("content rendered in wrong order: %s", html)
	}
}
