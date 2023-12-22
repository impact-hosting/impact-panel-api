import multiprocessing


wsgi_app = 'impact_api:app'
bind = 'localhost:5000'
workers = multiprocessing.cpu_count() * 2 + 1
reload = True
worker_class = 'eventlet'

capture_output = True
accesslog = "./logs/gunicorn_access.log"
errorlog = "./logs/gunicorn_error.log"
access_log_format = ''
loglevel = 'debug'
