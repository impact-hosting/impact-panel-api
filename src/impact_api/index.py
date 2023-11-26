from falcon import App

from impact_api.api.middleware import SessionMiddleware
from impact_api.api.controllers import UsersNamespace


MIDDLEWARE = [
    SessionMiddleware()
]
ROUTES = [
    ('/users', UsersNamespace)
]
app = App(cors_enable=True)

app.add_middleware(MIDDLEWARE)

for namespace_uri, namespace in ROUTES:
    for route_uri, controller in namespace.ROUTES:
        app.add_route(f'{namespace_uri}{route_uri}', controller)
