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