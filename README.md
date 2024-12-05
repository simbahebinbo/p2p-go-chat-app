This is a terminal based go chat application 
which uses libp2p package for establishing p2p connection.

启动一个节点 A
```shell
./go-libp2p-tutorial
My peer ID: 12D3KooWBYxrE3bavZRjdbGsje6W4KNRMDhrUmRpKCKwVrMtzrrt
My multiaddresses:
        /ip4/10.125.34.99/tcp/55546/p2p/12D3KooWBYxrE3bavZRjdbGsje6W4KNRMDhrUmRpKCKwVrMtzrrt
        /ip4/127.0.0.1/tcp/55546/p2p/12D3KooWBYxrE3bavZRjdbGsje6W4KNRMDhrUmRpKCKwVrMtzrrt
Enter peer multiaddr (or 'exit' to quit):
```

启动另一个节点 B
```shell
./go-libp2p-tutorial
My peer ID: 12D3KooWP8MtEh7dMPu4wVSzpdZKm9cBb5ztT5qm9HREv4NdNnq5
My multiaddresses:
        /ip4/10.125.34.99/tcp/55574/p2p/12D3KooWP8MtEh7dMPu4wVSzpdZKm9cBb5ztT5qm9HREv4NdNnq5
        /ip4/127.0.0.1/tcp/55574/p2p/12D3KooWP8MtEh7dMPu4wVSzpdZKm9cBb5ztT5qm9HREv4NdNnq5
Enter peer multiaddr (or 'exit' to quit):
```

将节点 A 的 一个 multiaddress 粘贴到 节点 B，则 节点 B 可以连接到节点 A，从而节点 B 到节点 A 发送消息;
```shell
Connected successfully!
Enter message (or 'exit' to quit):
```
将节点 B 的 一个 multiaddress 粘贴到 节点 A，则 节点 A 可以连接到节点 B，从而节点 A 到节点 B 发送消息.
```shell
Connected successfully!
Enter message (or 'exit' to quit):
```

