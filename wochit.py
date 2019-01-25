import os
import sys
import hashlib
import glob


def main():
    pathToScan = ""
    if len(sys.argv) > 1:
        pathToScan = sys.argv[1]
    else:
        pathToScan = os.getcwd()

    hashes = {}
    for dirName, subdirs, fileList in os.walk(pathToScan):    
        for filename in fileList:
            path = os.path.join(dirName, filename)
            theHash = hashFile(path)
           
            if theHash in hashes:
                hashes[theHash].append(path)
            else:
                hashes[theHash] = [path]
    
    for h, filenames in hashes.items():
        if len(filenames) > 1:
            print(f'{len(filenames)} identical files found with hash {h}')
            for f in filenames:
                print(f)


def hashFile(filename):
    BLOCKSIZE = 65536
    hasher = hashlib.sha1()
    with open(filename, 'rb') as afile:
        buf = afile.read(BLOCKSIZE)
        while len(buf) > 0:
            hasher.update(buf)
            buf = afile.read(BLOCKSIZE)
    return hasher.hexdigest()


if __name__ == "__main__":
    main()

