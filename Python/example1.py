"""Basic example script which accesses data through a REMS instance's REST API 
Uses the endpoint "my-applications" to get information about a user's own applications
Refer to https://rems-demo.rahtiapp.fi/swagger-ui/index.html#/applications/get_api_my_applications

Usage is just:

python example1.py
"""

import requests

apiUrl = "https://rems-demo.rahtiapp.fi/api"
apiEnd = "my-applications"
url = apiUrl + "/" + apiEnd
apiKey = "55"
userId = "RDapplicant1@funet.fi"

# to add to the request's headers:
headers = {'x-rems-api-key':apiKey, 'x-rems-user-id':userId}

# make the request:
print('sending request to', url)
r= requests.get(url, headers=headers)

# and check the response:
if not r.ok:
	# oops that didn't work
	print('The request failed with response status', r.status_code)
else:
	# print out a bit of information from the response
	json = r.json()
	print('There were', len(json), 'records returned')
	print('Here is the first of these, in json format:')
	print(json[0])
