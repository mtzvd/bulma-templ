package components

import (
	"strings"
	"testing"
)

// Test Modal Active state (critical conditional logic)
func TestModal_DefaultState(t *testing.T) {
	html := render(t,
		Modal(ModalProps{}, nil),
	)

	if !strings.Contains(html, `class="modal"`) {
		t.Fatalf("expected class='modal', got: %s", html)
	}
	if strings.Contains(html, `is-active`) {
		t.Fatalf("expected no is-active class by default, got: %s", html)
	}
}

func TestModal_ActiveState(t *testing.T) {
	html := render(t,
		Modal(ModalProps{Active: true}, nil),
	)

	if !strings.Contains(html, `is-active`) {
		t.Fatalf("expected is-active class when Active=true, got: %s", html)
	}
}

// Test ModalClose aria-label and size modifier
func TestModalClose_DefaultRendering(t *testing.T) {
	html := render(t,
		ModalClose(ModalCloseProps{}),
	)

	if !strings.Contains(html, `<button`) {
		t.Fatalf("expected <button> tag, got: %s", html)
	}
	if !strings.Contains(html, `modal-close`) {
		t.Fatalf("expected modal-close class, got: %s", html)
	}
	if !strings.Contains(html, `aria-label="close"`) {
		t.Fatalf("expected aria-label='close', got: %s", html)
	}
}

func TestModalClose_SizeModifier(t *testing.T) {
	html := render(t,
		ModalClose(ModalCloseProps{Size: "is-large"}),
	)

	if !strings.Contains(html, `is-large`) {
		t.Fatalf("expected is-large class, got: %s", html)
	}
}

// Smoke tests for simple wrapper components (no conditional logic)
func TestModalBackground_Rendering(t *testing.T) {
	html := render(t,
		ModalBackground(ModalBackgroundProps{}),
	)

	if !strings.Contains(html, `class="modal-background"`) {
		t.Fatalf("expected class='modal-background', got: %s", html)
	}
}

func TestModalContent_Rendering(t *testing.T) {
	html := render(t,
		ModalContent(ModalContentProps{}, nil),
	)

	if !strings.Contains(html, `class="modal-content"`) {
		t.Fatalf("expected class='modal-content', got: %s", html)
	}
}

func TestModalCard_Rendering(t *testing.T) {
	html := render(t,
		ModalCard(ModalCardProps{}, nil),
	)

	if !strings.Contains(html, `class="modal-card"`) {
		t.Fatalf("expected class='modal-card', got: %s", html)
	}
}

func TestModalCardHead_Rendering(t *testing.T) {
	html := render(t,
		ModalCardHead(ModalCardHeadProps{}, nil),
	)

	if !strings.Contains(html, `<header`) {
		t.Fatalf("expected <header> tag, got: %s", html)
	}
	if !strings.Contains(html, `class="modal-card-head"`) {
		t.Fatalf("expected class='modal-card-head', got: %s", html)
	}
}

func TestModalCardBody_Rendering(t *testing.T) {
	html := render(t,
		ModalCardBody(ModalCardBodyProps{}, nil),
	)

	if !strings.Contains(html, `<section`) {
		t.Fatalf("expected <section> tag, got: %s", html)
	}
	if !strings.Contains(html, `class="modal-card-body"`) {
		t.Fatalf("expected class='modal-card-body', got: %s", html)
	}
}

func TestModalCardFoot_Rendering(t *testing.T) {
	html := render(t,
		ModalCardFoot(ModalCardFootProps{}, nil),
	)

	if !strings.Contains(html, `<footer`) {
		t.Fatalf("expected <footer> tag, got: %s", html)
	}
	if !strings.Contains(html, `class="modal-card-foot"`) {
		t.Fatalf("expected class='modal-card-foot', got: %s", html)
	}
}
