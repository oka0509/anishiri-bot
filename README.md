# anishiri-bot
アニメの単語でしりとりをするtwitter-bot

# Requirement
* Dockerのインストール
* ElevatedにアップグレードされたTwitterのDeveloperアカウント

# Installation
```
git clone git@github.com:oka0509/shiritori-bot.git
cd shiritori-bot
```
直下に.envをおいてtwitter-apiのaccess_token等を設定
```
access_token=********
access_token_secret=*************
consumer_key=*************
consumer_key_secret=**********
```
立ち上げ
```
docker-compose up -d
```
