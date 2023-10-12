### Kafka Producers
> - A producer is the one which publishes or writes data to the topics within different partitions. Producers automatically know that, what data should be written to which partition and broker. The user does not require to specify the broker and the partition.

### How does the producer write data to a cluster
> - A producer uses following strategie//s to write data to the cluster: Message Keys and Acknowledgment

### Message Keys
> - Apache Kafka enables the concept of the key to send the messages in a specific order.
> - The key enables the producer with two choices, i.e., either to send data to each partition (automatically) or send data to a specific partition only.
> - Sending data to some specific partitions is possible with the message keys.
> - If the producers apply key over the data, that data will always be sent to the same partition always. But, if the producer does not apply the key while writing the data, it will be sent in a round-robin manner. This process is called load balancing.
> - In Kafka, load balancing is done when the producer writes data to the Kafka topic without specifying any key, Kafka distributes little-little bit data to each partition.
> - Therefore, a message key can be a string, number, or anything as we wish.

### There are two ways to know that the data is sent with or without a key:
> - If the value of key=NULL, it means that the data is sent without a key. Thus, it will be distributed in a round-robin manner (i.e., distributed to each partition).
> - If the value of the key!=NULL, it means the key is attached with the data, and thus all messages will always be delivered to the same partition.

### Let's see an example
> - Consider a scenario where a producer writes data to the Kafka cluster, and the data is written without specifying the key. So, the data gets distributed among each partition of Topic-T under each broker, i.e., Broker 1, Broker2, and Broker 3.
>> - <img width="350" alt="image" src="https://github.com/Maniabhishek/Kafka/assets/31520295/6337cd70-cac0-4fbc-b2e2-b8534855c1a8">
> - Consider another scenario where a producer specifies a key as Prod_id. So, data of Prod_id_1(say) will always be sent to partition 0 under Broker 1, and data of Prod_id_2 will always be in partition 1 under Broker 2. Thus, the data will not be distributed to each partition after applying the key (as saw in the above scenario).
>> - <img width="350" alt="image" src="https://github.com/Maniabhishek/Kafka/assets/31520295/0473f26e-a825-43c1-a782-950ed1239d76">


### Acknowledgment
> - In order to write data to the Kafka cluster, the producer has another choice of acknowledgment. It means the producer can get a confirmation of its data writes by receiving the following acknowledgments:
>> - acks=0: This means that the producer sends the data to the broker but does not wait for the acknowledgement. This leads to possible data loss because without confirming that the data is successfully sent to the broker or may be the broker is down, it sends another one.
>> -acks=1: This means that the producer will wait for the leader's acknowledgement. The leader asks the broker whether it successfully received the data, and then returns feedback to the producer. In such case, there is limited data loss only.
>> - acks=all: Here, the acknowledgment is done by both the leader and its followers. When they successfully acknowledge the data, it means the data is successfully received. In this case, there is no data loss.

### Let's see an example
> - Suppose, a producer writes data to Broker1, Broker 2, and Broker 3.
>> - Case1: Producer sends data to each of the Broker, but not receiving any acknowledgment. Therefore, there can be a severe data loss, and the correct data could not be conveyed to the consumers.
>> - <img width="421" alt="image" src="https://github.com/Maniabhishek/Kafka/assets/31520295/92c725b9-481a-4a18-b9de-fd6887b805f6">
>> - Case 2: The producers send data to the brokers. Broker 1 holds the leader. Thus, the leader asks Broker 1 whether it has successfully received data. After receiving the Broker's confirmation, the leader sends the feedback to the Producer with ack=1.
>> - <img width="504" alt="image" src="https://github.com/Maniabhishek/Kafka/assets/31520295/2c3e0380-c733-4474-8a49-bf0a8439a0e2">
>> - Case3: The producers send data to each broker. Now, the leader and its replica/ISR will ask their respective brokers about the data. Finally, acknowledge the producer with the feedback.
>> - <img width="524" alt="image" src="https://github.com/Maniabhishek/Kafka/assets/31520295/3dfd328a-45af-4d62-90ed-82b739109754">

