from client.ws import WebSocketRequest
from sys import stdin, stdout, exit
from argparse import ArgumentParser, FileType, ArgumentError
from django.core.management.base import BaseCommand
import array


class Command(BaseCommand):

    def add_arguments(self, parser: ArgumentParser):
        parser.add_argument('input', nargs='?', type=FileType('r'),
                            default=stdin)
        parser.add_argument("--server-addr")

    def handle(self, *args, **options):
        input = options['input']
        server = options['server_addr']
        
        if not server:
            print("Param server-addr is required")
            exit(2)

        self.request = WebSocketRequest(server, 'ws')

        if input != stdin:
            lines = self.lines_from_file(input)
        elif not input.isatty():
            lines = input.readlines()
        else:
            lines = []

        if lines:
            self.process_from_array(lines)
        else:
            self.process_from_console()
        
    def lines_from_file(self, input):
        lines = []

        for line in input:
            lines.append(line.strip())

        return lines

    def process_from_array(self, array_input):
        array_input.sort()

        for line in array_input:
            try:
                message = line.strip()
                print(f"client> {message}")
                response = self.request.send_message(message)
                print(f"server> {response}")
            except Exception as error:
                print(f"Error: {error}")

    def process_from_console(self):
        while True:
            line = input('client> ')
            if 'exit' == line.strip():
                exit(0)
            else:
                try:
                    response = self.request.send_message(line)
                    print(f"server> {response}")
                except Exception as exception:
                    print(f"Error: {exception}")
