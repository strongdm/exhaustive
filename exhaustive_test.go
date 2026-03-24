package sdmexhaustive

import (
	"regexp"
	"testing"

	"golang.org/x/tools/go/analysis/analysistest"
)

func TestExhaustive(t *testing.T) {
	runTest := func(t *testing.T, pattern string, setup ...func()) {
		t.Helper()
		t.Run(pattern, func(t *testing.T) {
			resetFlags()
			// default to checking switch and map for test.
			fCheck = stringsFlag{
				[]string{
					string(elementSwitch),
					string(elementMap),
				},
				nil,
			}
			for _, f := range setup {
				f()
			}
			analysistest.Run(t, analysistest.TestData(), Analyzer, pattern)
		})
	}

	if !testing.Short() {
		// Analysis of code that uses complex packages, such as package os and
		// package reflect, should not fail.
		runTest(t, "complexpkg/...")
	}

	// Enum discovery, enum types.
	runTest(t, "enum/...")

	// Tests for the -check-generated flag.
	runTest(t, "generated-file/check-generated-off/...")
	runTest(t, "generated-file/check-generated-on/...", func() { fCheckGenerated = true })

	// Tests for the -default-signifies-exhaustive flag.
	// (For tests with this flag off, see other testdata packages
	// such as "general/...".)
	runTest(t, "default-signifies-exhaustive/default-absent/...", func() { fDefaultSignifiesExhaustive = true })
	runTest(t, "default-signifies-exhaustive/default-present/...", func() { fDefaultSignifiesExhaustive = true })

	// These tests exercise the default-case-required flag and its escape comment
	runTest(t, "default-case-required/default-required/...", func() { fDefaultCaseRequired = true })
	runTest(t, "default-case-required/default-not-required/...", func() { fDefaultCaseRequired = false })

	// Tests for -ignore-enum-members and -ignore-enum-types flags.
	runTest(t, "ignore-pattern/...", func() {
		fIgnoreEnumMembers = regexpFlag{
			regexp.MustCompile(`_UNSPECIFIED$|^general/y\.Echinodermata$|^ignore-pattern\.User$`),
		}
		fIgnoreEnumTypes = regexpFlag{
			regexp.MustCompile(`label|^reflect\.Kind$|^time\.Duration$`),
		}
	})

	// Tests for -package-scope-only flag.
	runTest(t, "scope/allscope/...")
	runTest(t, "scope/pkgscope/...", func() { fPackageScopeOnly = true })

	// Program elements with ignore comment should not be
	// checked during implicitly exhaustive mode.
	runTest(t, "ignore-comment/...")

	// Program elements without enforce comment should not be
	// checked in explicitly exhaustive mode.
	runTest(t, "enforce-comment/...", func() {
		fExplicitExhaustiveSwitch = true
		fExplicitExhaustiveMap = true
	})

	// Package-level enforce: //exhaustive:enforce on package clause.
	// In explicit mode, package-level enforce activates checking for all
	// switches/maps in the package; per-switch ignore overrides.
	runTest(t, "package-enforce/explicit-switch/...", func() {
		fExplicitExhaustiveSwitch = true
		fExplicitExhaustiveMap = true
	})
	runTest(t, "package-enforce/explicit-map/...", func() {
		fExplicitExhaustiveSwitch = true
		fExplicitExhaustiveMap = true
	})
	// In implicit mode, package-level enforce is redundant but harmless.
	runTest(t, "package-enforce/implicit-switch/...")
	// Without package-level enforce, explicit mode requires per-switch enforce.
	runTest(t, "package-enforce/no-directive/...", func() {
		fExplicitExhaustiveSwitch = true
		fExplicitExhaustiveMap = true
	})

	// Package-level ignore: //exhaustive:ignore on package clause.
	// In non-explicit mode, package-level ignore skips all switches/maps
	// in the package; per-statement enforce overrides.
	runTest(t, "package-ignore/implicit-switch/...")
	runTest(t, "package-ignore/implicit-map/...")

	// Conflicting package-level directives should produce a diagnostic.
	runTest(t, "package-conflict/...")

	// To satisfy exhaustiveness, it is sufficient for each unique constant
	// value of the members to be listed, not each member by name.
	runTest(t, "duplicate-enum-value/...")

	runTest(t, "typealias/...")
	runTest(t, "typeparam/...")

	// mixture of general tests.
	runTest(t, "general/...")
}

func assertNoError(t *testing.T, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("got %s, want nil error", err)
	}
}
