# Alfred Yubico Auth

An Alfred Workflow for auto filling authentication codes stored on your Yubikey.

## Notes

This is definitely a work in progress. There are a lot of rough edges yet to be polished, but here it goes.

* Requires some to be installed with a package manager
* There is no way to input your key password through the UI yet. Do that with `make set-password` and then it should work fine.
* Error handling is terrible right now. If things don't work, check the debug log in Alfred

## Installation

Clone this repo

```bash
git clone https://git.iamthefij.com/iamthefij/alfred-yubico-auth.git
```

Either install your dependencies manually or, if you have MacPorts, you can use:

```bash
make install-ports
```

Otherwise you need to install `swig swig-python ykpers libu2f-host libusb` some other way.

Finally up the virtualenv and install to your Alfred with

```bash
make install
```

## Credits

Uses the amazing [deanishe/alfred-workflow](https://github.com/deanishe/alfred-workflow) package
