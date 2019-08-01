package com.ukelink.imservice;

//import org.springframework.beans.factory.annotation.Autowired;
import com.ukelink.imservice.server.NettyServer;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

@SpringBootApplication
public class ImServiceApplication implements CommandLineRunner {
    @Autowired
    private NettyServer nettyServer;
    public static void main(String[] args) {
        SpringApplication.run(ImServiceApplication.class, args);
    }
    @Override
    public void run(String... args) {
//        System.out.println(this.helloWorldService.getHelloMessage());
        nettyServer.init();
        System.out.println("hello world===============================");
    }
}
