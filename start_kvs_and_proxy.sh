#!/bin/bash

NODEID=$1
if [ $NODEID = "1" ]; then
    ./gossip-port-forward both -a 127.0.0.1 -f 20000 -l 20002 &
    ./gordolctl -l 127.0.0.1:20000 -p 127.0.0.1:20002
elif [ $NODEID = "2" ]; then
    ./gossip-port-forward both -a 127.0.0.1 -f 20004 -l 20006 &
    ./gordolctl -l 127.0.0.1:20004 -n 127.0.0.1:20000 -p 127.0.0.1:20006
elif [ $NODEID = "3" ]; then
    ./gossip-port-forward both -a 127.0.0.1 -f 20008 -l 20010 &
    ./gordolctl -l 127.0.0.1:20008 -n 127.0.0.1:20004 -p 127.0.0.1:20010
fi