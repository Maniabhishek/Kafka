### Topic
> -  Topic in general means the heading or name given to something 
> - In Kafka, the word topic refers to a category or a common name used to store and publish a particular stream of data.
> -  Basically, topics in Kafka are similar to tables in the database, but not containing all constraints , we can create as many topics as we want.
> -  A producer publishes data to the topics, and a consumer reads that data from the topic by subscribing it.

### Partitions
> - A topic is split into several parts which are known as the partitions of the topic. These partitions are separated in an order.
> - The data content gets stored in the partitions within the topic.
> - Therefore, while creating a topic, we need to specify the number of partitions(the number is arbitrary and can be changed later).
> - Each message gets stored into partitions with an incremental id known as its Offset value.
> - The order of the offset value is guaranteed within the partition only and not across the partition. The offsets for a partition are infinite.
<img width="350" alt="partition" src="https://github.com/Maniabhishek/Kafka/assets/31520295/84bf4778-c7a6-4d90-8c56-dde7f6647b6d">
> - Suppose, a topic containing three partitions 0,1 and 2. Each partition has different offset numbers. The data is distributed among each offset in each partition where data in offset 1 of Partition 0 does not have any relation with the data in offset 1 of Partition1. But, data in offset 1of Partition 0 is inter-related with the data contained in offset 2 of Partition0

### Broker
Role of Apache Kafka 
> - A Kafka cluster is comprised of one or more servers which are known as brokers or Kafka brokers.
> - A broker is a container that holds several topics with their multiple partitions.
> - The brokers in the cluster are identified by an integer id only.
> - Kafka brokers are also known as Bootstrap brokers because connection with any one broker means connection with the entire cluster.
> - Although a broker does not contain whole data, but each broker in the cluster knows about all other brokers, partitions as well as topics.
> - <img width="350" alt="Broker" src="https://github.com/Maniabhishek/Kafka/assets/31520295/2668efb5-07d1-43d4-92e7-b29be439eba6">
This is how a broker looks like in the figure containing a topic with n number of partitions.

### Example: Brokers and Topics
> - Suppose, a Kafka cluster consisting of three brokers, namely Broker 1, Broker 2, and Broker 3.
> - <img width="350" alt="broker with partition" src="https://github.com/Maniabhishek/Kafka/assets/31520295/243113b6-13fd-4749-ac1a-b56b1c006ee7">
> - Each broker is holding a topic, namely Topic-x with three partitions 0,1 and 2. Remember, all partitions do not belong to one broker only, it is always distributed among each broker (depends on the quantity). Broker 1 and Broker 2 contains another topic-y having two partitions 0 and 1. Thus, Broker 3 does not hold any data from Topic-y. It is also concluded that no relationship ever exists between the broker number and the partition number.
