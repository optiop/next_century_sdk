package enums

type ReadDownloadState string

/*
|Value 				|Description 							|
|-------------------|-------------------|
|UNINITIALIZED 		|Reads have not begun processing		|
|RUNNING 			|Reads are currently being processed	|
|UPLOADING			|Reads have finished processing and are being uploaded to a secure location	|
|COMPLETE			|Reads have successfully finished processing and are ready to be downloaded	|
*/
const (
	ReadDownloadStateUninitialized ReadDownloadState = "UNINITIALIZED"
	ReadDownloadStateRunning       ReadDownloadState = "RUNNING"
	ReadDownloadStateUploading     ReadDownloadState = "UPLOADING"
	ReadDownloadStateComplete      ReadDownloadState = "COMPLETE"
)
