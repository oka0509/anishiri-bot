# anishiri-bot
アニメの単語でしりとりをするtwitter-bot

# Requirement
* Dockerのインストール
* ElevatedにアップグレードされたTwitterのDeveloperアカウント

# Installation
```
git clone git@github.com:oka0509/anishiri-bot.git
cd anishiri-bot
```
直下にlocal.envファイルをおいてtwitter-apiのaccess_token等を設定
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
