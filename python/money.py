#/usr/bin/python
from enum import Enum
from typing import List, Tuple

import time

class TransactionType(Enum):
    INCOME = 1
    EXPENSE = 2

class Transaction:
    def __init__(
        self,
        amount: float,
        category: TransactionType
    ) -> None:
        self.amount = amount
        self.category = category

    def getAmount(self) -> float:
        return self.getAmount

    def __repr__(self) -> str:
        return f"${self.amount} is {self.category}"

class Ledger:
    def __init__(
        self,
        name: str,
        transactions: List[Transaction],
    ) -> None:
        self.name = name,
        self.transactions = transactions

    def getName(self) -> str:
        return self.name

    def getTransactions(self) -> List[str]:
        return self.transactions

    def getTransactionByType(
        self,
        category: TransactionType,
    ) -> float:
        filter = [t.amount for t in self.transactions if t.category is category]
        return {
            'total': sum(filter),
            'count': len(filter),
        }

    def getNetIncome(self) -> float:
        income = self.getTransactionByType(TransactionType.INCOME)
        expense = self.getTransactionByType(TransactionType.EXPENSE)

        return (income['total'] - expense['total'])

    def __repr__(self) -> str:
        count = len(self.transactions)
        income = self.getTransactionByType(TransactionType.INCOME)['total']
        expense = self.getTransactionByType(TransactionType.EXPENSE)['total']
        net = self.getNetIncome()
        return f"This ledger, of {count} transactions, reported a net income of ${net} (${income} income, ${expense} expenses)"

class Ledgers:
    def __init__(
        self,
        ledgers: List[Ledger],
    ) -> None:
        self.ledgers = ledgers

    def __getTopOrBottomLedger(self, top: bool = True) -> Tuple[str, float]:
        result = None
        for l in self.ledgers:
            if top:
                if result is None or l.getNetIncome() > result[1]:
                    result = (l.getName(), l.getNetIncome())
            else:
                if result is None or l.getNetIncome() < result[1]:
                    result = (l.getName(), l.getNetIncome())

        return result

    def getTopLedger(self) -> Tuple[str, float]:
        return self.__getTopOrBottomLedger(True)

    def getBottomLedger(self) -> Tuple[str, float]:
        return self.__getTopOrBottomLedger(False)

    def __repr__(self) -> str:
        top_ledger, top_income = self.getTopLedger()
        bottom_ledger, bottom_income = self.getBottomLedger()
        return f"{top_ledger[0]} was the best ledger (with ${top_income} net income), and {bottom_ledger[0]} was the worst (with ${bottom_income} net income)"
        
logs = Ledgers(ledgers=[
    Ledger(
        name="Ledger A",
        transactions=[
            Transaction(amount=13, category=TransactionType.INCOME),
            Transaction(amount=6, category=TransactionType.EXPENSE),
        ]
    ),
    Ledger(
        name="Ledger B",
        transactions=[
            Transaction(amount=12, category=TransactionType.INCOME),
            Transaction(amount=6, category=TransactionType.EXPENSE),
        ]
    )
])

print(logs)


