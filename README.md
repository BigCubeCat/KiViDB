# [<img src="potato.png" width="48"/>](potato.png) PotatoDB
🥔**PotatoDB** is nosql html database is written in Golang.
## Installation
>1. **Put** files from archive to some folder.
>2. **Create** folder for database, then in **.env** write its name e.g. 
>```env
>DIR_NAME=DB
>```
>3. **Create** cluster (new folder) in database folder.

## Routes
> /core<br>
> /filter<br>
> /cluster

## Request
### **/core**:
> GET: Returning object by Id.
>```json
>{
> "Cluster": "cluster_name",
>  "Id": "object_id",
>}
>```
> POST: Creating a new cluster. 
>```json
>{
> "Cluster": "cluster_name",
>  "Id": "optional_id",
>  "Value": "object_value"
>}
>```
> DELETE: Deleting cluster by id.
>```json
>{
>  "Cluster": "cluster_name",
>  "Id": "object_id"
>}
>```
### **/filter**:
> GET: Returning all objects, satisfying the RE.
>```json
>{
>  "Cluster": "cluster_name",
>  "Regex": "regular_expression"
>}
>```
> DELETE: Deleting all objects, satisfying the RE.
>```json
>{
>  "Cluster": "cluster_name",
>  "Regex": "regular_expression"
>}
>```
### **/cluster**:
> GET: Returning all objects in this cluster.
>```json
>{
>  "Cluster": "cluster_name"
>}
>```
> POST: Creating a new cluster.
>```json
>{
>  "Cluster": "cluster_name"
>}
>```
> DELETE: Deleting a cluster.
>```json
>{
>  "Cluster": "cluster_name"
>}
>```
