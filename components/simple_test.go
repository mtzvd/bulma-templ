package components

import (
	"strings"
	"testing"

	"github.com/mtzvd/bulma-templ/elements"
)

// Smoke tests for simple wrapper components without conditional logic.
// These tests provide regression confidence but limited technical value.

func TestSmokeBreadcrumb_Rendering(t *testing.T) {
	html := render(t,
		Breadcrumb(BreadcrumbProps{}, elements.Items{
			elements.Html("<li>Item</li>"),
		}),
	)

	if !strings.Contains(html, `breadcrumb`) {
		t.Fatalf("expected 'breadcrumb' class, got: %s", html)
	}
}

func TestSmokeMessage_Rendering(t *testing.T) {
	html := render(t,
		Message(MessageProps{}, elements.Items{
			elements.Html("Message content"),
		}),
	)

	if !strings.Contains(html, `message`) {
		t.Fatalf("expected 'message' class, got: %s", html)
	}
}

func TestSmokeMenu_Rendering(t *testing.T) {
	html := render(t,
		Menu(MenuProps{}, elements.Items{
			elements.Html("Menu content"),
		}),
	)

	if !strings.Contains(html, `menu`) {
		t.Fatalf("expected 'menu' class, got: %s", html)
	}
}

func TestSmokePanel_Rendering(t *testing.T) {
	html := render(t,
		Panel(PanelProps{}, elements.Items{
			elements.Html("Panel content"),
		}),
	)

	if !strings.Contains(html, `panel`) {
		t.Fatalf("expected 'panel' class, got: %s", html)
	}
}

func TestSmokeTabs_Rendering(t *testing.T) {
	html := render(t,
		Tabs(TabsProps{}, elements.Items{
			elements.Html("<ul><li>Tab</li></ul>"),
		}),
	)

	if !strings.Contains(html, `tabs`) {
		t.Fatalf("expected 'tabs' class, got: %s", html)
	}
}

func TestSmokeCard_Rendering(t *testing.T) {
	html := render(t,
		Card(CardProps{}, elements.Items{
			CardContent(CardContentProps{}, elements.Items{elements.Html("Content")}),
		}),
	)

	if !strings.Contains(html, `card`) {
		t.Fatalf("expected 'card' class, got: %s", html)
	}
}

func TestSmokeCardHeader_Rendering(t *testing.T) {
	html := render(t,
		CardHeader(CardHeaderProps{}, elements.Items{
			CardHeaderTitle(CardHeaderTitleProps{}, elements.Items{elements.Html("Title")}),
		}),
	)

	if !strings.Contains(html, `card-header`) {
		t.Fatalf("expected 'card-header' class, got: %s", html)
	}
}

func TestSmokeCardContent_Rendering(t *testing.T) {
	html := render(t,
		CardContent(CardContentProps{}, elements.Items{elements.Html("Content")}),
	)

	if !strings.Contains(html, `card-content`) {
		t.Fatalf("expected 'card-content' class, got: %s", html)
	}
}

func TestSmokeCardFooter_Rendering(t *testing.T) {
	html := render(t,
		CardFooter(CardFooterProps{}, elements.Items{
			CardFooterItem(CardFooterItemProps{}, elements.Items{elements.Html("Item")}),
		}),
	)

	if !strings.Contains(html, `card-footer`) {
		t.Fatalf("expected 'card-footer' class, got: %s", html)
	}
}

func TestSmokeBreadcrumbItem_Rendering(t *testing.T) {
	html := render(t,
		BreadcrumbItem(BreadcrumbItemProps{}, elements.Items{
			elements.Html("Item"),
		}),
	)

	if !strings.Contains(html, `<li`) {
		t.Fatalf("expected '<li' tag, got: %s", html)
	}
}

func TestSmokeBreadcrumbLink_Rendering(t *testing.T) {
	html := render(t,
		BreadcrumbLink(BreadcrumbLinkProps{}, elements.Items{
			elements.Html("Link"),
		}),
	)

	if !strings.Contains(html, `<a`) {
		t.Fatalf("expected '<a' tag, got: %s", html)
	}
}

func TestSmokeMenuLabel_Rendering(t *testing.T) {
	html := render(t,
		MenuLabel(MenuLabelProps{}, elements.Items{elements.Html("Label")}),
	)

	if !strings.Contains(html, `menu-label`) {
		t.Fatalf("expected 'menu-label' class, got: %s", html)
	}
}

func TestSmokeMenuList_Rendering(t *testing.T) {
	html := render(t,
		MenuList(MenuListProps{}, elements.Items{
			elements.Html("<li>Item</li>"),
		}),
	)

	if !strings.Contains(html, `menu-list`) {
		t.Fatalf("expected 'menu-list' class, got: %s", html)
	}
}

