import multiprocessing
import os


wsgi_app = 'impact_api:app'
bind = f'localhost:{os.getenv("HTTP_PORT")}'
workers = multiprocessing.cpu_count() * 2 + 1
reload = True
worker_class = 'eventlet'

capture_output = True
accesslog = os.getenv('LOG_FILE')
errorlog = os.getenv('ERROR_LOG_FILE')
access_log_format = ''
loglevel = 'debug'

keyfile = os.getenv('SSL_KEY_FILE')
certfile = os.getenv('SSL_CERT_FILE')
