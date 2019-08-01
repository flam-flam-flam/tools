package com.ukelink.imservice.server;

import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;

import java.util.List;

public class TestServerHandler extends ChannelInboundHandlerAdapter {
    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception {
        TestProtocol.RichMan req = (TestProtocol.RichMan) msg;
        System.out.println(req.getName() + "他有" + req.getCarsCount() + "量车:");
        List<TestProtocol.RichMan.Car> lists = req.getCarsList();
        if (null != lists) {

            for (TestProtocol.RichMan.Car car : lists) {
                System.out.println(car.getName() + "type" + car.getTypeValue());
            }
        }
    }
    @Override
    public void channelActive(ChannelHandlerContext ctx) throws Exception {
        super.channelActive(ctx);

        //logger.info("client connected! " + ctx.toString());
        System.out.println("client connected! " + ctx.toString());
//        linkTimeout.put(ctx, System.currentTimeMillis());
//        TopicChats.getInstance().joinTopic(ctx);
    }
    @Override
    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {
        cause.printStackTrace();
        ctx.close();
    }
}
