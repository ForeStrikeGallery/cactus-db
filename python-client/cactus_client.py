from urllib.parse import urlencode
from urllib.request import urlopen, Request
from cactus_exception import CactusException


def make_request(url, headers=None, data=None):
    headers = headers if headers is not None else {}

    print("post data", data)

    request = Request(url, headers=headers, data=data)
    try:
        with urlopen(request, timeout=10) as response:
            return response.read(), response
    except Exception as err:
        print(err)


class CactusClient:

    def __init__(self, connection_string=None):
        if connection_string:
            self.get_url = connection_string + "/get"
            self.put_url = connection_string + "/put"
        else:
            self.get_url = "http://localhost:3001/get"
            self.put_url = "http://localhost:3001/put"

    def get(self, key):
        try:
            with urlopen("{}/{}".format(self.get_url, key)) as response:
                return response.read()
        except Exception as e:
            raise CactusException("Get call failed", e)

    def put(self, key, value):
        new_record = {"key": key, "value": value}
        urlencoded_new_record = urlencode(new_record)
        post_data = urlencoded_new_record.encode("utf-8")

        try:
            body, response = make_request(self.put_url, data=post_data)
        except Exception as e:
            raise CactusException("Put call failed", e)

        return body
