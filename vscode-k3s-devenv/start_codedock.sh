#!/bin/bash

nohup dockerd &> /tmp/dockerd.out &
nohup k3s server --docker &> /tmp/k3s.out &
code-server --bind-addr=0.0.0.0:8080