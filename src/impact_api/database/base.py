from abc import ABC, abstractmethod


class BaseDatabase(ABC):
    @abstractmethod
    def __init__(self):
        pass

    @abstractmethod
    def get(self):
        pass

    @abstractmethod
    def create(self):
        pass

    @abstractmethod
    def create_many(self):
        pass

    @abstractmethod
    def update(self):
        pass

    @abstractmethod
    def delete(self):
        pass
