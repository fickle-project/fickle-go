# Contributing

## ブランチ

| ブランチ名 | 用途 |
|-|-|
| `feat/*` | 機能追加 |
| `fix/*` | バグ修正 |
| `docs/*` | ドキュメント類 |
| `chore/*` | 雑多 |

## コミット

```bash
# 例: 機能に関するコミット
git commit -m 'feat(機能名): 変更内容'
```

コミットメッセージのフォーマットは[onventional-changelog/commitlint](https://github.com/conventional-changelog/commitlint)の`README.md`を参照してください。

## ディレクトリ構成（フォルダ構成）

ドメイン駆動設計・レイヤードアーキテクチャを採用しています。

| path | レイヤー | 用途 |
|-|-|-|
|`presentation/` | プレゼンテーション層 | ユースケース層を扱い、インターフェイスを形成する |
|`application/` | アプリケーション層 | ドメイン層のユースケースを実装し、プレゼンテーション層への仲介を行う。 |
|`domain/` | ドメイン層 | ビジネスルール・ロジックを実装、リポジトリ・ファクトリのインターフェイスを定義 |
|`infrastructure/datasource/` | インフレストラクチャー層 | ドメイン層で定義したリポジトリインターフェイスを実装 |