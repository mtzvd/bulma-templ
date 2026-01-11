package form

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
)

func render(t *testing.T, c templ.Component) string {
	var b strings.Builder
	if err := c.Render(context.Background(), &b); err != nil {
		t.Fatalf("render failed: %v", err)
	}
	return b.String()
}

func TestInput_DefaultRendering(t *testing.T) {
	html := render(t,
		Input(InputProps{}),
	)

	if !strings.Contains(html, `<input`) {
		t.Fatalf("expected <input> tag, got: %s", html)
	}
	if !strings.Contains(html, `input`) {
		t.Fatalf("expected 'input' class, got: %s", html)
	}
}

func TestInput_TypeAttribute(t *testing.T) {
	html := render(t,
		Input(InputProps{Type: "email"}),
	)

	if !strings.Contains(html, `type="email"`) {
		t.Fatalf("expected type='email', got: %s", html)
	}
}

func TestInput_ValueAndPlaceholder(t *testing.T) {
	html := render(t,
		Input(InputProps{
			Value:       "test value",
			Placeholder: "Enter text",
		}),
	)

	if !strings.Contains(html, `value="test value"`) {
		t.Fatalf("expected value='test value', got: %s", html)
	}
	if !strings.Contains(html, `placeholder="Enter text"`) {
		t.Fatalf("expected placeholder='Enter text', got: %s", html)
	}
}

func TestInput_DisabledState(t *testing.T) {
	html := render(t,
		Input(InputProps{Disabled: true}),
	)

	if !strings.Contains(html, `disabled="disabled"`) {
		t.Fatalf("expected disabled='disabled' when Disabled=true, got: %s", html)
	}
}

func TestInput_ReadOnlyState(t *testing.T) {
	html := render(t,
		Input(InputProps{ReadOnly: true}),
	)

	if !strings.Contains(html, `readonly="readonly"`) {
		t.Fatalf("expected readonly='readonly' when ReadOnly=true, got: %s", html)
	}
}

func TestInput_NotDisabledByDefault(t *testing.T) {
	html := render(t,
		Input(InputProps{}),
	)

	if strings.Contains(html, `disabled`) {
		t.Fatalf("expected no disabled attribute by default, got: %s", html)
	}
}

func TestInput_NotReadOnlyByDefault(t *testing.T) {
	html := render(t,
		Input(InputProps{}),
	)

	if strings.Contains(html, `readonly`) {
		t.Fatalf("expected no readonly attribute by default, got: %s", html)
	}
}

func TestInput_ColorModifier(t *testing.T) {
	html := render(t,
		Input(InputProps{Color: "is-danger"}),
	)

	if !strings.Contains(html, `is-danger`) {
		t.Fatalf("expected is-danger class, got: %s", html)
	}
}

func TestInput_SizeModifier(t *testing.T) {
	html := render(t,
		Input(InputProps{Size: "is-large"}),
	)

	if !strings.Contains(html, `is-large`) {
		t.Fatalf("expected is-large class, got: %s", html)
	}
}
