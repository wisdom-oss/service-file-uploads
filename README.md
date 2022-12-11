# File Upload Service
![GitHub go.mod Go version (branch)](https://img.shields.io/github/go-mod/go-version/wisdom-oss/service-file-uploads/dev?style=for-the-badge&filename=src%2Fgo.mod)
<hr>
This microservice handles file uploads to the platform and updating/deleting those files.

## Functionality
The microservice accepts uploads as raw uploads, meaning the request body only contains the file contents which shall
be uploaded. All other information needs to be set via the headers. Further documentation on the headers
may be found in the [openapi.yaml](./openapi.yaml) file.

To successfully start the service you need to set values for the environment variables `PUBLIC_STORAGE_PATH` and 
`PRIVATE_STORAGE_PATH`. The storage paths function as root directory for all uploaded files and folders.

> **Warning**
> 
> When deploying the service in a docker environment make sure to mount the `PUBLIC_STORAGE_PATH` and the
> `PRIVATE_STORAGE_PATH` to either a docker volume or a path on your machine. If there is no mapping to either of the
> options

> **Note**
> 
> To be able to access the files under the `PUBLIC_STORAGE_PATH` needs to point to the docker volume which is also
> mounted to the reverse proxy in the project.

Files and folders are identified via a generated UUID. Folder UUIDs will correspond to the folder names in the storage 
paths. File names will be combined of the file uuid and the file name set during the upload, to allow machine discovery
as well as manually searching files.

> **Warning**
>   
>   Manually uploading files into the uploads folder is not supported by the service. Please use the documented api

<details>
    <summary>Example File Tree in <code>PUBLIC_STORAGE_PATH</code></summary>

    ├── uploads
    |   ├── 137b66ac-51e6-4427-976a-822571ee8c05_example.txt
    |   ├── 9a0ee649-1ea4-4b86-a81f-b0eca0330b25
    |   |   ├── 4d28b4b2-51f5-49ea-980f-669e12afca74_config.json
    |   ├── a1be3050-5013-469e-898c-da2f7afaedb3_report.pdf
    └── <other-static-files>
</details>
<details>
    <summary>Example File Tree in <code>PRIVATE_STORAGE_PATH</code></summary>


    ├── 137b66ac-51e6-4427-976a-822571ee8c05_example.txt
    ├── 9a0ee649-1ea4-4b86-a81f-b0eca0330b25
    |   ├── 4d28b4b2-51f5-49ea-980f-669e12afca74_config.json
    └── a1be3050-5013-469e-898c-da2f7afaedb3_report.pdf
</details>


