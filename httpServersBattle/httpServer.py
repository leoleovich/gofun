#!/usr/bin/env python3

from http.server import BaseHTTPRequestHandler, HTTPServer

class testHTTPServer_RequestHandler(BaseHTTPRequestHandler):
    def do_GET(self):
        message = "Hello, world!\n"
        # Send response status code
        self.send_response(200)

        # Send headers
        self.send_header('Content-type','text/html')
        self.send_header('Content-Length',len(message))
        self.end_headers()


        # Write content as utf-8 data
        self.wfile.write(bytes(message, "utf8"))
        return

    def log_message(self, format, *args):
        return

def run():
    server_address = ('127.0.0.1', 7000)
    httpd = HTTPServer(server_address, testHTTPServer_RequestHandler)
    httpd.serve_forever()

run()