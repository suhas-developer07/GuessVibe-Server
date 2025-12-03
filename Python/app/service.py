import llm_pb2 as pb
from app.pipelines import LLMEngine


class LLMServiceHandler:

    def __init__(self):
        self.engine = LLMEngine()

    def _history_to_text(self, history_list):
        lines = []
        for qa in history_list:
            lines.append(f"Q: {qa.question}\nA: {qa.answer}")
        return "\n".join(lines)

    def GenerateFirstQuestion(self, request):
        history = ""  
        question = self.engine.generate_next_question(history, 1)
        return pb.LLMQuestionResponse(question=question)

    def GenerateNextQuestion(self, request):
        history = self._history_to_text(request.history)
        question = self.engine.generate_next_question(history, request.questionCount)
        return pb.LLMQuestionResponse(question=question)

    def GenerateFinalGuess(self, request):
        history = self._history_to_text(request.history)
        guess = self.engine.generate_final_guess(history)
        return pb.LLMGuessResponse(
            guess=guess,
            confidence=0.92,
            alternatives=[]
        )
