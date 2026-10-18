使い方
```sh
cd testdata

md2csn < Universe.md
cat Universe.md | md2csn
cat Universe.md | md2csn -o Universe.csn
md2csn Universe.md
md2csn Universe.md -o Universe.csn
md2csn . -o output
```

```sh
mkdir -p csn

find markdown -maxdepth 1 -type f -name '*.md' -print0 |
  xargs -0 -P 4 -I {} sh -c '
    f="$1"
    out="csn/$(basename "${f%.md}").csn"
    md2csn "$f" -o "$out"
  ' _ {}
```

json
```sh
find csn -maxdepth 1 -type f -name '*.csn' -print0 |
  sort -z |
  perl -0ne 'chomp; push @f, $_; END { print join("\0", @f), "\0" }' |
  while IFS= read -r -d '' f; do
    jq -Rs --arg file "$f" '
      split("\n") as $lines |
      {
        title: ($lines[0] | gsub(" "; "_") | ascii_downcase),
        lines: $lines
      }
    ' "$f"
  done |
  jq -s '{pages: .}' > output.json
```