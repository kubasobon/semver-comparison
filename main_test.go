package main

import (
	"testing"

	masterminds "github.com/Masterminds/semver/v3"
	blang "github.com/blang/semver/v4"
	hashicorp "github.com/hashicorp/go-version"
)

type Case struct {
	Constraint    string
	Semvers       []string
	ExpectedMatch bool
}

var testCases = []Case{
	{Constraint: "1.0.0", Semvers: []string{"1.0.0"}, ExpectedMatch: true},
}

func TestHashicorp(t *testing.T) {
	for _, testCase := range testCases {
		constraint, err := hashicorp.NewConstraint(testCase.Constraint)
		if err != nil {
			t.Errorf("hashicorp lib could not parse %q constraint", testCase.Constraint)
		}
		for _, semverString := range testCase.Semvers {
			semver, err := hashicorp.NewSemver(semverString)
			if err != nil {
				t.Errorf("hashicorp lib could not parse %q as semver", semverString)
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
	for _, testCase := range testCases {
		constraint, err := blang.ParseRange(testCase.Constraint)
		if err != nil {
			t.Errorf("blang lib could not parse %q constraint", testCase.Constraint)
		}
		for _, semverString := range testCase.Semvers {
			semver, err := blang.Make(semverString)
			if err != nil {
				t.Errorf("blang lib could not parse %q as semver", semverString)
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
	for _, testCase := range testCases {
		constraint, err := masterminds.NewConstraint(testCase.Constraint)
		if err != nil {
			t.Errorf("Masterminds lib could not parse %q constraint", testCase.Constraint)
		}
		for _, semverString := range testCase.Semvers {
			semver, err := masterminds.NewVersion(semverString)
			if err != nil {
				t.Errorf("Masterminds lib could not parse %q as semver", semverString)
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
