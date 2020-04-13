cp $BACKEND_ENV ./.env
cp $API_SECRET ./impensa_be_api_secret.txt
cp $DB_CONN_STRING ./impensa_be_mgoconnstring.txt
scp ./.env ./impensa_be_api_secret.txt ./impensa_be_mgoconnstring.txt $HOST:~/
echo "init_backend.sh finished"
