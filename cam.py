import os
import time
import picamera

camera = None
filename = ""

#Start recording video into raw file
def start_record():
                print("Starting recording");
                reset_tmp()
                global filename
                filename = "vid/" + str(int(time.time()));
                
                global camera
                camera = picamera.PiCamera()
                camera.resolution = (1920, 1080)
                camera.start_recording(filename+".h264");

#Stop recording and convert h264 raw file to mp4 and remove raw file
def stop_record():
                print("Stopping recording");
                reset_tmp()
                global filename
                global camera
                camera.stop_recording()
                camera.close()
                os.system("MP4Box -fps 30 -add "+filename+".h264"+" "+filename+".mp4");
                os.system("rm "+filename+".h264");

def reset_tmp():
        os.system('echo "0" > tmp')

#Bootstrap      
while(True):
        with open('tmp', 'r') as myfile:
               status=myfile.read()
        if "1" in status:
                start_record();
        elif "2" in status:
                stop_record();
        time.sleep(0.3)
#Status codes: 0 Neutral, 1 Record, 2 Stop Recording.
