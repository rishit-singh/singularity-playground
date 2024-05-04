import os
import sys
import requests
import json

BaseURL = os.getenv("BASE_URL")
Credentials = os.getenv("CRED")
Username, Password = Credentials.split(':')

def PrettyPrint(data) -> str:
    return json.dumps(ListStorages(), indent=2)

def ExploreStorage(storage: str, path: str):
    response = requests.get(f"{BaseURL}/storage/{storage}/explore/{path}", auth=(Username, Password))
    
    return response.json() 

def ListStorages() -> list[str]:
    response = requests.get(f"{BaseURL}/storage", auth=(Username, Password))
    
    return response.json() 

# Fetch all the prep name
def SendMetadata(metadata: dict[str, str], storage: str):
    response = requests.patch(f"{BaseURL}/storage/{storage}", auth=(Username, Password), json={"config": {"metadata": metadata}})
    
    return response.content 

print(PrettyPrint(ExploreStorage(sys.argv[1], sys.argv[2])))
# print(SendMetadata(json.dumps({"": ""}, sys.argv[1])))
#print(DeletePreparations(sys.argv[1]))

