### Kafka ?
> - It is a software platform based on distributed streaming process
> - It is a publish-subscribe messaging system which let exchanging of data between applications, servers, and processors as well.

### What is messaging system?
> - A messaging system is a simple exchange of messages between two or more persons, devices, etc.
> - A publish-subscribe messaging system allows a sender to send/write the message and a receiver to read that message.
> - n Apache Kafka, a sender is known as a producer who publishes messages, and a receiver is known as a consumer who consumes that message by subscribing it.

### What is Streaming process
> - It is the processing of data in parallelly connected systems.
> - This process allows different applications to limit the parallel execution of the data

### Streams 
> - An event stream represents related events in motion.

> - where one record executes without waiting for the output of the previous record. Therefore, a distributed streaming platform enables the user to simplify the task of the streaming process and parallel execution. Therefore, a streaming platform in Kafka has the following key capabilities:
>> - As soon as the streams of records occur, it processes it.
>> - It works similar to an enterprise messaging system where it publishes and subscribes streams of records.

>> - It stores the streams of records in a fault-tolerant durable way.
<img width="400" alt="kafka cluster" src="https://github.com/Maniabhishek/Kafka/assets/31520295/a314c484-21f7-4673-a228-636973b91d4a">

### Four core apis
> - Producer API: This api allows to publish stream of records to one or more topics
> - Consumer API: This API allows an application to subscribe one or more topics and process the stream of records produced to them.
> - Streams API: This API allows an application to effectively transform the input streams to the output streams. It permits an application to act as a stream processor which consumes an input stream from one or more topics, and produce an output stream to one or more output topics.
> - Connector API: This API executes the reusable producer and consumer APIs with the existing data systems or applications.
