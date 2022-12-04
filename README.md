# anishiri-bot
アニメの単語でしりとりをするtwitter-bot

メンションを飛ばしたアカウントにアニメの単語で返信する処理を1分ごとに実行します。

("@{アカウントid} なんとかかんとか"でツイート→"かけぐるい"で返信を返す)
![zndxl-gvch7](https://user-images.githubusercontent.com/76876781/205491946-41d1cd51-eee7-44cf-953c-f94587773010.gif)

# 開発環境構築
* Dockerのインストール
* ElevatedにアップグレードされたTwitterのDeveloperアカウントを用意

    (参考:https://zenn.dev/eito_blog/articles/9a17a64ffa3c6b,  最後の"各種APIキーの取得"のところまで)
```
git clone git@github.com:oka0509/anishiri-bot.git
cd anishiri-bot
```
直下に.envファイルをおいてtwitter-apiのaccess_token等を設定
```
access_token=********
access_token_secret=*************
api_key=*************
api_key_secret=**********
twitter_id=********　(@は除外)
```
立ち上げ
```
docker-compose up -d
```
