package com.ukelink.imservice.server;

import io.netty.channel.ChannelHandlerContext;
import io.netty.channel.SimpleChannelInboundHandler;

import java.util.Calendar;
import java.util.TimeZone;

import static java.util.Calendar.*;

public class NettyServerHandler extends SimpleChannelInboundHandler<WorldClockProtocol.Locations> {
    @Override
    public void channelRead0(ChannelHandlerContext ctx, WorldClockProtocol.Locations locations) {
        long currentTime = System.currentTimeMillis();
        System.out.println("hello world read0000000 ------------------===============================");
        WorldClockProtocol.LocalTimes.Builder builder = WorldClockProtocol.LocalTimes.newBuilder();
        for (WorldClockProtocol.Location l: locations.getLocationList()) {
            TimeZone tz = TimeZone.getTimeZone(
                    toString(l.getContinent()) + '/' + l.getCity());
            System.out.println(toString(l.getContinent()) + '/' + l.getCity());
            Calendar calendar = getInstance(tz);
            calendar.setTimeInMillis(currentTime);

            builder.addLocalTime(WorldClockProtocol.LocalTime.newBuilder().
                    setYear(calendar.get(YEAR)).
                    setMonth(calendar.get(MONTH) + 1).
                    setDayOfMonth(calendar.get(DAY_OF_MONTH)).
                    setDayOfWeek(WorldClockProtocol.DayOfWeek.valueOf(calendar.get(DAY_OF_WEEK))).
                    setHour(calendar.get(HOUR_OF_DAY)).
                    setMinute(calendar.get(MINUTE)).
                    setSecond(calendar.get(SECOND)).build());
        }
        ctx.write(builder.build());
    }

    @Override
    public void channelReadComplete(ChannelHandlerContext ctx) {
        ctx.flush();
    }

    @Override
    public void exceptionCaught(ChannelHandlerContext ctx, Throwable cause) {
        cause.printStackTrace();
        ctx.close();
    }
    private static String toString(WorldClockProtocol.Continent c) {
        return c.name().charAt(0) + c.name().toLowerCase().substring(1);
    }
}
