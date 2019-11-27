#!/usr/bin/env bash


while [ -n "$1" ]; do # while loop starts

    case "$1" in

    -m)
        master="$2"
        echo "master_ip(-m) option passed, with value $master"
        shift
        ;;
    -t)
        token="$2"
        echo "token(-t) option passed, with value $token"
        shift
        ;;

    -h)
        hash="$2"
        echo "hash(-h) option passed, with value $hash"
        shift
        ;;

    --)
        shift # The double dash makes them parameters
        break
        ;;

    *) echo "Option $1 not recognized" ;;

    esac

    shift

done

kubeadm join $master:6443 --token $token  --discovery-token-ca-cert-hash $hash
#kubectl label node cb2.4xyz.couchbase.com[name]  node-role.kubernetes.io/worker=worker

