package main

import (
"fmt"
"github.com/stianeikeland/go-rpio"
"time"
"strconv"
"os"
"os/exec"
"log"
"io/ioutil"
"os/signal"
)

var img_btn, vid_switch, status_led, vid_led rpio.Pin
var btn_status rpio.State
var recording bool

func image_capture(){
        if (recording){
             fmt.Printf("Stopping recording") 
             stop_recording()
        }else{
                filename := "img/" + strconv.Itoa(int(time.Now().Unix())) + ".jpg";
                status_led.High()
                out, err := exec.Command("raspistill","-q","100","-th","400:400:100","--nopreview","-t","1","-o",filename).Output()
                status_led.Low()
                if err != nil{
                        log.Panic()
                }
                fmt.Printf("%s",out)        
        }
}

func initiate(){
        img_btn = rpio.Pin(4)
        status_led = rpio.Pin(17)
        vid_led = rpio.Pin(27)
        vid_switch = rpio.Pin(22)
        img_btn.PullDown()
        vid_switch.PullDown()
        status_led.Output()
        vid_led.Output()
}

func vidled(){
                if vid_switch.Read() == rpio.High{
                        vid_led.High()
                }else{
                        vid_led.Low()
                }
}

func cleanup(){
        status_led.Low()
        vid_led.Low()
        rpio.Close()
}

func start_recording(){
       ioutil.WriteFile("tmp", []byte("1"), 0644)
       recording = true    
}

func stop_recording(){
        ioutil.WriteFile("tmp", []byte("2"), 0644)
        recording = false
}

func main(){
        //Initiate our GPIO pins
        rpio.Open()
        defer rpio.Close()
        initiate()

        //Detect keyboard interrupt
        keyboardinterrupt := make(chan os.Signal, 1)
        signal.Notify(keyboardinterrupt, os.Interrupt)
        go func(){
            for sig := range keyboardinterrupt {
                fmt.Print(sig)
                cleanup()
                os.Exit(1)
            }
        }()

	//Bootstrap
        for{
                vidled()
                if img_btn.Read() != btn_status && img_btn.Read() == rpio.Low{
                        if vid_switch.Read() == rpio.Low{
                                fmt.Println("Capturing image")
                                image_capture()
                        }else{
			        if recording{
				        fmt.Println("Ending recording")
				        stop_recording()
				}else{
				        fmt.Println("Starting recording")
				        start_recording()
				}
                        }
                }
                btn_status = img_btn.Read()
                time.Sleep(80*time.Millisecond)
        }
}