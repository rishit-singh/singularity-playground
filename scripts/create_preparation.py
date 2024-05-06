import os
import re
import sys
import json
import requests

BaseURL = os.getenv("BASE_URL")
Credentials = os.getenv("CRED")
FtpCredentials = os.getenv("FTP_CRED")
AuthCredentials = tuple(Credentials.split(':'))
FtpUsername, FtpPassword = FtpCredentials.split(':')

def PrettyPrint(data) -> str:
    return json.dumps(data, indent=2)
# Fetch all the prep name
def ListPreparations() -> list[str]:
    response = requests.get(f"{BaseURL}/preparation", auth=AuthCredentials)
    
    return response.json() 

def DeletePreparations(prep: str, outputStorage: str) -> bool:
    response = requests.delete(f"{BaseURL}/preparation/{prep}/output/{outputStorage}", auth=AuthCredentials)  
    print(response.json())


def LoadMetadata(file: str) -> list[dict]:
    metadata = None

    with open(file, 'r') as fp:
        metadata = json.load(fp)
    
    return metadata

def ListStorages():
    response = requests.get(f"{BaseURL}/storage", auth=AuthCredentials)

    return response.json()

def StartScan(prep: str, source: str):
    response = requests.post(f"{BaseURL}/preparation/{prep}/source/{source}/start-scan", auth=AuthCredentials, json={})

    return response.json()

def CreatePreparation(storage: str):
    response = requests.post(f"{BaseURL}/preparation", auth=AuthCredentials, json={
        "name": f"prep-{"-".join(storage.split('-')[1:])}",
        "sourceStorages": [storage]
    })

    return response.json()

# print(PrettyPrint(CreateStorage(LoadMetadata(sys.argv[1])[0])))
# print(PrettyPrint(CreatePreparation(sys.argv[1])))
print(PrettyPrint(StartScan(sys.argv[1], sys.argv[2])))
# print(PrettyPrint(ListStorages())) 
# print(PrettyPrint(ListPreparations())) 
# print(PrettyPrint(DeletePreparations(sys.argv[1], sys.argv[2]))) 
