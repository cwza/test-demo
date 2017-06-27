
# Install docker client
set -x
VER="17.03.0-ce"
curl -L -o /tmp/docker-$VER.tgz https://get.docker.com/builds/Linux/x86_64/docker-$VER.tgz
tar -xz -C /tmp -f /tmp/docker-$VER.tgz
mv /tmp/docker/* /usr/bin

# Install heroku cli
curl -L -o /tmp/heroku.tar.gz https://cli-assets.heroku.com/heroku-cli/channels/stable/heroku-cli-linux-x64.tar.gz
tar -xz -C /tmp -f heroku.tar.gz
mkdir -p /usr/local/lib /usr/local/bin
mv /tmp/heroku-cli-v* /usr/local/lib/heroku
ln -s /usr/local/lib/heroku/bin/heroku /usr/local/bin/heroku

# push to heroku
heroku container:login
heroku container:push web --app immense-fortress-68091
