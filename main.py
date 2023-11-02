# import json
import time
import random

def get_index(request):
    angka = random.randint(0, 180)
    time.sleep(angka)

    x = '{ "message":"Success", "code": 0, "status":200, "latency": '+str(angka)+'}'
    
    headers = {
    	'Access-Control-Allow-Origin': '*',
    	'Access-Control-Allow-Methods': '*',
    	'Access-Control-Allow-Headers': '*',
    	'Access-Control-Max-Age': '3600'
    }

    # return (json.loads(x), 200)
    return (x, 200, headers)