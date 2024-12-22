# Atlas
A GitHub backup daemon to automate git commits and pushes for your config files and notes.

# Behaviour
## Worktree Updation
*Instantaneous:* Whenever any change is made within any repository managed by atlas, a _git add_ like functionaly is executed.
*Non-Modifiable:* User cannot modify this behaviour of atlas.
*Events:* It is done on the following events:
- WRITE
- REMOVE
- CREATE
