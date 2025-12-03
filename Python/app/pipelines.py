import os
from langchain_groq import ChatGroq
from langchain.prompts import PromptTemplate
from app.prompts import NEXT_QUESTION_PROMPT, FINAL_GUESS_PROMPT
from dotenv import load_dotenv
load_dotenv()

class LLMEngine:

    def __init__(self):
        self.llm = ChatGroq(
            groq_api_key=os.getenv("GROQ_API_KEY"),
            model_name="llama-3.1-8b-instant",
            temperature=0.2
        )

        self.next_q = PromptTemplate(
            template=NEXT_QUESTION_PROMPT,
            input_variables=["history", "question_num"]
        )

        self.final_guess = PromptTemplate(
            template=FINAL_GUESS_PROMPT,
            input_variables=["history"]
        )

    def generate_next_question(self, history, q_num):
        prompt = self.next_q.format(history=history, question_num=q_num)
        print("PROMPT SENT TO LLM:", prompt)
        response = self.llm.invoke(prompt)
        print("LLM RESPONSE:", response)
        return response.content.strip()

    def generate_final_guess(self, history):
        prompt = self.final_guess.format(history=history)
        response = self.llm.predict(prompt)
        return response.strip()
