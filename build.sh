#!/usr/bin/env bash

if [ $# -eq 0 ]; then
    echo "Please pass the server url address as an argument";
    echo "For example ./build.sh 192.168.0.10";
else
    echo "Start build."
    echo ""
    echo "Start building amd64 binaries and copy to bin/x86_64 folder."
    GOARCH=amd64 go build -o bin/x86_64/ipnotifier
    cp index.html bin/x86_64/
    cp ./scripts/install.sh bin/x86_64/
    cp ./scripts/ipnotifierclient.service bin/x86_64/
    sed -i "s/#IP#/$1/g" bin/x86_64/ipnotifierclient.service
    cp ./scripts/ipnotifierserver.service bin/x86_64/
    echo "Build for amd64 complete."
    echo ""

    echo "Start building arm64 binaries and copy to bin/arm64 folder."
    GOARCH=arm64 go build -o bin/arm64/ipnotifier
    cp index.html bin/arm64/
    cp ./scripts/install.sh bin/arm64/
    cp ./scripts/ipnotifierclient.service bin/arm64/
    sed -i "s/#IP#/$1/g" bin/arm64/ipnotifierclient.service
    cp ./scripts/ipnotifierserver.service bin/arm64/
    echo "Build for arm64 complete."
    echo ""
    echo "Binaries and installation files are located in bin folder."
    echo "Copy required folder arm64 or x86_64 to server and start install.sh"
fi