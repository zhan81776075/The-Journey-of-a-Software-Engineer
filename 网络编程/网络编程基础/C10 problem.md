链接：https://en.wikipedia.org/wiki/C10k_problem

# C10 Problem
## Description
The C10k problem is the problem of optimizing network sockets to handle a large number of clients at the same time.[1] The name C10k is a numeronym for concurrently handling ten thousand connections.[2] Handling many concurrent connections is a different problem to handling many requests per second: the latter requires high throughput (processing them quickly), while the former does not have to be fast, but requires efficient scheduling of connections.

The problem of socket server optimisation has been studied because a number of factors must be considered to allow a web server to support many clients. This can involve a combination of operating system constraints and web server software limitations. According to the scope of services to be made available and the capabilities of the operating system as well as hardware considerations such as multi-processing capabilities, a multi-threading model or a single threading model can be preferred. Concurrently with this aspect, which involves considerations regarding memory management (usually operating system related), strategies implied relate to the very diverse aspects of the I/O management.[2]

## History
The term C10k was coined in 1999 by software engineer Dan Kegel,[3][4] citing the Simtel FTP host, cdrom.com, serving 10,000 clients at once over 1 gigabit per second Ethernet in that year.[1] The term has since been used for the general issue of large number of clients, with similar numeronyms for larger number of connections, most recently "C10M" in the 2010s to refer to 10 million concurrent connection.[5]

By the early 2010s millions of connections on a single commodity 1U rackmount server became possible: over 2 million connections (WhatsApp, 24 cores, using Erlang on FreeBSD),[6][7] 10–12 million connections (MigratoryData, 12 cores, using Java on Linux).[5][8]

Common applications of very high number of connections include general public servers, or private servers of very big companies whose sparse offices are connected to those servers via their virtual private networks, that have to serve thousands or even millions of users at a time, such as file servers, FTP servers, proxy servers, web servers, load balancers, and so on.[9][5]
