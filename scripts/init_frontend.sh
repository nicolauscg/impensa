echo "$FRONTEND_ENV" | base64 -d > ./frontend/.env
ssh $HOST 'mkdir -p frontend'
scp -q ./frontend/.env $HOST:~/frontend
echo "init_frontend.sh finished"
