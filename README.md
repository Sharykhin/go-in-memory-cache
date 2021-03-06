# In-memory cache

Simple implementation of Redis-like in-memory cache

Desired features:
- Key-value storage with string, lists, dict support
- Per-key TTL
- Operations:
  - Get
  - Set
  - Remove
- Golang API client
- Telnet-like/HTTP-like API protocol
- Provide some tests, API spec, deployment docs without full coverage, just a few cases and some examples of telnet/http calls to the server.

Optional features:
- persistence to disk/db
- scaling(on server-side or on client-side, up to you)
- auth
- performance tests
- Operations:
  - Update
  - Keys
- wal

##### Todo list:
- [ ] Set string key
- [ ] Set lists with (https://redis.io/topics/data-types#lists)
- [ ] Set dict (https://redis.io/topics/data-types#hashes)
- [ ] Set TTL
- [ ] New Set on key with TTL should drop previous ttl
- [ ] Get string key
- [ ] Get lists, range, concrete index
- [ ] Get dict
- [ ] Remove key
- [ ] Get all keys
- [ ] Get stats
- [ ] Export to json
- [ ] Import from json
- [ ] Telnet like protocol
- [ ] Golang API client
- [ ] Persistence to disk/db
- [ ] Auth
- [ ] Scaling (Replica set, more space with more nodes)
- [ ] Benchmarks
- [ ] Unit Tests
- [ ] Configure docker to run all in one click
- [ ] Provide some use cases
