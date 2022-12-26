# process-parent-ext

To test this code, start an osquery shell and find the path of the osquery extension socket:

```sql
osqueryi --nodisable_extensions
osquery> select value from osquery_flags where name = 'extensions_socket';
+-----------------------------------+
| value                             |
+-----------------------------------+
| /Users/USERNAME/.osquery/shell.em |
+-----------------------------------+
```

Then start the Go extension and have it communicate with osqueryi via the extension socket that you retrieved above:

```bash
go run ./my_table_plugin.go --socket /Users/USERNAME/.osquery/shell.em
```
