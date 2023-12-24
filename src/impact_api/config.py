import re
import os
import logging

logger = logging.getLogger()


def load_env():
    with open('../../.env') as env_file:
        lines = env_file.readlines()
        for line in lines:
            if line.startswith('#'):
                logger.debug('load_env - Line is a comment. Skipping...')
                continue

            matches = re.match(r'^(?P<env_key>[0-9A-Z-_]+)=(?P<env_value>.+)$', re.IGNORECASE)
            if not matches:
                logger.debug('load_env - No valid match')
                continue

            key, value = matches.group('env_key', 'env_value')
            os.environ[key] = value

load_env()
