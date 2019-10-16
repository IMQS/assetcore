The AssetCore service wraps a very simple database with essentially one important table of "Assets", which all other services can link to.

## Why does this exist?
IMQS builds products for many different departments of an organization, and this leads to a natural fragmentation.

The fragmentation is caused by two major forces.

Firstly, the departments of the organization are different; they each have their own way of dealing with their information, they have their own processes, etc. The Water department of a city might have a very different database record for a road, versus the Roads department, who likely keeps more detailed information about the individual pieces that make up that road. The Sewer department might have yet another representation for that same road.

The second major force that drives this fragmentation is the different products at IMQS, each of which is built by a different team. One could try to force all teams to build their products in the same database, but we are believers in the notion that it's better to split databases up, and expose them via APIs - a notion made famous by Steve Yegge's post on Jeff Bezos at Amazon - see https://gist.github.com/chitchcock/1281611. Scroll down to the list of 7 things, in that blog post, for it's essence.

Unifying all of this information in one go has so far been beyond our reach. But there is a genuine need to unify these things to some degree, and this service aims to be that unifying central ground. Of course, nothing stops our `Project Control System` from talking directly to our `Maintenance Management` system's API, so that is certainly a form of unification, but it is labour intensive, and very loosely defined. In addition, simply mandating that teams "must talk to each other's services", is no guarantee that they actually will do so.

In October 2019, we felt that we would benefit from the creation of a single central table, which we call the Asset Core. This table has very few columns, namely `ID`, `Description`, `Geometry`, `Type`. Things that go into this table are things that are common to the different IMQS products, such as "building", "water pipe", "air conditioner external unit", "stop sign", or "land parcel".

The database schema is viewable here: https://dbdiagram.io/d/5d9f86ffff5115114db5209a

In addition to this central table, the Asset Core maintains a hierarchy which binds every single entry in the table, to a parent, from the same table. For example, the parent of a particular air conditioner external unit, might be the building that it's bolted onto. The parent of the building might be a hospital. The parent of the hospital could be "healthcare buildings". At some point you reach the top - these objects have no parent.

This service makes allowance for more than one such hierarchy. Different domains or departments might want to group their things differently, so we've made allowance for multiple hierarchies. At present, there is only a single hierarchy coded into the database, but the architecture is such that adding more hierarchies is easy, and will not break any code that is built around the first hierarchy.
