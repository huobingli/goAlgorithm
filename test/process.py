import psutil, time
from ctypes import *

def getProcessInfo(p):
    try:
        cpu = int(p.cpu_percent(interval=0))
        rss = 0     # 工作集内存
        vms = p.memory_info().rss
        name = p.name()
        handle = p.num_handles()
        pid = p.pid
        GDI = getGDICount(pid)
    except psutil.NoSuchProcess:
        name = "Closed_Process"
        pid = 0
        rss = 0
        vms = 0
        handle = 0
        cpu = 0
        GDI = 0
    if name == "Taskmgr.exe" and pid == 20048:
        print([name, pid, handle, GDI, rss, vms, cpu])

def getGDICount(PID):
    PH = windll.kernel32.OpenProcess(0x400, 0, PID)
    GDIcount = windll.user32.GetGuiResources(PH, 0)
    windll.kernel32.CloseHandle(PH)
    return GDIcount

def getAllProcessInfo():
    instances = []
    all_process = list(psutil.process_iter())
    # for proc in all_process:
    #     proc.cpu_percent()
        # proc.get_cpu_percent(interval = 0)
    while 1:
        time.sleep(1)
        for proc in all_process:
            instances.append(getProcessInfo(proc))

def getAllProcessID(): 
    for proc in psutil.process_iter():
        try:
            pinfo = proc.as_dict(attrs=['pid', 'name'])
        except psutil.NoSuchProcess:
            pass
        else:
            print(pinfo)

def getPID(processName):
    for proc in psutil.process_iter():
        try:
            if processName.lower() in proc.name().lower():
                return proc.pid
        except (psutil.NoSuchProcess, psutil.AccessDenied, psutil.ZombieProcess):
            pass
    return None;

def getGDIcount(PID):
    PH = windll.kernel32.OpenProcess(0x400, 0, PID)
    GDIcount = windll.user32.GetGuiResources(PH, 0)
    windll.kernel32.CloseHandle(PH)
    return GDIcount

PID = getPID('Taskmgr.exe')

def testGDI():
    while True:
        GDIcount = getGDIcount(PID)
        print(f"{time.ctime()}, {GDIcount}")
        time.sleep(1)

if __name__ == '__main__':
    testGDI()
    