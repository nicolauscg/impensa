echo "$BACKEND_ENV" | base64 -d > ./.env
echo "$API_SECRET" | base64 -d > ./impensa_be_api_secret.txt
echo "$DB_CONN_STRING" | base64 -d > ./impensa_be_mgoconnstring.txt
scp -q ./.env ./impensa_be_api_secret.txt ./impensa_be_mgoconnstring.txt $HOST:~/
echo "init_backend.sh finished"
