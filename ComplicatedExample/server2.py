from http.server import BaseHTTPRequestHandler, HTTPServer

class handler(BaseHTTPRequestHandler):
    def do_GET(self):
        self.send_response(200)
        self.end_headers()
        self.wfile.write(b"Hello from server2")

if __name__ == "__main__":
    print("Server 2 running on port 80")
    HTTPServer(('', 80), handler).serve_forever()