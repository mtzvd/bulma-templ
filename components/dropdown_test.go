package components

import (
	"strings"
	"testing"

	"github.com/mtzvd/bulma-templ/elements"
)

// Test Dropdown multiple boolean modifiers (critical conditional logic)
func TestDropdown_DefaultState(t *testing.T) {
	html := render(t,
		Dropdown(DropdownProps{}, nil),
	)

	if !strings.Contains(html, `class="dropdown"`) {
		t.Fatalf("expected class='dropdown', got: %s", html)
	}
	if strings.Contains(html, `is-active`) {
		t.Fatalf("expected no is-active class by default, got: %s", html)
	}
	if strings.Contains(html, `is-hoverable`) {
		t.Fatalf("expected no is-hoverable class by default, got: %s", html)
	}
	if strings.Contains(html, `is-right`) {
		t.Fatalf("expected no is-right class by default, got: %s", html)
	}
	if strings.Contains(html, `is-up`) {
		t.Fatalf("expected no is-up class by default, got: %s", html)
	}
}

func TestDropdown_ActiveState(t *testing.T) {
	html := render(t,
		Dropdown(DropdownProps{Active: true}, nil),
	)

	if !strings.Contains(html, `is-active`) {
		t.Fatalf("expected is-active class when Active=true, got: %s", html)
	}
}

func TestDropdown_HoverableState(t *testing.T) {
	html := render(t,
		Dropdown(DropdownProps{Hoverable: true}, nil),
	)

	if !strings.Contains(html, `is-hoverable`) {
		t.Fatalf("expected is-hoverable class when Hoverable=true, got: %s", html)
	}
}

func TestDropdown_RightModifier(t *testing.T) {
	html := render(t,
		Dropdown(DropdownProps{Right: true}, nil),
	)

	if !strings.Contains(html, `is-right`) {
		t.Fatalf("expected is-right class when Right=true, got: %s", html)
	}
}

func TestDropdown_UpModifier(t *testing.T) {
	html := render(t,
		Dropdown(DropdownProps{Up: true}, nil),
	)

	if !strings.Contains(html, `is-up`) {
		t.Fatalf("expected is-up class when Up=true, got: %s", html)
	}
}

func TestDropdown_MultipleModifiers(t *testing.T) {
	html := render(t,
		Dropdown(DropdownProps{
			Active:    true,
			Hoverable: true,
			Right:     true,
		}, nil),
	)

	if !strings.Contains(html, `is-active`) {
		t.Fatalf("expected is-active class, got: %s", html)
	}
	if !strings.Contains(html, `is-hoverable`) {
		t.Fatalf("expected is-hoverable class, got: %s", html)
	}
	if !strings.Contains(html, `is-right`) {
		t.Fatalf("expected is-right class, got: %s", html)
	}
}

// Test DropdownMenu nested structure (menu > content)
func TestDropdownMenu_NestedStructure(t *testing.T) {
	html := render(t,
		DropdownMenu(DropdownMenuProps{}, elements.Items{
			elements.Html("Item"),
		}),
	)

	if !strings.Contains(html, `dropdown-menu`) {
		t.Fatalf("expected dropdown-menu class, got: %s", html)
	}
	if !strings.Contains(html, `dropdown-content`) {
		t.Fatalf("expected nested dropdown-content class, got: %s", html)
	}
	if !strings.Contains(html, `Item`) {
		t.Fatalf("expected content 'Item', got: %s", html)
	}

	// Verify nesting order: menu wraps content
	menuIdx := strings.Index(html, "dropdown-menu")
	contentIdx := strings.Index(html, "dropdown-content")
	if menuIdx == -1 || contentIdx == -1 || menuIdx >= contentIdx {
		t.Fatalf("expected dropdown-menu to wrap dropdown-content, got: %s", html)
	}
}

// Test DropdownItem conditional href and active state
func TestDropdownItem_DefaultRendering(t *testing.T) {
	html := render(t,
		DropdownItem(DropdownItemProps{}, elements.Items{
			elements.Html("Item"),
		}),
	)

	if !strings.Contains(html, `<a`) {
		t.Fatalf("expected <a> tag, got: %s", html)
	}
	if !strings.Contains(html, `dropdown-item`) {
		t.Fatalf("expected dropdown-item class, got: %s", html)
	}
	if !strings.Contains(html, `Item`) {
		t.Fatalf("expected content 'Item', got: %s", html)
	}
	if strings.Contains(html, `is-active`) {
		t.Fatalf("expected no is-active class by default, got: %s", html)
	}
}

func TestDropdownItem_WithHref(t *testing.T) {
	html := render(t,
		DropdownItem(DropdownItemProps{Href: "/action"}, nil),
	)

	if !strings.Contains(html, `href="/action"`) {
		t.Fatalf("expected href='/action', got: %s", html)
	}
}

func TestDropdownItem_WithoutHref(t *testing.T) {
	html := render(t,
		DropdownItem(DropdownItemProps{}, nil),
	)

	// When Href is empty, no href attribute should be present
	// (action item handled by JS/Alpine)
	if strings.Contains(html, `href=""`) {
		t.Fatalf("expected no empty href attribute, got: %s", html)
	}
}

func TestDropdownItem_ActiveState(t *testing.T) {
	html := render(t,
		DropdownItem(DropdownItemProps{Active: true}, nil),
	)

	if !strings.Contains(html, `is-active`) {
		t.Fatalf("expected is-active class when Active=true, got: %s", html)
	}
}

// Smoke tests for simple wrapper components
func TestDropdownTrigger_Rendering(t *testing.T) {
	html := render(t,
		DropdownTrigger(DropdownTriggerProps{}, elements.Items{
			elements.Html("Trigger"),
		}),
	)

	if !strings.Contains(html, `dropdown-trigger`) {
		t.Fatalf("expected dropdown-trigger class, got: %s", html)
	}
	if !strings.Contains(html, `Trigger`) {
		t.Fatalf("expected content 'Trigger', got: %s", html)
	}
}

func TestDropdownDivider_Rendering(t *testing.T) {
	html := render(t,
		DropdownDivider(),
	)

	if !strings.Contains(html, `<hr`) {
		t.Fatalf("expected <hr> tag, got: %s", html)
	}
	if !strings.Contains(html, `dropdown-divider`) {
		t.Fatalf("expected dropdown-divider class, got: %s", html)
	}
}
