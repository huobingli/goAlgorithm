
def gatherIt(self,whoIt,whatIt,type,wiggle,process_info2):
 #this is the data gathering function thing
 data=0.0
 data1="wobble"
 if type=="counter":
 #gather data according to the attibutes
 try:
 data = win32pdhutil.FindPerformanceAttributesByName(whoIt, counter=whatIt)
 except:
 #a problem occoured with process not being there not being there....
 data1="N/A"

 elif type=="cpu":
 try:
 process_info={}#used in the gather CPU bassed on service
 for x in range(2):
 for procP in wiggle.Win32_PerfRawData_PerfProc_Process(name=whoIt):
 n1 = int(procP.PercentProcessorTime)
 d1 = int(procP.Timestamp_Sys100NS)
 #need to get the process id to change per cpu look...
 n0, d0 = process_info.get (whoIt, (0, 0)) 
 try:
 percent_processor_time = (float (n1 - n0)/float (d1 - d0)) *100.0
 #print whoIt, percent_processor_time
 except ZeroDivisionError:
 percent_processor_time = 0.0
 # pass back the n0 and d0
 process_info[whoIt] = (n1, d1)
 #end for loop (this should take into account multiple cpu's)
 # end for range to allow for a current cpu time rather that cpu percent over sampleint
 if percent_processor_time==0.0:
 data=0.0
 else:
 data=percent_processor_time
 except:
 data1="N/A"

 else:
 #we have done something wrong so data =0
 data1="N/A"
 #endif
 if data =="[]":
 data=0.0
 data1="N/A"
 if data =="" :
 data=0.0
 data1="N/A"
 if data =="":
 data=0.0
 data1="N/A"
 if data1!="wobble" and data==0.0:
 #we have not got the result we were expecting so add a n/a
 data=data1
 return data
