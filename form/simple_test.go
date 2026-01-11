package form

import (
	"strings"
	"testing"

	"github.com/mtzvd/bulma-templ/elements"
)

// Smoke tests for simple wrapper form components without conditional logic.

func TestSmokeField_Rendering(t *testing.T) {
	html := render(t,
		Field(FieldProps{}, elements.Items{
			elements.Html("Field content"),
		}),
	)

	if !strings.Contains(html, `field`) {
		t.Fatalf("expected 'field' class, got: %s", html)
	}
}

func TestSmokeControl_Rendering(t *testing.T) {
	html := render(t,
		Control(ControlProps{}, elements.Items{
			elements.Html("Control content"),
		}),
	)

	if !strings.Contains(html, `control`) {
		t.Fatalf("expected 'control' class, got: %s", html)
	}
}

func TestSmokeLabel_Rendering(t *testing.T) {
	html := render(t,
		Label(LabelProps{Text: "Label text"}),
	)

	if !strings.Contains(html, `label`) {
		t.Fatalf("expected 'label' class, got: %s", html)
	}
}

func TestSmokeHelp_Rendering(t *testing.T) {
	html := render(t,
		Help(HelpProps{Text: "Help text"}),
	)

	if !strings.Contains(html, `help`) {
		t.Fatalf("expected 'help' class, got: %s", html)
	}
}
