import falcon
from impact_api.api.context import get_context


@falcon.before(get_context)
class UsersController:
    def on_get(self, req, resp, ctx):
        pass


class UsersNamespace:
    def __init__(self):
        self.ROUTES = [
            ('/', UsersController())
        ]
