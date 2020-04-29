"# winlog" 


Package to subscribe to Windows Event Log.

Create test event on Windows:
```eventcreate /L Application /T SUCCESS /ID 777 /D TestEvent777```

Example fromat for xpath query:
```xquery = "Event/System[EventID=4]"```