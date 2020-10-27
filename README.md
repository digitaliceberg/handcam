# Handcam V1

Raspberry PI powered digital camera, broadcasts WIFI hotspot for users to view images/video captured.
![Alt text](https://i.imgur.com/CxQmDMc.jpg?raw=true "Front")
![Alt text](https://i.imgur.com/Qg5Jdca.jpg?raw=true "Front")

### Building steps
These are the steps I went through building the handcam.
https://imgur.com/a/uwzoy

### Program files

File specifications:

* `cam.go`: Main program, listens for user input via a button and switch on the camera design and sends status via LEDs.
* `cam.py`: Listener for video recording written in python, uses a tmp file to communicate with main to record and encode video since it can't be done with golang.
* `tmp`: File that stores one byte (status code) to trigger events in the python script, more info in comments.

### Usage
 Build the cam.go using go build and run the python script cam.py, preferably add them both to bootup.

### To Do
 Create webservice to download and view videos and images captured by the handcam
