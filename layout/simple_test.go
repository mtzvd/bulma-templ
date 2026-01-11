package layout

import (
	"context"
	"strings"
	"testing"

	"github.com/a-h/templ"
	"github.com/mtzvd/bulma-templ/elements"
)

func render(t *testing.T, c templ.Component) string {
	var b strings.Builder
	if err := c.Render(context.Background(), &b); err != nil {
		t.Fatalf("render failed: %v", err)
	}
	return b.String()
}

// Smoke tests for layout components without conditional logic.
// These tests provide regression confidence but limited technical value.

func TestSmokeContainer_Rendering(t *testing.T) {
	html := render(t,
		Container(ContainerProps{}, elements.Items{
			elements.Html("Content"),
		}),
	)

	if !strings.Contains(html, `container`) {
		t.Fatalf("expected 'container' class, got: %s", html)
	}
}

func TestSmokeSection_Rendering(t *testing.T) {
	html := render(t,
		Section(SectionProps{}, elements.Items{
			elements.Html("Section content"),
		}),
	)

	if !strings.Contains(html, `section`) {
		t.Fatalf("expected 'section' class, got: %s", html)
	}
}

func TestSmokeHero_Rendering(t *testing.T) {
	html := render(t,
		Hero(HeroProps{}, elements.Items{
			elements.Html("Hero content"),
		}),
	)

	if !strings.Contains(html, `hero`) {
		t.Fatalf("expected 'hero' class, got: %s", html)
	}
}

func TestSmokeFooter_Rendering(t *testing.T) {
	html := render(t,
		Footer(FooterProps{}, elements.Items{
			elements.Html("Footer content"),
		}),
	)

	if !strings.Contains(html, `footer`) {
		t.Fatalf("expected 'footer' class, got: %s", html)
	}
}

func TestSmokeLevel_Rendering(t *testing.T) {
	html := render(t,
		Level(LevelProps{}, elements.Items{
			elements.Html("Level content"),
		}),
	)

	if !strings.Contains(html, `level`) {
		t.Fatalf("expected 'level' class, got: %s", html)
	}
}

func TestSmokeMedia_Rendering(t *testing.T) {
	html := render(t,
		Media(MediaProps{}, elements.Items{
			elements.Html("Media content"),
		}),
	)

	if !strings.Contains(html, `media`) {
		t.Fatalf("expected 'media' class, got: %s", html)
	}
}
