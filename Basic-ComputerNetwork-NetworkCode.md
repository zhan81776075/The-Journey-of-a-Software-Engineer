- [send函数和write函数的区别是什么](https://github.com/zhan81776075/Learning/blob/main/%E7%BD%91%E7%BB%9C%E7%BC%96%E7%A8%8B.md#send%E5%87%BD%E6%95%B0%E5%92%8Cwrite%E5%87%BD%E6%95%B0%E7%9A%84%E5%8C%BA%E5%88%AB%E6%98%AF%E4%BB%80%E4%B9%88)

## send函数和write函数的区别是什么
在网络编程中，`send` 和 `write` 函数都可以用来发送数据，但它们之间有一些关键的区别。
1. **来源与用途**：
	* `write`：这是一个系统调用，通常用于向文件或标准输出/输入写入数据。当用于套接字（socket）时，其行为类似于 `send`。然而，`write` 在处理套接字时可能不如 `send` 灵活。
	* `send`：这是一个专门的网络编程接口，用于在套接字上发送数据。与 `write` 相比，`send` 提供了更多的控制和选项，例如指定标志来影响发送行为。
2. **返回值**：
	* 两者在成功时都返回发送的字节数，在失败时返回 -1。但 `send` 函数可以返回特殊的值来表示某些情况，例如当套接字缓冲区已满时。
3. **标志**：
	* `send` 函数允许你传递标志来改变发送行为。例如，你可以使用 `MSG_DONTWAIT` 标志来指示函数在无法立即发送数据时立即返回，而不是阻塞。这对于非阻塞套接字特别有用。
	* `write` 函数不接受此类标志，其行为更简单和直接。
4. **可移植性**：
	* `send` 是 POSIX 标准的一部分，因此在不同的操作系统和平台上通常具有一致的行为。
	* `write` 的行为可能因操作系统而异，特别是在处理套接字时。
5. **使用建议**：
	* 如果你正在进行网络编程，并且需要更多的控制和灵活性，建议使用 `send` 函数。
	* 如果你只是向文件或标准输出/输入写入数据，或者你的代码需要在多种操作系统上运行并且需要保持一致性，那么 `write` 可能是一个更好的选择。

总的来说，`send` 和 `write` 之间的选择取决于你的具体需求和上下文。在网络编程中，`send` 通常提供更多的控制和灵活性，而 `write` 则更简单和直接。
