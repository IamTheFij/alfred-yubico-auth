module git.iamthefij.com/iamthefij/alfred-yubico-auth

go 1.15

// Right now requires https://github.com/vividboarder/ykoath branch: validate
replace github.com/yawn/ykoath => ./ykoath

require (
	git.iamthefij.com/iamthefij/slog v1.0.0
	github.com/deanishe/awgo v0.27.1
	github.com/magefile/mage v1.10.0
	github.com/yawn/ykoath v1.0.4
)
