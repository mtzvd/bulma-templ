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

// Subcomponent tests for layout components

func TestSmokeHero_Subcomponents(t *testing.T) {
	// Test HeroHead
	head := render(t,
		HeroHead(HeroHeadProps{}, elements.Items{
			elements.Html("Header"),
		}),
	)
	if !strings.Contains(head, `hero-head`) {
		t.Fatalf("expected 'hero-head' class, got: %s", head)
	}

	// Test HeroBody
	body := render(t,
		HeroBody(HeroBodyProps{}, elements.Items{
			elements.Html("Body"),
		}),
	)
	if !strings.Contains(body, `hero-body`) {
		t.Fatalf("expected 'hero-body' class, got: %s", body)
	}

	// Test HeroFoot
	foot := render(t,
		HeroFoot(HeroFootProps{}, elements.Items{
			elements.Html("Footer"),
		}),
	)
	if !strings.Contains(foot, `hero-foot`) {
		t.Fatalf("expected 'hero-foot' class, got: %s", foot)
	}
}

func TestSmokeLevel_Subcomponents(t *testing.T) {
	// Test LevelLeft
	left := render(t,
		LevelLeft(LevelLeftProps{}, elements.Items{
			elements.Html("Left"),
		}),
	)
	if !strings.Contains(left, `level-left`) {
		t.Fatalf("expected 'level-left' class, got: %s", left)
	}

	// Test LevelRight
	right := render(t,
		LevelRight(LevelRightProps{}, elements.Items{
			elements.Html("Right"),
		}),
	)
	if !strings.Contains(right, `level-right`) {
		t.Fatalf("expected 'level-right' class, got: %s", right)
	}

	// Test LevelItem
	item := render(t,
		LevelItem(LevelItemProps{}, elements.Items{
			elements.Html("Item"),
		}),
	)
	if !strings.Contains(item, `level-item`) {
		t.Fatalf("expected 'level-item' class, got: %s", item)
	}
}

func TestSmokeMedia_Subcomponents(t *testing.T) {
	// Test MediaLeft
	left := render(t,
		MediaLeft(MediaLeftProps{}, elements.Items{
			elements.Html("Left"),
		}),
	)
	if !strings.Contains(left, `media-left`) {
		t.Fatalf("expected 'media-left' class, got: %s", left)
	}

	// Test MediaContent
	content := render(t,
		MediaContent(MediaContentProps{}, elements.Items{
			elements.Html("Content"),
		}),
	)
	if !strings.Contains(content, `media-content`) {
		t.Fatalf("expected 'media-content' class, got: %s", content)
	}

	// Test MediaRight
	right := render(t,
		MediaRight(MediaRightProps{}, elements.Items{
			elements.Html("Right"),
		}),
	)
	if !strings.Contains(right, `media-right`) {
		t.Fatalf("expected 'media-right' class, got: %s", right)
	}
}
