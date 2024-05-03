import os
import sys
import requests
import json

BaseURL = os.getenv("BASE_URL")
Credentials = os.getenv("CRED")
Username, Password = Credentials.split(':')

# Fetch all the prep name
def ListPreparations() -> list[str]:
    response = requests.get(f"{BaseURL}/preparation", auth=(Username, Password))
    
    return response.json() 

def DeletePreparations(outputStorage: str) -> bool:
    for prep in ListPreparations():
        prepName = prep["name"]
        print("Deleting {prepName}")
        response = requests.delete(f"{BaseURL}/preparation/{prep['name']}/output/{outputStorage}", auth=(Username, Password))  
        print(response.content)
    return True 

print(json.dumps(ListPreparations(), indent=2))
#print(DeletePreparations(sys.argv[1]))

