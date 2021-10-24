cp $BACKEND_ENV ./.env
cp $API_SECRET ./impensa_be_api_secret.txt
cp $DB_CONN_STRING ./impensa_be_mgoconnstring.txt
cp $MAILGUN_API ./impensa_be_mailgun_api.txt
cp $GOOGLE_OAUTH_CLIENT_ID ./impensa_be_google_oauth_client_id.txt
cp $GOOGLE_OAUTH_CLIENT_SECRET ./impensa_be_google_oauth_client_secret.txt
scp ./.env ./impensa_be_api_secret.txt ./impensa_be_mgoconnstring.txt ./impensa_be_mailgun_api.txt ./impensa_be_google_oauth_client_id.txt ./impensa_be_google_oauth_client_secret.txt $HOST:$FOLDER_PATH
echo "init_backend.sh finished"
