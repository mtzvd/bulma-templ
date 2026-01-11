package form

import (
	"strings"
	"testing"
)

func TestTextarea_DefaultRendering(t *testing.T) {
	html := render(t,
		Textarea(TextareaProps{}),
	)

	if !strings.Contains(html, `<textarea`) {
		t.Fatalf("expected <textarea> tag, got: %s", html)
	}
	if !strings.Contains(html, `textarea`) {
		t.Fatalf("expected 'textarea' class, got: %s", html)
	}
}

func TestTextarea_ValueAndPlaceholder(t *testing.T) {
	html := render(t,
		Textarea(TextareaProps{
			Value:       "test content",
			Placeholder: "Enter text",
		}),
	)

	if !strings.Contains(html, `test content`) {
		t.Fatalf("expected value 'test content', got: %s", html)
	}
	if !strings.Contains(html, `placeholder="Enter text"`) {
		t.Fatalf("expected placeholder='Enter text', got: %s", html)
	}
}

func TestTextarea_RowsAttribute(t *testing.T) {
	html := render(t,
		Textarea(TextareaProps{Rows: 5}),
	)

	if !strings.Contains(html, `rows="5"`) {
		t.Fatalf("expected rows='5', got: %s", html)
	}
}

func TestTextarea_DisabledState(t *testing.T) {
	html := render(t,
		Textarea(TextareaProps{Disabled: true}),
	)

	if !strings.Contains(html, `disabled="disabled"`) {
		t.Fatalf("expected disabled='disabled' when Disabled=true, got: %s", html)
	}
}

func TestTextarea_ReadOnlyState(t *testing.T) {
	html := render(t,
		Textarea(TextareaProps{ReadOnly: true}),
	)

	if !strings.Contains(html, `readonly="readonly"`) {
		t.Fatalf("expected readonly='readonly' when ReadOnly=true, got: %s", html)
	}
}

func TestTextarea_NotDisabledByDefault(t *testing.T) {
	html := render(t,
		Textarea(TextareaProps{}),
	)

	if strings.Contains(html, `disabled`) {
		t.Fatalf("expected no disabled attribute by default, got: %s", html)
	}
}

func TestTextarea_NotReadOnlyByDefault(t *testing.T) {
	html := render(t,
		Textarea(TextareaProps{}),
	)

	if strings.Contains(html, `readonly`) {
		t.Fatalf("expected no readonly attribute by default, got: %s", html)
	}
}

func TestTextarea_ColorModifier(t *testing.T) {
	html := render(t,
		Textarea(TextareaProps{Color: "is-success"}),
	)

	if !strings.Contains(html, `is-success`) {
		t.Fatalf("expected is-success class, got: %s", html)
	}
}

func TestTextarea_SizeModifier(t *testing.T) {
	html := render(t,
		Textarea(TextareaProps{Size: "is-medium"}),
	)

	if !strings.Contains(html, `is-medium`) {
		t.Fatalf("expected is-medium class, got: %s", html)
	}
}
