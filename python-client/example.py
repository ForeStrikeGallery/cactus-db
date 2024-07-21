from cactus_exception import CactusException
from cactus_client import CactusClient
import time 

cactus = CactusClient()    

def benchmark_reads(count): 
    print("Read benchmark test runnning...")

    start = time.time()
    for i in range(count):
        start_i = time.time()
        try:
            cactus.get(str(i))
            print("Read time for {}: {} ms".format(i, (time.time() - start_i) * 1000)) 
        except Exception as e: 
            print("Timed out; continuing...", e) 

    print("Time taken for {} GET requests: {}".format(i, time.time() - start))

def benchmark_writes(count):
    print("Write benchmark test runnning...")

    start = time.time()
    for i in range(count):
        start_i = time.time()
        try: 
            cactus.put(str(i), str(i))
            print("Write time for {}: {} ms".format(i, (time.time() - start_i) * 1000))
        except Exception as e: 
            print("Timed out; continuing...", e) 

    print("Time taken for {} PUT requests: {}".format(i, time.time() - start))

def run():
    global cactus
    try:
        cactus.put("yo", "hey!")
        print("Successfully put to CactusDB")

        value = cactus.get("yo")
        print("Retrieved from cactus: %s" % value)
        print(cactus.get("theree"))

    except CactusException as e:
        print("Call to CactusDB failed: ", e)

if __name__ == "__main__":
    run()
