package components

import (
	"strings"
	"testing"

	"github.com/mtzvd/bulma-templ/elements"
)

// Test Navbar ARIA attributes (required by Bulma)
func TestNavbar_AriaAttributes(t *testing.T) {
	html := render(t,
		Navbar(NavbarProps{}, nil),
	)

	if !strings.Contains(html, `role="navigation"`) {
		t.Fatalf("expected role='navigation', got: %s", html)
	}
	if !strings.Contains(html, `aria-label="main navigation"`) {
		t.Fatalf("expected aria-label='main navigation', got: %s", html)
	}
	if !strings.Contains(html, `<nav`) {
		t.Fatalf("expected <nav> tag, got: %s", html)
	}
}

func TestNavbar_ColorModifier(t *testing.T) {
	html := render(t,
		Navbar(NavbarProps{Color: "is-primary"}, nil),
	)

	if !strings.Contains(html, `is-primary`) {
		t.Fatalf("expected is-primary class, got: %s", html)
	}
}

func TestNavbar_TransparentModifier(t *testing.T) {
	html := render(t,
		Navbar(NavbarProps{Transparent: true}, nil),
	)

	if !strings.Contains(html, `is-transparent`) {
		t.Fatalf("expected is-transparent class, got: %s", html)
	}
}

func TestNavbar_FixedPosition(t *testing.T) {
	html := render(t,
		Navbar(NavbarProps{Fixed: "is-fixed-top"}, nil),
	)

	if !strings.Contains(html, `is-fixed-top`) {
		t.Fatalf("expected is-fixed-top class, got: %s", html)
	}
}

// Test NavbarBurger hardcoded spans and Active state
func TestNavbarBurger_RendersThreeSpans(t *testing.T) {
	html := render(t,
		NavbarBurger(NavbarBurgerProps{}),
	)

	// Count occurrences of <span aria-hidden="true">
	spanCount := strings.Count(html, `<span aria-hidden="true">`)
	if spanCount != 3 {
		t.Fatalf("expected 3 spans, got %d in: %s", spanCount, html)
	}
}

func TestNavbarBurger_DefaultState(t *testing.T) {
	html := render(t,
		NavbarBurger(NavbarBurgerProps{}),
	)

	if !strings.Contains(html, `class="navbar-burger"`) {
		t.Fatalf("expected class='navbar-burger', got: %s", html)
	}
	if strings.Contains(html, `is-active`) {
		t.Fatalf("expected no is-active class by default, got: %s", html)
	}
	// aria-expanded=false is omitted (templ boolean attribute behavior)
	if strings.Contains(html, `aria-expanded`) {
		t.Fatalf("expected no aria-expanded attribute when Active=false, got: %s", html)
	}
	if !strings.Contains(html, `role="button"`) {
		t.Fatalf("expected role='button', got: %s", html)
	}
	if !strings.Contains(html, `aria-label="menu"`) {
		t.Fatalf("expected aria-label='menu', got: %s", html)
	}
}

func TestNavbarBurger_ActiveState(t *testing.T) {
	html := render(t,
		NavbarBurger(NavbarBurgerProps{Active: true}),
	)

	if !strings.Contains(html, `is-active`) {
		t.Fatalf("expected is-active class when Active=true, got: %s", html)
	}
	// aria-expanded=true renders as boolean attribute (present without value or with value)
	if !strings.Contains(html, `aria-expanded`) {
		t.Fatalf("expected aria-expanded attribute when Active=true, got: %s", html)
	}
}

// Test NavbarMenu Active state
func TestNavbarMenu_DefaultState(t *testing.T) {
	html := render(t,
		NavbarMenu(NavbarMenuProps{}, nil),
	)

	if !strings.Contains(html, `class="navbar-menu"`) {
		t.Fatalf("expected class='navbar-menu', got: %s", html)
	}
	if strings.Contains(html, `is-active`) {
		t.Fatalf("expected no is-active class by default, got: %s", html)
	}
}

func TestNavbarMenu_ActiveState(t *testing.T) {
	html := render(t,
		NavbarMenu(NavbarMenuProps{Active: true}, nil),
	)

	if !strings.Contains(html, `is-active`) {
		t.Fatalf("expected is-active class when Active=true, got: %s", html)
	}
}

