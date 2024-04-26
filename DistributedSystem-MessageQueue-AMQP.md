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
在关闭Connection之前，每个peer都必须写入一个close frame。在写入后，peer应当在合理的时间内继续处理connection中的数据，直到收到了peer的close frame(或在合理的超时时间后关闭)。虽然close frame可以在任意一个channel上发送，但是依然建议在0号channel上面发送(如果是pipline形式，则必须是0号channel)。
![image](https://github.com/zhan81776075/The-Journey-of-a-Software-Engineer/assets/39268323/07aada56-a8d3-46e4-a5eb-06120fc3c7f8)

### 2.4.4	Simultaneous Close
可能存在两个endpoints出于各自的原因同时Close Connection。在这种情况下，从每个endpoint的角度来看，唯一可观察到的潜在差异就是表示关闭原因的代码。
![image](https://github.com/zhan81776075/The-Journey-of-a-Software-Engineer/assets/39268323/c68c5d7b-1d72-4d9f-a350-91b0b53bcc08)

### 2.4.5	Idle Time Out Of A Connection
Connection需要支持超时关闭，关闭超时的阈值在open frame里面设置。

空闲超时的使用是对任何网络协议级控制的补充。实施方案应尽可能使用 TCP keep-alive。

如果对等方需要满足发送流量的需要以防止空闲超时，但又没有东西要发送，那么它可以发送一个空帧，即一个只包含帧头而没有帧体的帧。该帧的Channel可以是Channel最大值以内的任何有效Channel，否则将被忽略。实现者应使用通道 0 来处理空帧，如果尚未协商通道最大值（即在接收到开放帧之前），则必须使用通道 0。除此以外，空帧没有任何意义。
空帧只能在发送开放帧之后发送。由于空帧是一个帧，因此不应在关闭帧发送后发送。

如果在操作过程中，peer超过了远程peer的空闲超时阈值（比如网络负载过重），那么它应该通过使用带有错误说明的关闭帧来优雅地关闭连接。

### 2.4.6	Connection States
| Connection States      | Description |
| ----------- | ----------- |
| START       | In this state a Connection exists, but nothing has been sent or received. This is the state an implementation would be in immediately after performing a socket connect or socket accept.       |
| HDR RCVD    | In this state the Connection header has been received from our peer, but we have not yet sent anything.        |
| HDR SENT    | In this state the Connection header has been sent to our peer, but we have not yet received anything.        |
| OPEN PIPE   | In this state we have sent both the Connection header and the open frame, but we have not yet received anything.        |
| OC PIPE     | In this state we have sent the Connection header, the open frame, any pipelined Connection traffic, and the close frame, but we have not yet received anything.       |
| OPEN RCVD   | In this state we have sent and received the Connection header, and received an open frame from our peer, but have not yet sent an open frame. |
| OPEN SENT   | In this state we have sent and received the Connection header, and sent an open frame to our peer, but have not yet received an open frame.        |
| CLOSE PIPE  | In this state we have send and received the Connection header, sent an open frame, any pipelined Connection traffic, and the close frame, but we have not yet received an open frame.       |
| OPENED      | In this state the Connection header and the open frame have both been sent and received.       |
| CLOSE RCVD  | In this state we have received a close frame indicating that our partner has initiated a close. This means we will never have to read anything more from this Connection, however we can continue to write frames onto the Connection. If desired, an implementation could do a TCP half-close at this point to shutdown the read side of the Connection.       |
| CLOSE SENT  | In this state we have sent a close frame to our partner. It is illegal to write any- thing more onto the Connection, however there may still be incoming frames. If desired, an implementation could do a TCP half-close at this point to shutdown the write side of the Connection.      |
| DISCARDING  | The DISCARDING state is a variant of the CLOSE SENT state where the close is triggered by an error. In this case any incoming frames on the connec- tion MUST be silently discarded until the peer’s close frame is received.        |
| END         | In this state it is illegal for either endpoint to write anything more onto the Connection. The Connection may be safely closed and discarded.        |

### 2.4.7	Connection State Diagram
The graph below depicts a complete state diagram for each endpoint. The boxes represent states, and the arrows represent state transitions. Each arrow is labeled with the action that triggers that particular transition.
![image](https://github.com/zhan81776075/The-Journey-of-a-Software-Engineer/assets/39268323/2b3aaa73-5c92-48d8-8814-39383115b7f3)

![image](https://github.com/zhan81776075/The-Journey-of-a-Software-Engineer/assets/39268323/b2adb550-674e-4189-a400-bd5b5f8d7d9c)

## 2.5 Sessions
Session是两个Containers之间的双向顺序对话(bidirectional sequential conversation)，为相关Links提供分组。Session是Link通信的上下文。一个给定的Session可以连接任意数量、任意方向的Link。不过，一个Link一次最多只能连接到一个Session。

Link上传输的信息在Session中按顺序标识。Session可被视为Link流量的复用，就像connection复用Session流量一样。不过，与connection上的Session不同，Session上的Link并非完全独立，因为它们共享一个适用于Session的共同传输序列（common delivery sequence）。这种共同序列允许endpoints有效地引用交付集(sets of deliveries regardless)，而不考虑源Link。当一个应用程序通过大量不同的Link接收信息时，这一点尤为重要。在这种情况下，Session可将原本独立的Link聚合为一个流，接收应用程序可有效地确认该流。

![image](https://github.com/zhan81776075/The-Journey-of-a-Software-Engineer/assets/39268323/3b8d965d-056c-446f-89a4-f3dffc1fc032)

### 2.5.1	Establishing A Session
建立Sessions的方法是创建一个SessionsEndpoint，将其分配给一个未使用的Channel编号，然后发送一个 begin，宣布SessionsEndpoint与传出Channel的关联。Remote Partner收到 begin 后，会检查远程Channel字段，发现该字段为空。这表明 begin 指的是远程发起的Sessions。因此，伙伴将为远程启动的Sessions分配一个未使用的outging Channel，并通过发送自己的 begin 将remote-channe字段设置为远程启动的Sessions的incoming channel表明这一点。

为便于监控 AMQP Sessions，建议实施方案始终分配可用的最低未用channel号。

对于本地发起的Sessions，begin 帧的远程Channel字段必须为空，而在宣布远程发起的Sessions所创建的Endpoint时，必须设置远程Channel字段。
![image](https://github.com/zhan81776075/The-Journey-of-a-Software-Engineer/assets/39268323/df403950-2031-4bcb-9170-19b22ada2de7)



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
