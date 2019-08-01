package io.netty.example.worldclock;

import io.netty.bootstrap.Bootstrap;
import io.netty.channel.Channel;
//import io.netty.channel.ChannelFuture;
import io.netty.channel.EventLoopGroup;
import io.netty.channel.nio.NioEventLoopGroup;
import io.netty.channel.socket.nio.NioSocketChannel;
import org.springframework.stereotype.Component;

import java.util.Arrays;
import java.util.List;

@Component
public class WorldClockClient {
//    static final List<String> CITIES = Arrays.asList(System.getProperty(
//            "cities", "Asia/Seoul,Europe/Berlin,America/Los_Angeles").split(","));
    public static void init() throws Exception
    {
        System.out.println("hello world init  client------------------===============================");
        EventLoopGroup group = new NioEventLoopGroup();
        try {
            Bootstrap b = new Bootstrap();
            b.group(group)
                    .channel(NioSocketChannel.class)
                    .handler(new WorldClockClientInitializer());

            // Make a new connection.

            Channel ch = b.connect("10.100.106.79", 8083).sync().channel();


            // Get the handler instance to initiate the request.
//            WorldClockClientHandler handler = ch.pipeline().get(WorldClockClientHandler.class);

            // Request and get the response.
//            List<String> response = handler.getLocalTimes(CITIES);

            // Close the connection.
            ch.close();

            // Print the response at last but not least.
//            for (int i = 0; i < CITIES.size(); i ++) {
//                System.out.format("%28s: %s%n", CITIES.get(i), response.get(i));
//            }
        } finally {
            group.shutdownGracefully();
        }
    }
}
