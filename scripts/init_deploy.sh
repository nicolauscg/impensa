which ssh-agent || ( apk add --update openssh-client )

eval $(ssh-agent -s)
chmod 400 $DEPLOYMENT_KEY
ssh-add $DEPLOYMENT_KEY
echo "Host *\n\tStrictHostKeyChecking no\n" >> ~/.ssh/config
mkdir -p ~/.ssh
chmod 700 ~/.ssh

scp -q ./docker-compose.yml ./scripts/restart_service.sh $HOST:~/
scp -q -r ./nginx-conf $HOST:~/

echo "init_deploy.sh finished"