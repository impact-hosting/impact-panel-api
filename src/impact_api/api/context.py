from impact_api.api.resources import (UsersResource)


class Context:
    def __init__(self, req, resp):
        pass

    @property()
    def users_resource(self):
        return UsersResource()


def get_context(req, resp, resource, params):
    return Context(req, resp)
