from cactus_exception import CactusException
from cactus_client import CactusClient
import time 

cactus = CactusClient()    

def benchmark_reads(count): 
    print("Write benchmark test runnning...")

    start = time.time()
    for i in range(count):
        start_i = time.time()
        cactus.get(str(i))
        print("Write time for {}: ", i, time.time() - start_i)

    print("Time taken for {} GET requests: {}".format(i, time.time() - start))

def benchmark_writes(count):
    print("Write benchmark test runnning...")

    start = time.time()
    for i in range(count):
        start_i = time.time()
        cactus.put(str(i), str(i))
        print("Write time for {}: ", i, time.time() - start_i)

    print("Time taken for {} PUT requests: {}".format(i, time.time() - start))

def run():
    try:
        cactus.put("yo", "hey!")
        print("Successfully put to CactusDB")

        value = cactus.get("yo")
        print("Retrieved from cactus: %s" % value)
        print(cactus_client.get("theree"))

    except CactusException as e:
        print("Call to CactusDB failed: ", e)

if __name__ == "__main__":
    run()
