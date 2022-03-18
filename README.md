# [<img src="potato.png" width="48"/>](potato.png) PotatoDB
**PotatoDB** is nosql html database is written in Golang.
## Installation
>1. **Put** files from archive to some folder.
>2. **Create** folder for database, then in **.env** write its name e.g. `DIR_NAME=DB`
>3. **Create** cluster (new folder) in database folder.

## Routes
> /core<br>
> /filter<br>
> /cluster

## Request
### **/core**:
> GET:
```json
{
 "Cluster": "cluster_name",
  "Id": "object_id",
}
```
> POST:
```json
{
 "Cluster": "cluster_name",
  "Id": "optional_id",
  "Value": "object_value"
}
```
> DELETE:
```json
{
  "Cluster": "cluster_name",
  "Id": "object_id"
}
```
### **/filter**:
> GET:
```json
{
  "Cluster": "cluster_name",
  "Regex": "Regular_expression"
}
```
> DELETE:
```json
{
  "Cluster": "cluster_name",
  "Regex": "Regular_expression"
}
```
### **/cluster**:
> GET:
```json
{
  "Cluster": "cluster_name"
}
```
