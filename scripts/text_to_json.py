import sys
import json 

def ParseTextFile(file: str):
    parseMap = {}

    with open(file, 'r') as fp:
        rawString = fp.read()
        split = [line.split('\t') for line in rawString.split('\n')]

        parsed = [] 

        for x in range(len(split)):
            if (len(split[x]) > 2):
                parsed.append(split[x])
        
        parseMap["keys"] = parsed.pop(0)
        parseMap["data"] = parsed
        
    return parseMap

def ExportToJson(parseMap: dict, file: str):
    with open(file, 'w') as fp:
        stack = []

        keys = parseMap["keys"]    

        for data in parseMap["data"]:
            jsonMap = {}

            for x in range(len(data)):
                jsonMap[keys[x]] = data[x]

            stack.append(jsonMap)
        
        json.dump(stack,fp, indent=2)

if (len(sys.argv) < 3):
    print("Usage: python3 text_to_json.py <input file> <output file>")
else:
    ExportToJson(ParseTextFile(sys.argv[1]), sys.argv[2])