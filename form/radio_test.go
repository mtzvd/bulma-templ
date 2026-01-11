package form

import (
	"strings"
	"testing"

	"github.com/a-h/templ"
	"github.com/mtzvd/bulma-templ/elements"
)

func TestRadio_DefaultUnchecked(t *testing.T) {
	html := render(t,
		Radio(RadioProps{}, elements.Items{
			elements.Html("Option A"),
		}),
	)

	if !strings.Contains(html, `<label`) {
		t.Fatalf("expected <label> tag, got: %s", html)
	}
	if !strings.Contains(html, `radio`) {
		t.Fatalf("expected 'radio' class, got: %s", html)
	}
	if !strings.Contains(html, `type="radio"`) {
		t.Fatalf("expected type='radio', got: %s", html)
	}
	if !strings.Contains(html, `Option A`) {
		t.Fatalf("expected label text 'Option A', got: %s", html)
	}
	if strings.Contains(html, `checked`) {
		t.Fatalf("expected no checked attribute by default, got: %s", html)
	}
}

func TestRadio_CheckedState(t *testing.T) {
	html := render(t,
		Radio(RadioProps{Checked: true}, nil),
	)

	if !strings.Contains(html, `checked="checked"`) {
		t.Fatalf("expected checked='checked' when Checked=true, got: %s", html)
	}
}

func TestRadio_DisabledState(t *testing.T) {
	html := render(t,
		Radio(RadioProps{Disabled: true}, nil),
	)

	if !strings.Contains(html, `disabled="disabled"`) {
		t.Fatalf("expected disabled='disabled' when Disabled=true, got: %s", html)
	}
}

func TestRadio_NotDisabledByDefault(t *testing.T) {
	html := render(t,
		Radio(RadioProps{}, nil),
	)

	if strings.Contains(html, `disabled`) {
		t.Fatalf("expected no disabled attribute by default, got: %s", html)
	}
}

func TestRadio_NameAttribute(t *testing.T) {
	html := render(t,
		Radio(RadioProps{
			Attr: templ.Attributes{
				"name": "choice",
			},
		}, nil),
	)

	if !strings.Contains(html, `name="choice"`) {
		t.Fatalf("expected name='choice' attribute, got: %s", html)
	}
}

func TestRadio_WithLabelText(t *testing.T) {
	html := render(t,
		Radio(RadioProps{}, elements.Items{
			elements.Html("I choose <em>option B</em>"),
		}),
	)

	if !strings.Contains(html, `I choose`) {
		t.Fatalf("expected label text, got: %s", html)
	}
	if !strings.Contains(html, `<em>option B</em>`) {
		t.Fatalf("expected HTML in label text, got: %s", html)
	}
}

// Smoke test for Radios wrapper
func TestRadios_Rendering(t *testing.T) {
	html := render(t,
		Radios(RadiosProps{}, elements.Items{
			Radio(RadioProps{}, elements.Items{elements.Html("Option 1")}),
		}),
	)

	if !strings.Contains(html, `radios`) {
		t.Fatalf("expected 'radios' class, got: %s", html)
	}
	if !strings.Contains(html, `Option 1`) {
		t.Fatalf("expected content 'Option 1', got: %s", html)
	}
}
