package elements

import (
	"strings"
	"testing"
)

// Smoke tests for thin wrapper components without conditional logic.
// These tests provide regression confidence but limited technical value.
// Protected by BaseElement tests.

func TestSmokeBox_Rendering(t *testing.T) {
	html := render(t,
		Box(BoxProps{}, Items{Html("Content")}),
	)

	if !strings.Contains(html, `box`) {
		t.Fatalf("expected 'box' class, got: %s", html)
	}
	if !strings.Contains(html, `Content`) {
		t.Fatalf("expected content, got: %s", html)
	}
}

func TestSmokeBlock_Rendering(t *testing.T) {
	html := render(t,
		Block(BlockProps{}, Items{Html("Block content")}),
	)

	if !strings.Contains(html, `block`) {
		t.Fatalf("expected 'block' class, got: %s", html)
	}
}

func TestSmokeContent_Rendering(t *testing.T) {
	html := render(t,
		Content(ContentProps{}, Items{Html("<p>Text</p>")}),
	)

	if !strings.Contains(html, `content`) {
		t.Fatalf("expected 'content' class, got: %s", html)
	}
}

func TestSmokeDelete_Rendering(t *testing.T) {
	html := render(t,
		Delete(DeleteProps{}),
	)

	if !strings.Contains(html, `delete`) {
		t.Fatalf("expected 'delete' class, got: %s", html)
	}
}

func TestSmokeIcon_Rendering(t *testing.T) {
	html := render(t,
		Icon(IconProps{IconClass: "fas fa-home"}),
	)

	if !strings.Contains(html, `icon`) {
		t.Fatalf("expected 'icon' class, got: %s", html)
	}
}

func TestSmokeImage_Rendering(t *testing.T) {
	html := render(t,
		Image(ImageProps{Ratio: "is-square"}),
	)

	if !strings.Contains(html, `image`) {
		t.Fatalf("expected 'image' class, got: %s", html)
	}
	if !strings.Contains(html, `is-square`) {
		t.Fatalf("expected ratio class, got: %s", html)
	}
}

func TestSmokeNotification_Rendering(t *testing.T) {
	html := render(t,
		Notification(NotificationProps{}, Items{Html("Alert")}),
	)

	if !strings.Contains(html, `notification`) {
		t.Fatalf("expected 'notification' class, got: %s", html)
	}
}

func TestSmokeTag_Rendering(t *testing.T) {
	html := render(t,
		Tag(TagProps{}, Items{Html("Label")}),
	)

	if !strings.Contains(html, `tag`) {
		t.Fatalf("expected 'tag' class, got: %s", html)
	}
}

func TestSmokeTitle_Rendering(t *testing.T) {
	html := render(t,
		Title(TitleProps{}, Items{Html("Heading")}),
	)

	if !strings.Contains(html, `title`) {
		t.Fatalf("expected 'title' class, got: %s", html)
	}
}

func TestSmokeSkeletonBlock_Rendering(t *testing.T) {
	html := render(t,
		SkeletonBlock(SkeletonBlockProps{}, nil),
	)

	if !strings.Contains(html, `skeleton-block`) {
		t.Fatalf("expected 'skeleton-block' class, got: %s", html)
	}
}

func TestSmokeTable_Rendering(t *testing.T) {
	html := render(t,
		Table(TableProps{}, Items{Html("<tr><td>Cell</td></tr>")}),
	)

	if !strings.Contains(html, `table`) {
		t.Fatalf("expected 'table' class, got: %s", html)
	}
}

func TestSmokeButtons_Rendering(t *testing.T) {
	html := render(t,
		Buttons(ButtonsProps{}, Items{
			Button(ButtonProps{}, Items{Html("Button 1")}),
			Button(ButtonProps{}, Items{Html("Button 2")}),
		}),
	)

	if !strings.Contains(html, `buttons`) {
		t.Fatalf("expected 'buttons' class, got: %s", html)
	}
}

func TestSmokeSkeletonLines_Rendering(t *testing.T) {
	html := render(t,
		SkeletonLines(SkeletonLinesProps{}, Items{
			Html("Line 1"),
			Html("Line 2"),
			Html("Line 3"),
		}),
	)
	if !strings.Contains(html, `skeleton-lines`) {
		t.Fatalf("expected 'skeleton-lines' class, got: %s", html)
	}
}

func TestSmokeTableContainer_Rendering(t *testing.T) {
	html := render(t,
		TableContainer(TableContainerProps{}, Items{
			Table(TableProps{}, Items{Html("Table content")}),
		}),
	)
	if !strings.Contains(html, `table-container`) {
		t.Fatalf("expected 'table-container' class, got: %s", html)
	}
}

func TestSmokeTags_Rendering(t *testing.T) {
	html := render(t,
		Tags(TagsProps{}, Items{
			Tag(TagProps{}, Items{Html("Tag 1")}),
			Tag(TagProps{}, Items{Html("Tag 2")}),
		}),
	)
	if !strings.Contains(html, `tags`) {
		t.Fatalf("expected 'tags' class, got: %s", html)
	}
}

func TestSmokeSubtitle_Rendering(t *testing.T) {
	html := render(t,
		Subtitle(SubtitleProps{}, Items{Html("Subtitle text")}),
	)
	if !strings.Contains(html, `subtitle`) {
		t.Fatalf("expected 'subtitle' class, got: %s", html)
	}
}
