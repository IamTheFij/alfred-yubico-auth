module git.iamthefij.com/iamthefij/alfred-yubico-auth

go 1.20

// Right now requires https://github.com/vividboarder/ykoath branch: validate
replace github.com/yawn/ykoath => ./ykoath

// Right now requires https://github.com/iamthefij/awgo branch: alfred-5
replace github.com/deanishe/awgo => github.com/iamthefij/awgo v0.29.1-pre1

require (
	git.iamthefij.com/iamthefij/slog v1.0.0
	github.com/deanishe/awgo v0.29.1
	github.com/magefile/mage v1.14.0
	github.com/yawn/ykoath v1.0.4
)

require (
	github.com/bmatcuk/doublestar v1.3.4 // indirect
	github.com/ebfe/scard v0.0.0-20190212122703-c3d1b1916a95 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	go.deanishe.net/env v0.5.1 // indirect
	go.deanishe.net/fuzzy v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9 // indirect
	golang.org/x/text v0.8.0 // indirect
	howett.net/plist v0.0.0-20201203080718-1454fab16a06 // indirect
)
