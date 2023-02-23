package main

import (
	"testing"

	masterminds "github.com/Masterminds/semver/v3"
	blang "github.com/blang/semver/v4"
	hashicorp "github.com/hashicorp/go-version"
)

type Case struct {
	// Constraint is a semantic version constraint, e.g. ">2.0.0".
	Constraint string
	// Semvers is a list of valid Semantic Versions to be checked using the
	// Constraint, e.g. "v1.0.0", "5.7", "v3.1.0-alpha".
	Semvers []string
	// ExpectedMatch is an expected outcome of comparing each of Semvers
	// against the Constraint.
	ExpectedMatch bool
}

var testCases = []Case{
	// equality cases (=, ==, no operator)
	{Constraint: "1.0.0", Semvers: []string{"1.0.0", "1.0", "1", "v1.0.0"}, ExpectedMatch: true},
	{Constraint: "=1.0.0", Semvers: []string{"1.0.0", "1.0", "1", "v1.0.0"}, ExpectedMatch: true},
	{Constraint: "==1.0.0", Semvers: []string{"1.0.0", "1.0", "1", "v1.0.0"}, ExpectedMatch: true},
	{Constraint: "1.0.0", Semvers: []string{"1.0.1", "2.0.0", "0.0.1"}, ExpectedMatch: false},
	// Does 'v' prefix make a difference?
	{Constraint: "v1.0.0", Semvers: []string{"1.0.0", "1.0", "1", "v1.0.0"}, ExpectedMatch: true},
	{Constraint: "v1.0.0", Semvers: []string{"1.0.1", "2.0.0", "0.0.1"}, ExpectedMatch: false},
	// less-than cases (<)
	{Constraint: "<1.2.3", Semvers: []string{"0.5.0", "1.2.2"}, ExpectedMatch: true},
	{Constraint: "<1.2.3", Semvers: []string{"2.5.0", "1.2.4", "1.2.3"}, ExpectedMatch: false},
	// less-or-equal cases (<=)
	{Constraint: "<=1.2.3", Semvers: []string{"0.5.0", "1.2.2", "1.2.3"}, ExpectedMatch: true},
	{Constraint: "<=1.2.3", Semvers: []string{"2.5.0", "1.2.4"}, ExpectedMatch: false},
	// greater-than cases (>)
	{Constraint: ">1.2.3", Semvers: []string{"1.2.4", "2.0.0", "19.0.0"}, ExpectedMatch: true},
	// not-equal (!, !=)
	{Constraint: "!1.0.0", Semvers: []string{"1.0.1", "0.9.9999"}, ExpectedMatch: true},
	{Constraint: "!=1.0.0", Semvers: []string{"1.0.1", "0.9.9999"}, ExpectedMatch: true},
	// tilde range/patch comparison (~)
	//   ~2.3 is equivalent to >= 2.3, < 2.4
	{Constraint: "~2.3", Semvers: []string{"2.3.0", "2.3.17"}, ExpectedMatch: true},
	{Constraint: "~2.3", Semvers: []string{"2.4", "2.4.0"}, ExpectedMatch: false},
	// caret range/major comparison (^)
	//   ^2.3 is equivalent to >= 2.3, < 3
	{Constraint: "^2.3", Semvers: []string{"2.3.0", "2.3.17", "2.4.0", "2.99.0"}, ExpectedMatch: true},
	{Constraint: "^2.3", Semvers: []string{"3.0", "3.0.0"}, ExpectedMatch: false},
	// wildcards (x, *, incomplete versions)
	{Constraint: "5.x.x", Semvers: []string{"5.0.0", "5.1.0", "5.3.17"}, ExpectedMatch: true},
	{Constraint: "5.*.*", Semvers: []string{"5.0.0", "5.1.0", "5.3.17"}, ExpectedMatch: true},
	{Constraint: "5.x", Semvers: []string{"5.0.0", "5.1.0", "5.3.17"}, ExpectedMatch: true},
	{Constraint: "5.*", Semvers: []string{"5.0.0", "5.1.0", "5.3.17"}, ExpectedMatch: true},
	{Constraint: "5", Semvers: []string{"5.0.0", "5.1.0", "5.3.17"}, ExpectedMatch: true},
	{Constraint: "5.*.*", Semvers: []string{"v5.0", "v5"}, ExpectedMatch: true},
	// ranges, hyphen ranges, AND, OR, exclusions
	{Constraint: "0.5.0 - 2.0.0", Semvers: []string{"0.5.0", "1.2.3", "2.0.0"}, ExpectedMatch: true},
	{Constraint: ">=0.5.0 <=2.0.0", Semvers: []string{"0.5.0", "1.2.3", "2.0.0"}, ExpectedMatch: true},
	{Constraint: ">0.5.0 <2.0.0", Semvers: []string{"0.5.0", "2.0.0"}, ExpectedMatch: false},
	{Constraint: "<1.0.0 || >3.0.0", Semvers: []string{"0.2.0", "3.1.0"}, ExpectedMatch: true},
	{Constraint: ">1.0.0 <2.0.0 || >3.0.0 !=4.2.1", Semvers: []string{"1.2.3", "1.9.9", "3.1.1"}, ExpectedMatch: true},
	{Constraint: ">1.0.0 <2.0.0 || >3.0.0 !=4.2.1", Semvers: []string{"4.2.1", "2.1.1"}, ExpectedMatch: false},
	{Constraint: ">1.0.0 <2.0.0 || >3.0.0 !4.2.1", Semvers: []string{"1.2.3", "1.9.9", "3.1.1"}, ExpectedMatch: true},
	{Constraint: ">1.0.0 <2.0.0 || >3.0.0 !4.2.1", Semvers: []string{"4.2.1", "2.1.1"}, ExpectedMatch: false},
}

