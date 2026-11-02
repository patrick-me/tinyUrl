#!/usr/bin/env bash
#
# Copies every key from one Redis-compatible container to another, keeping
# TTLs. Used to move tinyUrl data from the old Redis 8 image to Valkey: their
# data files are not compatible, and DUMP/RESTORE fails for the same reason,
# so keys are copied with GET/SET.
#
# tinyUrl stores plain strings (short code -> URL), so this only handles
# string keys whose values have no newlines or tabs.
#
# Usage: scripts/copy-redis-keys.sh <source-container> <target-container>

set -euo pipefail

if [[ $# -ne 2 ]]; then
  echo "usage: $0 <source-container> <target-container>" >&2
  exit 2
fi
src="$1"
dst="$2"

# Valkey ships valkey-cli, Redis ships redis-cli.
cli() { echo 'c=$(command -v valkey-cli || command -v redis-cli)'; }

# Source: one "key<TAB>pttl<TAB>value" line per key.
docker exec "$src" sh -c "$(cli); \$c --scan --count 1000 | while IFS= read -r k; do
    printf '%s\t%s\t%s\n' \"\$k\" \"\$(\$c pttl \"\$k\")\" \"\$(\$c --raw get \"\$k\")\"
  done" |
# Target: SET with PX when the key has a TTL.
docker exec -i "$dst" sh -c "$(cli); n=0; while IFS=\$(printf '\t') read -r k t v; do
    if [ \"\$t\" -gt 0 ] 2>/dev/null; then \$c set \"\$k\" \"\$v\" px \"\$t\" >/dev/null; else \$c set \"\$k\" \"\$v\" >/dev/null; fi
    n=\$((n+1))
  done; echo \"copied: \$n\""

# Every source key must now exist on the target with the same value. The
# target may hold more keys: links created after it went live. Values are
# compared by md5 so that a value the copy above could not carry (one with a
# newline or tab) shows up as a mismatch instead of passing unnoticed.
size() { docker exec "$1" sh -c "$(cli); \$c dbsize"; }
echo "source: $(size "$src") keys, target: $(size "$dst") keys"

bad=$(docker exec "$src" sh -c "$(cli); \$c --scan --count 1000 | while IFS= read -r k; do
    printf '%s\t%s\n' \"\$k\" \"\$(\$c --raw get \"\$k\" | md5sum | cut -d' ' -f1)\"
  done" |
docker exec -i "$dst" sh -c "$(cli); bad=0; while IFS=\$(printf '\t') read -r k h; do
    [ \"\$(\$c --raw get \"\$k\" | md5sum | cut -d' ' -f1)\" = \"\$h\" ] || bad=\$((bad+1))
  done; echo \$bad")

if [[ "$bad" != "0" ]]; then
  echo "ERROR: $bad source keys are missing or different on the target" >&2
  exit 1
fi
echo "ok: every source key is on the target with the same value"
