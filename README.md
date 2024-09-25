# CMF-SDK-GO 

This folder contains generated go sdk for CMF API, with each specific client library version nested under the respective directory.
See below on how to update a version of the CMF sdk.


## Updating cmf-sdk-go after changing an API spec. 

1. Download the OpenAPI generator tool
```shell
wget https://repo1.maven.org/maven2/org/openapitools/openapi-generator-cli/5.2.1/openapi-generator-cli-5.2.1.jar -O openapi-generator-cli.jar
```
2. Install mocker if you haven't already
```shell
 cd "path/to/client-go/v1"
```
3. Generate the updated SDK. Note that the path to the generator should be absolute.
```shell
./generate.sh </path/to>/openapi-generator-cli.jar ./v1  /path/to/cp-flink-fmc/src/main/resources/apispec-v1.yml
```
4. Run a git diff and ensure that the changes you see are only to the files you think should have been modified. If you see some minor changes across a bunch of files, likely something went wrong like specifying a wrong flag to openapi-generator. The generate.sh script may need to be updated.

5. Use git to commit and push the result. You should likely do this as a PR.

6. Make sure downstream dependencies like CLI get the updated SDK, and that the new version causes no breakages.