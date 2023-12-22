from .base import BaseResource
from impact_api.database import UsersDatabase


class UsersResource(BaseResource):
    def __init__(self):
        self.users_db = UsersDatabase()
