package grid

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

// Smoke test for grid without conditional logic.

func TestSmokeGrid_Rendering(t *testing.T) {
	html := render(t,
		Grid(GridProps{}, elements.Items{
			elements.Html("<div>Grid item</div>"),
		}),
	)

	if !strings.Contains(html, `grid`) {
		t.Fatalf("expected 'grid' class, got: %s", html)
	}
}
