# Alfred Yubico Auth

This workflow allows quick searching and filling, and copying of OTP codes from a supported Yubikey.

So far, it has been tested with a Yubikey NEO on a MacBook Pro running macOS Catalina. I have no other devices to test with, but bug reports and patches may still be reviewed.

## Building

Building requires [`mage`](https://magefile.org/)

To see all targets and their descriptions, run `mage -l`.

## Credits

This uses [deanishe/awgo](https://github.com/deanishe/awgo) to interface with Alfred and [yawn/ykoath](https://github.com/yawn/ykoath) for interracting with the Yubikey
