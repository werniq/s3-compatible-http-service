# cubbit_tech_task
# Firstly, I would like to thank you for such opportunity! I have really enjoyed while doing this project.

## I decided to do this task in unusual way. I have created it as Full-Stack web application.
 - To download it, run `git clone https://github.com/werniq/cubbit_tech_task/` </h2>
 - You can test it in two ways: either with graphic interaction, or just using tool like `Postman`

# Using Postman

<h3> Please, run <b>`go run cmd/api`</b> and feel free to test. </h3>

- `r.PUT("/buckets", app.CreateNewBucket)`: This route is used to create a new bucket. When a PUT request is made to "/buckets", the CreateNewBucket handler function is executed. It is responsible for creating a new bucket in the S3-compatible service. 
- `r.POST("/buckets/:bucket-name/files", app.UploadFiles)`: This route allows uploading files to a specific bucket. When a POST request is made to "/buckets/:bucket-name/files", the UploadFiles handler function is invoked. It handles the task of accepting the uploaded file and saving it to the specified bucket.
- `r.GET("/buckets/:bucket-name/files", app.ListFiles)`: This route is used to list the files in a specific bucket. A GET request to "/buckets/:bucket-name/files" triggers the ListFiles handler function. It retrieves and returns a list of files present in the specified bucket.
- `r.POST("/buckets/:bucket-name/files/:filename", app.DownloadFile)`: This route is responsible for downloading a specific file from a bucket. A POST request to "/buckets/:bucket-name/files/:filename" invokes the DownloadFile handler function. It retrieves and sends the requested file from the specified bucket.
- `r.POST("/check-bucket-name", app.CheckBucketName)`: This route is used to check the availability and validity of a bucket name. When a POST request is made to "/check-bucket-name", the CheckBucketName handler function is executed. It verifies whether the provided bucket name is valid and available for use.

<hr> 

# Using browser ( preferable )
<h3> Run `go run cmd/api`. Open another terminal, and run `go run cmd/client`</h3>
<h4> There is a navbar, so I hope you won't get lost <3 </h4>
 

 - `r.GET("/", app.HomeHandler)`: This route handles the GET request to the root URL ("/") and invokes the HomeHandler function. It is responsible for rendering the home page of the application.
 - `r.GET("/create-bucket", app.CreateBucketHandler)`: This route handles the GET request to "/create-bucket" and triggers the CreateBucketHandler function. It is responsible for rendering the page where users can create a new bucket.
 - `r.POST("/create-bucket", app.CreateBucketPostHandler)`: This route handles the POST request to "/create-bucket" and invokes the CreateBucketPostHandler function. It is responsible for processing the user's request to create a new bucket and generating the appropriate response.
 - `r.GET("/upload-files", app.UploadFilesHandler)`: This route handles the GET request to "/upload-files" and triggers the UploadFilesHandler function. It is responsible for rendering the page where users can upload files to a specific bucket.
 - `r.POST("/upload-files", app.UploadFilesPostHandler)`: This route handles the POST request to "/upload-files" and invokes the UploadFilesPostHandler function. It processes the user's request to upload files to a bucket and generates the corresponding response.
 - `r.GET("/list-files", app.ListFilesHandler`): This route handles the GET request to "/list-files" and triggers the ListFilesHandler function. It is responsible for rendering the page that displays the list of files in a specific bucket.
 - `r.POST("/list-files", app.ListFilesPostHandler)`: This route handles the POST request to "/list-files" and invokes the ListFilesPostHandler function. It processes the user's request to list the files in a bucket and generates the corresponding response.
 - `r.GET("/download-file", app.DownloadLastUploadedFileHandler)`: This route handles the GET request to "/download-file" and triggers the DownloadLastUploadedFileHandler function. It is responsible for rendering the page where users can download the last uploaded file from a specific bucket
 - `r.POST("/download-file", app.DownloadLastUploadedFilePostHandler)`: This route handles the POST request to "/download-file" and invokes the DownloadLastUploadedFilePostHandler function. It processes the user's request to download the last uploaded file from a bucket and generates the corresponding response.
 - `r.POST("/check-bucket-name", app.CheckBucketname)`: This route handles the POST request to "/check-bucket-name" and invokes the CheckBucketname function. It processes the user's request to check the validity and availability of a bucket name and generates the appropriate response.
