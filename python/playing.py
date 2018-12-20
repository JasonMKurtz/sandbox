#!/usr/bin/python3.6
from typing import List, NamedTuple

class Person:
    def __init__(
        self,
        name: str,
        age: int,
    ) -> None:
        self.name = name
        self.age = age


class ValidatorBase:
    def __init__(self):
        pass
    def validate(self):
        pass


class ValidateFullName(ValidatorBase):
    """Validates the person has a first and last name"""
    def validate(self, p: Person) -> bool:
        return " " in p.name

class ValidateSaneAge(ValidatorBase):
    """Validates the person has a sane age"""
    def validate(self, p: Person) -> bool:
        return p.age >= 0 and p.age <= 100


class Validator:
    def __init__(
        self,
        person: Person,
        checks: List[ValidatorBase],
    ) -> None:
        self.is_valid = []
        self.chain = []
        for check in checks:
            if check.__doc__ is not None:
                self.is_valid.append(check.validate(person))
                self.chain.append(
                    f"{check.__class__.__name__}(): {check.__doc__}"
                )

    def valid(self) -> bool:
        return len([v for v in self.is_valid if not v]) == 0

    def get_chain(self) -> List[str]:
        return "\n".join([doc for doc in self.chain])
        

checks = Validator(
    Person(name="Jason Kurtz", age=100), [
        ValidateFullName(),
        ValidateSaneAge()
    ]
)


print(checks.valid())
print(checks.get_chain())

"""
$ ./playing.py
True
ValidateFullName(): Validates the person has a first and last name
ValidateSaneAge(): Validates the person has a sane age
"""
