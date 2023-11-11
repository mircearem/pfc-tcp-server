# How the protocol works

* The system requires one program running on the 
IEC61131 devices that reads data from RAM:

> Files that signal change in IO data model:
> /tmp/var/codesys.kbus.data
> /tmp/var/codesys.kbus.lock
> Files that signal that an update has to be made
> to one or more variables in the model
> /tmp/var/codesys.kbus.update
> /tmp/var/codesys.knus.lock