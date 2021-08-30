import http.client
import json
import sys

args = sys.argv
if len(args) != 1 + 2:
  print("Need 2 arguments, got", args)
  exit(1)

endpointsite = args[1].split("@")
if len(endpointsite) != 2:
  print("Syntax err, check 1st arg")
  exit(1)

api_site = endpointsite[1]
endpointId = endpointsite[0]
hardwareId = args[2]

api_site_parts = api_site.split(".")
if len(api_site_parts) < 2 or api_site_parts[-2] != "neuralgalaxy":
  print("Bad API hostname")
  exit(1)

headers = {
  "x-ng-endpointid": endpointId,
  "x-ng-endpointtype": 4,
  "Content-type": "application/json",
  "Accept": "text/json"
}
body_obj = {
  "endpointId": endpointId,
  "hardwareId": hardwareId
}
body = json.dumps(body_obj)
conn = http.client.HTTPSConnection(api_site)
conn.request("POST", "/api/auth/noninteractive/login", body, headers)
res = conn.getresponse()
if res.status == 201:
  exit(0)
exit(1)
