#!/bin/bash
USER=$(id -u)

chown -R ${USER} ./volumes/grafana/ && \
chmod -R 777 ./volumes/grafana/