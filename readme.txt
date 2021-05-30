             .--~~,__
:-....,-------`~~'._.'   paste.hauru.club
 `-,,,  ,_      ;'~U'
  _,-' ,'`-__; '--.
 (_/'~~      ''''(;

Welcome to paste.hauru.club, simple service for hosting
text files.

* _WARNING_: paste.hauru.club is still work in progress. Do not
expect your files to live longer than ~1 hour. *

Below you'll find instructions how to use paste.hauru.club for
posting your files with help of cURL command line program.

If you want to post your file to paste.hauru.club you'll
have to execute below command.

  $ curl -XPOST https://paste.hauru.club --data--binary @-

Now you can start typing. If you want to stop typing and
send file, enter: ^D (end of file). Now you'll receive
short instructions how to access your file and how to remove
it from service permanently.

You can send file to paste.hauru.club in different ways. For
example:

- you can redirect file to cURL,

  $ curl -XPOST https://paste.hauru.club --data-binary @- < file.txt

- send multiple files,

  $ cat file1.txt file2.txt | curl -XPOST https://paste.hauru.club --data-binary @-

- send file by pointing to it by its by name.

  $ curl -XPOST https://paste.hauru.club --data-binary @file.txt

*Remember!* There is no way to retrieve access key again for
your file, so make sure that you copied it somewhere and keep
it safe, if you're willing to access your file later. Same rules
work for delete key. If you lost it, you'll never be able to delete
this file (there will be garbage colleter one day, but it's work in
progress).
