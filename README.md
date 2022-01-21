# File Manager
## Basic
This application is for managing files. Create files and folders, rename it, delete, copy and paste, sorting. You can upload files from your device or drag and drop them. Also you have search system based on tags.

## Extensions
* Video player
* Image preview
* Download manager
* Disk usage
* Paste file from buffer

## Tag system
Each file can contain custom tags, and you can search files with these tags. But for correct working you must set file with hash name. Then you can put tag, for example "good: 1". After that you go to search and write "good > 0.5" and see all files when such tag is greater than 0.5.

## Download manager

You can send data from external resources to this application. It uses **Signal Api**.

*download_photo* - Accept image url and system path

*download_blob_photo* - Accept image base64 and system path

Send API requests to http://%DESKTOP_APP%/api/main/applicationData with data

```
appId: "maldan/fileman",
data: JSON.stringify({
  type: "download_photo",
  path: `/home/hello/`,
  url: PHOTO_DATA,
}),
```

But only these sites (xvideos,simpy-hentai,vk,rule34,discordapp) support such system. In application, you can just put url to download manager extension.