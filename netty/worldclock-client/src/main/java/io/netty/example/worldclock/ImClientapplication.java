package io.netty.example.worldclock;

import io.netty.channel.ChannelHandlerContext;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.CommandLineRunner;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import java.util.concurrent.ConcurrentHashMap;
import java.util.concurrent.ConcurrentLinkedDeque;

import static io.netty.example.worldclock.WorldClockClient.init;

@SpringBootApplication
public class ImClientapplication implements CommandLineRunner {
    @Autowired
    private WorldClockClient worldClockClient;
    public static void main(String[] args) {
        //System.out.println(cn.wildfirechat.proto.WFCMessage.getDescriptor().toProto().toString());
        SpringApplication.run(ImClientapplication.class, args);
    }
//    private String[][] conDetails = new String[][]{
//            new String[] {"Mars", "0", "STN Discuss"},
//            new String[] {"Mars", "1", "Xlog Discuss"},
//            new String[] {"Mars", "2", "SDT Discuss"}
//    };
//    public ConcurrentHashMap<String, ConcurrentLinkedDeque<String>> topicJoiners = new ConcurrentHashMap<>();
//    private static ImClientapplication topicChats = new ImClientapplication();
//
//    private ImClientapplication() {
//
//        for (int i = 0; i < conDetails.length; i++) {
//            ConcurrentLinkedDeque<String> ctxs = new ConcurrentLinkedDeque<>();
//            topicJoiners.put(conDetails[i][1], ctxs);
//        }
//        System.out.println("The set2 is: " + topicJoiners.keySet());
//        System.out.println("The topicJoiners is: " + topicJoiners);
//
//        for (String topicName : topicJoiners.keySet()) {
//            System.out.println("The  topicJoiners.get(topicName) is: " + topicJoiners.get(topicName) + ";topicname:" +topicName);
//            if (!topicJoiners.get(topicName).contains("5")) {
//                topicJoiners.get(topicName).offer("5");
//            }
//            System.out.println("The topicJoiners.get(topicName)222 is: " + topicJoiners.get(topicName));
//        }
//        System.out.println("The topicJoiners2 is: " + topicJoiners);
//    }
    public void run(String... args)  throws Exception{
        worldClockClient.init();
//        ConcurrentHashMap<String, String> hash_map
//                = new ConcurrentHashMap<String, String>();
//
//        // Mapping string values to int keys
//        hash_map.put("10", "Geeks");
//        hash_map.put("15", "4");
//        hash_map.put("20", "Geeks");
//        hash_map.put("25", "Welcomes");
//        hash_map.put("30", "You");
//
//        // Displaying the HashMap
//        System.out.println("Initial Mappings are: "
//                + hash_map);
//
//        // Using keySet() to get the set view of keys
//        System.out.println("The set is: "
//                + hash_map.keySet());
//        System.out.println("hello world,I am client ===============================");
//
//        System.out.println("Is the value 'World' present? "
//                + hash_map.contains("Welcomes"));
//        System.out.println("Is the value 'World' present? "
//                + hash_map.contains("4"));
    }
}
