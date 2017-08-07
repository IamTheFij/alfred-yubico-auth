.PHONY: default
default: run

# Simple execution of the workflow to see all results
.PHONY: run
run: install-requirements
	./src/main.py

# Runs workflow and prompts for Yubikey password
.PHONY: set-password
set-password: install-requirements
	./src/main.py set-password

.PHONY: install-requirements
install-requirements: src/vendor src/libusb-1.0.dylib

# Installs libusb from /opt/local, where MacPorts installs it
src/libusb-1.0.dylib:
	cp /opt/local/lib/libusb-1.0.dylib ./src/

# Installs 3rd party packages into vendor directory
src/vendor:
	mkdir -p .pip-cache
	pip install -r ./requirements.txt -t ./src/vendor --cache-dir .pip-cache
	cp vendor.py src/vendor/__init__.py

# Creates virtualenv for testing using MacPorts Python
virtualenv:
	virtualenv --python=/opt/local/bin/python2.7 virtualenv
	./virtualenv_run/bin/pip install -r ./requirements.txt

# Runs workflow using virtualenv Python
.PHONY: virtualenv_run
virtualenv_run: virtualenv
	./virtualenv/bin/python src/main.py

# Clears the virtualenv and other installed files
.PHONY: clean
clean:
	rm -fr virtualenv src/vendor src/libusb-1.0.dylib
	find . -name '*.pyc' -delete
	find . -name '__pycache__' -delete

# Installs required MacPorts
.PHONY: install-ports
install-ports:
	sudo port install swig swig-python ykpers libu2f-host libusb

# Install precommit hooks
.PHONY: intall-hooks
install-hooks:
	tox -e pre-commit -- install -f --install-hooks
