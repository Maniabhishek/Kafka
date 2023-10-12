### Replication
> - In kafka each broker contains some data , what if one broker fails down, the data will be lost , Apache Kafka enables a feature of replication to secure data loss even when a broker fails down.
> - To do so, a replication factor is created for the topics contained in any particular broker.
> - A replication factor is the number of copies of data over multiple brokers.
> - The replication factor value should be greater than 1 always (between 2 or 3). This helps to store a replica of the data in another broker from where the user can access it.
> - A broker can only host a single replica for a partition. So if your cluster has 3 brokers, the maximum replication factor you can have is 3.
> - For example, suppose we have a cluster containing three brokers say Broker 1, Broker 2, and Broker 3. A topic, namely Topic-X is split into Partition 0 and Partition 1 with a replication factor of 2.
> - <img width="520" alt="image" src="https://github.com/Maniabhishek/Kafka/assets/31520295/df50e3a0-8f3d-4b2d-82fa-3009cbb31f58">

