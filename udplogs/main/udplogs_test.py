import socket
import time
import json


data = {
	"filename": "1xx.log",
	"time": 12312312,
	"content": "hello logs"
}

print(json.dumps(data))

sock = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
for i in range(10):
	data["filename"] = "xx%s.log" % i
	sock.sendto(json.dumps(data), ('127.0.0.1', 13333))

sock.close()
print("done!")

