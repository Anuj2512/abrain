### This document contains specification for knowledge store and its implementation.

### knowledge schema transformer
1. A config file will contain schema for knowledge store.
2. Above function will use config file to read schema specification and converted input data to desired schema.
3. Also it will map ouput to desired form

### knowledge sync
1. This function will be used to sync knowledge store with the current data.
2. This will take data from schema transformer and send it to knowledge store for storing the data.

### knowledge apis.
1. Based on config file it will transfer generic apis to specialized schema based apis to fetch knowledge from knowledge store.
2. The apis should be made in such a way that it binds with schema present in config file.
