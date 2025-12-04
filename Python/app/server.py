import grpc
from concurrent import futures
import llm_pb2_grpc as pb_grpc
from app.service import LLMServiceHandler
import os


class LLMServer(pb_grpc.LLMServiceServicer):
    
    def __init__(self):
        self.service = LLMServiceHandler()

    def GenerateFirstQuestion(self, request, context):
        return self.service.GenerateFirstQuestion(request)

    def GenerateNextQuestion(self, request, context):
        return self.service.GenerateNextQuestion(request)

    def GenerateFinalGuess(self, request, context):
        return self.service.GenerateFinalGuess(request)

def serve():
    server = grpc.server(futures.ThreadPoolExecutor(max_workers=10))
    pb_grpc.add_LLMServiceServicer_to_server(LLMServer(), server)
    server.add_insecure_port("0.0.0.0:50051")
    print("🚀 Python LLM gRPC running on 50051...")
    server.start()
    server.wait_for_termination()

