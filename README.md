# cubbit_tech_task
# Firstly, I would like to thank you for such opportunity! I have really enjoyed while doing this project.
<hr>

## I decided to do this task in unusual way. I have created it as Full-Stack web application.
 - To download it, run `git clone https://github.com/werniq/cubbit_tech_task/` </h2>
 - You can test it in two ways: either with browser, or just using tool like `Postman`

# Using Postman

<h3> Please, run <b>`go run cmd/api`</b> and feel free to test. </h3>
<h3> Path: <b>http://localhost:4001/</b> </h3>

- `r.PUT("/buckets", app.CreateNewBucket)`: This route is used to create a new bucket. When a PUT request is made to "/buckets", the CreateNewBucket handler function is executed. It is responsible for creating a new bucket in the S3-compatible service. 
- `r.POST("/buckets/:bucket-name/files", app.UploadFiles)`: This route allows uploading files to a specific bucket. When a POST request is made to "/buckets/:bucket-name/files", the UploadFiles handler function is invoked. It handles the task of accepting the uploaded file and saving it to the specified bucket.
- `r.GET("/buckets/:bucket-name/files", app.ListFiles)`: This route is used to list the files in a specific bucket. A GET request to "/buckets/:bucket-name/files" triggers the ListFiles handler function. It retrieves and returns a list of files present in the specified bucket.
- `r.POST("/buckets/:bucket-name/files/:filename", app.DownloadFile)`: This route is responsible for downloading a specific file from a bucket. A POST request to "/buckets/:bucket-name/files/:filename" invokes the DownloadFile handler function. It retrieves and sends the requested file from the specified bucket.
- `r.POST("/check-bucket-name", app.CheckBucketName)`: This route is used to check the availability and validity of a bucket name. When a POST request is made to "/check-bucket-name", the CheckBucketName handler function is executed. It verifies whether the provided bucket name is valid and available for use.

<hr> 

# Using browser ( preferable )
<h3> Run `go run cmd/api`. Open another terminal, and run `go run cmd/client`</h3>
<h4> There is a navbar, so I hope you won't get lost <3 </h4>
 <h3> Path: <b>`http://localhost:4000/`</b> </h3>
 

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

 <hr>
 
 # Running
 <h1>To run it using Docker: </h1>
 
   - `docker build -t image-name -f .\client.Dockerfile . ` 
   - `docker run -p 4001:4001 image-name `  
   - After that, you can use Postman (or any another tool for making requests), and play with it
 
 # Using browser
 <h1> To run and play it with browser </h1>
  
- I kindly ask you to use 
  - 1. `cd .\cubbit_tech_task\` 
  - 2. `go run cmd/api` 
  - 3. `go run cmd/client`
 
 
## Using docker it is hard to configure them on one host.. Firslty run client, change urls in code, run using those port.. A lot of things. </li>
 
## Using Postman:  
![photo_2023-05-23_10-22-58](https://github.com/werniq/cubbit_tech_task/assets/73220736/63a22163-16be-424a-909d-cb0f9235f8b3)
![photo_2023-05-23_10-23-04](https://github.com/werniq/cubbit_tech_task/assets/73220736/5e8bb3d0-308f-459d-ba49-d02fd99667c7)
![photo_2023-05-23_10-23-07](https://github.com/werniq/cubbit_tech_task/assets/73220736/ebe34e1b-a44d-46d4-9549-4035ab542364)
 
 <hr>

## Using browser
 ![photo_2023-05-23_10-22-41](https://github.com/werniq/cubbit_tech_task/assets/73220736/1ef1806f-54d7-4c9f-93b0-e4020d8daf9e)
![photo_2023-05-23_10-22-49](https://github.com/werniq/cubbit_tech_task/assets/73220736/fd1e1798-42e7-44ca-a44c-3c3f01b4bf8f)
![photo_2023-05-23_10-22-52](https://github.com/werniq/cubbit_tech_task/assets/73220736/cd13f11e-354d-435a-a1a9-9bc0f20ba283)
![photo_2023-05-23_10-22-55](https://github.com/werniq/cubbit_tech_task/assets/73220736/275d0d58-6b02-4e44-a525-2e7662405911)
![photo_2023-05-23_10-23-10](https://github.com/werniq/cubbit_tech_task/assets/73220736/20ae569a-e8f8-4216-b28a-099781a0f636)
