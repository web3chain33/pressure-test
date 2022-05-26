#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import json
import os
import datetime
import time

print("开始记录高度")

while 1:
    time.sleep(1)

    val1 = os.popen('./ycc-cli net peer --rpc_laddr="http://127.0.0.1:7905"').read()
    val2 = json.loads(val1)
    peers = val2.get("peers")

    if peers is None:
        continue

    now = datetime.datetime.now().strftime("%Y-%m-%dT%H:%M:%S")
    s = f"时间：{now} \n"

    for p in peers:
        addr = p.get("addr", "空地址")
        height = p.get("header", {}).get("height", 0)
        s += f"节点地址：{addr}，高度：{height} \n"

    s += "\n"
    with open("addrheight.txt", mode="a", encoding="utf-8") as f:
        f.write(s)

