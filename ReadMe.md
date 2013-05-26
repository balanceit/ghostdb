one file which is append only, call it 'storage'.

it's contents will look like:

GUIDLODDATA

where GUID is a unique id of the data
LOD is the length of the DATA this will always be of a defined length (32bytes??)
and DATA is the data

This file will be a binary file and is only appended to.

But there will be a daemon (compress) which will run everyonce in a while.
This daemon will remove data which has been superseeded

-------------------------------------------------------------------

another file 'index_storage' which will contain pointers to the GUID which represents the data

GUIDLOG

where GUID is the id of some data
LOG is the location of the 'active' piece of data

this file is append only.

when new data is written to 'storage' it is not considered the 'active' data until the new location is written to 'index_storage'.

-------------------------------------------------------------------

another file 'index_outdated' which will contain the locations of data guids which can be removed from storage.

this file will be appened to at the same time as 'index_storage'

-------------------------------------------------------------------

the process:
-get length of data to be written
-get location of end of 'storage' file
-append guid to 'storage'
-append length of data to 'storage'
-append data to 'storage'
-?lock? 'index_storage' and 'index_outdated'
-append old location to 'index_outdated'
-append guid and new location to 'index_storage'
-?unlock?


-------------------------------------------------------------------

the 'compress' daemon process
-