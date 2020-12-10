module git.iamthefij.com/iamthefij/alfred-yubico-auth

go 1.15

// Right now requires github.com/vividboarder/ykauth branch: validate
replace github.com/yawn/ykoath => ../ykoath

require (
	git.iamthefij.com/iamthefij/slog v1.0.0
	github.com/deanishe/awgo v0.27.1
	github.com/magefile/mage v1.10.0
	github.com/yawn/ykoath v1.0.4
	golang.org/x/crypto v0.0.0-20201208171446-5f87f3452ae9
)
