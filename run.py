import subprocess
from sys import argv, platform
from os import getcwd, path
import processes as ps

if "-h" in argv or "--help" in argv:
    print("""
  python run.py [options]

  * tui - run tui for configuring facecloud")
  * server - run facecloud server
""")

try:
    if (len(argv) == 1):
        ps.run_sync_multi([
            ["go", "get"],
            ["go", "run", "main.go"]
        ], "tui")
        ps.run_sync(["go", "get"], "server")
        if platform == "win32": #get rid of windows network access pop up
            ps.run_sync_multi([
                ["go", "build"],
                "server/fcserver.exe"
            ], "server")
        else:
            ps.run_sync(["go", "run", "main.go"], "server")

    if "server" in argv:
        ps.run_sync(["go", "get"], "server")
        if platform == "win32": #get rid of windows network access pop up
            ps.run_sync_multi([
                ["go", "build"],
                "server/fcserver.exe"
            ], "server")
        else:
            ps.run_sync(["go", "run", "main.go"], "server")

    if "tui" in argv:
        ps.run_sync_multi([
            ["go", "get"],
            ["go", "run", "main.go"]
        ], "tui")
except KeyboardInterrupt:
    pass
