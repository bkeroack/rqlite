## 3.2.0 (Unreleased)
- [PR #137](https://github.com/rqlite/rqlite/pull/137): Use resolved version of joining node's address.
- [PR #136](https://github.com/rqlite/rqlite/pull/136): Better errors on join failures.
- [PR #133](https://github.com/rqlite/rqlite/pull/133): Add Peers to status output.
- [PR #132](https://github.com/rqlite/rqlite/pull/132): Support removing a node from a cluster.
- [PR #131](https://github.com/rqlite/rqlite/pull/131): Only convert []byte from database to string if "text". Thanks @bkeroackdsc
- [PR #129](https://github.com/rqlite/rqlite/pull/129): Verify all statements sent to query endpoint begin with "SELECT".

## 3.1.0 (May 4th 2016)
- [PR #118](https://github.com/rqlite/rqlite/pull/118): New rqlite CLI. Thanks @mkideal
- [PR #125](https://github.com/rqlite/rqlite/pull/125): Add Go runtime details to status endpoint.

## 3.0.1 (May 1st 2016)
- [PR #117](https://github.com/rqlite/rqlite/pull/117): Use Raft advertise address, if exists, during join.

## 3.0.0 (May 1st 2016)
**The Raft log format in 3.0 is not compatible with the 2.0 series**. To upgrade from an earlier version to this release you should backup your 2.0 master node, and replay the database dump into a new 3.0 cluster. The HTTP API remains unchanged however.

- [PR #116](https://github.com/rqlite/rqlite/pull/116): Allow HTTP advertise address to be set.
- [PR #115](https://github.com/rqlite/rqlite/pull/115): Support advertising address different than Raft bind address.
- [PR #113](https://github.com/rqlite/rqlite/pull/113): Switch to in-memory SQLite databases by default.
- [PR #109](https://github.com/rqlite/rqlite/pull/109): Nodes broadcast meta to cluster via Raft.
- [PR #109](https://github.com/rqlite/rqlite/pull/109), [PR #111](https://github.com/rqlite/rqlite/pull/111): Leader redirection
- [PR #104](https://github.com/rqlite/rqlite/pull/104): Handle the `-join` option sensibly when already member of cluster.

## 2.2.2 (April 24th 2016)
- [PR #96](https://github.com/rqlite/rqlite/pull/96): Add build time to status output.
- [PR #101](https://github.com/rqlite/rqlite/pull/101): Fix restore to in-memory databases.

## 2.2.1 (April 19th 2016)
- [PR #95](https://github.com/rqlite/rqlite/pull/95): Correctly set HTTP authentication.

## 2.2.0 (April 18th 2016)
- [PR #84](https://github.com/rqlite/rqlite/pull/84): Encrypted API (HTTPS) now supported.
- [PR #85](https://github.com/rqlite/rqlite/pull/85): BasicAuth support.
- [PR #85](https://github.com/rqlite/rqlite/pull/85): User-level permissions support.
- Print rqlited logo on start-up.
- End-to-end single-node and multi-node unit tests.

## 2.1 (April 9th 2016)
- [PR #76](https://github.com/rqlite/rqlite/pull/76): Obey timing information display at database level.
- [PR #77](https://github.com/rqlite/rqlite/pull/77): Add version information to binary.

## 2.0 (April 5th 2016)
- `timings` URL param to control presence of timing information in response.
- [PR #74](https://github.com/rqlite/rqlite/pull/74): Use SQLite connection directly. Thanks @zmedico.
- Update operations return last-inserted ID.
- Column-oriented API responses.
- Types in API response body.
- Query times in addition to sum of query times.
- New Raft consensus module, built on Hashsicorp's implementation.
- Hot backup support.
- Selectable read-consistency levels of none, weak, and strong.
- SQLite file size added to HTTP API status endpoint.
- expvar support added to HTTP server.

## 1.0 (June 23rd 2015)
Check out [this tag](https://github.com/rqlite/rqlite/releases/tag/v1.0) for full details.

