# semver-comparison

Comparing [Masterminds/semver](https://github.com/Masterminds/semver),
[blang/semver](https://github.com/blang/semver), and
[hashicorp/go-version](https://github.com/hashicorp/go-version).

Result:
```
$ go test -v

=== RUN   TestHashicorp
    main_test.go:82: 3: hashicorp lib could not parse "==1.0.0" constraint
    main_test.go:82: 12: hashicorp lib could not parse "!1.0.0" constraint
    main_test.go:82: 14: hashicorp lib could not parse "~2.3" constraint
    main_test.go:82: 15: hashicorp lib could not parse "~2.3" constraint
    main_test.go:82: 16: hashicorp lib could not parse "~2.3.0" constraint
    main_test.go:82: 17: hashicorp lib could not parse "~2.3.0" constraint
    main_test.go:82: 18: hashicorp lib could not parse "^2.3" constraint
    main_test.go:82: 19: hashicorp lib could not parse "^2.3" constraint
    main_test.go:82: 20: hashicorp lib could not parse "^2.3.0" constraint
    main_test.go:82: 21: hashicorp lib could not parse "^2.3.0" constraint
    main_test.go:82: 22: hashicorp lib could not parse "5.x.x" constraint
    main_test.go:82: 23: hashicorp lib could not parse "5.*.*" constraint
    main_test.go:82: 24: hashicorp lib could not parse "5.*.*" constraint
    main_test.go:82: 25: hashicorp lib could not parse "5.x" constraint
    main_test.go:82: 26: hashicorp lib could not parse "5.*" constraint
    main_test.go:82: 30: hashicorp lib could not parse "0.5.0 - 2.0.0" constraint
    main_test.go:82: 31: hashicorp lib could not parse ">=0.5.0 <=2.0.0" constraint
    main_test.go:82: 32: hashicorp lib could not parse ">0.5.0 <2.0.0" constraint
    main_test.go:82: 33: hashicorp lib could not parse "<1.0.0 || >3.0.0" constraint
    main_test.go:82: 34: hashicorp lib could not parse ">1.0.0 <2.0.0 || >3.0.0 !=4.2.1" constraint
    main_test.go:82: 35: hashicorp lib could not parse ">1.0.0 <2.0.0 || >3.0.0 !=4.2.1" constraint
    main_test.go:82: 36: hashicorp lib could not parse ">1.0.0 <2.0.0 || >3.0.0 !4.2.1" constraint
    main_test.go:82: 37: hashicorp lib could not parse ">1.0.0 <2.0.0 || >3.0.0 !4.2.1" constraint
--- FAIL: TestHashicorp (0.00s)
=== RUN   TestBlang
    main_test.go:113: 1: blang lib could not parse "1.0" as semver
    main_test.go:113: 1: blang lib could not parse "1" as semver
    main_test.go:113: 1: blang lib could not parse "v1.0.0" as semver
    main_test.go:113: 2: blang lib could not parse "1.0" as semver
    main_test.go:113: 2: blang lib could not parse "1" as semver
    main_test.go:113: 2: blang lib could not parse "v1.0.0" as semver
    main_test.go:113: 3: blang lib could not parse "1.0" as semver
    main_test.go:113: 3: blang lib could not parse "1" as semver
    main_test.go:113: 3: blang lib could not parse "v1.0.0" as semver
    main_test.go:107: 5: blang lib could not parse "v1.0.0" constraint
    main_test.go:107: 6: blang lib could not parse "v1.0.0" constraint
    main_test.go:107: 14: blang lib could not parse "~2.3" constraint
    main_test.go:107: 15: blang lib could not parse "~2.3" constraint
    main_test.go:107: 16: blang lib could not parse "~2.3.0" constraint
    main_test.go:107: 17: blang lib could not parse "~2.3.0" constraint
    main_test.go:107: 18: blang lib could not parse "^2.3" constraint
    main_test.go:107: 19: blang lib could not parse "^2.3" constraint
    main_test.go:107: 20: blang lib could not parse "^2.3.0" constraint
    main_test.go:107: 21: blang lib could not parse "^2.3.0" constraint
    main_test.go:119: matching constraint "5.x.x" with semver "5.1.0": expected true, got false
    main_test.go:119: matching constraint "5.x.x" with semver "5.3.17": expected true, got false
    main_test.go:107: 23: blang lib could not parse "5.*.*" constraint
    main_test.go:107: 24: blang lib could not parse "5.*.*" constraint
    main_test.go:107: 26: blang lib could not parse "5.*" constraint
    main_test.go:113: 27: blang lib could not parse "5.0" as semver
    main_test.go:113: 27: blang lib could not parse "5" as semver
    main_test.go:107: 28: blang lib could not parse "5.0" constraint
    main_test.go:107: 29: blang lib could not parse "5" constraint
    main_test.go:119: matching constraint "0.5.0 - 2.0.0" with semver "0.5.0": expected true, got false
    main_test.go:119: matching constraint "0.5.0 - 2.0.0" with semver "1.2.3": expected true, got false
    main_test.go:119: matching constraint "0.5.0 - 2.0.0" with semver "2.0.0": expected true, got false
--- FAIL: TestBlang (0.00s)
=== RUN   TestMasterminds
    main_test.go:132: 3: Masterminds lib could not parse "==1.0.0" constraint
    main_test.go:132: 12: Masterminds lib could not parse "!1.0.0" constraint
    main_test.go:132: 36: Masterminds lib could not parse ">1.0.0 <2.0.0 || >3.0.0 !4.2.1" constraint
    main_test.go:132: 37: Masterminds lib could not parse ">1.0.0 <2.0.0 || >3.0.0 !4.2.1" constraint
--- FAIL: TestMasterminds (0.00s)
FAIL
exit status 1
FAIL	github.com/kubasobon/semver-comparison	0.002s
```
