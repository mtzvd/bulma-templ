package columns

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

// Smoke tests for columns without conditional logic.

func TestSmokeColumns_Rendering(t *testing.T) {
	html := render(t,
		Columns(ColumnsProps{}, elements.Items{
			Column(ColumnProps{}, elements.Items{
				elements.Html("Column content"),
			}),
		}),
	)

	if !strings.Contains(html, `columns`) {
		t.Fatalf("expected 'columns' class, got: %s", html)
	}
}

func TestSmokeColumn_Rendering(t *testing.T) {
	html := render(t,
		Column(ColumnProps{}, elements.Items{
			elements.Html("Content"),
		}),
	)

	if !strings.Contains(html, `column`) {
		t.Fatalf("expected 'column' class, got: %s", html)
	}
}
