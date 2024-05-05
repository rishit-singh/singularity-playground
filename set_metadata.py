import os
import sys
import json
import requests

BaseURL = os.getenv("BASE_URL")
Credentials = os.getenv("CRED")
Username, Password = Credentials.split(':')

def GetMetadataByPath():
    return

def PrettyPrint(data) -> str:
    return json.dumps(data, indent=2)

def SetMetadata(storage: str, metadata: list[dict]):
    response = requests.patch(f"{BaseURL}/storage/{storage}", auth=(Username, Password), json={"config": metadata})    

    return response.json() 

print(PrettyPrint(SetMetadata(sys.argv[1], {"META-MD5": "6e923b819d61b608cd66c3af2f801499"})))