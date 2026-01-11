package components

import (
	"strings"
	"testing"

	"github.com/mtzvd/bulma-templ/elements"
)

// Smoke tests for simple wrapper components without conditional logic.
// These tests provide regression confidence but limited technical value.

func TestSmokeBreadcrumb_Rendering(t *testing.T) {
	html := render(t,
		Breadcrumb(BreadcrumbProps{}, elements.Items{
			elements.Html("<li>Item</li>"),
		}),
	)

	if !strings.Contains(html, `breadcrumb`) {
		t.Fatalf("expected 'breadcrumb' class, got: %s", html)
	}
}

func TestSmokeMessage_Rendering(t *testing.T) {
	html := render(t,
		Message(MessageProps{}, elements.Items{
			elements.Html("Message content"),
		}),
	)

	if !strings.Contains(html, `message`) {
		t.Fatalf("expected 'message' class, got: %s", html)
	}
}

func TestSmokeMenu_Rendering(t *testing.T) {
	html := render(t,
		Menu(MenuProps{}, elements.Items{
			elements.Html("Menu content"),
		}),
	)

	if !strings.Contains(html, `menu`) {
		t.Fatalf("expected 'menu' class, got: %s", html)
	}
}

func TestSmokePanel_Rendering(t *testing.T) {
	html := render(t,
		Panel(PanelProps{}, elements.Items{
			elements.Html("Panel content"),
		}),
	)

	if !strings.Contains(html, `panel`) {
		t.Fatalf("expected 'panel' class, got: %s", html)
	}
}

func TestSmokeTabs_Rendering(t *testing.T) {
	html := render(t,
		Tabs(TabsProps{}, elements.Items{
			elements.Html("<ul><li>Tab</li></ul>"),
		}),
	)

	if !strings.Contains(html, `tabs`) {
		t.Fatalf("expected 'tabs' class, got: %s", html)
	}
}
