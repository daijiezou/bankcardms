#!/bin/bash
#syslog
syslogd
# start service
/usr/bank-card-ms -env release -config /usr/config/config.yml
