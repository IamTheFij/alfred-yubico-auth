Alfred Yubico Auth
==================

An Alfred Workflow for auto filling authentication codes stored on your Yubikey.

Notes
-----

This is definitely a work in progress. There are a lot of rough edges yet to be polished, but here it goes.

* There is no way to input your key password through the UI yet. Do that with `make set-password` and then it should work fine.
* Does not bundle dylibs for you. Hanve't quite figured out what needs to be done to make this portable
* Requires MacPorts to install dependencies
* Error handling is terrible right now. If things don't work, check the debug log in Alfred

Credits
-------

Uses the amazing [deanishe/alfred-workflow](https://github.com/deanishe/alfred-workflow) package
