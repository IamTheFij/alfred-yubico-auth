.PHONY: default
default: run

.PHONY: build
build: venv install-ports

.PHONY: install
install: venv
	./replace-workflow.sh

Yauth.alfredWorkflow: venv
	mkdir Yauth.alfredWorkflow
	cp -r alfred_yauth Yauth.alfredworkflow/
	cp -r venv Yauth.alfredWorkflow/
	cp info.plist Yauth.alfredWorkflow/
	cp icon.png Yauth.alfredWorkflow/

# Installs required MacPorts
.PHONY: install-ports
install-ports:
	sudo port install swig swig-python ykpers libu2f-host libusb

# Creates venv using MacPorts Python (Required for it to refrence libusb)
venv:
	virtualenv --python=/opt/local/bin/python2.7 venv
	./venv/bin/pip install -r ./requirements.txt

# Simple execution of the workflow to see all results
.PHONY: run
run: venv
	@./venv/bin/python -m alfred_yauth.main

# Runs workflow and prompts for Yubikey password
.PHONY: set-password
set-password: venv
	@./venv/bin/python -m alfred_yauth.main set-password

# Clears the virtualenv and other installed files
.PHONY: clean
clean:
	rm -fr venv Yauth.alfredWorkflow
	find . -name '*.pyc' -delete
	find . -name '__pycache__' -delete

# Install precommit hooks
.PHONY: intall-hooks
install-hooks:
	tox -e pre-commit -- install -f --install-hooks