func TestSmokeMessage_Subcomponents(t *testing.T) {
	// Test MessageHeader
	header := render(t,
		MessageHeader(MessageHeaderProps{}, elements.Items{
			elements.Html("Header"),
		}),
	)
	if !strings.Contains(header, `message-header`) {
		t.Fatalf("expected 'message-header' class, got: %s", header)
	}

	// Test MessageBody
	body := render(t,
		MessageBody(MessageBodyProps{}, elements.Items{
			elements.Html("Body"),
		}),
	)
	if !strings.Contains(body, `message-body`) {
		t.Fatalf("expected 'message-body' class, got: %s", body)
	}
}

func TestSmokePanel_Subcomponents(t *testing.T) {
	// Test PanelHeading
	heading := render(t,
		PanelHeading(PanelHeadingProps{}, elements.Items{
			elements.Html("Heading"),
		}),
	)
	if !strings.Contains(heading, `panel-heading`) {
		t.Fatalf("expected 'panel-heading' class, got: %s", heading)
	}

	// Test PanelBlock
	block := render(t,
		PanelBlock(PanelBlockProps{}, elements.Items{
			elements.Html("Block"),
		}),
	)
	if !strings.Contains(block, `panel-block`) {
		t.Fatalf("expected 'panel-block' class, got: %s", block)
	}

	// Test PanelTabs
	tabs := render(t,
		PanelTabs(PanelTabsProps{}, elements.Items{
			elements.Html("<a>Tab</a>"),
		}),
	)
	if !strings.Contains(tabs, `panel-tabs`) {
		t.Fatalf("expected 'panel-tabs' class, got: %s", tabs)
	}
}

func TestSmokeTabs_Subcomponents(t *testing.T) {
	// Test TabsList
	list := render(t,
		TabsList(TabsListProps{}, elements.Items{
			elements.Html("<li>Tab</li>"),
		}),
	)
	if !strings.Contains(list, `<ul`) {
		t.Fatalf("expected '<ul' tag, got: %s", list)
	}

	// Test TabsItem
	item := render(t,
		TabsItem(TabsItemProps{}, elements.Items{
			elements.Html("Item"),
		}),
	)
	if !strings.Contains(item, `<li`) {
		t.Fatalf("expected '<li' tag, got: %s", item)
	}

	// Test TabsLink
	link := render(t,
		TabsLink(TabsLinkProps{}, elements.Items{
			elements.Html("Link"),
		}),
	)
	if !strings.Contains(link, `<a`) {
		t.Fatalf("expected '<a' tag, got: %s", link)
	}
}

func TestSmokeBreadcrumbList_Rendering(t *testing.T) {
	html := render(t,
		BreadcrumbList(BreadcrumbListProps{}, elements.Items{
			elements.Html("<li>Item</li>"),
		}),
	)

	if !strings.Contains(html, `<ul`) {
		t.Fatalf("expected '<ul' tag, got: %s", html)
	}
}

func TestSmokeBreadcrumbIcon_Rendering(t *testing.T) {
	html := render(t,
		BreadcrumbIcon(BreadcrumbIconProps{
			Icon: elements.IconProps{},
		}),
	)

	if !strings.Contains(html, `<span`) {
		t.Fatalf("expected '<span' tag, got: %s", html)
	}
}

func TestSmokeCardHeaderIcon_Rendering(t *testing.T) {
	html := render(t,
		CardHeaderIcon(CardHeaderIconProps{}, elements.Items{
			elements.Html("<i class=\"fa\"></i>"),
		}),
	)

	if !strings.Contains(html, `card-header-icon`) {
		t.Fatalf("expected 'card-header-icon' class, got: %s", html)
	}
}

func TestSmokeMenuItem_Rendering(t *testing.T) {
	html := render(t,
		MenuItem(MenuItemProps{}, elements.Items{
			elements.Html("Menu Item"),
		}),
	)

	if !strings.Contains(html, `<a`) {
		t.Fatalf("expected '<a' tag, got: %s", html)
	}
}

func TestSmokeNavbarStart_Rendering(t *testing.T) {
	html := render(t,
		NavbarStart(NavbarStartProps{}, elements.Items{
			elements.Html("Start"),
		}),
	)

	if !strings.Contains(html, `navbar-start`) {
		t.Fatalf("expected 'navbar-start' class, got: %s", html)
	}
}

func TestSmokeNavbarEnd_Rendering(t *testing.T) {
	html := render(t,
		NavbarEnd(NavbarEndProps{}, elements.Items{
			elements.Html("End"),
		}),
	)

	if !strings.Contains(html, `navbar-end`) {
		t.Fatalf("expected 'navbar-end' class, got: %s", html)
	}
}

func TestSmokePaginationEllipsis_Rendering(t *testing.T) {
	html := render(t,
		PaginationEllipsis(),
	)

	if !strings.Contains(html, `pagination-ellipsis`) {
		t.Fatalf("expected 'pagination-ellipsis' class, got: %s", html)
	}
}
