module git.iamthefij.com/iamthefij/alfred-yubico-auth

go 1.20

// Right now requires https://github.com/vividboarder/ykoath branch: validate
replace github.com/yawn/ykoath => ./ykoath

require (
	git.iamthefij.com/iamthefij/slog v1.0.0
	github.com/deanishe/awgo v0.27.1
	github.com/magefile/mage v1.10.0
	github.com/yawn/ykoath v1.0.4
)

require (
	github.com/bmatcuk/doublestar v1.3.1 // indirect
	github.com/ebfe/scard v0.0.0-20190212122703-c3d1b1916a95 // indirect
	github.com/pkg/errors v0.8.1 // indirect
	go.deanishe.net/env v0.5.1 // indirect
	go.deanishe.net/fuzzy v1.0.0 // indirect
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9 // indirect
	golang.org/x/text v0.3.3 // indirect
	howett.net/plist v0.0.0-20200419221736-3b63eb3a43b5 // indirect
)
