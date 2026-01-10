package elements

import (
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func TestWrap_ClassMerging(t *testing.T) {
	html := render(t,
		Wrap(
			WrapProps{
				Tag:     "div",
				Classes: []string{"a", "b"},
				Attr: templ.Attributes{
					"class": "c d",
				},
			},
			Items{},
		),
	)

	if !strings.Contains(html, `class="a b c d"`) {
		t.Fatalf("expected merged class attribute, got: %s", html)
	}

	if strings.Count(html, "class=") != 1 {
		t.Fatalf("class attribute duplicated: %s", html)
	}
}

func TestWrap_BooleanAttributes(t *testing.T) {
	html := render(t,
		Wrap(
			WrapProps{
				Tag: "button",
				Attr: templ.Attributes{
					"disabled": true,
					"hidden":   false,
				},
			},
			Items{},
		),
	)

	if !strings.Contains(html, "disabled") {
		t.Fatalf("expected disabled attribute, got: %s", html)
	}

	if strings.Contains(html, "hidden") {
		t.Fatalf("false boolean attribute must be omitted: %s", html)
	}
}

func TestWrap_EscapesAttributes(t *testing.T) {
	html := render(t,
		Wrap(
			WrapProps{
				Tag: "div",
				Attr: templ.Attributes{
					"title": `a"b<c>d`,
				},
			},
			Items{},
		),
	)

	if !strings.Contains(html, `title="a&quot;b&lt;c&gt;d"`) {
		t.Fatalf("attribute not escaped correctly: %s", html)
	}
}

func TestWrap_OmitsNilAttributes(t *testing.T) {
	html := render(t,
		Wrap(
			WrapProps{
				Tag: "div",
				Attr: templ.Attributes{
					"data-x": nil,
				},
			},
			Items{},
		),
	)

	if strings.Contains(html, "data-x") {
		t.Fatalf("nil attribute must be omitted: %s", html)
	}
}
