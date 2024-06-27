import json
import requests

url = "http://127.0.0.1:8000/api/stations"

with open("locations.json") as f:
    mock = json.loads(f.read())
    for loc in mock:
        print(json.dumps(loc))
        requests.post(url, json.dumps(loc))