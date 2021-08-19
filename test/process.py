import psutil, time, ctypes

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
    if name == "Taskmgr.exe" and pid == 6108:
        print([name, pid, handle, GDI, rss, vms, cpu])

def getGDICount(PID):
    PH = ctypes.windll.kernel32.OpenProcess(0x400, 0, PID)
    GDIcount = ctypes.windll.user32.GetGuiResources(PH, 0)
    ctypes.windll.kernel32.CloseHandle(PH)
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

if __name__ == '__main__':
    getAllProcessInfo()
    