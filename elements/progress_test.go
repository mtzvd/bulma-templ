package elements

import (
	"strings"
	"testing"
)

func TestProgress_DefaultRendering(t *testing.T) {
	html := render(t,
		Progress(ProgressProps{}, nil),
	)

	if !strings.Contains(html, `<progress`) {
		t.Fatalf("expected <progress> tag, got: %s", html)
	}
	if !strings.Contains(html, `class="progress"`) {
		t.Fatalf("expected class='progress', got: %s", html)
	}
}

func TestProgress_WithValue(t *testing.T) {
	html := render(t,
		Progress(ProgressProps{Current: 60}, nil),
	)

	if !strings.Contains(html, `value="60"`) {
		t.Fatalf("expected value='60', got: %s", html)
	}
}

func TestProgress_WithMax(t *testing.T) {
	html := render(t,
		Progress(ProgressProps{Max: 100}, nil),
	)

	if !strings.Contains(html, `max="100"`) {
		t.Fatalf("expected max='100', got: %s", html)
	}
}

func TestProgress_UnsetSentinel(t *testing.T) {
	// ProgressUnset (-1) should omit the attribute
	html := render(t,
		Progress(ProgressProps{
			Max:     ProgressUnset,
			Current: ProgressUnset,
		}, nil),
	)

	if strings.Contains(html, `max=`) {
		t.Fatalf("expected no max attribute when ProgressUnset, got: %s", html)
	}
	if strings.Contains(html, `value=`) {
		t.Fatalf("expected no value attribute when ProgressUnset, got: %s", html)
	}
}

func TestProgress_IndeterminateState(t *testing.T) {
	// When Current is unset, progress bar is indeterminate
	html := render(t,
		Progress(ProgressProps{
			Max:     100,
			Current: ProgressUnset,
		}, nil),
	)

	if !strings.Contains(html, `max="100"`) {
		t.Fatalf("expected max='100', got: %s", html)
	}
	if strings.Contains(html, `value=`) {
		t.Fatalf("expected no value attribute for indeterminate state, got: %s", html)
	}
}

func TestProgress_ColorModifiers(t *testing.T) {
	html := render(t,
		Progress(ProgressProps{Color: "is-primary"}, nil),
	)

	if !strings.Contains(html, `is-primary`) {
		t.Fatalf("expected is-primary class, got: %s", html)
	}
}

func TestProgress_SizeModifiers(t *testing.T) {
	html := render(t,
		Progress(ProgressProps{Size: "is-large"}, nil),
	)

	if !strings.Contains(html, `is-large`) {
		t.Fatalf("expected is-large class, got: %s", html)
	}
}

func TestProgress_ZeroValue(t *testing.T) {
	// Zero is a valid value (0% complete), not a sentinel
	html := render(t,
		Progress(ProgressProps{
			Max:     100,
			Current: 0,
		}, nil),
	)

	if !strings.Contains(html, `value="0"`) {
		t.Fatalf("expected value='0' for zero progress, got: %s", html)
	}
	if !strings.Contains(html, `max="100"`) {
		t.Fatalf("expected max='100', got: %s", html)
	}
}
