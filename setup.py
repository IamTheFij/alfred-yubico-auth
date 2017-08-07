#! /usr/bin/env python
# -*- coding: utf-8 -*-
from __future__ import absolute_import
from __future__ import unicode_literals

import os

from setuptools import find_packages
from setuptools import setup


def read_version():
    version_path = 'version'
    try:
        return open(
            os.path.join(os.path.dirname(__file__), version_path)
        ).read()
    except:
        pass


setup(
    name='Yubico Auth Workflow',
    version=read_version(),
    description='Yubico Auth workflow for Alfred',
    author='Ian Fijolek',
    author_email='ian@iamthefij.com.com',
    url='',
    packages=find_packages(exclude=['tests*']),
    install_requires=[],
    license='MIT',
)
