import json
from websocket import create_connection


class WebSocketRequest:
    def __init__(self, server_name, api_uri):
        self.ws = create_connection(f"ws://{server_name}/{api_uri}")

    def send_message(self, message):
        value = {'Text': message}
        self.ws.send(json.dumps(value))
        response = json.loads(self.ws.recv())
        return response['Count']
