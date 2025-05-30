package check

import (
	"testing"
)

func TestFrontMatterCheck(t *testing.T) {
	testCases := []struct {
		Name              string
		Source            string
		Options           *FrontMatterOptions
		ExpectError       bool
		ExpectSubcategory string
	}{
		{
			Name:   "empty source",
			Source: ``,
		},
		{
			Name: "valid YAML with default options",
			Source: `
description: |-
  Example description
layout: "example"
page_title: Example Page Title
subcategory: Example Subcategory
`,
			ExpectSubcategory: "Example Subcategory",
		},
		{
			Name: "valid YAML section and Markdown with default options",
			Source: `
---
description: |-
  Example description
layout: "example"
page_title: Example Page Title
subcategory: Example Subcategory
---

# Markdown here we go!
`,
			ExpectSubcategory: "Example Subcategory",
		},
		{
			Name: "invalid YAML",
			Source: `
description: |-
  Example description
Extraneous newline
`,
			ExpectError: true,
		},
		{
			Name: "allowed subcategory option matching",
			Source: `
description: |-
  Example description
layout: "example"
page_title: Example Page Title
subcategory: Example Subcategory
`,
			ExpectSubcategory: "Example Subcategory",
			Options: &FrontMatterOptions{
				AllowedSubcategories: []string{"Example Subcategory"},
			},
		},
		{
			Name: "allowed subcategory option not matching",
			Source: `
description: |-
  Example description
layout: "example"
page_title: Example Page Title
subcategory: Example Subcategory
`,
			Options: &FrontMatterOptions{
				AllowedSubcategories: []string{"Another Subcategory"},
			},
			ExpectError: true,
		},
		{
			Name: "no description option",
			Source: `
description: |-
  Example description
layout: "example"
page_title: Example Page Title
subcategory: Example Subcategory
`,
			Options: &FrontMatterOptions{
				NoDescription: true,
			},
			ExpectError: true,
		},
		{
			Name: "no layout option",
			Source: `
description: |-
  Example description
layout: "example"
page_title: Example Page Title
subcategory: Example Subcategory
`,
			Options: &FrontMatterOptions{
				NoLayout: true,
			},
			ExpectError: true,
		},
		{
			Name: "no page_title option",
			Source: `
description: |-
  Example description
layout: "example"
page_title: Example Page Title
subcategory: Example Subcategory
`,
			Options: &FrontMatterOptions{
				NoPageTitle: true,
			},
			ExpectError: true,
		},
		{
			Name: "no sidebar_current option",
			Source: `
description: |-
  Example description
layout: "example"
page_title: Example Page Title
sidebar_current: "example_resource"
subcategory: Example Subcategory
`,
			Options: &FrontMatterOptions{
				NoSidebarCurrent: true,
			},
			ExpectError: true,
		},
		{
			Name: "no subcategory option",
			Source: `
description: |-
  Example description
layout: "example"
page_title: Example Page Title
subcategory: Example Subcategory
`,
			Options: &FrontMatterOptions{
				NoSubcategory: true,
			},
			ExpectError: true,
		},
		{
			Name: "require description option",
			Source: `
layout: "example"
page_title: Example Page Title
subcategory: Example Subcategory
`,
			Options: &FrontMatterOptions{
				RequireDescription: true,
			},
			ExpectError: true,
		},
		{
			Name: "require layout option",
			Source: `
description: |-
  Example description
page_title: Example Page Title
subcategory: Example Subcategory
`,
			Options: &FrontMatterOptions{
				RequireLayout: true,
			},
			ExpectError: true,
		},
		{
			Name: "require page_title option",
			Source: `
description: |-
  Example description
layout: "example"
subcategory: Example Subcategory
`,
			Options: &FrontMatterOptions{
				RequirePageTitle: true,
			},
			ExpectError: true,
		},
		{
			Name: "require subcategory option",
			Source: `
description: |-
  Example description
layout: "example"
page_title: Example Page Title
`,
			Options: &FrontMatterOptions{
				RequireSubcategory: true,
			},
			ExpectError: true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.Name, func(t *testing.T) {
			subcategory, err := NewFrontMatterCheck(testCase.Options).Run([]byte(testCase.Source))

			if err == nil && testCase.ExpectError {
				t.Errorf("expected error, got no error")
			}

			if err != nil && !testCase.ExpectError {
				t.Errorf("expected no error, got error: %s", err)
			}

			if got, want := subcategory, testCase.ExpectSubcategory; want != "" && *got != want {
				t.Errorf("expected subcategory %q, got: %q", want, *got)
			}
		})
	}
}
