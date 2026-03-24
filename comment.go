package sdmexhaustive

import (
	"fmt"
	"go/ast"
	"go/token"
	"strings"
)

const (
	exhaustiveComment                 = "//exhaustive:"
	ignoreComment                     = "ignore"
	enforceComment                    = "enforce"
	ignoreDefaultCaseRequiredComment  = "ignore-default-case-required"
	enforceDefaultCaseRequiredComment = "enforce-default-case-required"
)

type directive int64

const (
	ignoreDirective = 1 << iota
	enforceDirective
	ignoreDefaultCaseRequiredDirective
	enforceDefaultCaseRequiredDirective
)

type directiveSet int64

func parseDirectives(commentGroups []*ast.CommentGroup) (directiveSet, error) {
	var out directiveSet
	for _, commentGroup := range commentGroups {
		if commentGroup == nil {
			continue
		}
		for _, comment := range commentGroup.List {
			commentLine := comment.Text
			if !strings.HasPrefix(commentLine, exhaustiveComment) {
				continue
			}
			directive := commentLine[len(exhaustiveComment):]
			if whiteSpaceIndex := strings.IndexAny(directive, " \t"); whiteSpaceIndex != -1 {
				directive = directive[:whiteSpaceIndex]
			}
			switch directive {
			case ignoreComment:
				out |= ignoreDirective
			case enforceComment:
				out |= enforceDirective
			case ignoreDefaultCaseRequiredComment:
				out |= ignoreDefaultCaseRequiredDirective
			case enforceDefaultCaseRequiredComment:
				out |= enforceDefaultCaseRequiredDirective
			default:
				return out, fmt.Errorf("invalid directive %q", directive)
			}
		}
	}
	return out, out.validate()
}

func (d directiveSet) has(directive directive) bool {
	return int64(d)&int64(directive) != 0
}

func (d directiveSet) validate() error {
	enforceConflict := ignoreDirective | enforceDirective
	if d&(directiveSet(enforceConflict)) == directiveSet(enforceConflict) {
		return fmt.Errorf("conflicting directives %q and %q", ignoreComment, enforceComment)
	}
	defaultCaseRequiredConflict := ignoreDefaultCaseRequiredDirective | enforceDefaultCaseRequiredDirective
	if d&(directiveSet(defaultCaseRequiredConflict)) == directiveSet(defaultCaseRequiredConflict) {
		return fmt.Errorf("conflicting directives %q and %q", ignoreDefaultCaseRequiredComment, enforceDefaultCaseRequiredComment)
	}
	return nil
}

func fileCommentMap(fset *token.FileSet, file *ast.File) ast.CommentMap {
	return ast.NewCommentMap(fset, file, file.Comments)
}

// packageDirective represents the package-level directive state.
type packageDirective int

const (
	packageDirectiveNone    packageDirective = iota
	packageDirectiveEnforce                  // //exhaustive:enforce on package clause
	packageDirectiveIgnore                   // //exhaustive:ignore on package clause
)

// packageLevelDirective returns the package-level directive found on
// package clauses across all files in the package. A directive on any
// file's package clause applies to the entire package.
func packageLevelDirective(files []*ast.File) (packageDirective, error) {
	var out directiveSet
	for _, file := range files {
		if file.Doc == nil {
			continue
		}
		d, err := parseDirectives([]*ast.CommentGroup{file.Doc})
		if err != nil {
			return packageDirectiveNone, err
		}
		out |= d
	}
	if err := out.validate(); err != nil {
		return packageDirectiveNone, err
	}
	switch {
	case out.has(enforceDirective):
		return packageDirectiveEnforce, nil
	case out.has(ignoreDirective):
		return packageDirectiveIgnore, nil
	default:
		return packageDirectiveNone, nil
	}
}
