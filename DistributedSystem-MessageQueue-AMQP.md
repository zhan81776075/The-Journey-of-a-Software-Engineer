# AMQP论文总结
## Introduction
## Transport
AMQP网络是由通过link连接的Node组成，Nodes是具有命名的实体，负责安全存储(Safe storage)和传递消息（delivery message）。message可以通过node发送，终止或者传递。
Link是连接两个node的单向通道(unidirectional route)。

### 2.4 Connections
AMQP Connections分为多个单向Channels。一个Connection Endpoint包含两种Channel endpoints：incoming and outgoing。Connection Endpoint会根据incoming Channel number，将除了open和close之外的incoming Frames映射到incoming Channel endpoint，并中继由outgoing Channel endpoints产生的帧，在发送之前用相关的outgoing Channel Number标记这些帧。

这就要求Connection Endpoint包含两个映射。一个是从incoming Channel number到incoming Channel endpoint的映射，另一个是从outgoing Channel Endpoint到outgoing Channel Number的映射。
![image](https://github.com/zhan81776075/The-Journey-of-a-Software-Engineer/assets/39268323/444fd451-1647-46c0-95f6-93bcc7c9a5d2)


通道是单向的，因此在每个连接端点，传入和传出通道是完全不同的。通道编号的作用域是相对于方向而言的，因此输入和输出通道之间没有因果关系，而这些通道恰好由 "相同 "的编号标识。这意味着，如果一个双向端点是由一个传入通道端点和一个传出通道端点构成的，则用于传入帧的通道编号不一定与用于传出帧的通道编号相同。
虽然严格来说，开始和结束报文不是针对连接端点的，但连接端点截获这些报文可能很有用，因为这些报文是会话标记特定通道上通信开始和结束的方式（请参阅第 2.5 节会话）。

# AMQP问题
## Q: AMQP协议的目标是什么?
AMQP是用于业务消息传递的Internet协议

高级消息队列协议（AMQP）是一个用于在应用程序或组织之间传递业务消息的开放标准。它连接系统，为业务流程提供它们所需的信息，并可可靠地传输实现其目标的指令。

**关键功能**
AMQP跨以下方面进行连接：
组织 - 不同组织中的应用程序
技术 - 不同平台上的应用程序
时间 - 系统不需要同时可用
空间 - 在远距离或劣质网络上可靠运行