func TestHashicorp(t *testing.T) {
	for i, testCase := range testCases {
		constraint, err := hashicorp.NewConstraint(testCase.Constraint)
		if err != nil {
			t.Errorf("%d: hashicorp lib could not parse %q constraint", i+1, testCase.Constraint)
			continue
		}
		for _, semverString := range testCase.Semvers {
			semver, err := hashicorp.NewSemver(semverString)
			if err != nil {
				t.Errorf("%d: hashicorp lib could not parse %q as semver", i+1, semverString)
				continue
			}

			outcome := constraint.Check(semver)
			if outcome != testCase.ExpectedMatch {
				t.Errorf(
					"matching constraint %q with semver %q: expected %v, got %v",
					testCase.Constraint, semverString, testCase.ExpectedMatch, outcome,
				)
			}
		}
	}
}

func TestBlang(t *testing.T) {
	for i, testCase := range testCases {
		constraint, err := blang.ParseRange(testCase.Constraint)
		if err != nil {
			t.Errorf("%d: blang lib could not parse %q constraint", i+1, testCase.Constraint)
			continue
		}
		for _, semverString := range testCase.Semvers {
			semver, err := blang.Make(semverString)
			if err != nil {
				t.Errorf("%d: blang lib could not parse %q as semver", i+1, semverString)
				continue
			}

			outcome := constraint(semver)
			if outcome != testCase.ExpectedMatch {
				t.Errorf(
					"matching constraint %q with semver %q: expected %v, got %v",
					testCase.Constraint, semverString, testCase.ExpectedMatch, outcome,
				)
			}
		}
	}
}

func TestMasterminds(t *testing.T) {
	for i, testCase := range testCases {
		constraint, err := masterminds.NewConstraint(testCase.Constraint)
		if err != nil {
			t.Errorf("%d: Masterminds lib could not parse %q constraint", i+1, testCase.Constraint)
			continue
		}
		for _, semverString := range testCase.Semvers {
			semver, err := masterminds.NewVersion(semverString)
			if err != nil {
				t.Errorf("%d: Masterminds lib could not parse %q as semver", i+1, semverString)
				continue
			}

			outcome := constraint.Check(semver)
			if outcome != testCase.ExpectedMatch {
				t.Errorf(
					"matching constraint %q with semver %q: expected %v, got %v",
					testCase.Constraint, semverString, testCase.ExpectedMatch, outcome,
				)
			}
		}
	}
}
