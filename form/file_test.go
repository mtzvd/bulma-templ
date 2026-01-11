package form

import (
	"strings"
	"testing"

	"github.com/a-h/templ"
	"github.com/mtzvd/bulma-templ/elements"
)

func TestFile_DefaultRendering(t *testing.T) {
	html := render(t,
		File(FileProps{}, nil),
	)

	if !strings.Contains(html, `<div`) {
		t.Fatalf("expected <div> tag, got: %s", html)
	}
	if !strings.Contains(html, `file`) {
		t.Fatalf("expected 'file' class, got: %s", html)
	}
}

func TestFile_FullWidthModifier(t *testing.T) {
	html := render(t,
		File(FileProps{FullWidth: true}, nil),
	)

	if !strings.Contains(html, `is-fullwidth`) {
		t.Fatalf("expected is-fullwidth class when FullWidth=true, got: %s", html)
	}
}

func TestFile_BoxedModifier(t *testing.T) {
	html := render(t,
		File(FileProps{Boxed: true}, nil),
	)

	if !strings.Contains(html, `is-boxed`) {
		t.Fatalf("expected is-boxed class when Boxed=true, got: %s", html)
	}
}

func TestFile_HasNameModifier(t *testing.T) {
	html := render(t,
		File(FileProps{HasName: true}, nil),
	)

	if !strings.Contains(html, `has-name`) {
		t.Fatalf("expected has-name class when HasName=true, got: %s", html)
	}
}

func TestFile_ColorAndSize(t *testing.T) {
	html := render(t,
		File(FileProps{
			Color: "is-warning",
			Size:  "is-large",
		}, nil),
	)

	if !strings.Contains(html, `is-warning`) {
		t.Fatalf("expected is-warning class, got: %s", html)
	}
	if !strings.Contains(html, `is-large`) {
		t.Fatalf("expected is-large class, got: %s", html)
	}
}

func TestFile_AlignModifier(t *testing.T) {
	html := render(t,
		File(FileProps{Align: "is-right"}, nil),
	)

	if !strings.Contains(html, `is-right`) {
		t.Fatalf("expected is-right class, got: %s", html)
	}
}

func TestFileInput_DefaultRendering(t *testing.T) {
	html := render(t,
		FileInput(FileInputProps{}),
	)

	if !strings.Contains(html, `<input`) {
		t.Fatalf("expected <input> tag, got: %s", html)
	}
	if !strings.Contains(html, `type="file"`) {
		t.Fatalf("expected type='file', got: %s", html)
	}
	if !strings.Contains(html, `file-input`) {
		t.Fatalf("expected file-input class, got: %s", html)
	}
}

func TestFileInput_DisabledState(t *testing.T) {
	html := render(t,
		FileInput(FileInputProps{Disabled: true}),
	)

	if !strings.Contains(html, `disabled="disabled"`) {
		t.Fatalf("expected disabled='disabled' when Disabled=true, got: %s", html)
	}
}

func TestFileInput_NotDisabledByDefault(t *testing.T) {
	html := render(t,
		FileInput(FileInputProps{}),
	)

	if strings.Contains(html, `disabled`) {
		t.Fatalf("expected no disabled attribute by default, got: %s", html)
	}
}

func TestFileInput_MultipleAttribute(t *testing.T) {
	html := render(t,
		FileInput(FileInputProps{
			Attr: templ.Attributes{
				"multiple": "multiple",
			},
		}),
	)

	if !strings.Contains(html, `multiple="multiple"`) {
		t.Fatalf("expected multiple attribute from Attr, got: %s", html)
	}
}

func TestFileInput_AcceptAttribute(t *testing.T) {
	html := render(t,
		FileInput(FileInputProps{
			Attr: templ.Attributes{
				"accept": ".pdf,.docx",
			},
		}),
	)

	if !strings.Contains(html, `accept=".pdf,.docx"`) {
		t.Fatalf("expected accept attribute from Attr, got: %s", html)
	}
}

// Smoke tests for simple wrapper components
func TestFileLabel_Rendering(t *testing.T) {
	html := render(t,
		FileLabel(FileLabelProps{}, elements.Items{
			elements.Html("Content"),
		}),
	)

	if !strings.Contains(html, `<label`) {
		t.Fatalf("expected <label> tag, got: %s", html)
	}
	if !strings.Contains(html, `file-label`) {
		t.Fatalf("expected file-label class, got: %s", html)
	}
}

func TestFileCTA_Rendering(t *testing.T) {
	html := render(t,
		FileCTA(FileCTAProps{}, elements.Items{
			elements.Html("Choose file"),
		}),
	)

	if !strings.Contains(html, `file-cta`) {
		t.Fatalf("expected file-cta class, got: %s", html)
	}
}

func TestFileName_Rendering(t *testing.T) {
	html := render(t,
		FileName(FileNameProps{}, elements.Items{
			elements.Html("document.pdf"),
		}),
	)

	if !strings.Contains(html, `file-name`) {
		t.Fatalf("expected file-name class, got: %s", html)
	}
	if !strings.Contains(html, `document.pdf`) {
		t.Fatalf("expected file name content, got: %s", html)
	}
}
