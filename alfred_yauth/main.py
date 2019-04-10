# -*- coding: utf-8 -*-
import sys
from getpass import getpass
from time import time

from controller import APDUError
from controller import Controller
from controller import DeviceNotFoundError
from workflow import ICON_ACCOUNT
from workflow import ICON_ERROR
from workflow import Workflow3


YUBIKEY_CREDS_KEYCHAIN = 'yubico-auth-creds'


def cred_to_item_kwargs(cred):
    if cred.get('hidden'):
        return None
    return {
        'icon': ICON_ACCOUNT,
        'title': cred['name'],
        'subtitle': 'Copy to clipboard',
        'copytext': cred['code'],
        'arg': cred['code'],
        'valid': True,
    }


class YubicoAuth(Workflow3):
    _controller = None

    def get_controller(self):
        if not self._controller:
            self._controller = Controller()
            self._controller.refresh()
        return self._controller

    def ask_yubikey_password(self):
        """Prompts the user for their Yubikey password and stores it"""
        self.logger.debug('Set password')
        password_key = self.get_controller().derive_key(
            getpass('Yubikey Password:')
        )
        self.save_password(YUBIKEY_CREDS_KEYCHAIN, password_key)

        self.get_controller().refresh_credentials(time(), password_key)
        self.add_item(
            'Yubikey password set successfully',
            '',
            icon=ICON_ACCOUNT,
        )

    def get_yubikey_password(self):
        """Returns stored Yubikey password from keychain"""
        return self.get_password(YUBIKEY_CREDS_KEYCHAIN)

    def _get_positional_arg(self, position):
        """Safely return a positional argument"""
        if len(self.args) > position:
            return self.args[position]
        return None

    def get_command(self):
        """Get command out of the args as first parameter"""
        return self._get_positional_arg(0)

    def get_query(self):
        """Get query out of the args after first parameter"""
        if len(self.args) < 2:
            return None
        return ' '.join(self.args[1:])

    def _validate(self, command):
        """Validates that we can handle the current command"""
        # if self.get_api_key() is None:
        #     self.add_item(
        #         title='Missing API key',
        #         subtitle='Set variable in settings',
        #         icon=ICON_ACCOUNT,
        #         valid=False,
        #     )
        #     return False
        # if command == COMMAND_LOVE and self.get_recipient() is None:
        #     self.add_item(
        #         title='Recipient is required',
        #         icon=ICON_ERROR,
        #         valid=False,
        #     )
        #     return False
        return True

    def _add_cred_to_results(self, cred):
        self.logger.debug('Read {}'.format(cred.get('name')))
        item_args = cred_to_item_kwargs(cred)
        if item_args:
            self.add_item(**item_args)

    def list_credentials(self):
        password_key = self.get_yubikey_password()
        for cred in self.get_controller().list_credentials(password_key):
            self._add_cred_to_results(cred)

    def refresh_credentials(self):
        key = self.get_yubikey_password()
        for cred in self.get_controller().refresh_credentials(time(), key):
            self._add_cred_to_results(cred)

    def main(self):
        self.logger.debug('Starting...')
        command = self.get_command()

        if not self._validate(command):
            self.send_feedback()
            return

        command_action = None
        if command == 'set-password':
            command_action = self.ask_yubikey_password
        elif command == 'list':
            command_action = self.list_credentials
        else:
            command_action = self.refresh_credentials

        try:
            command_action()
        except DeviceNotFoundError:
            self.add_item(
                'Could not find device',
                'Is your Yubikey plugged in?',
                icon=ICON_ERROR,
            )
        except APDUError:
            self.add_item(
                'Could not communicate with device',
                'Is your Yubikey password set correctly?',
                icon=ICON_ERROR,
            )

        self.send_feedback()


def no_wf():
    controller = Controller()
    print(controller.get_features())
    print(controller.count_devices())
    print(controller.refresh())

    password = getpass('YubiKey password?')
    password_key = controller.derive_key(password)
    timestamp = time()
    print(controller.refresh_credentials(timestamp, password_key))
    creds = controller.list_credentials(password_key)

    print(creds)


def main(wf=None):
    if wf is None:
        no_wf()
    else:
        wf.main()


if __name__ == '__main__':
    # main()
    wf = YubicoAuth()
    sys.exit(wf.run(main))
