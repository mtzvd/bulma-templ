package form

import (
	"strings"
	"testing"

	"github.com/a-h/templ"
	"github.com/mtzvd/bulma-templ/elements"
)

func TestSelect_DefaultRendering(t *testing.T) {
	html := render(t,
		Select(SelectProps{}, elements.Items{
			elements.Html("<option>Option 1</option>"),
		}),
	)

	if !strings.Contains(html, `<div`) {
		t.Fatalf("expected wrapper <div> tag, got: %s", html)
	}
	if !strings.Contains(html, `select`) {
		t.Fatalf("expected 'select' class on wrapper, got: %s", html)
	}
	if !strings.Contains(html, `<select`) {
		t.Fatalf("expected nested <select> tag, got: %s", html)
	}
	if !strings.Contains(html, `Option 1`) {
		t.Fatalf("expected option content, got: %s", html)
	}
}

func TestSelect_MultipleAttribute(t *testing.T) {
	html := render(t,
		Select(SelectProps{Multiple: true}, nil),
	)

	// Multiple affects both wrapper class and select attribute
	if !strings.Contains(html, `is-multiple`) {
		t.Fatalf("expected is-multiple class on wrapper when Multiple=true, got: %s", html)
	}
	if !strings.Contains(html, `multiple="multiple"`) {
		t.Fatalf("expected multiple='multiple' attribute on select when Multiple=true, got: %s", html)
	}
}

func TestSelect_NotMultipleByDefault(t *testing.T) {
	html := render(t,
		Select(SelectProps{}, nil),
	)

	if strings.Contains(html, `is-multiple`) {
		t.Fatalf("expected no is-multiple class by default, got: %s", html)
	}
	if strings.Contains(html, `multiple`) {
		t.Fatalf("expected no multiple attribute by default, got: %s", html)
	}
}

func TestSelect_DisabledState(t *testing.T) {
	html := render(t,
		Select(SelectProps{Disabled: true}, nil),
	)

	if !strings.Contains(html, `disabled="disabled"`) {
		t.Fatalf("expected disabled='disabled' on select when Disabled=true, got: %s", html)
	}
}

func TestSelect_NotDisabledByDefault(t *testing.T) {
	html := render(t,
		Select(SelectProps{}, nil),
	)

	if strings.Contains(html, `disabled`) {
		t.Fatalf("expected no disabled attribute by default, got: %s", html)
	}
}

func TestSelect_ColorModifier(t *testing.T) {
	html := render(t,
		Select(SelectProps{Color: "is-info"}, nil),
	)

	if !strings.Contains(html, `is-info`) {
		t.Fatalf("expected is-info class on wrapper, got: %s", html)
	}
}

func TestSelect_SizeModifier(t *testing.T) {
	html := render(t,
		Select(SelectProps{Size: "is-large"}, nil),
	)

	if !strings.Contains(html, `is-large`) {
		t.Fatalf("expected is-large class on wrapper, got: %s", html)
	}
}

func TestSelect_FullWidthModifier(t *testing.T) {
	html := render(t,
		Select(SelectProps{FullWidth: true}, nil),
	)

	if !strings.Contains(html, `is-fullwidth`) {
		t.Fatalf("expected is-fullwidth class on wrapper when FullWidth=true, got: %s", html)
	}
}

func TestSelect_WrapperAttributes(t *testing.T) {
	html := render(t,
		Select(SelectProps{
			WrapperAttr: templ.Attributes{
				"data-testid": "my-select-wrapper",
			},
		}, nil),
	)

	if !strings.Contains(html, `data-testid="my-select-wrapper"`) {
		t.Fatalf("expected WrapperAttr on wrapper div, got: %s", html)
	}
}

func TestSelect_ControlAttributes(t *testing.T) {
	html := render(t,
		Select(SelectProps{
			Attr: templ.Attributes{
				"name": "country",
			},
		}, nil),
	)

	if !strings.Contains(html, `name="country"`) {
		t.Fatalf("expected Attr on select element, got: %s", html)
	}
}

func TestSelect_NestedStructure(t *testing.T) {
	html := render(t,
		Select(SelectProps{}, elements.Items{
			elements.Html("<option>Test</option>"),
		}),
	)

	// Verify wrapper wraps select
	wrapperIdx := strings.Index(html, `class="select`)
	selectIdx := strings.Index(html, `<select`)
	if wrapperIdx == -1 || selectIdx == -1 || wrapperIdx >= selectIdx {
		t.Fatalf("expected .select wrapper to wrap <select>, got: %s", html)
	}
}
