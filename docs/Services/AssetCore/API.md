The AssetCore APIs are accessible from _/assetcore/api/_

For example, `https://demo.imqs.co.za/assetcore/api/asset/list?type_ids=33,34`

Any functions that consume or return geometry use Well Known Text, and the coordinate system is EPSG 4326.

When dealing with hierarchies, one must always specify the particular hierarchy that you're interested in. The only hierarchy presently available is hierarchy `1`. When specifying a parent of an asset, you refer to hierarchy `1` as `Parent1`. Likewise, if an API takes a `hierarchy_id` as a parameter, then that value must be `1`.

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
		Geometry: "POINTZ(33.21,-18.45, 178.2)",
		Parent1: 9876
	}
]
```

# GET /type/list
List asset types

... TBC ...
# POST /type/add
Create a new asset type

... TBC ...

# GET /hierarchy/:hierarchy_id/get_children/:asset_id
Fetch all children of the given asset, in the given hierarchy

### Query Parameters
|            |                   |
| ---------- | ----------------- |
| depth      | Maximum depth of hierarchy traversal. Default = 1 |

# GET /hierarchy/:hierarchy_id/get_parent/:asset_id
Fetch the parent of the given asset, in the given hierarchy

# GET /hierarchy/:hierarchy_id/get_tree
Fetch the entire hierarchy tree, as a recursive array of arrays, eventually ending in the IDs

Example response:
```json
[
	1: [
		6: [10,13]
	],
	2: [
		9: [23,40]
	]
]
```
