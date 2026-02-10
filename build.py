import subprocess
from sys import argv

if "-h" in argv or "--help" in argv:
    print("""
  python build.py [options]

  * tui - build tui for configuring facecloud")
  * server - build facecloud server
""")

if (len(argv) == 1):
    subprocess.run(["go", "build"], cwd="server")
    subprocess.run(["go", "build"], cwd="tui")

if "server" in argv:
    subprocess.run(["go", "build"], cwd="server")

if "tui" in argv:
    subprocess.run(["go", "build"], cwd="tui")
