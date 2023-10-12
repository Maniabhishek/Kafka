### Replication
> - In kafka each broker contains some data , what if one broker fails down, the data will be lost , Apache Kafka enables a feature of replication to secure data loss even when a broker fails down.
> - To do so, a replication factor is created for the topics contained in any particular broker.
> - A replication factor is the number of copies of data over multiple brokers.
> - The replication factor value should be greater than 1 always (between 2 or 3). This helps to store a replica of the data in another broker from where the user can access it.
> - A broker can only host a single replica for a partition. So if your cluster has 3 brokers, the maximum replication factor you can have is 3.
> - For example, suppose we have a cluster containing three brokers say Broker 1, Broker 2, and Broker 3. A topic, namely Topic-X is split into Partition 0 and Partition 1 with a replication factor of 2.
> - <img width="520" alt="image" src="https://github.com/Maniabhishek/Kafka/assets/31520295/df50e3a0-8f3d-4b2d-82fa-3009cbb31f58">
> - Thus, we can see that Partition 0 of Topic-x is having its replicas in Broker 1 and Broker 2. Also, Partition1 of Topic-x is having its replication in Broker 2 and Broker 3.
> - It is obvious to have confusion when both the actual data and its replicas are present. The cluster may get confuse that which broker should serve the client request. To remove such confusion, the following task is done by Kafka:

- how kafka solves this confusion:
> - It chooses one of the broker's partition as a leader, and the rest of them becomes its followers.
> - The followers(brokers) will be allowed to synchronize the data. But, in the presence of a leader, none of the followers is allowed to serve the client's request. These replicas are known as ISR(in-sync-replica). So, Apache Kafka offers multiple ISR(in-sync-replica) for the data.
> - Therefore, only the leader is allowed to serve the client request. The leader handles all the read and writes operations of data for the partitions. The leader and its followers are determined by the zookeeper(discussed later).
> - If the broker holding the leader for the partition fails to serve the data due to any failure, one of its respective ISR replicas will takeover the leadership. Afterward, if the previous leader returns back, it tries to acquire its leadership again.

### Let's see an example
> - Suppose, a cluster with the following three brokers 1,2, and 3. A topic x is present having two partitions and with replication factor=2.
> - <img width="501" alt="image" src="https://github.com/Maniabhishek/Kafka/assets/31520295/72253215-f58e-4d87-867d-eadbc6031c97">
