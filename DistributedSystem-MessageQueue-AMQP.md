# AMQP论文总结
## Introduction
## Transport
AMQP网络是由通过link连接的Node组成，Nodes是具有命名的实体，负责安全存储(Safe storage)和传递消息（delivery message）。message可以通过node发送，终止或者传递。
Link是连接两个node的单向通道(unidirectional route)。

### 2.4 Connections
AMQP Connections分为多个单向Channels。一个Connection Endpoint包含两种Channel endpoints：incoming and outgoing。Connection Endpoint会根据incoming Channel number，将除了open和close之外的incoming Frames映射到incoming Channel endpoint，并中继由outgoing Channel endpoints产生的帧，在发送之前用相关的outgoing Channel Number标记这些帧。

这就要求Connection Endpoint包含两个映射。一个是从incoming Channel number到incoming Channel endpoint的映射，另一个是从outgoing Channel Endpoint到outgoing Channel Number的映射。
![image](https://github.com/zhan81776075/The-Journey-of-a-Software-Engineer/assets/39268323/444fd451-1647-46c0-95f6-93bcc7c9a5d2)

Channel是单向的，因此在每个Connection Endpoint，incoming和outgoing Channel是完全不同的。Channel number的作用域是相对于方向而言的，因此incoming和outgoing Channel是之间没有因果关系(可以相同number)，
![image](https://github.com/zhan81776075/The-Journey-of-a-Software-Engineer/assets/39268323/a7e4cbbb-f8c0-493e-844d-d09e66a35152)

虽然严格来说，begin or end frame不是针对Connection Endpoint的，但Connection Endpoint截获这些frame可能很有用，因为这些报文是session标记特定channel上通信开始和结束的方式（请参阅第 2.5 节session）。

### 2.4.1	Opening A Connection
每个 AMQP Connection在开始其它的报文发送之前, 会通过open frame来描述这条connect的能力和限制(如MIN-MAX-FRAME-SIZE，MAX-CHANNEL-NUM), 因此，open frame只能在0号channel进行，双方在收到open Frame之后开始加入下一个状态。
![image](https://github.com/zhan81776075/The-Journey-of-a-Software-Engineer/assets/39268323/181030b5-533c-4e2b-a762-c4b84e598548)

### 2.4.2	Pipelined Open
对于使用许多short-lived的app来说，可能需要将连接协商过程管道化(pipeline the Connection negotiation process)。这种情况只要后续的报文满足这对connect的capabilities and limitations即可。
![image](https://github.com/zhan81776075/The-Journey-of-a-Software-Engineer/assets/39268323/a25610d2-773e-4d4e-b4d3-c4505cb41e31)

### 2.4.3	Closing A Connection


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
