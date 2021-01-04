# Alfred Yubico Auth

This workflow allows quick searching and filling, and copying of OTP codes from a supported Yubikey.

So far, it has been tested with a Yubikey NEO on a MacBook Pro running macOS Catalina. I have no other devices to test with, but bug reports and patches may still be reviewed.

## Cloning

Currently this package depends on a fork of [yawn/ykoath](https://github.com/yawn/ykoath). To allow this to be built directly from this repo, the fork is added as a git submodule. This can be cloned using `git clone --recurse-submodules` or cloning normally and then executing `git submodule update --init`. Once the change has been merged upstream, the submodule and this notice will go away.

## Building

Building requires [`mage`](https://magefile.org/)

To see all targets and their descriptions, run `mage -l`. The most basic ones are as follows:

* `mage install`: Build and install into your local machine for testing
* `mage dist`: Build bundle for distribution


## Credits

This uses [deanishe/awgo](https://github.com/deanishe/awgo) to interface with Alfred and [yawn/ykoath](https://github.com/yawn/ykoath) for interracting with the Yubikey
