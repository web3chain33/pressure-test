#!/usr/bin/env bash

function init() {
    main_jrpc="http://192.168.11.123:7905"
    echo "=========== # start set wallet 1 ============="
    echo "=========== # save seed to wallet ============="
    result=$(./ycc-cli --rpc_laddr=http://localhost:7905 seed generate -l 0)
    result=$(./ycc-cli --rpc_laddr=http://localhost:7905 seed save -p 1314fuzamei -s "${result}" | jq ".isok")
    if [ "${result}" = "false" ]; then
        echo "save seed to wallet error seed, result: ${result}"
        exit 1
    fi

    sleep 1

    echo "=========== # unlock wallet ============="
    result=$(./ycc-cli --rpc_laddr=http://localhost:7905 wallet unlock -p 1314fuzamei -t 0 | jq ".isok")
    if [ "${result}" = "false" ]; then
        exit 1
    fi

    sleep 1

    echo "=========== # create new key for transfer ============="
    transfer_addr=$(./ycc-cli --rpc_laddr=http://localhost:7905 account create -l transfer | jq ".acc.addr"| sed -r 's/"//g')
    echo "${transfer_addr}"
    if [ -z "${transfer_addr}" ]; then
        exit 1
    fi

    echo "=========== # get transfer key ============="
    transfer_prikey=$(./ycc-cli --rpc_laddr=http://localhost:7905 account dump_key -a ${transfer_addr} | jq ".data" | sed -r 's/"//g')
    echo "${transfer_prikey}"
    if [ -z "${transfer_prikey}" ]; then
        exit 1
    fi

    echo "=========== # create new key for mining ============="
    mining_addr=$(./ycc-cli --rpc_laddr=http://localhost:7905 account create -l mining | jq ".acc.addr" | sed -r 's/"//g')
    echo "${mining_addr}"
    if [ -z "${mining_addr}" ]; then
        exit 1
    fi

    echo "=========== # get mining key ============="
    mining_prikey=$(./ycc-cli --rpc_laddr=http://localhost:7905 account dump_key -a ${mining_addr} | jq ".data" | sed -r 's/"//g')
    echo "${mining_prikey}"
    if [ -z "${mining_prikey}" ]; then
        exit 1
    fi

    sleep 1
    echo "=========== # send to transfer_addr  ============="
    result=$(./ycc-cli --rpc_laddr=${main_jrpc} send coins transfer -a=100003000 -t ${transfer_addr} -k CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944)
    echo "${result}"
    if [ -z "${result}" ]; then
        exit 1
    fi

    echo "=========== # send to mining_addr  ============="
    result=$(./ycc-cli --rpc_laddr=${main_jrpc} send coins transfer -a=100003000 -t ${mining_addr} -k CC38546E9E659D15E6B4893F0AB32A06D103931A8230B0BDE71459D2B27D6944)
    echo "${result}"
    if [ -z "${result}" ]; then
        exit 1
    fi

    sleep 1
    echo "=========== # transfer_addr deposit to pos33  ============="
    deposit_hash=$( ./ycc-cli --rpc_laddr=${main_jrpc} send coins transfer -a=100000000 -t 1Wj2mPoBwJMVwAQLKPNDseGpDNibDt9Vq -k ${transfer_prikey})
    echo "${deposit_hash}"
    if [ -z "${deposit_hash}" ]; then
        exit 1
    fi
   
    sleep 1
    echo "=========== # transf_account entrust mining  ============="
    entrust_hash=$(./ycc-cli --rpc_laddr=${main_jrpc} send pos33 entrust -a 1000000 -e ${mining_addr}  -r ${transfer_addr} -k ${transfer_prikey})
    echo "${entrust_hash}"
    if [ -z "${entrust_hash}" ]; then
        exit 1
    fi

    sleep 1
    echo "=========== # check deposit and entrust result  ============="
    transfer_account=$(./ycc-cli --rpc_laddr=${main_jrpc} account balance -a ${transfer_addr}  | jq ".execAccount[0]")
    echo "${transfer_account}"
    if [[${transfer_account}==""]||[${transfer_account}|jq ".execer"!="pos33"]||[${transfer_account}|jq ".account.frozen"<1000000]]; then
        exit 1
    fi


    echo "=========== # end  ============="

}

init
