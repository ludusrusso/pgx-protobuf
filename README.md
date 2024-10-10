## Reproduction of PGX and Protobut not calling Scanner interface

This is the output of `go run main.go`:

```
❯ go run main.go
INFO[0000] calling valuer
FATA[0000] can't get tests: can't scan into dest[1]: json: cannot unmarshal string into Go struct field Test.timestamp of type timestamppb.Timestamp
exit status 1
```

It seems that the `Scanner` interface is not being called when scanning a `timestamp` field into a `timestamppb.Timestamp` field.

This leads to pgx to try to unmarshal timestamp using the `json` package, which fails. It should use the `protojson` package instead.
