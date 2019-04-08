# GoAutoCopy

Just a basic program that automatically copies files from one directory to another during an file system update event; such as a file being created, a file being updated, etc..

Let's say you want to copy EVERY file from 'C:\src' to 'C:\dest', you would enter the line below to the 'config.xml' file:
```xml
<action source="C:\src" destination="C:\dest" filetypes="" />
```

Now let's say you want to only copy over files of type '.txt':
```xml
<action source="C:\src" destination="C:\dest" filetypes="txt" />
```

Now let's say you want to copy over files of type '.txt' or '.bat':
```xml
<action source="C:\src" destination="C:\dest" filetypes="txt,bat" />
```