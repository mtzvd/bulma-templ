package form

import (
	"strings"
	"testing"

	"github.com/a-h/templ"
	"github.com/mtzvd/bulma-templ/elements"
)

func TestCheckbox_DefaultUnchecked(t *testing.T) {
	html := render(t,
		Checkbox(CheckboxProps{}, elements.Items{
			elements.Html("Accept terms"),
		}),
	)

	if !strings.Contains(html, `<label`) {
		t.Fatalf("expected <label> tag, got: %s", html)
	}
	if !strings.Contains(html, `checkbox`) {
		t.Fatalf("expected 'checkbox' class, got: %s", html)
	}
	if !strings.Contains(html, `type="checkbox"`) {
		t.Fatalf("expected type='checkbox', got: %s", html)
	}
	if !strings.Contains(html, `Accept terms`) {
		t.Fatalf("expected label text 'Accept terms', got: %s", html)
	}
	if strings.Contains(html, `checked`) {
		t.Fatalf("expected no checked attribute by default, got: %s", html)
	}
}

func TestCheckbox_CheckedState(t *testing.T) {
	html := render(t,
		Checkbox(CheckboxProps{Checked: true}, nil),
	)

	if !strings.Contains(html, `checked="checked"`) {
		t.Fatalf("expected checked='checked' when Checked=true, got: %s", html)
	}
}

func TestCheckbox_DisabledState(t *testing.T) {
	html := render(t,
		Checkbox(CheckboxProps{Disabled: true}, nil),
	)

	if !strings.Contains(html, `disabled="disabled"`) {
		t.Fatalf("expected disabled='disabled' when Disabled=true, got: %s", html)
	}
}

func TestCheckbox_NotDisabledByDefault(t *testing.T) {
	html := render(t,
		Checkbox(CheckboxProps{}, nil),
	)

	if strings.Contains(html, `disabled`) {
		t.Fatalf("expected no disabled attribute by default, got: %s", html)
	}
}

func TestCheckbox_WithLabelText(t *testing.T) {
	html := render(t,
		Checkbox(CheckboxProps{}, elements.Items{
			elements.Html("I agree to the <strong>terms</strong>"),
		}),
	)

	if !strings.Contains(html, `I agree to the`) {
		t.Fatalf("expected label text, got: %s", html)
	}
	if !strings.Contains(html, `<strong>terms</strong>`) {
		t.Fatalf("expected HTML in label text, got: %s", html)
	}
}

func TestCheckbox_WrapperAttributes(t *testing.T) {
	html := render(t,
		Checkbox(CheckboxProps{
			WrapperAttr: templ.Attributes{
				"data-testid": "my-checkbox",
			},
		}, nil),
	)

	if !strings.Contains(html, `data-testid="my-checkbox"`) {
		t.Fatalf("expected WrapperAttr on label element, got: %s", html)
	}
}

// Smoke test for Checkboxes wrapper
func TestCheckboxes_Rendering(t *testing.T) {
	html := render(t,
		Checkboxes(CheckboxesProps{}, elements.Items{
			Checkbox(CheckboxProps{}, elements.Items{elements.Html("Option 1")}),
		}),
	)

	if !strings.Contains(html, `checkboxes`) {
		t.Fatalf("expected 'checkboxes' class, got: %s", html)
	}
	if !strings.Contains(html, `Option 1`) {
		t.Fatalf("expected content 'Option 1', got: %s", html)
	}
}
