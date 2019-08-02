package com.ukelink.demokafka;

import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.stereotype.Component;

@SpringBootApplication
//@ComponentScan("com.ukelink.demokafka")
//@ComponentScan("com.ukelink.demokafka")
public class DemokafkaApplication {

	public static void main(String[] args) {
		SpringApplication.run(DemokafkaApplication.class, args);
	}

}
