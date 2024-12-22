# Atlas
A GitHub backup daemon to automate git commits and pushes for your config files and notes.

## Behaviour
### Staging Changes
- **Instantaneous:** Whenever any change is made within any repository managed by atlas, a _git add_ like functionaly is executed.
- **Non-Modifiable:** User cannot modify this behaviour of atlas.
- **Events:** It is done on the following events:
- - WRITE
- - REMOVE
- - CREATE

### Commiting Changes
- **Variable:** Staged changes are commited every _30 minutes_, when a _shutdown_ event occurs, or _as defined by the user._
- **Modifiable:** The user can provide a time interval for atlas to commit changes by using the _atlas set --interval="timeInMinutes"_ command. [read more](www.tomavilius.in)
- **Events:** Performed during the following events:
- - TIMEOUT (user defined)
- - DEFAULT (30 min)
- - SHUTDOWN (systemd)
