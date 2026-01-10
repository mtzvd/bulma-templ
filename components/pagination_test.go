package components

import (
	"strings"
	"testing"

	"github.com/mtzvd/bulma-templ/elements"
)

func TestPagination_FreeOrdering(t *testing.T) {
	html := render(t,
		Pagination(
			PaginationProps{
				Content: elements.Items{
					PaginationNext(
						elements.Items{elements.Html("Next")},
					),
					PaginationList(
						elements.Items{
							PaginationItem(
								elements.Items{
									PaginationLink(
										PaginationLinkProps{},
										elements.Items{elements.Html("1")},
									),
								},
							),
						},
					),
					PaginationPrev(
						elements.Items{elements.Html("Prev")},
					),
				},
			},
		),
	)

	// The rendered order MUST match the explicit Items order.
	nextIdx := strings.Index(html, "pagination-next")
	listIdx := strings.Index(html, "pagination-list")
	prevIdx := strings.Index(html, "pagination-previous")

	if !(nextIdx < listIdx && listIdx < prevIdx) {
		t.Fatalf("pagination block order not preserved: %s", html)
	}
}

func TestPagination_OmittedBlocks(t *testing.T) {
	html := render(t,
		Pagination(
			PaginationProps{
				Content: elements.Items{
					PaginationPrev(elements.Items{}),
					PaginationList(elements.Items{}),
				},
			},
		),
	)

	// Empty blocks MUST NOT be rendered.
	if strings.Contains(html, "pagination-previous") {
		t.Fatalf("empty PaginationPrev must not be rendered: %s", html)
	}

	if strings.Contains(html, "pagination-list") {
		t.Fatalf("empty PaginationList must not be rendered: %s", html)
	}
}

func TestPaginationLink_ActiveState(t *testing.T) {
	html := render(t,
		PaginationLink(
			PaginationLinkProps{
				Active: true,
			},
			elements.Items{elements.Html("1")},
		),
	)

	if !strings.Contains(html, "is-current") {
		t.Fatalf("active pagination link must have is-current: %s", html)
	}
}

func TestPagination_NoImplicitLogic(t *testing.T) {
	html := render(t,
		Pagination(
			PaginationProps{
				Content: elements.Items{
					PaginationList(
						elements.Items{
							PaginationItem(
								elements.Items{
									PaginationLink(
										PaginationLinkProps{},
										elements.Items{elements.Html("X")},
									),
								},
							),
						},
					),
				},
			},
		),
	)

	// Pagination MUST NOT infer active or disabled state implicitly.
	if strings.Contains(html, "is-current") {
		t.Fatalf("pagination must not infer active state: %s", html)
	}
	if strings.Contains(html, "disabled") {
		t.Fatalf("pagination must not infer disabled state: %s", html)
	}
}
