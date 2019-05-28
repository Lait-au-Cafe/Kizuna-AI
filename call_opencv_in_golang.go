package main

// #cgo pkg-config: opencv
// #include <opencv/cv.h>
// #include <opencv/highgui.h>
import "C"
import "unsafe"

func main(){
	
	capture := C.cvCreateCameraCapture(C.int(0))
	defer C.cvReleaseCapture(&capture)

	C.cvSetCaptureProperty(capture, C.CV_CAP_PROP_FRAME_WIDTH, 1240)
	C.cvSetCaptureProperty(capture, C.CV_CAP_PROP_FRAME_HEIGHT, 720)

	windowName := "test"
	C.cvNamedWindow(C.CString(windowName), C.CV_WINDOW_AUTOSIZE)
	defer C.cvDestroyWindow(C.CString(windowName))

	for C.cvWaitKey(30) != 113 {
		frame := C.cvQueryFrame(capture)
		C.cvShowImage(C.CString(windowName), unsafe.Pointer(frame))
	}
}
