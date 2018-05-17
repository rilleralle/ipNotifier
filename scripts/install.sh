#!/usr/bin/env bash

# Copy files to opt folder
mkdir -p /opt/ipnotifier
cp ipnotifier /opt/ipnotifier/
cp index.html /opt/ipnotifier/

if [ $# -eq 0 ]; then
    echo "Please pass argument 'client' or 'server'";
elif [ $1 == "client" ]; then
    echo "Install IpNotifier Client";
    cp ipnotifierclient.service /etc/systemd/system/
    systemctl enable ipnotifierclient.service
    systemctl start ipnotifierclient.service
elif [ $1 == "server" ]; then
    echo "Install IpNotifier Server";
    cp ipnotifierserver.service /etc/systemd/system/
    systemctl enable ipnotifierserver.service
    systemctl start ipnotifierserver.service
else
    echo "Unknown argument. Please pass argument 'client' or 'server'";
fi

