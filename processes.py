import subprocess

def run_async(cmd, cwd):
    subprocess.Popen(cmd, cwd=cwd)

def run_sync(cmd, cwd):
    subprocess.Popen(cmd, cwd=cwd).wait()

def run_async_multi(cmds: list, cwd):
    ps = []
    for cmd in cmds:
        ps.append(subprocess.Popen(cmd, cwd=cwd))
    for p in ps:
        errc = p.wait()
        if errc != 0:
            raise subprocess.CalledProcessError(errc, p)

def run_sync_multi(cmds: list, cwd):
    ps = []
    for cmd in cmds:
        p = subprocess.Popen(cmd, cwd=cwd)
        errc = p.wait()
        if errc != 0:
            raise subprocess.CalledProcessError(errc, cmd)
