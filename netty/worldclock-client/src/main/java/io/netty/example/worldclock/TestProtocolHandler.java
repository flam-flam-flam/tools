package io.netty.example.worldclock;

import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.ChannelInboundHandlerAdapter;

import java.util.ArrayList;
import java.util.List;

public class TestProtocolHandler extends ChannelInboundHandlerAdapter {
    @Override
    public void channelRead(ChannelHandlerContext ctx, Object msg) throws Exception {
        System.out.println("====================channelRead===================");
    }

    @Override
    public void channelActive(ChannelHandlerContext ctx) {
        System.out.println("====================channelActive===================");
        TestProtocol2.RichMan.Builder builder = TestProtocol2.RichMan.newBuilder();
        builder.setName("高创");
        builder.setId(1);
        builder.setEmail("gaochuang@163.com");

        List<TestProtocol2.RichMan.Car> cars = new ArrayList<TestProtocol2.RichMan.Car>();
        TestProtocol2.RichMan.Car car1 = TestProtocol2.RichMan.Car.newBuilder().setName("上海大众超跑").setType(TestProtocol2.RichMan.CarType.DASAUTO).build();
        TestProtocol2.RichMan.Car car2 = TestProtocol2.RichMan.Car.newBuilder().setName("Aventador").setType(TestProtocol2.RichMan.CarType.LAMBORGHINI).build();
        TestProtocol2.RichMan.Car car3 = TestProtocol2.RichMan.Car.newBuilder().setName("奔驰SLS级AMG").setType(TestProtocol2.RichMan.CarType.BENZ).build();

        cars.add(car1);
        cars.add(car2);
        cars.add(car3);

        builder.addAllCars(cars);
        ctx.writeAndFlush(builder.build());
    }
    @Override
    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {
        cause.printStackTrace();
        ctx.close();
    }
}
