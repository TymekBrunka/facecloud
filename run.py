import subprocess
from sys import argv

if "-h" in argv or "--help" in argv:
    print("""
  python run.py [options]

  * tui - run tui for configuring facecloud")
  * server - run facecloud server
""")

if (len(argv) == 1):
    subprocess.run(["go", "run", "main.go"], cwd="server")
    subprocess.run(["go", "run", "main.go"], cwd="tui")

if "server" in argv:
    subprocess.run(["go", "run", "main.go"], cwd="server")

if "tui" in argv:
    subprocess.run(["go", "run", "main.go"], cwd="tui")
