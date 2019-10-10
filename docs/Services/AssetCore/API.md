The AssetCore APIs are accessible from _/api/assetcore/_

For example, `https://demo.imqs.co.za/api/assetcore/asset/list?type_ids=33,34`

Any functions that consume or return geometry use Well Known Text, and the coordinate system is EPSG 4326.

# GET /ping
Ping the service to see if it is alive. Returns a JSON object with the current unix time in seconds, such as `{"Timestamp": 1509373918}`

# GET /asset/list
List assets. Returns a JSON array with the assets that match the criteria specified in the request.

### Query Parameters
|            |                   |
| ---------- | ----------------- |
| type_ids   | Comma separated list of type IDs |
| parent_ids | Find assets that are children of particular objects (comma separated IDs) |

# POST /asset/add
Insert one or more assets into the database. The assets must be a JSON array of the form:
```json
[
	{
		TypeID: 123,
		Description: "East wing pillar 3",
		Geometry: "POINTZ(33.21,-18.45, 178.2)"
	}
]
```

# POST /type/add
Create a new asset type

... TBC ...