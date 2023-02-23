# semver-comparison

Comparing [Masterminds/semver](https://github.com/Masterminds/semver),
[blang/semver](https://github.com/blang/semver), and
[hashicorp/go-version](https://github.com/hashicorp/go-version).

Result:
```
$ go test -v

=== RUN   TestHashicorp
    main_test.go:72: 3: hashicorp lib could not parse "==1.0.0" constraint
    main_test.go:72: 12: hashicorp lib could not parse "!1.0.0" constraint
    main_test.go:72: 14: hashicorp lib could not parse "~2.3" constraint
    main_test.go:72: 15: hashicorp lib could not parse "~2.3" constraint
    main_test.go:72: 16: hashicorp lib could not parse "^2.3" constraint
    main_test.go:72: 17: hashicorp lib could not parse "^2.3" constraint
    main_test.go:72: 18: hashicorp lib could not parse "5.x.x" constraint
    main_test.go:72: 19: hashicorp lib could not parse "5.*.*" constraint
    main_test.go:72: 20: hashicorp lib could not parse "5.x" constraint
    main_test.go:72: 21: hashicorp lib could not parse "5.*" constraint
    main_test.go:84: matching constraint "5" with semver "5.1.0": expected true, got false
    main_test.go:84: matching constraint "5" with semver "5.3.17": expected true, got false
    main_test.go:72: 23: hashicorp lib could not parse "5.*.*" constraint
    main_test.go:72: 24: hashicorp lib could not parse "0.5.0 - 2.0.0" constraint
    main_test.go:72: 25: hashicorp lib could not parse ">=0.5.0 <=2.0.0" constraint
    main_test.go:72: 26: hashicorp lib could not parse ">0.5.0 <2.0.0" constraint
    main_test.go:72: 27: hashicorp lib could not parse "<1.0.0 || >3.0.0" constraint
    main_test.go:72: 28: hashicorp lib could not parse ">1.0.0 <2.0.0 || >3.0.0 !=4.2.1" constraint
    main_test.go:72: 29: hashicorp lib could not parse ">1.0.0 <2.0.0 || >3.0.0 !=4.2.1" constraint
    main_test.go:72: 30: hashicorp lib could not parse ">1.0.0 <2.0.0 || >3.0.0 !4.2.1" constraint
    main_test.go:72: 31: hashicorp lib could not parse ">1.0.0 <2.0.0 || >3.0.0 !4.2.1" constraint
--- FAIL: TestHashicorp (0.00s)
=== RUN   TestBlang
    main_test.go:103: 1: blang lib could not parse "1.0" as semver
    main_test.go:103: 1: blang lib could not parse "1" as semver
    main_test.go:103: 1: blang lib could not parse "v1.0.0" as semver
    main_test.go:103: 2: blang lib could not parse "1.0" as semver
    main_test.go:103: 2: blang lib could not parse "1" as semver
    main_test.go:103: 2: blang lib could not parse "v1.0.0" as semver
    main_test.go:103: 3: blang lib could not parse "1.0" as semver
    main_test.go:103: 3: blang lib could not parse "1" as semver
    main_test.go:103: 3: blang lib could not parse "v1.0.0" as semver
    main_test.go:97: 5: blang lib could not parse "v1.0.0" constraint
    main_test.go:97: 6: blang lib could not parse "v1.0.0" constraint
    main_test.go:97: 14: blang lib could not parse "~2.3" constraint
    main_test.go:97: 15: blang lib could not parse "~2.3" constraint
    main_test.go:97: 16: blang lib could not parse "^2.3" constraint
    main_test.go:97: 17: blang lib could not parse "^2.3" constraint
    main_test.go:109: matching constraint "5.x.x" with semver "5.1.0": expected true, got false
    main_test.go:109: matching constraint "5.x.x" with semver "5.3.17": expected true, got false
    main_test.go:97: 19: blang lib could not parse "5.*.*" constraint
    main_test.go:97: 21: blang lib could not parse "5.*" constraint
    main_test.go:97: 22: blang lib could not parse "5" constraint
    main_test.go:97: 23: blang lib could not parse "5.*.*" constraint
    main_test.go:109: matching constraint "0.5.0 - 2.0.0" with semver "0.5.0": expected true, got false
    main_test.go:109: matching constraint "0.5.0 - 2.0.0" with semver "1.2.3": expected true, got false
    main_test.go:109: matching constraint "0.5.0 - 2.0.0" with semver "2.0.0": expected true, got false
--- FAIL: TestBlang (0.00s)
=== RUN   TestMasterminds
    main_test.go:122: 3: Masterminds lib could not parse "==1.0.0" constraint
    main_test.go:122: 12: Masterminds lib could not parse "!1.0.0" constraint
    main_test.go:122: 30: Masterminds lib could not parse ">1.0.0 <2.0.0 || >3.0.0 !4.2.1" constraint
    main_test.go:122: 31: Masterminds lib could not parse ">1.0.0 <2.0.0 || >3.0.0 !4.2.1" constraint
--- FAIL: TestMasterminds (0.00s)
FAIL
exit status 1
FAIL	github.com/kubasobon/semver-comparison	0.003s
```
