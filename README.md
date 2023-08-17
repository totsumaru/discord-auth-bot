# discord-auth-bot

## サービスURL

https://scan-discord.app

`stg`環境のURL: https://stg.scan-discord.app

## インフラ

### FrontEnd

[Vercel](https://vercel.com/)を使用しています。

[ダッシュボード](https://vercel.com/totsumaru/discord-auth-bot-fe/deployments)

- `prd`環境: `main`ブランチにpushすることでデプロイされます。
- `stg`環境: `stg`ブランチにpushすることでデプロイされます。

### Discord

導入URLは権限の変更があるため、プログラム内部を参照してください。

- 設定項目
    - Redirects(RedirectURL) -> supabaseのURL
    - Client ID -> supabaseに登録
    - Client Secret -> supabaseに登録
- `prd`環境: SCAN-dev botを使用します。
- `stg`環境: SCAN botを使用します。

### Supabase

認証,DBはsupabaseを使用します。

ログイン方法はDiscordのみとします。

- 設定項目
    - Client ID -> discordのbot(app)の情報を登録
    - Client Secret -> discordのbot(app)の情報を登録

### BackEnd

[render](https://render.com/)を使用します。

`stg`と`prd`でプロジェクトを分けて管理します。

## 支払い

- vercel: $20/month
- render(stg) : $7/month
- render(prd) : $7/month
