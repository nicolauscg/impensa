which ssh-agent || ( apk add --update openssh-client )

eval $(ssh-agent -s)
chmod 400 $DEPLOYMENT_KEY
ssh-add $DEPLOYMENT_KEY
mkdir -p ~/.ssh
chmod 700 ~/.ssh
[[ -f /.dockerenv ]] && echo -e "Host *\n\tStrictHostKeyChecking no\n\n" > ~/.ssh/config

scp ./docker-compose.yml ./scripts/restart_service.sh $HOST:~/
scp -r ./nginx-conf $HOST:~/

echo "init_deploy.sh finished"
