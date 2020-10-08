# order-proxy-be

# 環境

Go:1.13.1
Google Cloud SDK 311.0.0

# セットアップ

$ gcloud init
Choose the account you would like to use to perform operations for 
this configuration:
 [1] ********@gmail.com
 [2] Log in with a new account

 You are logged in as: [********@gmail.com].

Pick cloud project to use: 
 [1] project-name
 [2] Create a new project
Please enter numeric choice or text value (must exactly match list 
item):  

# デプロイ

$ gcloud functions deploy GetHello --runtime go113 --region us-central1 --trigger-http --allow-unauthenticated
