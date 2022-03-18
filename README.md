# [<img src="potato.png" width="48"/>](potato.png) PotatoDB
**PotatoDB** is nosql html database is written in Golang.
## Installation
> 1. **Put** files from archive to some folder.
> 2. **Create** folder for database, then in **.env** write its name `DIR_NAME=DB`
> 3. **Create** cluster (new folder) in database folder.

## Routes
> /core
> /filter
> /cluster

## Request
### **/core**:
> GET: `{"Cluster": "cluster_name", "Id": "object_id"}`<br>
> POST: `{"Cluster": "cluster_name", "Id": "optional_id", "Value": "object_value"}`<br>
> DELETE: `{"Cluster": "cluster_name", "Id": "object_id"}`
### **/filter**: GET, DELETE
### **/cluster**: GET
