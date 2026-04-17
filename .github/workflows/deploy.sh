name: Auto Deploy GO App to GCP

on:
  push:
    branches:
      - main

jobs:
  runs on: ubuntu-22.04

  steps:
    - name: Checkout the code:
    - uses: actions/checkout@v4

    - name: Deploy via SSH
    - uses: appleboy/ssh-action@v1.0.3

  with:
    host: ${{ secrets.GCP_SERVER_HOST }}
    user_name: ${{ secrets.GCP_SERVER_USERNAME }}
    key: ${{ secrets.GCP_SSH_PRIVATE_KEY }}

  scripts:
    cd ~/hng-stage-1
    git pull og main
    docker compose up -d --build
