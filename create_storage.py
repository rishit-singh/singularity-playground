import os
import re
import sys
import json
import requests

BaseURL = os.getenv("BASE_URL")
Credentials = os.getenv("CRED")
FtpCredentials = os.getenv("FTP_CRED")
Username, Password = Credentials.split(':')
FtpUsername, FtpPassword = FtpCredentials.split(':')

def PrettyPrint(data) -> str:
    return json.dumps(data, indent=2)

def LoadMetadata(file: str) -> list[dict]:
    metadata = None

    with open(file, 'r') as fp:
        metadata = json.load(fp)
    
    return metadata

def CreateStorage(metadata: dict):
    url: str = metadata["url"]

    host: str = re.search(r'(?<=://)[^/]+', url).group()
    split: str = url[url.find(host) + len(host):].split('/')[:-2]

    path: list = "/".join(split)

    nameSplit = split[6].split('_')

    print(PrettyPrint({
        "name": f"ftp-{split[4]}-{'-'.join(nameSplit)}",
        "path": path,
        "config": {
            "user": "anonymous",
            "pass": "pR2ciGTcNxC_KyfQOsUjOAs",
            "concurrency": "0",
            "idleTimeout": "300",
            # "metadata": json.dumps(metadata)
        }
    }
    ))

    response = requests.post(f"{BaseURL}/storage/ftp", auth=(Username, Password), json={
        "name": f"ftp-{split[4]}-{'-'.join(nameSplit)}",
        "path": path,
        "config": {
            "user": FtpUsername,
            "pass": FtpPassword,
            "host": host,
            "concurrency": "0",
            "idleTimeout": "300",
            "metadata": json.dumps(metadata)
        }
    })

    return response.json()

print(PrettyPrint(CreateStorage(LoadMetadata(sys.argv[1])[0])))