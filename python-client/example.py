from cactus_exception import CactusException
from cactus_client import CactusClient


def run():
    cactus_client = CactusClient()

    try:
        cactus_client.put("hello", "there")
        value = cactus_client.get("hello")
        print("Retrieved from cactus %d" % value)
    except CactusException as e:
        print("Call to CactusDB failed", e)


if __name__ == "__main__":
    run()
