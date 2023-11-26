class UsersController:
    def on_get(self, req, resp, ctx):
        pass


class UsersNamespace:
    def __init__(self):
        self.ROUTES = [
            ('/', UsersController())
        ]
