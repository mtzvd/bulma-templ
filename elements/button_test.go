package elements

import (
	"strings"
	"testing"
)

func TestButton_DefaultState(t *testing.T) {
	html := render(t,
		Button(
			ButtonProps{},
			Items{Html("Click")},
		),
	)

	if !strings.HasPrefix(html, "<button") {
		t.Fatalf("expected <button>, got: %s", html)
	}

	if !strings.Contains(html, `class="button"`) {
		t.Fatalf("expected base button class, got: %s", html)
	}

	if strings.Contains(html, "disabled") {
		t.Fatalf("disabled must not be present by default: %s", html)
	}
}

func TestButton_HrefSwitchesToAnchor(t *testing.T) {
	html := render(t,
		Button(
			ButtonProps{
				Href: "/x",
			},
			Items{Html("Link")},
		),
	)

	if !strings.HasPrefix(html, "<a") {
		t.Fatalf("expected <a>, got: %s", html)
	}

	if !strings.Contains(html, `href="/x"`) {
		t.Fatalf("href missing: %s", html)
	}
}

func TestButton_DisabledButton(t *testing.T) {
	html := render(t,
		Button(
			ButtonProps{
				Disabled: true,
			},
			Items{Html("Disabled")},
		),
	)

	if !strings.Contains(html, "disabled") {
		t.Fatalf("expected disabled attribute: %s", html)
	}

	if !strings.Contains(html, "is-disabled") {
		t.Fatalf("expected is-disabled class: %s", html)
	}
}

func TestButton_DisabledAnchorIsVisualOnly(t *testing.T) {
	html := render(t,
		Button(
			ButtonProps{
				Href:     "/x",
				Disabled: true,
			},
			Items{Html("Disabled link")},
		),
	)

	// Anchor must NOT have a native disabled attribute.
	if strings.Contains(html, " disabled") || strings.Contains(html, `disabled="`) {
		t.Fatalf("anchor must not have disabled attribute: %s", html)
	}

	// Visual disabled state MUST be present.
	if !strings.Contains(html, "is-disabled") {
		t.Fatalf("expected visual disabled state: %s", html)
	}
}

func TestButton_Modifiers(t *testing.T) {
	html := render(t,
		Button(
			ButtonProps{
				Color:     "is-primary",
				Size:      "is-large",
				State:     "is-loading",
				FullWidth: true,
			},
			Items{Html("Go")},
		),
	)

	for _, cls := range []string{
		"button",
		"is-primary",
		"is-large",
		"is-loading",
		"is-fullwidth",
	} {
		if !strings.Contains(html, cls) {
			t.Fatalf("missing class %q in: %s", cls, html)
		}
	}
}

func TestButton_TypeAttribute(t *testing.T) {
	html := render(t,
		Button(
			ButtonProps{
				Type: "submit",
			},
			Items{Html("Submit")},
		),
	)

	if !strings.Contains(html, `type="submit"`) {
		t.Fatalf("type attribute missing: %s", html)
	}
}
