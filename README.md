# An Api endpoint to evaluate js and Build Go Templates

## Description
The solution enables you to evaluate JS code and also build out Go Templates via a REST API


## When to use evaluate-js-templates-api
- When you use an automation pipeline or tool and need to evaluate JS.
- When you use an automation pipeline or tool and need to build template.
---
## Requirements
* you will need to install go 1.8 [https://go.dev/doc/install](https://go.dev/doc/install) to run and install evaluate-js-templates-api

---

## Project folders
Below is a description evaluate-js-templates-api project folders and what they contain
|   Folder        | Description  | 
|-----------|---|
| docs      | Contains the swagger documents |
| documents | Contains cli and api markdown files  |
| modules   | Contains evaluate-js-templates-api modules and code  |
|           |   |

---

## Installation

```yaml
go install github.com/Mrpye/evaluate-js-templates-api
```

## Run as a container

<details>
<summary>1. Clone the repository</summary>

```
git clone https://github.com/Mrpye/evaluate-js-templates-api.git
```
</details>

<details>
<summary>2. Build the container as API endpoint</summary>

```
sudo docker build . -t  evaluate-js-templates-api:v1.0.0 -f Dockerfile
```

</details>

<details>
<summary>3. Run the container</summary>

```
sudo docker run -d -p 8080:8080 --name=evaluate-js-templates-api --restart always --env=WEB_IP=0.0.0.0  --env=WEB_PORT=8080 -t evaluate-js-templates-api:1.0.0
```

</details>

---

### Environment  variables
- WEB_IP (set the listening ip address 0.0.0.0 allow from everywhere)
- WEB_PORT (set the port to listen on)


---
## How to use evaluate-js-templates-api CLI
Check out the CLI documentation [here](./documents/evaluate-js-templates-api.md)

---

# Using the API

## Run web server
```bash
    evaluate-js-templates-api
```
---

## Examples

<details>
<summary>Add two numbers</summary>

``` bash
curl --location 'localhost:8080/evaluate' \
--header 'Content-Type: application/json' \
--data '{
    "code":"10+125",
    "model":null
}'

```
Result:
```
135
```

</details>

<details>
<summary>Manipulate data model and return the model</summary>

``` bash
curl --location 'localhost:8080/evaluate' \
--header 'Content-Type: application/json' \
--data '{
    "code":"model.test=\"mydata\";model.test2=model.test2+10;model",
    "model":{
        "test":"test1",
        "test2":5
    }
}'

```
Result:
```
{
    "test": "mydata",
    "test2": 15
}
```

</details>

<details>
<summary>Build a template simple</summary>

``` bash
curl --location 'localhost:8080/template' \
--header 'Content-Type: application/json' \
--data '{
    "template":"Message: {{.test1}} Message: {{.test2}}",
    "model":{
        "test1":"This is Test1",
        "test2":5
    }
}'

```
Result:
```
Message: This is Test1
 A number: 5
 list of numbers:  1 2 3 4 5
 list of names and age:  TONY:22  FRED:45 
```

</details>

<details>
<summary>Build a template complex</summary>

``` bash
curl --location 'localhost:8080/template' \
--header 'Content-Type: application/json' \
--data '{
    "template":"Message: {{.test1}}\n A number: {{.test2}}\n list of numbers: {{range $val := .loop1}} {{$val}}{{end}}\n list of names and age: {{range $val := .loop2}} {{uc $val.name}}:{{$val.age}} {{end}}",
    "model":{
        "test1":"This is Test1",
        "test2":5,
        "loop1":[1,2,3,4,5],
        "loop2":[
            {
                "name":"tony",
                "age":22
            },
            {
                "name":"fred",
                "age":45
            }
        ]
    }
}'

```
Result:
```
Message: This is Test1
 A number: 5
 list of numbers:  1 2 3 4 5
 list of names and age:  TONY:22  FRED:45 
```

</details>


<details>
<summary>Test is alive</summary>

```bash
curl --location --request GET 'localhost:8080/'
```

Result:
```
OK
```

</details>

---
## evaluate-js-templates-api Helm chart
This guide will show you how to build the helm chart package for evaluate-js-templates-api, you will need to have helm installed to build the package.

<details>
<summary>1. Build the helm chart package for evaluate-js-templates-api</summary>

```bash
# change into the chart directory
cd charts
# Package the evaluate-js-templates-api chart
helm package evaluate-js-templates-api

```

the helm chart package will be saved under the charts folder evaluate-js-templates-api-0.1.0.tgz

</details>

---

## Update the swagger document
The code below shows you how to update the swagger API documents.

If you need more helm on using these tools please refer to the links below
- gin-swagge [https://github.com/swaggo/gin-swagger](https://github.com/swaggo/gin-swagger)
- swag [https://github.com/swaggo/swag](https://github.com/swaggo/swag)

<details>
<summary>1. Install swag</summary>

```bash
#Install swag
go install github.com/swaggo/swag/cmd/swag
```
</details>

<details>
<summary>2. Update APi document</summary>

```bash
#update the API document
swag init
```
</details>
<details>
<summary>3. Update the api.md</summary>

```bash
swagger generate markdown -f .\docs\swagger.json --output .\documents\api.md 
```
</details>

---

# To Do
- Nothing at the moment

---

# 3rd party Libraries
https://github.com/dop251/goja

---
# license
cimplex is Apache 2.0 licensed.