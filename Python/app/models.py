from dataclasses import dataclass
from typing import List

@dataclass
class QA:
    question: str
    answer: str

@dataclass
class SessionStateModel:
    session_id: str
    question_count: int
    history: List[QA]
