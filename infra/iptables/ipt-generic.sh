#!/bin/sh

# iptables configuration for DNS, HTTP(S), and SSH ports
# for a multi-purpose server:
# - access to SSH for editing
# - access to DNS port for remote calls
# - access to HTTP to forward to HTTPS
# - access to HTTPS for multiple services
#

# * HOST SPECIFIC SSH PORT: 7822 *

# TODO Make script host agnostic through script argument

DNS_PORT=53
HTTPS_PORT=443
HTTP_PORT=80

# HOST SPECIFIC uses 7822, default is 22
SSH_PORT=7822


# Initial setup

# Flush the previous settings
iptables -F
iptables -t nat -F
iptables -t mangle -F
iptables -X

# Drop null packets
iptables -A INPUT -p tcp --tcp-flags ALL NONE -j DROP

# Drop invalid packets
iptables -A INPUT -m state --state INVALID -j DROP

# Drop syn-flood packets
iptables -A INPUT -p tcp ! --syn -m state --state NEW -j DROP

# Drop christmas-tree packets
iptables -A INPUT -p tcp --tcp-flags ALL ALL -j DROP

# Reduce pings
iptables -A INPUT -p icmp -m icmp --icmp-type address-mask-request -j DROP
iptables -A INPUT -p icmp -m icmp --icmp-type timestamp-request -j DROP
iptables -A INPUT -p icmp -m icmp --icmp-type any -m limit --limit 1/second -j ACCEPT 


# Multipurpose server configuration

# Accept SSH input
iptables -A INPUT -p tcp --dport ${SSH_PORT} -m state --state NEW,ESTABLISHED -j ACCEPT

# Accept established SSH output
iptables -A OUTPUT -p tcp --sport ${SSH_PORT} -m state --state ESTABLISHED -j ACCEPT

# Accept DNS input
iptables -A INPUT -p udp --sport ${DNS_PORT} -m state --state NEW,ESTABLISHED -j ACCEPT
iptables -A INPUT -p tcp --sport ${DNS_PORT} -m state --state NEW,ESTABLISHED -j ACCEPT

# Accept DNS output
iptables -A OUTPUT -p udp --dport ${DNS_PORT} -m state --state NEW,ESTABLISHED -j ACCEPT
iptables -A OUTPUT -p tcp --dport ${DNS_PORT} -m state --state NEW,ESTABLISHED -j ACCEPT

# Accept established input on port 80
iptables -A INPUT -p tcp --dport ${HTTP_PORT} -m state --state ESTABLISHED -j ACCEPT
iptables -A INPUT -p tcp --sport ${HTTP_PORT} -m state --state NEW,ESTABLISHED -j ACCEPT

# Accept output on port 80
iptables -A OUTPUT -p tcp --dport ${HTTP_PORT} -m state --state NEW,ESTABLISHED -j ACCEPT
iptables -A OUTPUT -p tcp --sport ${HTTP_PORT} -m state --state ESTABLISHED -j ACCEPT

# Accept input on port 443
iptables -A INPUT -p tcp --dport ${HTTPS_PORT} -m state --state NEW,ESTABLISHED -j ACCEPT
iptables -A INPUT -p tcp --sport ${HTTPS_PORT} -m state --state ESTABLISHED -j ACCEPT

# Accept output on port 443
iptables -A OUTPUT -p tcp --dport ${HTTPS_PORT} -m state --state NEW,ESTABLISHED -j ACCEPT
iptables -A OUTPUT -p tcp --sport ${HTTPS_PORT} -m state --state ESTABLISHED -j ACCEPT

# Set loopback
iptables -A INPUT -i lo -j ACCEPT
iptables -A OUTPUT -o lo -j ACCEPT


# Drop unmentioned packets by default

# Drop all other packets not excplicitly accepted
iptables -P INPUT DROP
iptables -P OUTPUT DROP
iptables -P FORWARD DROP