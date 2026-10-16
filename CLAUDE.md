# Overview

- md2csnは、Markdown記法で書かれた文書をCosense記法で書かれた文書に変換するCLIツールである。

- goldmark v2はmarkdownをAST解析できる。AST解析したあと、Cosense記法に変換するGo製の高品質なCLIツールをつくりたい。引数は、md2csn (inputfile | inputfolder) -o (outputfile | outputfolder)のような感じであり、インプットはMarkdownファイルまたはフォルダを想定。複数OK。よくあるX2Y コンバータの慣習にしたがう。

- Cosense記法は仕様が明文化されているわけではない。MarkdownのASTノードと、Cosense記法の変換ルールを手作業で積み上げいく必要がある。sb-spec/scrapbox-converter/packages/md2sbは参考にはなると思う。（ただしこれは正規表現ベースなので正確ではない）

## Rules

- 後方互換性は維持しなくてよい。
- データベースのマイグレーションはPocketBaseのWEB UIから行うのでマイグレーションコードを作成する必要はない。
- When fixing bugs, add a failing regression test first.
- All errors are user-facing, so messages should be clear.
- Keep functions small and focused.
- Module files should re-export what's needed, hide implementation details.
- 変更内容を Codex形式(Search/Replace形式)で出力してください。
例）
```
mathweb/flask/app.py
<<<<<<< SEARCH
from flask import Flask
=======
import math
from flask import Flask
>>>>>>> REPLACE
```
- ファイルを削除・移動するときはrm・mvコマンドで提示する。
- 全コードを書き直すときは、古いファイルをSearch/Replaceせずに削除して、新規ファイルとして出力する。
例)
internal/parser/links_test.go (new file)
```go
package parser
import (
	"slices"
	"testing"
)
```


## Tech Stack
- Go 1.26.0 (1.27はライブラリ非対応)
- goldmark v2