// Test NavbarItem tag switching (critical conditional logic)
func TestNavbarItem_AsDiv(t *testing.T) {
	html := render(t,
		NavbarItem(NavbarItemProps{}, nil),
	)

	if !strings.HasPrefix(strings.TrimSpace(html), "<div") {
		t.Fatalf("expected <div> when Href is empty, got: %s", html)
	}
	if strings.Contains(html, `href=`) {
		t.Fatalf("expected no href attribute when Href is empty, got: %s", html)
	}
}

func TestNavbarItem_AsLink(t *testing.T) {
	html := render(t,
		NavbarItem(NavbarItemProps{Href: "/home"}, nil),
	)

	if !strings.HasPrefix(strings.TrimSpace(html), "<a") {
		t.Fatalf("expected <a> when Href is set, got: %s", html)
	}
	if !strings.Contains(html, `href="/home"`) {
		t.Fatalf("expected href='/home', got: %s", html)
	}
}

func TestNavbarItem_ActiveState(t *testing.T) {
	html := render(t,
		NavbarItem(NavbarItemProps{Active: true}, nil),
	)

	if !strings.Contains(html, `is-active`) {
		t.Fatalf("expected is-active class when Active=true, got: %s", html)
	}
}

func TestNavbarItem_TabModifier(t *testing.T) {
	html := render(t,
		NavbarItem(NavbarItemProps{Tab: true}, nil),
	)

	if !strings.Contains(html, `is-tab`) {
		t.Fatalf("expected is-tab class when Tab=true, got: %s", html)
	}
}

func TestNavbarItem_HoverableModifier(t *testing.T) {
	html := render(t,
		NavbarItem(NavbarItemProps{Hoverable: true}, nil),
	)

	if !strings.Contains(html, `is-hoverable`) {
		t.Fatalf("expected is-hoverable class when Hoverable=true, got: %s", html)
	}
}

// Test NavbarLink Arrowless modifier
func TestNavbarLink_DefaultState(t *testing.T) {
	html := render(t,
		NavbarLink(NavbarLinkProps{}, nil),
	)

	if !strings.Contains(html, `class="navbar-link"`) {
		t.Fatalf("expected class='navbar-link', got: %s", html)
	}
	if strings.Contains(html, `is-arrowless`) {
		t.Fatalf("expected no is-arrowless class by default, got: %s", html)
	}
}

func TestNavbarLink_ArrowlessModifier(t *testing.T) {
	html := render(t,
		NavbarLink(NavbarLinkProps{Arrowless: true}, nil),
	)

	if !strings.Contains(html, `is-arrowless`) {
		t.Fatalf("expected is-arrowless class when Arrowless=true, got: %s", html)
	}
}

// Test NavbarDropdown Right modifier
func TestNavbarDropdown_DefaultState(t *testing.T) {
	html := render(t,
		NavbarDropdown(NavbarDropdownProps{}, nil),
	)

	if !strings.Contains(html, `class="navbar-dropdown"`) {
		t.Fatalf("expected class='navbar-dropdown', got: %s", html)
	}
	if strings.Contains(html, `is-right`) {
		t.Fatalf("expected no is-right class by default, got: %s", html)
	}
}

func TestNavbarDropdown_RightModifier(t *testing.T) {
	html := render(t,
		NavbarDropdown(NavbarDropdownProps{Right: true}, nil),
	)

	if !strings.Contains(html, `is-right`) {
		t.Fatalf("expected is-right class when Right=true, got: %s", html)
	}
}

// Smoke tests for simple wrappers (no conditional logic)
func TestNavbarBrand_Rendering(t *testing.T) {
	html := render(t,
		NavbarBrand(NavbarBrandProps{}, elements.Items{elements.Html("Logo")}),
	)

	if !strings.Contains(html, `class="navbar-brand"`) {
		t.Fatalf("expected class='navbar-brand', got: %s", html)
	}
	if !strings.Contains(html, `Logo`) {
		t.Fatalf("expected content 'Logo', got: %s", html)
	}
}

func TestNavbarDivider_Rendering(t *testing.T) {
	html := render(t,
		NavbarDivider(NavbarDividerProps{}),
	)

	if !strings.Contains(html, `<hr`) {
		t.Fatalf("expected <hr> tag, got: %s", html)
	}
	if !strings.Contains(html, `class="navbar-divider"`) {
		t.Fatalf("expected class='navbar-divider', got: %s", html)
	}
}
