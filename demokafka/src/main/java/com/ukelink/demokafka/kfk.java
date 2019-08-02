package com.ukelink.demokafka;

import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Component;

import java.util.Optional;

/**
 * @ClassName kfk
 * @Description TODO
 * @Author chuang.gao
 * @Date 2019-8-2 18:56
 * @Version 1.0
 */
@Component
public class kfk {
    @KafkaListener(topics = {"countload"})
    public void listen(ConsumerRecord<String,String> record){
//        System.out.println("record0000: {}" + record);
        Optional<String> kafkaMessage = Optional.ofNullable(record.value());
        if(kafkaMessage.isPresent()) {
            Object message = kafkaMessage.get();
            System.out.println("record: {}" + record);
            System.out.println("message: {}" + message);
        }
    }
}

